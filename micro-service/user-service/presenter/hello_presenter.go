package presenter

import (
	"context"
	"errors"
	"fmt"
	hellopb "github.com/YukiOnishi1129/auto-stock-trading-system/user-service/grpc"
	"io"
	"time"
)

type HelloPresenter struct {
	hellopb.UnimplementedGreetingServiceServer
}

func NewHelloPresenter() *HelloPresenter {
	return &HelloPresenter{}
}

// Hello is a simple unary RPC
func (s *HelloPresenter) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	// get the name from the request
	return &hellopb.HelloResponse{Message: fmt.Sprintf("Hello, %s!", req.GetName())}, nil
}

// HelloServerStream is a server streaming RPC
func (s *HelloPresenter) HelloServerStream(req *hellopb.HelloRequest, stream hellopb.GreetingService_HelloServerStreamServer) error {
	resCount := 5
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
func (s *HelloPresenter) HelloClientStream(stream hellopb.GreetingService_HelloClientStreamServer) error {
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
func (s *HelloPresenter) HelloBiStreams(stream hellopb.GreetingService_HelloBiStreamsServer) error {
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
