package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc"
	pb "grpc-gateway-demo/proto"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: defaultName})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
