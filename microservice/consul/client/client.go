package main

import (
	"consul/pb"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	grpcConn, _ := grpc.NewClient("127.0.0.1:8800", grpc.WithTransportCredentials(insecure.NewCredentials()))

	grpcClient := pb.NewHelloClient(grpcConn)

	resp, err := grpcClient.SayHello(context.Background(), &pb.Person{Name: "lily", Age: 18})
	if err != nil {
		fmt.Println("Call err: ", err)
		return
	}
	fmt.Println("Response: ", resp)
}
