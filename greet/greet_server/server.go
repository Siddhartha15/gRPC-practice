package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Siddhartha15/gRpc-practice/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Println("Greet server invoked")

	first_name := req.GetGreeting().GetFirstName()
	last_name := req.GetGreeting().GetLastName()

	result := "hello " + first_name + " " + last_name
	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("hello")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalln(" Failed Serving ", err)
	}
}
