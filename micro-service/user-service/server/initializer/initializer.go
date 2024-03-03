package initializer

import (
	"fmt"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/database"
	pb "github.com/YukiOnishi1129/auto-stock-trading-system/user-service/grpc"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/infrastructure/mysql/repository"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/presenter"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func Init() {
	//	database
	db, dbErr := database.Init()
	if dbErr != nil {
		fmt.Printf("Error connecting to DB\n")
		return
	}

	//	repository
	ur := repository.NewUserRepository(db)

	//	usecase
	uu := usecase.NewUserUsecase(db, ur)

	//	presenter
	up := presenter.NewUserPresenter(db, uu)

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
	pb.RegisterGreetingServiceServer(s, presenter.NewHelloPresenter())
	pb.RegisterUserServiceServer(s, up)

	// register reflection service on gRPC server
	reflection.Register(s)

	// start the gRPC server port at 8080
	go func() {
		log.Printf("start gRPC server listening on port: %d", port)
		err := s.Serve(listener)
		if err != nil {
			return
		}
	}()

	// wait for Control C to exit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down gRPC server...")
	s.GracefulStop()
}
