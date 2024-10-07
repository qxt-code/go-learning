package main

import (
	"fmt"
	"net/rpc/jsonrpc"
	inter "rpc/interface"
)

func main01() {
	// 1.用rpc连接服务器 Dial
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("Dial err: ", err)
		return
	}
	defer conn.Close()

	// 2.调用远程函数
	var reply string
	conn.Call("hello.HelloWorld", "I am Libai", &reply)
	if err != nil {
		fmt.Println("Call err: ", err)
		return
	}
	fmt.Println("Reply: ", reply)
}

// 结合03
func main() {
	c := inter.InitClient("127.0.0.1:8800")
	var resp string
	if err := c.HelloWorld("Dufu", &resp); err != nil {
		fmt.Println("HlloWorld err:", err)
		return
	}
	fmt.Println("resp: ", resp)

}
