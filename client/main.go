package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"grpcdemo/pb"
)

/**
  Author: jisd
  Created: 2021-01-13 11:37:13
*/

var (
	serverAddr string
)

func init() {
	flag.StringVar(&serverAddr, "server", "127.0.0.1:8100", "grpc server")
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(2 * time.Second)))
	defer cancel()

	resp, err := client.Greet(ctx, &pb.Request{Name: "zhangsan"})
	if err != nil {
		fmt.Println("Greet err", err)
	} else {
		fmt.Println("Greet resp...", resp)
	}
}