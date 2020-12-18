package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "go-resk/src"
	_ "go-resk/src/apis"
	"go-resk/src/infra"
)

func main() {
	path := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(path)
	app := infra.NewBootApplication(conf)
	app.Start()
	//select {}
}
