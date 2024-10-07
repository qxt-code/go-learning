package main

import (
	"fmt"
	"net"
)

func main() {

}

func handleConnect(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read failed")
			return
		}
	}

}
