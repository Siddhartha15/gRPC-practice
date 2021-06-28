package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Siddhartha15/gRpc-practice/greet/greetpb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
	// defaultName = "world"
)

func main() {
	// fmt.Println("hel")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	// fmt.Println("created clien : ", c)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "john",
			LastName:  "wick",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalln("error while requesting ", err)
	}

	fmt.Println("success response: ", res)

}
