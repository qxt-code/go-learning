package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// http包
	client := http.Client{}

	// func (c *client) Get(url string) (resp *Response, err error)
	resp, err := client.Get("https://www.aliyun.com")
	if err != nil {
		fmt.Println("client.Get:", err)
		return
	}

	ct := resp.Header.Get("Contest-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")

	fmt.Println("content-Type:", ct)
	fmt.Println("date:", date)
	fmt.Println("server:", server)

	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status

	fmt.Println("url:", url)
	fmt.Println("code:", code)
	fmt.Println("status:", status)

	body := resp.Body
	fmt.Println("body:", body)
	readBodyStr, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("body str:", string(readBodyStr))

}
