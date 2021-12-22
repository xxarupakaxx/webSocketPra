package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/xxarupakaxx/webSocketPra/chat"
	"google.golang.org/grpc"
	"net/http"
)

func run() error {
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	endpoint := fmt.Sprintf("localhost:9000")

	err := chat.RegisterChatServiceHandlerFromEndpoint(ctx,mux,endpoint,opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":5000", mux)
}


func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}