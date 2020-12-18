package test

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
	"go-resk/src/dao/mapper"
	"go-resk/src/entity/po"
	"go-resk/src/infra/base"
	_ "go-resk/src/test"
	"go-resk/src/utils"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestAccountDao_GetByUserId(t *testing.T) {
	err := base.DbxDatabase().Tx(func(tx *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(tx)
		Convey("TestAccountDao_GetByUserId", t, func() {
			account := accountDao.GetByUserId("test", 1)
			if account != nil {
				So(account.Id, ShouldNotEqual, 0)
			}
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

func TestAccountDao_GetOne(t *testing.T) {
	err := base.DbxDatabase().Tx(func(tx *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(tx)
		Convey("TestAccountDao_GetOne", t, func() {
			account := accountDao.GetOne("test")
			if account != nil {
				So(account.Id, ShouldNotEqual, 0)
			}
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

func TestAccountDao_Insert(t *testing.T) {
	err := base.DbxDatabase().Tx(func(tx *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(tx)
		Convey("TestAccountDao_Insert", t, func() {
			rand.Seed(time.Now().UnixNano())
			x := rand.Int()
			random := strconv.Itoa(x)
			a := &po.Account{
				Id:          int64(x),
				Balance:     decimal.NewFromFloat(100),
				Status:      1,
				AccountNo:   utils.GetUUID(),
				AccountName: random,
				UserId:      random,
				Username:    sql.NullString{String: "测试用户", Valid: true},
			}
			id, err := accountDao.Insert(a)
			if err != nil {
				So(id, ShouldEqual, -1)
			} else {
				So(id, ShouldNotEqual, -1)
			}
			//tx.Rollback()
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

func TestAccountDao_Update(t *testing.T) {
	err := base.DbxDatabase().Tx(func(tx *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(tx)
		Convey("TestAccountDao_Update", t, func() {
			rand.Seed(time.Now().UnixNano())
			x := rand.Int()
			random := strconv.Itoa(x)
			a := &po.Account{
				Id:          1,
				Balance:     decimal.NewFromFloat(100),
				Status:      1,
				AccountNo:   "test",
				AccountName: random,
				UserId:      random,
				Username:    sql.NullString{String: "测试用户", Valid: true},
			}
			id, err := accountDao.Update(a)
			if err != nil {
				So(id, ShouldEqual, -1)
			} else {
				So(id, ShouldNotEqual, -1)
			}
			//tx.Rollback()
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

func TestAccountDao_UpdateBalance(t *testing.T) {
	err := base.DbxDatabase().Tx(func(tx *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(tx)
		Convey("TestAccountDao_UpdateBalance", t, func() {
			newBanlace, err := decimal.NewFromString("100")
			rows, err := accountDao.UpdateBalance("test", newBanlace)
			if err != nil {
				So(rows, ShouldEqual, 0)
			} else {
				So(rows, ShouldNotEqual, 0)
			}
			tx.Rollback()
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

func TestAccountDao_UpdateStatus(t *testing.T) {
	err := base.DbxDatabase().Tx(func(tx *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(tx)
		Convey("TestAccountDao_UpdateStatus", t, func() {
			rows, err := accountDao.UpdateStatus("test", 2)
			if err != nil {
				So(rows, ShouldEqual, 0)
			} else {
				So(rows, ShouldNotEqual, 0)
			}
			tx.Rollback()
		})
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}

func TestAccountDao_Insert2(t *testing.T) {
	err := base.DbxDatabase().Tx(func(tx *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(tx)
		account := &po.Account{
			Id:        1,
			AccountNo: "test",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now()}
		id, err := accountDao.Update(account)
		if err != nil {
			return err
		}
		logrus.Info(id)
		return nil
	})
	if err != nil {
		logrus.Error(err)
	}
}
