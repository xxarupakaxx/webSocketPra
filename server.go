package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/xxarupakaxx/webSocketPra/helloworld"
	chat2 "github.com/xxarupakaxx/webSocketPra/server/chat"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":50051"
)

type server struct{
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	panic("implement me")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	chats := chat2.Server{}
	hellos := server{}


	grpcServer := grpc.NewServer()

	chat2.RegisterChatServiceServer(grpcServer,&chats)
	helloworld.RegisterGreeterServer(grpcServer,&hellos)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server L%s",err)
	}

}
