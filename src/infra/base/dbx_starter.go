package base

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"go-resk/src/infra"
)

type DbxStarter struct {
	infra.BaseStarter
}

//dbx 数据库实例
var database *dbx.Database

func DbxDatabase() *dbx.Database {
	return database
}

func (s *DbxStarter) Setup(ctx infra.StarterContext) {
	conf := ctx.Props()
	//数据库配置
	settings := dbx.Settings{}
	err := kvs.Unmarshal(conf, &settings, "mysql")
	if err != nil {
		panic(err)
	}
	settings.Options = map[string]string{
		"charset":   "utf8",
		"parseTime": "True",
		"loc":       "Asia%2FShanghai",
	}
	logrus.Info("mysql.conn url:", settings.ShortDataSourceName())
	dbx, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	logrus.Info(dbx.Ping())
	database = dbx
}
