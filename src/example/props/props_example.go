package main

import (
	"fmt"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

func main() {
	path := kvs.GetCurrentFilePath("config.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(path)
	appName := conf.GetDefault("app.name", "test")
	appPort := conf.GetIntDefault("app.server.port", 8080)
	appEnable := conf.GetBoolDefault("app.enable", false)
	fmt.Println(appName)
	fmt.Println(appPort)
	fmt.Println(appEnable)
}
