package main

import (
	"context"
	"fmt"
	"grpc-foo/calculator/calcpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting a calculator client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer cc.Close()
	c := calcpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(client calcpb.CalculatorServiceClient) {
	fmt.Println("Starting a Unary RPC...")
	req := &calcpb.SumRequest{
		FirstNumber:  22,
		SecondNumber: 33,
	}
	resp, _ := client.Sum(context.Background(), req)
	log.Printf("Response from RPC: %v \n", resp)
}
