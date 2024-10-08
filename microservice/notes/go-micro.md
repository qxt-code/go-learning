
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

### consul安装

下载二进制文件放入/usr/local/bin/

### consul命令

* consul agent
  * -bind=0.0.0.0       指定consul所在机器的ip地址
  * -http-prot=8500     自带的web访问端口
  * -client=127.0.0.1   表明哪些机器可以访问consul。127:默认本机；0.0.0.0:所有机器
  * -config-dir=foo     所有主动注册服务的描述信息
  * -data-dir=path      储存所有注册过来的机器的详细信息
  * -dev                开发者模式，直接以默认配置启动consul
  * -node=hostname      服务发现的名字
  * -rejoin             consul启动时，允许加入当前集群？
  * -server             以服务方式开启consul，允许其他的consul连接（形成集群）。如果不加-server，表示以“客户端”方式开启，不能被连接
  * -ui                 可以使用web页面来查看服务发现的详情
* consul members        查看集群中的成员
* consul info           查看当前consul的Ip信息
* consul leave          优雅地关闭consul。——不优雅：ctrl+c
* 示例
  
```shell
consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul -node=n1 -bind=192.168.198.155 -ui -rejoin -config-dir=/etc/consul.d/ -client 0.0.0.0
```

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

### 编码实现

* 用到的函数

   ```go
   // 从consul服务发现上 获取健康服务
   func (h *Health) Service(service, tag string, passingOnly bool, q *QueryOptions) ([]*ServiceEntry, *QueryMeta, error)
   // 参数
   service : 服务名。 ——注册服务时，指定的string
   tag: 外名/别名。 如果有多个，任选一个
   passingOnly: 是否通过健康检查。 true
   q： 查询参数。 通常传nil
   // 返回值
   ServiceEntgry: 存储服务的切片
   QueryMeta: 额外查询返回值。nil
   error： 错误信息
  ```
