package main

import (
	"os"
	"github.com/sirupsen/logrus"
	"sync"
	"io"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	easy "github.com/t-tomalak/logrus-easy-formatter"

	chat "github.com/Chubanova/ProxyService/proxygrpc"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10001, "The server port")
)

type messangerProxyServer struct {
	chat.UnimplementedMessangerProxyServer
}



func newServer() *messangerProxyServer {
	s := &messangerProxyServer{}
	return s
}


func main() {
	fmt.Print("Start proxy service")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	chat.RegisterMessangerProxyServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func loggerFormatter() *logrus.Logger {
	log := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05.12",
			LogFormat:       "%time% [%lvl%]:- %msg%\n",
		},
	}
	return log
}


func buildServiceName() string {
	host := "localhost"
	port := "10000"

	return fmt.Sprintf("%s:%s", host, port)
}

func (s *messangerProxyServer) GetInfo(ctx context.Context, in *chat.GetInfoRequest) (*chat.GetInfoResponce, error) {
	return &chat.GetInfoResponce{}, nil
}
func (s *messangerProxyServer) JoinChannel( in *chat.JoinChannelRequest, rsp chat.MessangerProxy_JoinChannelServer) error {
	return nil

}
func (s *messangerProxyServer) StartChannel(req chat.MessangerProxy_StartChannelServer) error {
	return nil

}
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

