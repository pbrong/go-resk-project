package resk

import (
	"github.com/sirupsen/logrus"
	"go-resk/src/infra"
	"go-resk/src/infra/base"
)

// 全局Starter初始化
func init() {
	logrus.Info("全局Starter初始化")
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxStarter{})
	infra.Register(&base.LogStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.WebApiStarter{})
	infra.Register(&base.IrisStarter{})
}
