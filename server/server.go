package main

import (
	context "context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"math/rand"
	"net"
	"sync"
)

func getToken() string {
	tkn := make([]byte, 0)
	rand.Read(tkn)
	return fmt.Sprintf("%x", tkn)
}

type server struct {
	clients   sync.Map
	streamChs sync.Map
	exitCh    chan struct{}
	wg        *sync.WaitGroup
	UnimplementedChatServiceServer
}

func (s *server) run(ctx context.Context) error {
	log.Println("server run")
	defer func() {
		close(s.exitCh)
		s.wg.Wait()
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	grpcServer := grpc.NewServer()
	defer func() {
		grpcServer.Stop()
		log.Println("shuttle down")
	}()

	RegisterChatServiceServer(grpcServer, s)
	reflection.Register(grpcServer)

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			return
		}
	}()

	<-ctx.Done()

	return nil
}
func (s *server) Login(ctx context.Context, user *User) (*User, error) {
	log.Println("try to logged in.")

	clientExists := false
	s.clients.Range(func(_, client interface{}) bool {
		if value, ok := client.(string); ok && value == user.GetName() {
			clientExists = true
			return false
		}
		return true
	})
	if clientExists {
		return &User{}, fmt.Errorf("\"%s\" is already in use.", user.GetName())
	}

	user.Token = getToken()
	s.clients.Store(user.GetToken(), user.GetName())

	log.Printf("%s logged in . \n", user.GetName())
	return user, nil
}

func (s *server) Logout(ctx context.Context, user *User) (*emptypb.Empty, error) {
	log.Println("Try to logged out.")

	s.clients.Delete(user.GetToken())
	s.deleteStreamCh(user.GetToken())
	log.Printf("%s logged out.\n", user.GetName())
	return &empty.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, message *Message) (*Message, error) {
	log.Println("try to broadcast message .")
	s.broadcast(message)
	log.Printf("sent %s,\n", message.GetContent())
	return message, nil
}

func (s *server) GetMessage(user *User, messageServer ChatService_GetMessageServer) error {
	s.wg.Add(1)
	defer s.wg.Done()
	streamCh := s.createStreamCh(user.GetToken())
	defer s.deleteStreamCh(user.GetToken())

	for true {
		select {
		case msg, ok := <-streamCh:
			if !ok {
				return nil
			}
			if err := messageServer.Send(msg); err != nil {
				log.Println("sending error.")
				return err
			}
		case <-s.exitCh:
			log.Printf("%s exit.\n", user.GetName())
			return nil
		}
	}
	return nil
}

func (s *server) deleteStreamCh(token string) {
	log.Println("Try to delete stream.")
	if msg, ok := s.streamChs.Load(token); ok {
		if value, ok := msg.(chan *Message); ok {
			close(value)
		}
		s.streamChs.Delete(token)
	}
}

func (s *server) createStreamCh(token string) chan *Message {
	log.Println("Try to create stream.")
	ch := make(chan *Message, 1)
	s.streamChs.Store(token, ch)
	return ch
}

func (s *server) broadcast(msg *Message) {
	wg := new(sync.WaitGroup)
	s.streamChs.Range(func(_, ch interface{}) bool {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if value, ok := ch.(chan *Message); ok {
				value <- msg
			}
		}()

		return true
	})
	wg.Wait()
}
