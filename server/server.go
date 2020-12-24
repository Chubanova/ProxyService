package main

import(
	"fmt"
	chat "github.com/Chubanova/ProxyService/proxygrpc"
	"flag"
	"net"
	"log"
	"google.golang.org/grpc"

)

var(
	port       = flag.Int("port", 10000, "The server port")
)

func main(){
	fmt.Print("hello world")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	
	grpcServer := grpc.NewServer(opts...)
	
	// pb.Reg(grpcServer, newServer())
	// grpcServer.Serve(lis)
}