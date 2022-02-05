package helloworld

import (
	"context"
	"log"
)

type Server struct {
	UnimplementedGreeterServer
	ServerName string
}

func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName() + ", from " + s.ServerName}, nil
}
