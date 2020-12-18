package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	conf := viper.New()
	conf.SetConfigName("config")
	conf.SetConfigType("yaml")
	conf.AddConfigPath("./src/example/viper")
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	serverName := conf.GetString("server.name")
	serverEnable := conf.GetBool("server.enable")
	serverPort := conf.GetInt("server.port")
	fmt.Println(serverName)
	fmt.Println(serverPort)
	fmt.Println(serverEnable)
}
