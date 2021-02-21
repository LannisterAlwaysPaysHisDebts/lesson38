package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	go_micro_service_user "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/03user/proto/user"

	"github.com/micro/go-micro/v2"
)

func main() {
	// 创建micro服务
	service := micro.NewService(micro.Name("go.micro.client.user"), micro.Version("latest"))
	service.Init()

	// rpc 初始化service
	userService := go_micro_service_user.NewUserService("go.micro.service.user", service.Client())

	// 执行业务逻辑
	ctx := context.TODO()

	var (
		userName  = "chaoliu" + strconv.FormatInt(time.Now().Unix(), 10)
		firstName = "liu"
		pwd       = "999999"
	)

	// 注册
	registerReq := go_micro_service_user.UserRegisterRequest{
		FirstName: firstName,
		UserName:  userName,
		Pwd:       pwd,
	}
	res, err := userService.Register(ctx, &registerReq)
	if err != nil {
		fmt.Printf("register error: %+v\n", err.Error())
	} else {
		fmt.Printf("register: %+v\n", res.Message)
	}
	// 登陆
	loginReq := go_micro_service_user.UserLoginRequest{UserName: userName, Pwd: pwd}
	logRes, err := userService.Login(ctx, &loginReq)
	if err != nil {
		fmt.Printf("login error : %+v \n", err.Error())
	} else {
		fmt.Printf("login: %+v\n", logRes)
	}
	// 获取用户信息
	infoReq := go_micro_service_user.UserInfoRequest{UserName: userName}
	infoRes, err := userService.GetUserInfo(ctx, &infoReq)
	if err != nil {
		fmt.Printf("getUserInfo error: %+v\n", err.Error())
	} else {
		fmt.Printf("getUserInfo: %+v\n", infoRes)
	}

}
