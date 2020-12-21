package mapper

import (
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
	"go-resk/src/entity/po"
	"go-resk/src/infra/base"
	_ "go-resk/src/test" // 测试引入
	"go-resk/src/utils"
	"testing"
)

func TestNewUserDao(t *testing.T) {
	Convey("TestNewUserDao", t, func() {
		err := base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
			userDao := NewUserDao(runner)
			So(userDao, ShouldNotBeNil)
			So(userDao.runner, ShouldNotBeNil)
			return nil
		})
		if err != nil {
			logrus.Error(err)
		}
	})
}

func TestUserDao_DeleteByUserId(t *testing.T) {
	Convey("TestUserDao_DeleteByUserId", t, func() {
		base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
			userDao := NewUserDao(runner)
			user := &po.User{
				UserId:       utils.GetUUID(),
				UserName:     utils.GetUUID(),
				UserPassword: "TestUserDao_DeleteByUserId",
			}
			id, err := userDao.Insert(user)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			// 删除
			rows, err := userDao.DeleteByUserId(user.UserId)
			So(rows, ShouldBeGreaterThan, 0)
			So(err, ShouldBeNil)
			// 回查校验
			userRS, err := userDao.GetUserByUserId(user.UserId)
			So(err, ShouldBeNil)
			So(userRS, ShouldBeNil)
			// 数据回滚
			runner.Rollback()
			return nil
		})
	})
}

func TestUserDao_GetUserByUserId(t *testing.T) {
	Convey("TestUserDao_GetUserByUserId", t, func() {
		base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
			userDao := NewUserDao(runner)
			user := &po.User{
				UserId:       utils.GetUUID(),
				UserName:     utils.GetUUID(),
				UserPassword: "TestUserDao_GetUserByUserId",
			}
			id, err := userDao.Insert(user)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			// 回查校验
			userRS, err := userDao.GetUserByUserId(user.UserId)
			So(err, ShouldBeNil)
			So(userRS, ShouldNotBeNil)
			So(userRS.UserId, ShouldEqual, user.UserId)
			So(userRS.UserName, ShouldEqual, user.UserName)
			So(userRS.UserPassword, ShouldEqual, user.UserPassword)
			// 数据回滚
			runner.Rollback()
			return nil
		})
	})
}

func TestUserDao_GetUserByUsername(t *testing.T) {
	Convey("TestUserDao_GetUserByUsername", t, func() {
		base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
			userDao := NewUserDao(runner)
			user := &po.User{
				UserId:       utils.GetUUID(),
				UserName:     utils.GetUUID(),
				UserPassword: "TestUserDao_GetUserByUsername",
			}
			id, err := userDao.Insert(user)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			// 回查校验
			userRS, err := userDao.GetUserByUsername(user.UserName)
			So(err, ShouldBeNil)
			So(userRS, ShouldNotBeNil)
			So(userRS.UserId, ShouldEqual, user.UserId)
			So(userRS.UserName, ShouldEqual, user.UserName)
			So(userRS.UserPassword, ShouldEqual, user.UserPassword)
			// 数据回滚
			runner.Rollback()
			return nil
		})
	})
}

func TestUserDao_Insert(t *testing.T) {
	Convey("TestUserDao_Insert", t, func() {
		base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
			userDao := NewUserDao(runner)
			user := &po.User{
				UserId:       utils.GetUUID(),
				UserName:     utils.GetUUID(),
				UserPassword: "TestUserDao_Insert",
			}
			id, err := userDao.Insert(user)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			// 回查校验
			userRS, err := userDao.GetUserById(id)
			So(err, ShouldBeNil)
			So(userRS, ShouldNotBeNil)
			So(userRS.UserId, ShouldEqual, user.UserId)
			So(userRS.UserName, ShouldEqual, user.UserName)
			So(userRS.UserPassword, ShouldEqual, user.UserPassword)
			// 数据回滚
			runner.Rollback()
			return nil
		})
	})
}

func TestUserDao_UpdateUser(t *testing.T) {
	Convey("TestUserDao_UpdateUser", t, func() {
		base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
			userDao := NewUserDao(runner)
			user := &po.User{
				UserId:       utils.GetUUID(),
				UserName:     utils.GetUUID(),
				UserPassword: "TestUserDao_Insert",
			}
			// 先插入数据
			id, err := userDao.Insert(user)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			// 回查数据
			userRS, err := userDao.GetUserById(id)
			So(err, ShouldBeNil)
			So(userRS, ShouldNotBeNil)
			So(userRS.UserId, ShouldEqual, user.UserId)
			So(userRS.UserName, ShouldEqual, user.UserName)
			So(userRS.UserPassword, ShouldEqual, user.UserPassword)
			// 获取id、user_id用以测试更新数据
			Convey("通过Id更新用户数据", func() {
				userUpdate := &po.User{
					Id:       2,
					UserName: utils.GetUUID(),
				}
				rows, err := userDao.UpdateUser(userUpdate)
				So(rows, ShouldBeGreaterThan, 0)
				So(err, ShouldBeNil)
				// 回查校验
				userRS, err := userDao.GetUserById(2)
				So(err, ShouldBeNil)
				So(userRS.UserName, ShouldEqual, userUpdate.UserName)
			})
			Convey("通过UserId更新用户数据", func() {
				userUpdate := &po.User{
					UserId:   user.UserId,
					UserName: utils.GetUUID(),
				}
				idRs, err := userDao.UpdateUser(userUpdate)
				logrus.Info(idRs)
				So(err, ShouldBeNil)
				// 回查校验
				userRS, err = userDao.GetUserByUserId(user.UserId)
				So(err, ShouldBeNil)
				So(userRS.UserName, ShouldEqual, userUpdate.UserName)
			})
			// 数据回滚
			runner.Rollback()
			return nil
		})
	})
}

func TestUserDao_GetUserByUserNameAndPassword(t *testing.T) {
	Convey("TestUserDao_GetUserByUserNameAndPassword", t, func() {
		base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
			user := &po.User{
				UserName:     utils.GetUUID(),
				UserPassword: "GetUserByUserNameAndPassword",
			}
			// 插入测试用户
			userDao := NewUserDao(runner)
			rows, err := userDao.Insert(user)
			So(err, ShouldBeNil)
			So(rows, ShouldBeGreaterThan, 0)
			// 取用户看是否存在
			userDTO := userDao.GetUserByUserNameAndPassword(user.UserName, user.UserPassword)
			So(userDTO, ShouldNotBeNil)
			runner.Rollback()
			return nil
		})
	})
}
