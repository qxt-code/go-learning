package main

import (
	"context"
	"fmt"
	"grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:8800", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Dial failed : ", err)
		return
	}
	defer conn.Close()

	c := pb.NewTeachClient(conn)

	t := pb.Teacher{
		Name: "Lily",
		Age:  18,
	}

	r, err := c.Teaching(context.Background(), &t)

	if err != nil {
		fmt.Println("Teaching failed : ", err)
		return

	}
	fmt.Println("Request : ", t.Age, " , ", t.Name)
	fmt.Println("Response : ", r.Age, " , ", r.Name)

}
