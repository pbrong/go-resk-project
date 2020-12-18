package test

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"go-resk/src/infra"
	"go-resk/src/infra/base"
)

func init() {
	logrus.Info("测试Starter初始化")
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxStarter{})
	infra.Register(&base.LogStarter{})
	infra.Register(&base.ValidatorStarter{})
	path := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(path)
	app := infra.NewBootApplication(conf)
	app.Start()
}
