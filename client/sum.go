package main

import (
	"context"
	"log"

	pb "github.com/puffyguy/grpcUnary/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.Request{
		A: 21,
		B: 4,
	})
	if err != nil {
		log.Fatalf("Could not Sum %v\n", err)
	}

	log.Printf("Sum is %v\n", res.Result)
}
