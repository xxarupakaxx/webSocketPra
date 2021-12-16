package main

import (
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis,err := net.Listen("tcp",":9000")
	if err != nil {
		log.Fatalf("failed to listen :%v",err)
	}

	grpcServer := grpc.NewServer()

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server L%s",err)
	}

}
