package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	pb "grpc-gateway-demo/proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var grpcport = flag.Int("grpc", 50051, "the port grpc serve on")
var httpport = flag.Int("http", 8080, "the port http serve on")

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + in.Name }, nil
}

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterGreeterServer(s, &server{})
	// Serve gRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *grpcport))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func () {
		log.Println("Serving gRPC on 0.0.0.0" + fmt.Sprintf(":%d", *grpcport))
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	go func() {
		log.Println("Serving http on 0.0.0.0" + fmt.Sprintf(":%d", *httpport))
		conn, err := grpc.NewClient(fmt.Sprintf("localhost:%v", *grpcport), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer conn.Close()
		mux := runtime.NewServeMux()
		err = pb.RegisterGreeterHandler(context.Background(), mux, conn)
		if err != nil {
			log.Fatalln("Failed to register gateway:", err)
		}
		if err := http.ListenAndServe(fmt.Sprintf(":%v", *httpport), mux); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

 	<-sigChan
	 fmt.Println("Exiting program...")
}
