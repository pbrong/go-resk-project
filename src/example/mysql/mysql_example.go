package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
)

var (
	id   int
	name string
	age  int
)

func main() {
	test1()
}

func test1() {
	// database/sql
	path := kvs.GetCurrentFilePath("mysql.ini", 1)
	conf := ini.NewIniFileCompositeConfigSource(path)
	/**
	driverName = mysql
	host = 127.0.0.1
	port = 3306
	dbName = test
	username = root
	password = 123
	*/
	driver := conf.GetDefault("mysql.driverName", "mysql")
	dbName := conf.GetDefault("mysql.dbName", "nil")
	host := conf.GetDefault("mysql.host", "127.0.0.1")
	port := conf.GetIntDefault("mysql.port", 3306)
	username := conf.GetDefault("mysql.username", "root")
	password := conf.GetDefault("mysql.password", "123")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true&loc=Local",
		username, password, host, port, dbName)
	db, err := sql.Open(driver, dataSource)
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 查询所有数据
	rows, err := db.Query("select * from golang_test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("get name = %v, age = %v \n", name, age)
	}
}
