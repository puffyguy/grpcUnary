package main

import (
	"context"
	"log"

	pb "github.com/puffyguy/grpcUnary/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.Response{
		Result: in.A + in.B,
	}, nil
}
