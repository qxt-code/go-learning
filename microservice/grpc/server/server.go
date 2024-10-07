package main

import (
	"context"
	"fmt"
	"grpc/pb"
	"net"

	"google.golang.org/grpc"
)

// 定义类对象
type Lily struct{}

// 绑定类方法
func (l *Lily) Teaching(ctx context.Context, in *pb.Teacher) (*pb.Teacher, error) {
	in.Name = "lalala" + in.Name
	return &pb.Teacher{Name: "Teacher " + in.Name, Age: in.Age}, nil
}

func main() {
	// 1. 初始grpc对象
	s := grpc.NewServer()

	// 2. 注册服务
	pb.RegisterTeachServer(s, &Lily{})

	// 3. 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	// 4. 启动服务
	if err := s.Serve(listener); err != nil {
		fmt.Println("failed to serve: ", err)
		return
	}
}
