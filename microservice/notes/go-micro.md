
# micro

## micro 简介

* github.com/micro
  * go-micro: 核心库
  * micro: 运行环境、命令、创建微服务空项目
  * go-plugins: 微服务插件
  * examples: 案例
  * protoc-gen-micro: 生成micro代码

## 服务发现

* 微服务开发的核心

![alt text](image.png)

### 有服务发现后，client、server工作流程

1. 每个server启动是，将自己的ip、port注册和服务名注册给“服务发现”
2. 当client西哪个服务发现发起服务请求时，“服务发现”会自动找一个可用的服务，将其ip/port/服务名返回给client
3. client 再借助服务发现，访问server

### 服务发现种类

* consul: 常被用于go-micro中
* mdns: go-micro中默认自带的服务发现
* etcd: k8s内嵌的服务发现
* zookeeper: java中比较常用

### consul关键特性

1. 服务发现: 服务端主动向consul发起注册
2. 健康检查: 定时发送消息，类似于"心跳包"
3. 键值存储: consul提供，但是一般用redis
4. 多数据中心: 可以轻松搭建集群

### 注册服务到consul

步骤：
1. 创建/etc/consul.d
2. 创建服务文件/etc/consul.d/web.json
  ```json
  {
        "service":{
            "name": "Faceid",
            "tags": ["rails", "subway"],
            "port": 8800
        }
  }
  ```
3. 重新启动
4. 查询服务
   1. 浏览器: IP:8500
   2. 命令行查询: `curl -s https://127.0.0.1:8500/v1/catalog/service/Faceid`

### 健康检查

1. 配置文件中加入
```json
{
    "service":{
    "name": "Faceid",
    "tags": ["rails", "subway"],
    "port": 8800,
    "check": {
        "id": "api",
        "name": "HTTP API on port 8800",
        "http": "http://10.8.56.240:8800",
        "interval": "10s",
        "timeout": "1s"
    }
  }
}
```
2. 执行consul reload重新加载配置文件或者重启
3. 会显示不健康
   1. 因为当前没有注册好的服务
4. 健康检查方式："script"、'"tcp"、"ttl"

## consul与grpc结合

安装consul 源码包：

```shell
go get -u -v github.com/hashicorp/consul
```

### 使用整体流程

1. 创建proto文件，指定rpc服务
2. 启动Consul服务发现
3. 启动server
   1. 获取consul对象
   2. 使用consul对象，将server信息注册给consul
   3. 启动Consul服务发现
4. 启动client
   1. 获取consul对象
   2. 使用consul对象，从consul上获取健康的服务
   3. 再访问服务（grpc远程调用）