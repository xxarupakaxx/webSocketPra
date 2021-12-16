package chat

import (
	context "context"
	"github.com/labstack/gommon/log"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client:%w", in.Body)
	return &Message{Body: "Hello From the server!"}, nil
}
