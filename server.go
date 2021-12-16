package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/xxarupakaxx/webSocketPra/chat"
	"google.golang.org/grpc"
	"net"
)

func main() {
	fmt.Println("Go gRPC Beginners Tutorial!")
	lis,err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatalf("failed to listen :%v",err)
	}
	s:= chat.Server{}


	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer,&s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server L%s",err)
	}

}
