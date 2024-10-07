package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 创建监听
	ip := "127.0.0.1"
	port := 8848
	adress := fmt.Sprintf("%s:%d", ip, port)

	// func Listen(network, adress string) (Listener, error)
	// net.Listen("tcp", ":8848") 简写，:前默认是本机：127.0.0.1
	listener, err := net.Listen("tcp", adress)

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// Accept() (conn, error)
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	fmt.Println("connect establish successfully")

	// 创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024)

	// Read(b []byte) (n int, err error)
	// cnt: 真正读取client发来的数据长度
	cnt, err := conn.Read(buf)

	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("Client =======> Server, 长度：", cnt, "，数据：", string(buf))

	// 服务器堆客户端进行响应
	//将数据转换成大写 "hello" ====> HELLO
	upperData := strings.ToUpper(string(buf))

	// Write(b []byte) (n int, err error)
	cnt, err = conn.Write([]byte(upperData))
	fmt.Println("Client <======== Server,长度：", cnt, ",数据：", upperData)

	// 关闭连接
	conn.Close()

}
