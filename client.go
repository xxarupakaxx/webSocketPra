package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/xxarupakaxx/webSocketPra/client"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn,err := grpc.Dial(":9000",grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect :%s",err)
		return
	}
	defer conn.Close()

	c := client.NewChatServiceClient(conn)

	res,err := c.SayHello(context.Background(),&client.Message{Body: "Hello form client"})
	if err != nil {
		log.Fatalf("error when calling sayHello:%s",err)
		return
	}

	log.Printf("Response from server :%s",res.Body)

}