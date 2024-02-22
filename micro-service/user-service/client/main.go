package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	hellopb "github.com/YukiOnishi1129/auto-stock-trading-system/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


var (
	scanner *bufio.Scanner
	client hellopb.GreetingServiceClient
)
func main() {
	fmt.Println("start, gRPC Client!")

	// create a scanner to read from the terminal
	scanner  = bufio.NewScanner(os.Stdin)

	// create a connection to the gRPC server
	// "localhost:8080" is the address of the gRPC server
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // don't use SSL/TLS connection
		grpc.WithBlock(), // block until the connection is established
	)
	if err != nil {
		fmt.Printf("Error connecting: %v", err)
		return
	}
	defer conn.Close()

	// create a gRPC client
	client = hellopb.NewGreetingServiceClient(conn)

	for {
		fmt.Println("1: send Request")
		fmt.Println("2: HelloServerStream")
		fmt.Println("3: HelloClientStream")
		fmt.Println("4: HelloBiStreams")
		fmt.Println("5: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
			case "1":
				Hello()

			case "2":
				HelloServerStream()

			case "3":
				HelloClientStream()

			case "4":
				HelloBiStream()

			case "5":
				fmt.Println("bye.")
				goto M
		}
	}
	M:
}

func Hello() {
	fmt.Println("Please enter your name:")
	scanner.Scan()
	name := scanner.Text()

	// create a request message
	req := &hellopb.HelloRequest{Name: name}

	// send the request to the gRPC server
	// Hello is the method name of the gRPC server
	res, err := client.Hello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	} else {
		// print the response message
		fmt.Println(res.GetMessage())
	}
}

func HelloServerStream() {
	fmt.Println("Please enter your name:")
	scanner.Scan()
	name := scanner.Text()

	// create a request message
	req := &hellopb.HelloRequest{Name: name}

	// send the request to the gRPC server
	// receive stream from the gRPC server
	stream, err := client.HelloServerStream(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// receive the response message from the gRPC server
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("all the responses have already received")
			break
		}

		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(res)
	}
}

func HelloClientStream() {
	// create a client stream to the gRPC server
	stream, err := client.HelloClientStream(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	sendCount := 3
	fmt.Printf("Please enter %d names.\n", sendCount)

	for i := 0; i < sendCount; i++ {

		scanner.Scan()
		name := scanner.Text()

		// send the request message to the gRPC server
		if err :=stream.Send(&hellopb.HelloRequest{Name: name}); err != nil {
			fmt.Println(err)
			return
		}
	}

	// close the client stream
	res, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetMessage())
	}
}

func HelloBiStream() {
	stream, err := client.HelloBiStreams(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	sendNum := 5
	fmt.Printf("Please enter %d names.\n", sendNum)

	var sendEnd, recvEnd bool

	sendCount := 0

	for !(sendEnd && recvEnd) {

		// send the request message to the gRPC server
		if !sendEnd {
			scanner.Scan()
			name := scanner.Text()

			if err := stream.Send(&hellopb.HelloRequest{Name: name}); err != nil {
				fmt.Println(err)
				return
			}
			
			if sendCount == sendNum {
				sendEnd = true
				if err := stream.CloseSend(); err != nil {
					fmt.Println(err)
				}
			}
		}

		// recieve the response message from the gRPC server
		if !recvEnd {
			if res, err := stream.Recv(); err != nil {
				if errors.Is(err, io.EOF) {
					recvEnd = true
				} 
				recvEnd = true
			} else {
				fmt.Println(res)
			}
		}
	}
}