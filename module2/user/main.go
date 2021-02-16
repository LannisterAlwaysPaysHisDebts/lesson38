package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/LannisterAlwaysPaysHisDebts/lesson38/common/redis"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module2/user/conf"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module2/user/dao"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module2/user/endpoint"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module2/user/service"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module2/user/transport"
)

func main() {
	var (
		// 服务器地址与服务器名
		servicePort = flag.Int("service.port", 10086, "service port")
	)
	flag.Parse()

	ctx := context.Background()

	errChan := make(chan error)
	err := dao.InitMysql(conf.InitLocalDb())
	if err != nil {
		log.Fatal(err)
	}

	redisConf := conf.InitLocalRedis()
	err = redis.InitRedis(redisConf.Host, redisConf.Port, redisConf.Passwd)
	if err != nil {
		log.Fatal(err)
	}

	userService := service.MakeUserServiceImpl(&dao.UserDAOImpl{})

	userEndpoints := &endpoint.UserEndpoints{
		RegisterEndpoint: endpoint.MakeRegisterEndpoint(userService),
		LoginEndpoint:    endpoint.MakeLoginEndpoint(userService),
	}

	r := transport.MakeHttpHandler(ctx, userEndpoints)

	go func() {
		errChan <- http.ListenAndServe(":"+strconv.Itoa(*servicePort), r)
	}()

	go func() {
		// 监控系统信号，等待 ctrl + c 系统信号通知服务关闭
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	e := <-errChan
	log.Println(e)
}
