package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"user/domain/repository"
	service2 "user/domain/service"
	"user/handler"
	users "user/proto"
)

func main() {
	srv := micro.NewService(
		micro.Name("service.user"),
		micro.Version("latest"),
		)
	srv.Init()
	//创建数据库
	db,err := gorm.Open("mysql",
		"gostudy:gostudy@tcp(39.107.65.116)/gostudy?charset=utf8&parseTime=True&loc=Local",
		)
	if err!=nil{
		fmt.Println(err)
	}
	defer db.Close()
	//执行完所有后关闭数据库
	db.SingularTable(true)

	rp:=repository.NewUserRepository(db)
	rp.Initable()

	userDataservice := service2.NewUserDataService(repository.NewUserRepository(db))
	//创建实例
	err = users.RegisterUserHandler(srv.Server(),&handler.User{
		UserDataService: userDataservice,
	})
	//绑定实例
	if err!=nil{
		fmt.Println(err)
	}
	if err:=srv.Run();err!=nil{
		fmt.Println(err)
	}
}