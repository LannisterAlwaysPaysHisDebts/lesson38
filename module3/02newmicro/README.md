1. proto定义Cap服务的SayHello方法；
2. server实现SayHello方法；
3. client通过CapService调用SayHello方法(rpc)；

```/bin/bash
cd proto
protoc -I ./ --go_out=./ --micro_out=./ ./*.proto

// server
go run server/server.go

// client
go run client.go

```