package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type User struct {
	name string
	id   string
	msg  chan string
}

// 创建一个全局的map结构， 用于保存所有的用户
var allUsers = make(map[string]User)

// 定义一个全局的message通道
var message = make(chan string, 10)

func main() {
	// 创建服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	// 全局唯一的go程，负责监听message通道
	go broadcast()

	fmt.Println("server launch successfully....")

	for { // 监听
		fmt.Println("======>主go程监听中")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}

		// 建立连接
		fmt.Println("建立连接成功！")

		// 启动处理业务的go程
		go hander(conn)
	}
}

// 处理具体业务
func hander(conn net.Conn) {
	fmt.Println("启动业务...")
	// 客户端与服务器建立连接的时候，会有ip和port ===> 当成user id
	clientAddr := conn.RemoteAddr().String()

	// 创建 user
	newUser := User{
		id:   clientAddr,
		name: clientAddr,
		msg:  make(chan string, 10), // 注意分配空间
	}

	// 启动go程，负责将msg信息返回给客户端
	go writeBackToClient(&newUser, conn)

	// 添加到map
	allUsers[newUser.id] = newUser

	// 定义退出信号
	isQuit := make(chan bool)

	// 重置计时器信号
	restTimer := make(chan bool)

	// 监听退出信号
	go watch(&newUser, conn, isQuit, restTimer)

	// 向message写入数据，通知所有人当前用户上线
	loginInfo := fmt.Sprintf("%s:%s ===> 上线了！！！\n", newUser.id, newUser.name)
	message <- loginInfo

	for {
		buf := make([]byte, 1024)

		// 读取
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("receive: ", string(buf[:cnt-2]), ", cnt:", cnt)
		// ---------业务逻辑处理---------
		// 1. 查询当前所有的用户 who
		// 	 a. 先判断接收的数据是不是who ===> 长度&& 字符串
		userInput := string(buf[:cnt-2])
		if len(userInput) == 4 && userInput == "\\who" {
			//   b. 遍历allUsers， 将id 和name拼接程字符串，返回给客户端
			fmt.Println("用户即将查询所有用户信息")

			// 包含所有用户信息的切片
			var userInfos []string
			for _, user := range allUsers {
				userInfo := fmt.Sprintf("userid:%s, username:%s", user.id, user.name)
				userInfos = append(userInfos, userInfo)
			}

			// 最终写到管道中，一定是一个字符串
			r := strings.Join(userInfos, "\n")

			// 将查询结果返回给查询用户
			newUser.msg <- r

		} else if len(userInput) > 9 && userInput[:7] == "\\rename" {
			// 2. 使用|进行分割
			newUser.name = strings.Split(userInput, "|")[1]
			allUsers[newUser.id] = newUser
			// 4. 通知客户端，更新成功
			newUser.msg <- "rename successfully"

		} else if userInput == "\\quit" {
			isQuit <- true
			return
		} else {
			// 普通聊天信息直接广播
			message <- userInput
		}
		restTimer <- true

	}
}

// 向所有用户广播消息，启动一个全局唯一的go程
func broadcast() {
	fmt.Println("广播go程启动成功")
	defer fmt.Println("broadcast 程序退出!")

	for {
		// 1. 从message中读取数据
		fmt.Println("broadcast 监听中......")
		info := <-message
		fmt.Println("message receive : ", info)
		// 2. 将数据写入到每一个用户的msg管道中
		for _, user := range allUsers {
			// 如果msg是非缓冲的，那么会在这里阻塞
			user.msg <- info
		}
	}
}

// 每个用户有一个监听发送给自己msg管道的go程
func writeBackToClient(user *User, conn net.Conn) {
	fmt.Printf("user: %s 监听msg管道:\n", user.name)
	for data := range user.msg {
		fmt.Printf("user : %s 写回给客户端：%s\n", user.name, data)
		_, _ = conn.Write([]byte(data))
	}
}

// 监听退出信号，并进行清理
func watch(user *User, conn net.Conn, isQuit, restTimer <-chan bool) {
	fmt.Println("监听退出信号：")
	defer fmt.Println("watch go 退出")
	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("%s 退出", user.name)
			fmt.Println("删除当前用户：", user.name)
			delete(allUsers, user.id)
			message <- logoutInfo

			conn.Close()
			return
		case <-time.After(5 * time.Second):
			logoutInfo := fmt.Sprintf("%s timeout exit", user.name)
			fmt.Println("删除当前用户：", user.name)
			delete(allUsers, user.id)
			message <- logoutInfo

			conn.Close()
			return
		case <-restTimer:
			fmt.Printf("%s 重置计时器\n", user.name)
		}
	}
}
