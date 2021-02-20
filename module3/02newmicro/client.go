package main

import (
	"context"
	"fmt"

	go_micro_service_imooc "newmicro/proto"

	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("cap.imooc.client"),
	)

	service.Init()

	capImooc := go_micro_service_imooc.NewCapService(
		"cap.imooc.server", service.Client())

	res, err := capImooc.SayHello(context.TODO(),
		&go_micro_service_imooc.SayRequest{
			Message: "你好你好",
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("server: ", res.Answer)
}
