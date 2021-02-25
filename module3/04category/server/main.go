package main

import (
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/04category/common"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/04category/domain/repository"
	service2 "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/04category/domain/service"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/04category/handler"
	category "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/04category/proto/category"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	consulHost := "192.168.199.198"

	//配置中心
	consulConf, err := common.GetConsulConfig(consulHost, 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}

	//注册中心
	consulRegister := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulHost + ":8500",
		}
	})

	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address(""),
		micro.Registry(consulRegister),
	)

	mysqlInfo := common.GetMysqlFromConsul(consulConf, "mysql")
	db, err := gorm.Open("mysql",
		mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+":"+strconv.FormatInt(mysqlInfo.Port, 10)+")/"+mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	db.SingularTable(true)

	rp := repository.NewCategoryRepository(db)
	err = rp.InitTable()
	if err != nil {
		log.Error(err)
	}

	service.Init()

	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))

	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{Srv: categoryDataService})
	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
