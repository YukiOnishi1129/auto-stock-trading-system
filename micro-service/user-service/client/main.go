package main

import (
	"bufio"
	"context"
	"fmt"
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
		fmt.Println("2: exit")
		fmt.Print("please enter >")

		scanner.Scan()
		in := scanner.Text()

		switch in {
			case "1":
				Hello()

			case "2":
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