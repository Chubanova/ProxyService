package main

import (
	"feed/feedgrpc/feed"
	"flag"
	"fmt"
	"log"
	"net"

	chat "github.com/Chubanova/ProxyService/proxygrpc"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type messangerServer struct {
	chat.UnimplementedMessangerServer
}

func newServer() *messangerServer {
	s := &messangerServer{}
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
	chat.RegisterMessangerServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

func (s *messangerServer) GetInfo(req *feed.SubscriptionRequest, rsp FrontFeedService_SubscribeFeedServer) error {
}
func (s *messangerServer) JoinChannel(req *feed.SubscriptionRequest, rsp FrontFeedService_SubscribeFeedServer) error {
}
func (s *messangerServer) StartChannel(req *feed.SubscriptionRequest, rsp FrontFeedService_SubscribeFeedServer) error {
}
func (s *messangerServer) JoinChat(req *feed.SubscriptionRequest, rsp FrontFeedService_SubscribeFeedServer) error {
}
