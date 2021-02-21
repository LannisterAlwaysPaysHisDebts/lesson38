package main

import (
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/03user/conf"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/03user/domain/repository"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/03user/domain/service"
	"github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/03user/handler"
	user "github.com/LannisterAlwaysPaysHisDebts/lesson38/module3/03user/proto/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
)

func main() {
	// 启动micro服务， 命名
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:20999"),
	)
	srv.Init()

	// 建立数据库连接
	db, err := gorm.Open("mysql", conf.NewDbArgs().DSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SingularTable(true)

	// 初始化数据表，只执行一次
	rp := repository.NewUserRepository(db)
	err = rp.InitTable()
	if err != nil {
		panic(err)
	}

	// 初始化service
	userSrv := service.NewUserDataService(rp)

	// 注册handler
	err = user.RegisterUserHandler(srv.Server(),
		&handler.User{UserDataService: userSrv})
	if err != nil {
		panic(err)
	}

	if err := srv.Run(); err != nil {
		panic(err)
	}
}
