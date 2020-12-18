package base

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	irisrecover "github.com/kataras/iris/middleware/recover"
	"github.com/sirupsen/logrus"
	"go-resk/src/infra"
	"time"
)

type IrisStarter struct {
	infra.BaseStarter
}

var irisServer *iris.Application

func IrisServer() *iris.Application {
	return irisServer
}

// 采用 main goroutine 阻塞式启动
func (i *IrisStarter) StartBlocking() bool {
	return true
}

func (i *IrisStarter) Init(ctx infra.StarterContext) {
	irisServer = initIrisServer()
	// 使用logrus配置日志模块
	logger := irisServer.Logger()
	logger.Install(logrus.StandardLogger())
}

func (i *IrisStarter) Start(ctx infra.StarterContext) {
	// 打印展示路由信息
	routes := IrisServer().GetRoutes()
	for _, r := range routes {
		logrus.Info(r.Trace())
	}
	//启动iris
	port := ctx.Props().GetDefault("app.server.port", "8080")
	IrisServer().Run(iris.Addr(":" + port))
}

func initIrisServer() *iris.Application {
	app := iris.New()
	app.Use(irisrecover.New())
	// 主要中间件的配置:recover,日志输出中间件的自定义
	cfg := logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
		Query:  true,
		LogFunc: func(now time.Time, latency time.Duration,
			status, ip, method, path string,
			message interface{},
			headerMessage interface{}) {
			app.Logger().Infof("| %s | %s | %s | %s | %s | %s | %s | %s",
				now.Format("2006-01-02.15:04:05.000000"),
				latency.String(), status, ip, method, path, headerMessage, message,
			)
		},
	}
	app.Use(logger.New(cfg))
	return app
}
