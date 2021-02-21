package common

import "github.com/micro/go-micro/v2/config"

type MysqlConf struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConf {
	conf := &MysqlConf{}
	config.Get(path...).Scan(conf)
	return conf
}
