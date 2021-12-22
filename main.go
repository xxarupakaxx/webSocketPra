package main

import (
	"fmt"
	"github.com/xxarupakaxx/webSocketPra/chat"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main()  {
	fmt.Println("gRPCサーバー")

	lis,err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatalf("errro:%v",err)
		return
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}


