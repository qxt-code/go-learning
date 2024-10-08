package main

import (
	"consul/pb"
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/api"
	"google.golang.org/grpc"
)

type Child struct{}

func (c *Child) SayHello(ctx context.Context, p *pb.Person) (*pb.Person, error) {
	p.Name = "hello " + p.Name
	return p, nil
}

func main() {
	// grpc注册到consul

	// 1. 初始化consul配置
	consulConfig := api.DefaultConfig()

	// 2. 创建consul 对象
	consuClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("Client failed: ", err)
	}

	// 3. 告诉consul即将注册的服务信息
	registerService := api.AgentServiceRegistration(
		ID: "server-1",
		Tags: []string{"grpc-hello", "consul-hello"},
		Name: "grpc and consul",
		Address: "127.0.0.1",
		Port: "8800",
		Check: &api.AgentServiceCheck{
			CheckID: "consul hello check",
			TCP: "127.0.0.1:8800",
			Timeout: "1s",
			Interval: "5s",
		},
	)

	// 4. 注册grpc服务到consul上
	consulConfig.Agent().ServiceRegister(&registerService)


	// grpc服务
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServer(grpcServer, new(Child))

	lintener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Listen err: ", err)
		return
	}

	grpcServer.Serve(lintener)
}
