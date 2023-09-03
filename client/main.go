package main

import (
	"log"

	pb "github.com/puffyguy/grpcUnary/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5091"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect at %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	doSum(c)
}

// func doSum(c pb.CalculatorServiceClient) {
// 	res, err := c.Sum(context.Background(), &pb.Request{
// 		A: 25,
// 		B: 25,
// 	})
// 	if err != nil {
// 		log.Fatalf("Could not sum %v\n", err)
// 	}

// 	log.Printf("Sum is %v\n", res.Result)
// }
