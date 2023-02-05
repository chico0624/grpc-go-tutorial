package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"mygrpc/cmd/server/hello"
	hellopb "mygrpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	// gRPCサーバにGreetingServiceを登録する
	hellopb.RegisterGreetingServiceServer(s, hello.NewMyServer())

	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port :%v", port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stoping gRPC server...")
	s.GracefulStop()
}
