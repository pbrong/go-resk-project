package base

import (
	"go-resk/src/infra"
)

type WebApiStarter struct {
	infra.BaseStarter
}

// 配置路由信息
func (w *WebApiStarter) Setup(ctx infra.StarterContext) {
	// 获取WebApiContainer并初始化所有WebApi
	webApiContainer := infra.GetWebApiContainer()
	for _, api := range webApiContainer.WebApis {
		api.Init()
	}
}
