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
	port = flag.Int("port", 10000, "The server port")
)

type messangerProxyServer struct {
	chat.UnimplementedMessangerProxyServer
}



// type routeGuideServer struct {
// 	pb.UnimplementedRouteGuideServer
// 	savedFeatures []*pb.Feature // read-only after initialized

// 	mu         sync.Mutex // protects routeNotes
// 	routeNotes map[string][]*pb.RouteNote
// }


func newServer() *messangerProxyServer {
	s := &messangerProxyServer{}
	return s
}


func main() {
	fmt.Print("hello world")
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


	// Connect to resource
	//resourceConn, err := grpc.Dial(buildServiceName("RESOURCE"), grpc.WithTransportCredentials(creds))
	resourceConn, err := grpc.Dial(buildServiceName("RESOURCE"), grpc.WithInsecure())
	if err != nil {
		return err
	}
	// Close connection when work finished
	defer resourceConn.Close()

	// Init client
	res := resource.NewResourceServiceClient(resourceConn)

	// Subscribe
	cl, err := res.SubscribeOnChangeById(stream.Context())
	if err != nil {
		return err
	}

	// Init feed stream
	go func() {
		defer wg.Done()

		for {
			select {
			case <-syncChannel:
				log.Infof("Resource stream was closed. Closing feed stream...")
				return
			default:
				msg, err := stream.Recv()
				if err != nil {
					if err != io.EOF {
						log.Errorf("SubscribeOnChangeById feed stream has failed with error: %v", err.Error())
					}
					syncChannel <- 1
					return
				}
				log.Debugf("Received message from feed stream %v", msg)

				err = cl.Send(msg)
				if err != nil {
					log.Errorf("SubscribeOnChangeById feed stream has failed with error: %v", err.Error())
					syncChannel <- 1
					return
				}
			}

		}
	}()

	// Init resource stream
	go func() {
		defer wg.Done()

		for {
			select {
			case <-syncChannel:
				log.Infof("Feed stream was closed. Closing resource stream...")
				return
			default:
				msg, err := cl.Recv()
				if err != nil {
					log.Errorf("SubscribeOnChangeById resource stream has failed with error: %v", err.Error())
					stream.Context().Done()
					syncChannel <- 1
					return
				}
				log.Debugf("Received message from resource stream %v", msg)

				err = stream.Send(msg)
				if err != nil {
					log.Errorf("SubscribeOnChangeById resource stream has failed with error: %v", err.Error())
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

