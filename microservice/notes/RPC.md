
# RPC协议

## 什么是RPC

Remote Procedure Call Protocol -- 远程过程调用协议
IPC: 进程间通信
RPC: 远程进程间通信 ——应用层协议(与HTTP同层)。底层使用TCP实现

回顾

* OSI 7层：物、数、网、传、会、表、应
* TCP/IP层：链、网、传、应

* 理解RPC：
  * 像调用本地函数一样，去调用远程函数：
    * 通过RPC协议，传递：函数名、函数参数，达到在本地，调用远端函数的返回值到本地的目标

* 为什么微服务使用RPC：
  1. 每个服务都被封装成进程，彼此独立
  2. 进程和进程之间，可以使用不同的语言实现

### RPC入门使用

远程 —— 网络

```GO
回顾：Go 语言 socket通信
server端：
    listener = net.Listen()
    conn = listener.Accept()
    conn.Read()
    conn.Write()
    defer conn.Close() /listener.Clos()
client端：
    conn = net.Dial()
    conn.Write()
    conn.Read()
    defer conn.Clos()
```

### RPC使用步骤

--- 服务端

1. 注册RPC服务对象。给对象绑定方法（1.定义类，2.绑定类方法）

    ```go
    rpc.RegisterName("服务名", 回调对象)
    ```

2. 创建监听器

    ```go
    listener, err := net.Listen()
    ```

3. 建立连接

    ```go
    conn, err := listener.Accept()
    ```

4. 连接绑定RPC服务

    ```go
    rpc.ServeConn(conn)
    ```

--- 客户端

1. 用RPC连接服务器

    ```go
    conn, err := rpc.Dial()
    ```

2. 调用远程函数

    ```go
    conn.Call("服务名.方法名", 传入参数, 传出参数)
    ```

## RPC相关函数

1、注册RPC服务

```go
func (server *Server) RegisterName(name string, rcvr interface{}) error
    参1:服务名。字符串类型
    参2:对应的rpc对象。该对象绑定的方法满足如下条件：
        1. 方法必须是导出的 --包外可见（首字母大写）
        2. 方法必须有两个参数， 都是导出类型、内建类型
        3. 方法的第二个参数必须是“指针”（传出参数）
        4. 方法只有一个error接口类型的返回值

举例：
type World struct {
}
func (this *World) HelloWorld(name string, resp *string) error {
}
rpc.RegisterName("服务名", new(World))
```

2、绑定RPC服务

```go
func (server *Server) ServeConn(conn io.ReadWriteCloser)
    conn:建立好的socket(connect)
```

3、调用远程函数

```go
func (client *Clientt) Call(serveceMethod string, args interface{}, reply interface{}) error
    serviceMethod：服务名.方法名
    args：传入参数
    reply：传出参数。定义var 变量，&变量名 完成传参
```

## json版rpc

* 使用nc -l 127.0.0.1 8800 充当服务器
* 02-client.go充当客户端发起通信 ——乱码
  * 因为RPC使用go语言特有的序列化gob，其他编程语言不能解析
* 解决乱码方法：使用通用的序列化、反序列化工具——json 、 protobuf
  * go中自带相关的库：net/rpc/jsonrpc

## rpc封装

### 服务端封装

定义接口

```go
type xxx interface {
    方法名(传入参数, 传出参数) error
}
例
type MyInterface interface {
    HelloWorld(string, *string) error
}
```

封装注册服务方法

```go
func RegisterService (i MyInterface) {
    rpc.RegisterNmae("hello", i)
}
```

### 客户端封装

```go
// 定义类
type MyClient struct {
    c *rpc.Client
}
```

```go
// 绑定类方法
func (mc *MyClient) HelloWorld (a string, b *string) error {
    return mc.c.Call("hello.HelloWorld", a, b)
}
```

```go
// 初始化客户端
func InitClient(addr string) MyClient {
    conn, _ := jsonrpc.Dial("tcp", addr)
    return MyClient{c:conn}
}
```

