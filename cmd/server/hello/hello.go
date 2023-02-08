package hello

import (
	"context"
	"errors"
	"fmt"
	"io"
	hellopb "mygrpc/pkg/grpc"
	"time"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

// Unary RPC
func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

// Server Stream RPC
func (s *myServer) HelloServerStream(req *hellopb.HelloRequest, stream hellopb.GreetingService_HelloServerStreamServer) error {
	resCount := 5
	for i := 0; i < resCount; i++ {
		if err := stream.Send(&hellopb.HelloResponse{
			Message: fmt.Sprintf("[%d] Hello %s!", i, req.GetName()),
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
	}

	return nil
}

// Client Stream RPC
func (s *myServer) HelloClientStream(stream hellopb.GreetingService_HelloClientStreamServer) error {
	nameList := make([]string, 0)

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			massage := fmt.Sprintf("Hello %v", nameList)

			return stream.SendAndClose(
				&hellopb.HelloResponse{
					Message: massage,
				},
			)
		}
		if err != nil {
			return err
		}
		nameList = append(nameList, req.GetName())
	}
}

// 双方向ストリーミング
func (s *myServer) HelloBiStreams(stream hellopb.GreetingService_HelloBiStreamsServer) error {
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return err
		}

		message := fmt.Sprintf("hello %v", req.GetName())

		if err := stream.Send(&hellopb.HelloResponse{
			Message: message,
		}); err != nil {
			return err
		}
	}
}
