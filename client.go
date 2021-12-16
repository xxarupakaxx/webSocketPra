package main

import (
	"context"
	"github.com/labstack/gommon/log"
	"github.com/xxarupakaxx/webSocketPra/helloworld"
	"google.golang.org/grpc"
	"os"
	"time"
)

func main() {

	conn,err := grpc.Dial(":50051",grpc.WithInsecure())
	if err != nil {
		log.Printf("err")
		return
	}
	defer conn.Close()

	c := helloworld.NewGreeterClient(conn)

	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()

	r,err := c.SayHello(ctx,&helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("errrrrrr")
		return
	}
	log.Printf("greeting :%s",r.Message)
	/*var conn *grpc.ClientConn
	conn,err := grpc.Dial(":9000",grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect :%s",err)
		return
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	res,err := c.SayHello(context.Background(),&chat.Message{Body: "Hello form client"})
	if err != nil {
		log.Fatalf("error when calling sayHello:%s",err)
		return
	}

	log.Printf("Response from server :%s",res.Body)*/

}