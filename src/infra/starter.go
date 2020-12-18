package infra

import (
	"github.com/tietang/props/kvs"
)

// conf对应的key
var propsKey = "propsKey"

// 启动器上下文
type StarterContext map[string]interface{}

func (s StarterContext) Props() kvs.ConfigSource {
	conf := s[propsKey]
	if conf == nil {
		panic("配置文件未被指定")
	}
	return conf.(kvs.ConfigSource)
}

// Starter接口定义
type Starter interface {
	// 初始化
	Init(StarterContext)
	// 资源安装
	Setup(StarterContext)
	// 启动资源
	Start(StarterContext)
	// 是否阻塞
	StartBlocking() bool
	// 停止资源
	Stop(StarterContext)
}

// Starter注册器接口
type starterRegister struct {
	allnonBlockingStarters []Starter
	blockingStarters       []Starter
}

// 获得全部Starter
func (s *starterRegister) AllStarters() []Starter {
	starters := make([]Starter, 0)
	starters = append(starters, s.allnonBlockingStarters...)
	starters = append(starters, s.blockingStarters...)
	return starters
}

// 注册Starter
func (s *starterRegister) Register(starter Starter) {
	if starter.StartBlocking() {
		s.blockingStarters = append(s.blockingStarters, starter)
	} else {
		s.allnonBlockingStarters = append(s.allnonBlockingStarters, starter)
	}
}

// 定义全局Register用于存储Starter列表
var StarterRegister = &starterRegister{}

// 注册starter
func Register(starter Starter) {
	StarterRegister.Register(starter)
}

// 获取所有注册的starter
func GetStarters() []Starter {
	return StarterRegister.AllStarters()
}

var _ Starter = new(BaseStarter)

type BaseStarter struct {
}

func (b *BaseStarter) Init(ctx StarterContext)  {}
func (b *BaseStarter) Setup(ctx StarterContext) {}
func (b *BaseStarter) Start(ctx StarterContext) {}
func (b *BaseStarter) StartBlocking() bool      { return false }
func (b *BaseStarter) Stop(ctx StarterContext)  {}
