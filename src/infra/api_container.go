package infra

// WebApi的容器，启动阶段被被WebStarter的Init函数获取并执行每个WebApi的Init方法，注册路由
type WebApi interface {
	Init()
}

type webApiContainer struct {
	WebApis []WebApi
}

var apiContainer = new(webApiContainer)

// webApiContainer全局唯一暴露点
func GetWebApiContainer() *webApiContainer {
	return apiContainer
}

func RegisterApi(api WebApi) {
	apiContainer.WebApis = append(apiContainer.WebApis, api)
}
