package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"mygrpc/cmd/server/hello"
	"mygrpc/cmd/server/interceptor"
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

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.MyUnaryServerInterceptor1,
			interceptor.MyUnaryServerInterceptor2,
		),
		grpc.ChainStreamInterceptor(
			interceptor.MyStreamServerInterceptor1,
			interceptor.MyStreamServerInterceptor2,
		),
	)

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
