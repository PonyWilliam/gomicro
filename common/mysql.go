package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct{
	Host string `json:"host"`
	Port int64 `json:"port"`
	User string `json:"user"`
	Pwd string `json:"pwd"`
	DataBase string `json:"database"`
}

//获取mysql的配置
func GetMysqlFromConsul(config config.Config,path ...string) *MysqlConfig{
	mysqlConfig := &MysqlConfig{}
	config.Get(path...).Scan(mysqlConfig)//扫描进mysqlconfig
	return mysqlConfig
}