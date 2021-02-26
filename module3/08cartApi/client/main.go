package main

import (
	"context"
	"fmt"
	common "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/05common"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/08cartApi/handler"
	cart "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/08cartApi/proto/cart"
	cartApi "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/08cartApi/proto/cartApi"
	"github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	"github.com/opentracing/opentracing-go"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
)

func main() {
	consulHost := "192.168.199.198"

	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":8500",
		}
	})

	//链路追踪
	t, io, err := common.NewTracer("go.micro.api.cartApi", consulHost+":6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	//启动端口
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "12399"), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8086"),
		//添加 consul 注册中心
		micro.Registry(consul),
		//添加链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//添加熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		//添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	cartService := cart.NewCartService("go.micro.service.cart", service.Client())
	_, err = cartService.AddCart(context.TODO(), &cart.CartInfo{
		UserId:    3,
		ProductId: 4,
		SizeId:    5,
		Num:       5,
	})
	if err != nil {
		panic(err)
	}

	// Register Handler
	if err := cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartSrv: cartService}); err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//run 正常执行
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(e error) error {
		fmt.Println(e)
		return e
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}
