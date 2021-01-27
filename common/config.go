package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)
//设置配置中心
func GetConsualConfig(host string,port int64,prefix string)(config.Config,error){
	consualSource := consul.NewSource(
		//设置配置中心地址
		consul.WithAddress(host + ":" + strconv.FormatInt(port, 10)),
		//设置前缀
		consul.WithPrefix(prefix),
		//设置是否移除前缀,表示可以不带前缀直接获取对应的地址
		consul.StripPrefix(true),
		)
	//配置初始化`
	config,err := config.NewConfig()
	if err != nil{
		return config,err
	}
	//加载配置
	err = config.Load(consualSource)
	return config,err
}