package mapper

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"go-resk/src/entity/dto"
	"go-resk/src/entity/po"
)

type UserDao struct {
	runner *dbx.TxRunner
}

func NewUserDao(runner *dbx.TxRunner) *UserDao {
	return &UserDao{runner: runner}
}

// 新增用户
func (dao *UserDao) Insert(user *po.User) (int64, error) {
	rs, err := dao.runner.Insert(user)
	if err != nil {
		logrus.Error(
			fmt.Sprintf("新增用户失败, err = %+v", err))
		return -1, err
	}
	return rs.LastInsertId()
}

// 根据UserId删除用户
func (dao *UserDao) DeleteByUserId(userId string) (affectd int64, error error) {
	sql := "delete from user where user_id = ?"
	rs, err := dao.runner.Exec(sql, userId)
	if err != nil {
		logrus.Error(
			fmt.Sprintf("根据UserId删除用户失败, err = %+v", err))
		return -1, err
	}
	return rs.RowsAffected()
}

// 根据修改用户数据
func (dao *UserDao) UpdateUser(user *po.User) (affected int64, err error) {
	rs, err := dao.runner.Update(user)
	if err != nil {
		logrus.Error(
			fmt.Sprintf("修改用户数据失败, err = %+v", err))
		return -1, err
	}
	return rs.RowsAffected()
}

// 根据UserId查询数据
func (dao *UserDao) GetUserByUserId(userId string) (*po.User, error) {
	user := &po.User{}
	sql := "select * from user where user_id = ?"
	ok, err := dao.runner.Get(user, sql, userId)
	if err != nil || !ok {
		logrus.Error(
			fmt.Sprintf("根据UserId查询用户数据失败, err = %+v", err))
		return nil, err
	}
	return user, nil
}

// 根据Username查询数据
func (dao *UserDao) GetUserByUsername(username string) (*po.User, error) {
	user := &po.User{}
	sql := "select * from user where user_name = ?"
	ok, err := dao.runner.Get(user, sql, username)
	if err != nil || !ok {
		logrus.Error(
			fmt.Sprintf("根据Username查询用户数据失败, err = %+v", err))
		return nil, err
	}
	return user, nil
}

// 根据id查询数据
func (dao *UserDao) GetUserById(id int64) (*po.User, error) {
	user := &po.User{}
	sql := "select * from user where id = ?"
	ok, err := dao.runner.Get(user, sql, id)
	if err != nil || !ok {
		logrus.Error(
			fmt.Sprintf("根据Id查询用户数据失败, err = %+v", err))
		return nil, err
	}
	return user, nil
}

// 根据账户和密码获取用户
func (dao *UserDao) GetUserByUserNameAndPassword(username string, password string) *dto.UserDTO {
	user := &po.User{}
	sql := "select * from user where user_name = ? and user_password = ?"
	ok, err := dao.runner.Get(user, sql, username, password)
	if err != nil || !ok {
		logrus.Error(
			fmt.Sprintf("根据username和password查询用户数据失败, err = %+v", err))
		return nil
	}
	if user == nil {
		// 不存在
		return nil
	}
	return user.ToDTO()
}
