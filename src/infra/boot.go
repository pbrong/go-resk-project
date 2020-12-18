package infra

import "github.com/tietang/props/kvs"

// 全局资源启动器
type BootApplication struct {
	conf kvs.ConfigSource
	ctx  StarterContext
}

func NewBootApplication(conf kvs.ConfigSource) *BootApplication {
	application := &BootApplication{conf: conf, ctx: StarterContext{}}
	application.ctx[propsKey] = conf
	return application
}

func (b *BootApplication) Start() {
	// 初始化Starter
	b.init()
	// 安装配置Starter
	b.setup()
	// 启动所有Starter
	b.start()
}

func (b *BootApplication) init() {
	for _, s := range GetStarters() {
		s.Init(b.ctx)
	}
}

func (b *BootApplication) setup() {
	for _, s := range GetStarters() {
		s.Setup(b.ctx)
	}
}

func (b *BootApplication) start() {
	for index, s := range GetStarters() {
		if !s.StartBlocking() {
			// 非阻塞式启动
			if index+1 == len(GetStarters()) {
				s.Start(b.ctx)
			} else {
				go s.Start(b.ctx)
			}
		} else {
			// 阻塞式启动
			s.Start(b.ctx)
		}
	}
}

func (b *BootApplication) Stop() {
	for _, s := range GetStarters() {
		s.Stop(b.ctx)
	}
}
