package main

import (
	"context"
	"fmt"

	imooc "newmicro/proto"

	"github.com/micro/go-micro/v2"
)

// 创建capServer，实现proto里定义的接口
type CapServer struct{}

// 需要实现的方法，参数是固定的
func (c CapServer) SayHello(ctx context.Context, request *imooc.SayRequest, response *imooc.SayResponse) error {
	fmt.Println("client send: ", request.Message)
	response.Answer = "hello"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("cap.imooc.server"),
	)
	service.Init()

	// 注册服务
	err := imooc.RegisterCapHandler(service.Server(), new(CapServer))
	if err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
