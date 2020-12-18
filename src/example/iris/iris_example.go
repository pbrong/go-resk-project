package main

import (
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func main() {
	app := iris.Default()
	app.Get("/hello", func(ctx iris.Context) {
		panic("出错了")
		ctx.WriteString("hello,world! iris")
	})
	v1 := app.Party("/v1")
	v1.Use(func(context iris.Context) {
		logrus.Info("自定义中间件")
		context.Next()
	})
	v1.Get("/users/{id:uint64 min(2)}",
		func(ctx iris.Context) {
			id := ctx.Params().GetUint64Default("id", 0)
			ctx.WriteString(strconv.Itoa(int(id)))
		})
	v1.Get("/orders/{action:string prefix(a_)}", func(ctx iris.Context) {
		a := ctx.Params().Get("action")
		ctx.WriteString(a)
	})
	app.OnAnyErrorCode(func(context iris.Context) {
		context.WriteString("看起来服务器出错了！")
	})
	app.OnErrorCode(http.StatusNotFound, func(context iris.Context) {
		context.WriteString("访问的路径不存在。")
	})
	app.Run(iris.Addr(":8080"))
}
