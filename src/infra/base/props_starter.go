package base

import (
	"github.com/tietang/props/kvs"
	"go-resk/src/infra"
)

// 实现BaseStarter
type PropsStarter struct {
	infra.BaseStarter
}

var conf kvs.ConfigSource

func (p *PropsStarter) Props() kvs.ConfigSource {
	return conf
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	conf = ctx.Props()
}
