package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 注册路由 router
	// xxx/user ===> func1
	// xxx/name ===> func2
	// xxx/id ===> func3

	// func是回调函数，用于路由的响应，这个回调函数的原型是固定的
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// request: 客户端发来的数据
		fmt.Println("用户请求详情：")
		fmt.Println("request:", r)

		// writer: 将数据返回给客户端
		_, _ = io.WriteString(w, "这是/user请求返回的数据")

	})
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		// request: 客户端发来的数据
		fmt.Println("用户请求详情：")
		fmt.Println("request:", r)

		// writer: 将数据返回给客户端
		_, _ = io.WriteString(w, "这是/name请求返回的数据")

	})
	http.HandleFunc("/iid", func(w http.ResponseWriter, r *http.Request) {
		// request: 客户端发来的数据
		fmt.Println("用户请求详情：")
		fmt.Println("request:", r)

		// writer: 将数据返回给客户端
		_, _ = io.WriteString(w, "这是/id请求返回的数据")

	})

	// func ListenAndServe(addr string, handler Handler) error
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("http start failed, err:", err)
		return
	}

}
