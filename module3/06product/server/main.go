package main

import (
	"strconv"

	common "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/05common"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/06product/domain/repository"
	service2 "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/06product/domain/service"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/06product/handler"
	product "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/06product/proto/product"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	openTracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/common/log"
)

func main() {
	consulHost := "192.168.199.198"
	serviceName := "go.micro.service.product"
	tracerAddr := "192.168.199.198:6831"

	// 配置中心
	consulConfig, err := common.GetConsulConfig(consulHost, 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}

	// 注册服务
	consul := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer(serviceName, tracerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 数据库
	mysqlInfo := common.GetMysqlFromConsul(consulConfig, "mysql")
	dbArgs := mysqlInfo.User + ":" + mysqlInfo.Pwd + "@tcp(" + mysqlInfo.Host + ":" + strconv.FormatInt(mysqlInfo.Port, 10) + ")/" + mysqlInfo.Database + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbArgs)
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	err = repository.NewProductRepository(db).InitTable()
	if err != nil {
		log.Fatal(err)
	}

	productDataService := service2.NewProductDataService(repository.NewProductRepository(db))

	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
		micro.Address(":8085"),
		micro.Registry(consul),
		micro.WrapHandler(openTracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	srv.Init()

	err = product.RegisterProductHandler(srv.Server(), &handler.Product{ProductDataService: productDataService})
	if err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
