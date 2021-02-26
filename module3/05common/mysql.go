package common

import (
	"strconv"

	"github.com/micro/go-micro/v2/config"
)

type MysqlConf struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Database string `json:"database"`
	Port     int64  `json:"port"`
}

func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConf {
	conf := &MysqlConf{}
	_ = config.Get(path...).Scan(conf)
	return conf
}

func GetMysqlDSN(mysqlInfo MysqlConf) string {
	return mysqlInfo.User + ":" + mysqlInfo.Pwd + "@tcp(" + mysqlInfo.Host + ":" + strconv.FormatInt(mysqlInfo.Port, 10) + ")/" + mysqlInfo.Database + "?charset=utf8&parseTime=True&loc=Local"
}
