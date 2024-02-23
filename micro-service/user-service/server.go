package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	hellopb "github.com/YukiOnishi1129/auto-stock-trading-system/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func main() {
	// crate a listener on TCP port 3001
	port := 3001
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// create a new gRPC server
	s := grpc.NewServer()

	// register the greeting service with the gRPC server
	hellopb.RegisterGreetingServiceServer(s, NewMyServer())

	// register reflection service on gRPC server
	reflection.Register(s)

	// start the gRPC server port at 8080
	go func() {
		log.Printf("start gRPC server listening on port: %d", port)
		s.Serve(listener)
	}()

	// wait for Control C to exit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down gRPC server...")
	s.GracefulStop()
}

// Hello is a simple unary RPC
func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	// get the name from the request
	return &hellopb.HelloResponse{Message: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}

// HelloServerStream is a server streaming RPC
func (s *myServer) HelloServerStream(req *hellopb.HelloRequest, stream hellopb.GreetingService_HelloServerStreamServer) error {
	resCount := 5;
	for i := 0; i < resCount; i++ {
		// send a message to the client
		if err := stream.Send(&hellopb.HelloResponse{
			Message: fmt.Sprintf("[%d] Hello, %s!", i, req.GetName()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}
	// return nil to indicate that the stream has finished
	return nil
}

// HelloClientStream is a client streaming RPC
func (s *myServer) HelloClientStream(stream hellopb.GreetingService_HelloClientStreamServer) error {
	nameList := make([]string, 0)
	for {
		// get request from the client
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			// finish reading the client stream
			message := fmt.Sprintf("Hello, %v!", nameList)
			return stream.SendAndClose(&hellopb.HelloResponse{Message: message})
		}
		if err != nil {
			return err
		}
		nameList = append(nameList, req.GetName())
	}
}

// HelloBiStreams is a bidirectional streaming RPC
func (s *myServer) HelloBiStreams(stream hellopb.GreetingService_HelloBiStreamsServer) error {
	for {
		// recieve request from the client
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		// send a message to the client
		if err := stream.Send(&hellopb.HelloResponse{
			Message: fmt.Sprintf("Hello, %s!", req.GetName()),
		}); err != nil {
			return err
		}
	}
}