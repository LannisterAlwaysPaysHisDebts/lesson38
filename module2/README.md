# go-kit
Go Kit 是 Go 语言工具包的集合，可帮助工程师构建强大、可靠和可维护的微服务。它提供了用于实现系统监控和弹性模式组件的库，例如日志记录、跟踪、限流和熔断等，这些库可以协助工程师提高微服务架构的性能和稳定性。

![Go Kit 框架分层](https://s0.lgstatic.com/i/image/M00/2E/D4/CgqCHl8Fm7mAc6jtAACZVV3YQYA093.png)

基于 Go Kit 的应用程序架构由三个主要部分组成：传输层、接口层和服务层。
 - 传输层用于网络通信，服务通常使用 HTTP 或 gRPC 等网络传输方式，或使用 NATS 等发布订阅系统相互通信。除此之外，Go Kit 还支持使用 AMQP 和 Thrift 等多种网络通信模式。
 - 接口层是服务器和客户端的基本构建块。在 Go Kit 中，服务中的每个对外提供的接口方法都会被定义为一个端点，以便在服务器和客户端之间进行网络通信。每个端点通过使用 HTTP 或 gRPC 等具体通信模式对外提供服务。
 -    服务层是具体的业务逻辑实现，包括核心业务逻辑。它不会也不应该进行 HTTP 或 gRPC 等具体网络传输，或者请求和响应消息类型的编码和解码。

# Go Micro 框架
Go Micro 是基于 Go 语言实现的插件化 RPC 微服务框架。它提供了服务发现、负载均衡、同步传输、异步通信以及事件驱动等机制，并尝试去简化分布式系统间的通信，让开发者可以专注于自身业务逻辑的开发。

Go Micro 框架的结构可以描述为三层堆栈，如下图所示：
![Go Micro 三层堆栈](https://s0.lgstatic.com/i/image/M00/2E/D4/CgqCHl8Fm8aACnuYAAB4Q4ubZZg969.png)

可以看到，Go Micro 框架模型的上层由 Service 和 Client-Server 模型抽象组成。Server 服务器是用于编写服务的构建块；Client 客户端提供了向服务请求的接口；底层由代理、编解码器和注册表等类型的插件组成。Go Micro 是组件化的框架，每一个基础功能都有对应的接口抽象，方便扩展。

另外，Go Micro 具有可插拔的特点，其提供的组件可帮助我们快速构建应用系统，并且可以定制所需要的插件功能。（相关插件可在仓库 github.com/micro/go-plugins 中找到。）

# Go Kit 与 Go Micro 的对比
Go Kit 是一个微服务的标准库。它可以提供独立的包，通过这些包，开发者可以用来组建自己的应用程序。微服务架构意味着构建分布式系统，这带来了许多挑战，Go Kit 可以为多数业务场景下实施微服务软件架构提供指导和解决方案。

Go Micro 是一个面向微服务的可插拔 RPC 框架，可以快速启动微服务的开发。Go Micro 框架提供了许多功能，无须重新“造轮子”，所以开发者可以花更多的时间在需要关注的业务逻辑上。但是 Go Micro 在快速启动微服务开发的同时，也牺牲了灵活性，并且将 gRPC 强制为默认通信类型，更换组件不如 Go Kit 简便。



# user
文件结构分为
1. dao包: 提供 MySQL 数据层持久化能力；
2. endpoint包： 负责接收请求，并调用 service 包中的业务接口处理请求后返回响应；
3. redis包： 提供 Redis 数据层操作能力；
4. service包： 提供主要业务实现接口；
5. transport包： 对外暴露项目的服务接口；
6. main: 应用主入口。






