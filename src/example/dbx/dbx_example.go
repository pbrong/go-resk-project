package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tietang/dbx"
	"time"
)

type GolangTest struct {
	Id   int    `db:"id,id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	// dbx
	settings := dbx.Settings{
		DriverName: "mysql",
		User:       "arong",
		Password:   "123",
		Database:   "test",
		Host:       "10.211.55.30:3306",
		Options: map[string]string{
			"charset":   "utf8",
			"parseTime": "true",
		},
		ConnMaxLifetime: 7 * time.Hour,
		MaxOpenConns:    2,
		MaxIdleConns:    5,
		LoggingEnabled:  false,
	}
	db, err := dbx.Open(settings)
	if err != nil {
		fmt.Println(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	var res []GolangTest
	err = db.Find(&res, "select * from golang_test")
	if err != nil {
		fmt.Println("无法映射数据")
	} else {
		for _, v := range res {
			fmt.Println(v)
		}
	}
}
