# Proxy server for bidectional stream

D:\Soft\protoc\bin\protoc.exe -I .\proto\ .\proto\chat.proto --go_out=plugins=grpc:proxygrpc        
D:\Soft\protoc\bin\protoc.exe -I .\proto\ .\proto\proxy.proto --go_out=plugins=grpc:proxygrpc

# This project shows how to develop proxy service on golang for biderectioanl stream
## Этот проект реализован прокси серверна голанг для двунаправленного стрима

### 
Сложность написания прокси для двунаправленного стрима - необходимость запустить два потока которые будут жить пока его контакт не разорвет соединение и пока жив второй поток
А весь метод должен ждать окончания обоих потоков.

Я не нашла ни одного примера такой реализации, поэтому расписываю как это работает.

Объявляем переменую sync.WaitGroup. Описание пакета sync и пример использования вы смотрите тут https://golang.org/pkg/sync/ .
```
	var wg sync.WaitGroup
```

Функция WaitGroup  - дождаться окончания рутину. WaitGroup принимает - сколько рутин должно закончится.
```
	wg.Add(2)
```

Потом объявляем chan c int. chan - структура для получения и отправки сообщений и удостоверяемся что он закроется когда метод закончится
```
	syncChannel := make(chan int)
	defer close(syncChannel)
```


```
func (s *messangerProxyServer) JoinChat(stream chat.MessangerProxy_JoinChatServer) error {
	// Init logger
	log := loggerFormatter()
	log.Info("Received JoinChat request")

	// WaitGroup for goroutines sync
	var wg sync.WaitGroup
	wg.Add(2)

	// Channel for streams closing countdown
	syncChannel := make(chan int)
	defer close(syncChannel)


	// Connect to server
	chatConn, err := grpc.Dial(buildServiceName(), grpc.WithInsecure())
	if err != nil {
		return err
	}
	// Close connection when work finished
	defer chatConn.Close()

	// Init client
	res := chat.NewMessangerClient(chatConn)

	// Subscribe
	cl, err := res.JoinChat(stream.Context())
	if err != nil {
		return err
	}

	// Init Client stream
	go func() {
		defer wg.Done()

		for {
			select {
			case <-syncChannel:
				log.Infof("Server stream was closed. Closing Client stream...")
				return
			default:
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
			}

		}
	}()

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

	// Wait for all streams to close
	wg.Wait()
	return err
}
```