# protobuf

## archlinux安装

* protobuf安装
  * 所有语言通用，使用pacman安装
  * `sudo pacman -S protobuf`
* go编译插件安装
  * `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
* grpc编译插件安装
  * `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0`

——Google

## 语法

```protobuf
// 默认是proto2
syntax = "proto3";

// 指定所在包包名
package demo1; // 包名（疑似被下面的方法替代）
option go_package = "./;demo"; // 分号前是生成路径，分号后是包名

// 定义枚举类型
enum Week {
    Monday = 0; // 枚举值必须从-开始
    Tuesday = 1;
}

// 定义消息体
message Student {
    int32 age = 1; // 可以不从1开始，但是不能重复。——不能使用19000 - 19999，这些值为协议保留
    string name = 2;
    People p = 3;
    repeated int32 score = 4;
    // 枚举
    Week w = 5;
    // 联合体
    oneof data {
        string teacher = 6;
        string class = 7;
    }
}

// 消息体可以嵌套
message People {
    int32 weight = 1;
}
```

## 编写的注意事项

1. message成员编号可以不从1开始，但是不能重复。——不能使用19000 - 19999，这些值为协议保留
2. 可以使用message嵌套
3. 定义数组、切片使用repeated关键字
4. 可以使用枚举enum
5. 可以使用联合体oneof关键字，成员编号也不能重复。

## 编译protobuf

> 回顾：C++编译命令：
> protoc --cpp_out=./ *.proto

* go语言中编译命令
```protoc --go_out=./ *.proto```

## 添加rpc服务

* 语法:
  
```protobuf
service 服务名{
    rpc 函数名(参数:消息体) returns (返回值: 消息)
}
例：
message People {
    string name = 1;
}
message Student {
    int32 age = 2;
}
service Hello {
    rpc HelloWorld (People) returns (Student);
}
```

* 知识点：
  * protobuf编译期间，默认不编译服务，需要使用gRPC。
  * gRPC编译指令：
    * `protoc --go_out=plugins=ggrpc:./ *.proto`\
    * 新版`protoc --go-grpc_out=./ *.proto`
    * `protoc -I$GOPATH/src -I. --go-grpc_out=require_unimplemented_servers=false:$GOPATH/src *.proto`
* 生成的.rpc.go与自己写的对比

```go
// 客户端
type helloClient struct{}
func NewHelloClient(cc grpc.ClientConnInterface) HelloClient{}
func (c *helloClient) GoToSchool(ctx context.Context, in *People, opts ...grpc.CallOption) (*Student, error) {}
// 服务端
type HelloServer interface {}
func RegisterHelloServer(s grpc.ServiceRegistrar, srv HelloServer) {}
```

## 示例

* 服务端
  
```go
// 定义类对象
type Lily struct{}

// 绑定类方法
func (l *Lily) Teaching(ctx context.Context, in *pb.Teacher) (*pb.Teacher, error) {

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
```

* 客户端

```go
func main() {
    // 1. 连接grpc服务 老版：conn, err := grpc.Dial("127.0.0.1:8800", grpc,WithInsecure())
    conn, err := grpc.NewClient("127.0.0.1:8800", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        fmt.Println("Dial failed : ", err)
        return
    }
    defer conn.Close()

    // 2. 初始化grpc客户端
    c := pb.NewTeachClient(conn)

    // 初始化Teacher对象
    var tarcher pb.Teacher
    teacher.Name = 
    teacher.Age = 

    // 3. 调用远程服务
    r, err := c.Teaching(context.Background(), &pb.Teacher{Name: "LL", Age: 18})
    r, err := c.Teaching(context.Background(), &teacher) // TODO：虽然传的指针，但是服务端的修改不会同步？
    if err != nil {
        fmt.Println("Teaching failed : ", err)
        return
    }
    fmt.Println("Response : ", r.Age, " , ", r.Name)

}
```
