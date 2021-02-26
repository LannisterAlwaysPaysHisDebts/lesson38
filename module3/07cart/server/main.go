package main

import (
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/05common"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/07cart/domain/repository"
	service2 "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/07cart/domain/service"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/07cart/handler"
	cart "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/07cart/proto/cart"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	rateLimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	openTracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var QPS = 100

func main() {
	consulHost := "192.168.199.198"

	//配置中心
	consulConfig, err := common.GetConsulConfig(consulHost, 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	//注册中心
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":8500",
		}
	})

	//链路追踪
	t, io, err := common.NewTracer("go.micro.service.cart", "192.168.199.198:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	//数据库连接
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	//创建数据库连接
	db, err := gorm.Open("mysql", common.GetMysqlDSN(*mysqlInfo))
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	//禁止副表
	db.SingularTable(true)

	//第一次初始化
	err = repository.NewCartRepository(db).InitTable()
	if err != nil {
		log.Error(err)
	}

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		//暴露的服务地址
		micro.Address("0.0.0.0:8087"),
		//注册中心
		micro.Registry(consul),
		//链路追踪
		micro.WrapHandler(openTracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		micro.WrapHandler(rateLimit.NewHandlerWrapper(QPS)),
	)
	// Initialise service
	service.Init()

	cartDataService := service2.NewCartDataService(repository.NewCartRepository(db))

	// Register Handler
	_ = cart.RegisterCartHandler(service.Server(), &handler.Cart{CartDataService: cartDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
