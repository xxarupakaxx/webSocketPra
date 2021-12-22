package chat

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *MessageRequest) (*MessageResponse, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &MessageResponse{Body: "Hello From the Server!"}, nil
}
