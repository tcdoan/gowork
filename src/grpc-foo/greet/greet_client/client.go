package main

import (
	"context"
	"fmt"
	"grpc-foo/greet/greetpb"
	"io"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting a client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)

	// doUnary(c)
	//doServerStreaming(c)
	doClientStreaming(c)
}

func doServerStreaming(client greetpb.GreetServiceClient) {
	fmt.Println("Starting a Server streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Thanh",
			LastName:  "Doan",
		},
	}

	stream, _ := client.GreetManyTimes(context.Background(), req)

	for resp, err := stream.Recv(); err != io.EOF; resp, err = stream.Recv() {
		if err != nil {
			log.Fatalf("Fatal error while receving from streaming server %v", err)
		}

		log.Printf("Response from RPC: %v \n", resp.Result)
	}
}

func doClientStreaming(client greetpb.GreetServiceClient) {
	fmt.Println("Starting client streaming RPC...")

	stream, _ := client.LongGreet(context.Background())

	for reqID := 1; reqID < 10; reqID++ {
		req := &greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Quang " + strconv.Itoa(reqID),
			},
		}
		err := stream.Send(req)
		time.Sleep(1 * time.Second)
		if err != nil {
			log.Fatalf("Error sending client streaming request %v \n", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response for the streaming request %v \n", err)
	}

	log.Println(resp.Result)
}

func doUnary(client greetpb.GreetServiceClient) {
	fmt.Println("Starting a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Thanh",
			LastName:  "Doan",
		},
	}
	resp, _ := client.Greet(context.Background(), req)
	log.Printf("Response from RPC: %v \n", resp)
}
