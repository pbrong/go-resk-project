package mapper

import (
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
	"go-resk/src/infra/base"
	"testing"
)

func TestAccountLogDao_GetByTradeNo(t *testing.T) {
	err := base.DbxDatabase().Tx(func(tx *dbx.TxRunner) error {
		dao := NewAccountLogDao(tx)
		Convey("TestAccountLogDao_GetByTradeNo", t, func() {
			accountLog := dao.GetByTradeNo("test")
			if accountLog != nil {
				So(accountLog.Id, ShouldNotEqual, 0)
			}
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

func TestAccountLogDao_GetOne(t *testing.T) {

}

func TestAccountLogDao_Insert(t *testing.T) {

}

func TestNewAccountLogDao(t *testing.T) {

}
