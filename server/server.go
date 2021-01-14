package main

import (
	"io"
	"context"
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
	fmt.Print("Start server chat")
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


func (s *messangerServer) GetInfo(ctx context.Context, in *chat.GetInfoRequest) (*chat.GetInfoResponce, error) {
	return &chat.GetInfoResponce{}, nil
}
func (s *messangerServer) JoinChannel( in *chat.JoinChannelRequest, rsp chat.Messanger_JoinChannelServer) error {
	return nil

}
func (s *messangerServer) StartChannel(req chat.Messanger_StartChannelServer) error {
	return nil

}
func (s *messangerServer) JoinChat(stream chat.Messanger_JoinChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err := stream.Send(in); err != nil {
			return err
		}
	}
}
