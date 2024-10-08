package main

import (
	"consul/pb"
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. 初始化consul 配置
	consulConfig := api.DefaultConfig()

	// 2. 创建 consul对象 ——(可以重新指定consul属性：IP/PORT，也可以使用默认)
	consulClient, err := api.NewClient(consulConfig)
	if err != nil {
		fmt.Println("NewCLient err: ", err)
		return
	}

	// 3. 服务发现， 从consul上获取健康的服务
	services, _, err := consulClient.Health().Service("grpc and consul", "grpc-hello", true, nil)
	if err != nil {
		fmt.Println("Service err: ", err)
		return
	}

	addr := services[0].Service.Address + ":" + strconv.Itoa(services[0].Service.Port)

	// grpc远程调用
	// grpcConn, _ := grpc.NewClient("127.0.0.1:8800", grpc.WithTransportCredentials(insecure.NewCredentials()))
	grpcConn, _ := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	grpcClient := pb.NewHelloClient(grpcConn)

	resp, err := grpcClient.SayHello(context.Background(), &pb.Person{Name: "lily", Age: 18})
	if err != nil {
		fmt.Println("Call err: ", err)
		return
	}
	fmt.Println("Response: ", resp)
}
