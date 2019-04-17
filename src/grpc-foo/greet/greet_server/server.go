package main

import (
	"context"
	"fmt"
	"grpc-foo/greet/greetpb"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v \n", req)
	result := "Hello " + req.GetGreeting().GetFirstName()
	resp := &greetpb.GreetResponse{Result: result}
	return resp, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("Greet function was invoked with %v \n", req)

	for i := 0; i < 10; i++ {
		greeting := "Hello " + req.GetGreeting().GetFirstName() + " " + strconv.Itoa(i)
		resp := &greetpb.GreetManyTimesResponse{
			Result: greeting,
		}
		stream.Send(resp)
		time.Sleep(1 * time.Second)
	}

	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			resp := &greetpb.LongGreetResponse{Result: result}
			return stream.SendAndClose(resp)
		}
		result += "Hello " + req.GetGreeting().GetFirstName() + "! "
	}
}

func main() {
	const PORT = 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	fmt.Printf("Listening on port %d...\n", PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
