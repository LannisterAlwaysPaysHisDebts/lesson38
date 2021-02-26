package main

import (
	common "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/05common"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/09order/domain/repository"
	service2 "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/09order/domain/service"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/09order/handler"
	order "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/09order/proto/order"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	rateLimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	openTracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

var (
	QPS         = 1000
	ConsulHost  = "192.168.199.198"
	ServiceName = "go.micro.service.order"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig(ConsulHost, 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}

	// 服务注册
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			ConsulHost + ":8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer(ServiceName, ":6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 数据库
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql", common.GetMysqlDSN(*mysqlInfo))
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	// 初始化数据表
	orderRepository := repository.NewOrderRepository(db)
	err = orderRepository.InitTable()
	if err != nil {
		log.Error(err)
	}

	// 创建实例
	orderSrv := service2.NewOrderDataService(orderRepository)

	// 暴露监控地址
	common.PrometheusBoot(9092)

	// 启动service
	srv := micro.NewService(
		micro.Name(ServiceName),
		micro.Version("latest"),
		micro.Address(":9085"),
		micro.Registry(consul),

		micro.WrapHandler(openTracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(rateLimit.NewHandlerWrapper(QPS)),
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	srv.Init()

	err = order.RegisterOrderHandler(srv.Server(), &handler.Order{OrderDataService: orderSrv})
	if err != nil {
		log.Error(err)
	}

	if err = srv.Run(); err != nil {
		log.Error(err)
	}
}
