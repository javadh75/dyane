package controller

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/javadh75/dyane/protoc"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:5000"
	defaultName = "world"
)

func Run() {
	fmt.Println("Running Controller")
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
