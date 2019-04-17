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
	// doServerStreaming(c)
	// doClientStreaming(c)
	doBiDirectionalStreaming(c)
}

func doBiDirectionalStreaming(client greetpb.GreetServiceClient) {
	fmt.Println("Starting a bi-directional streaming RPC...")

	requests := []greetpb.GreetEveryoneRequest{
		greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "John",
			},
		},
		greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Max",
			},
		},
		greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "David",
			},
		},
		greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "James",
			},
		},
		greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Anwell",
			},
		},
	}

	// We create a stream by invoking client.GreetEveryone(...)
	stream, _ := client.GreetEveryone(context.Background())

	waitChanel := make(chan struct{})

	// We send a bunch of messages to the client (go routine)
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending request %v \n", req)
			err := stream.Send(&req)
			time.Sleep(1 * time.Second)
			if err != nil {
				fmt.Printf("Error while sending %v \n", err)
				continue
			}
		}
		stream.CloseSend()
	}()

	// We receive a bunch of messages from the client (go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Error while sending %v \n", err)
				break
			}
			log.Printf("Received: %s\n", res.GetResult())
		}
		close(waitChanel)
	}()

	// Block until everything is done
	<-waitChanel
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
