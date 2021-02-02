# Proxy server for bidectional stream

D:\Soft\protoc\bin\protoc.exe -I .\proto\ .\proto\chat.proto --go_out=plugins=grpc:proxygrpc        
D:\Soft\protoc\bin\protoc.exe -I .\proto\ .\proto\proxy.proto --go_out=plugins=grpc:proxygrpc

# This project shows how to develop proxy service on golang for biderectioanl stream
## Этот проект реализован прокси серверна голанг для двунаправленного стрима

### 
I cannot find any examples of proxy service for biderctional stream, so i decided if i will found a solution I published it with explanation.

The difficult part of this task - launch two different go routine that can finish their work if one of them broke or finished.
Сложность написания прокси для двунаправленного стрима - необходимость запустить два потока которые будут жить пока его контакт не разорвет соединение и пока жив второй поток
А весь метод должен ждать окончания обоих потоков.

Я не нашла ни одного примера такой реализации, поэтому расписываю как это работает.

Initialize variable sync.WaitGroup. You can read about this package here https://golang.org/pkg/sync/ 
Объявляем переменую sync.WaitGroup. Описание пакета sync и пример использования вы смотрите тут https://golang.org/pkg/sync/ .
```
	var wg sync.WaitGroup
```

The main task of wg do not finish this functioun untill go routine is done. We must say wg how many routine should end.
Функция WaitGroup  - дождаться окончания рутину. WaitGroup принимает - сколько рутин должно закончится.
```
	wg.Add(2)
```


Next - create channel. "Channels are a typed conduit through which you can send and receive values with the channel operator, <-." (c) More about channels - https://tour.golang.org/concurrency/2
With defer we make sure that syncChannel will be close when function finished.


Потом объявляем chan c int. chan - структура для получения и отправки сообщений и удостоверяемся что он закроется когда метод закончится
```
	syncChannel := make(chan int)
	defer close(syncChannel)
```
Connect to server

Следующий шаг - подключиться к серверу, проверить что нет ошибок и прописать что подключение закрывается когда функция заканчивается
```
	// Connect to server
	chatConn, err := grpc.Dial(buildServiceName(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	// Close connection when work finished
	defer chatConn.Close()
```
Initialize client and subsciribe on chat
Инициализируем клиента и подписываемся на чат

```
	// Init client
	res := chat.NewMessangerClient(chatConn)

	// Subscribe
	cl, err := res.JoinChat(stream.Context())
	if err != nil {
		return err
	}
```
Now we need to process all message from client and from server and redirect client's message to server and server's message to client.
Launch two go routine. One for client and one for server. It will be the same, so we can watch only the first routine. 

Теперь нам нужно обрабатывать приходящиие сообщения из обоих каналов - и пересылать их в соседний канал. При этом когда один из каналов прекратит работу (отпишется или получит ошибку надо уведомить и закрыть канал и закрыть всю функцию)
для этого запустим 2 го рутины - клиентскую и серверную. Логика в них будет одинаковая поэтому рассмотрим только одну из них.

When this routine ends it calls Done. wg.Wait blocks until all goroutines have finished.
Сначала мы сообщаем что как только go routine закончится необходимо закрыть sync.WaitGroup, когда оба sync.WaitGroup закроются - закончится и наш основной метод
```
defer wg.Done()
```

Endless cycle starts. At select we wait for message from syncChannel - if we received this message, we finished cycle and close routine.
После этого мы запускаем бесконечны цикл в котором проверяем сначала что к нам не прилетело уведомление из канала syncChannel  - если прилетело - выходим из цикла и go  routine завершается.
```
case <-syncChannel:
	log.Infof("Server stream was closed. Closing Client stream...")
	return
```
If we don't have any message we listen client channel, 
Message from client is error - we send message to syncChannel and finished cycle and close routine.
Message without error - we send message to server. If we have error after this step - we send message to syncChannel and finished cycle and close routine.
We use syncChannel to inform another routine that it need to shout down.


Если ничего в syncChannel не прилетело - получаем сообщение из канала.
Сообщение с ошибкой - кладем сообщение в syncChannel и выходим из цикла.
Сообщение без ошибки отправлем его клиенту. 
Если отправилось не удачно - пишем в лог,  кладем сообщение в syncChannel и выходим из цикла.
Как вы помните syncChannel нужен как раз для того чтоб сообщить второй go routine что нужно закрываться.
```
msg, err := stream.Recv()
if err != nil {
	if err != io.EOF {
		log.Errorf("JoinChat Client stream has failed with error: %v", err.Error())
	}
	syncChannel <- 1
	return
}
log.Debugf("Received message from Client stream %v", msg)

err = cl.Send(msg)
if err != nil {
	log.Errorf("JoinChat Client stream has failed with error: %v", err.Error())
	syncChannel <- 1
	return
}
```
Second go routine is the same. 

Вторую go routine обрабатываем анологично.
```
// Init Server stream
go func() {
	defer wg.Done()
	for {
		select {
		case <-syncChannel:
			log.Infof("Client stream was closed. Closing Server stream...")
			return
		default:
			msg, err := cl.Recv()
			if err != nil {
				log.Errorf("JoinChat Server stream has failed with error: %v", err.Error())
				stream.Context().Done()
				syncChannel <- 1
				return
			}
			log.Debugf("Received message from Server stream %v", msg)

			err = stream.Send(msg)
			if err != nil {
				log.Errorf("JoinChat Server stream has failed with error: %v", err.Error())
				stream.Context().Done()
				syncChannel <- 1
				return
			}
		}
	}
}()
```