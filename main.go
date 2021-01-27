package main

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/micro/go-plugins/registry/consul/v2"
	"strconv"
	"time"
	"user/common"
	"user/domain/repository"
	service2 "user/domain/service"
	"user/handler"
	users "user/proto"
)

func main() {
	//配置中心
	consulConfig,err := common.GetConsualConfig("127.0.0.1",8500,"/micro/config")
	if err != nil{
		//记录日志
		log.Error(err)
	}
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options){
		options.Addrs = []string{"127.0.0.1"}
		options.Timeout = 10 * time.Second
	})
	//NEW service
	srv := micro.NewService(
		micro.Name("service.user"),
		micro.Version("latest"),
		//这里设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8081"),
		//添加consul 作为注册中心
		micro.Registry(consulRegistry),
		)
	//路径中不用代前缀
	mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	//创建数据库
	db,err := gorm.Open("mysql",
		mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host + ":"+ strconv.FormatInt(mysqlInfo.Port,10) +")/"+mysqlInfo.DataBase+"?charset=utf8&parseTime=True&loc=Local",
		)
	if err!=nil{
		log.Error(err)
	}
	defer db.Close()
	//执行完所有后关闭数据库
	db.SingularTable(true)

	srv.Init()



	rp:=repository.NewUserRepository(db)
	rp.I()

	userDataservice := service2.NewUserDataService(repository.NewUserRepository(db))
	//创建实例
	err = users.RegisterUserHandler(srv.Server(),&handler.User{
		UserDataService: userDataservice,
	})
	//绑定实例
	if err!=nil{
		log.Fatal(err)
	}
	if err:=srv.Run();err!=nil{
		log.Fatal(err)
	}
}