package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"grpcdemo/pb"
)

/**
  Author: jisd
  Created: 2021-01-13 11:37:26
*/
var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8100, "grpc port")
}

type GreetServer struct {
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &GreetServer{})
	grpcServer.Serve(lis)
}

func (server *GreetServer) Greet(ctx context.Context, in *pb.Request) (*pb.Response, error) {

	time.Sleep(time.Second*3)

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("context.DeadlineExceeded")
		return nil, status.Errorf(codes.Canceled, "GreetServer.Greet canceled")
	}

	fmt.Println("GreetServer Greet...", in.Name)
	return &pb.Response{Greet: "lisi"}, nil
}
