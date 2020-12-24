package mapper

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"go-resk/src/entity/po"
)

type AccountDao struct {
	runner *dbx.TxRunner
}

func NewAccountDao(runner *dbx.TxRunner) *AccountDao {
	return &AccountDao{runner: runner}
}

// 新增账户
func (dao *AccountDao) Insert(account *po.Account) (id int64, err error) {
	rs, err := dao.runner.Insert(account)
	if err != nil {
		logrus.Error(
			fmt.Sprintf("创建账户失败, err = %+v", err))
		return -1, err
	}
	return rs.LastInsertId()
}

// 根据UserId删除账户
func (dao *AccountDao) DeleteByUserId(userId string) (affectd int64, error error) {
	sql := "delete from account where user_id = ?"
	rs, err := dao.runner.Exec(sql, userId)
	if err != nil {
		logrus.Error(
			fmt.Sprintf("根据UserId删除账户失败, err = %+v", err))
		return -1, err
	}
	return rs.RowsAffected()
}

// 根据AccountNo删除账户
func (dao *AccountDao) DeleteByAccountNo(accountNo string) (affectd int64, error error) {
	sql := "delete from account where account_no = ?"
	rs, err := dao.runner.Exec(sql, accountNo)
	if err != nil {
		logrus.Error(
			fmt.Sprintf("根据accountNo删除账户失败, err = %+v", err))
		return -1, err
	}
	return rs.RowsAffected()
}

// 更新账户（以Id、UserId、AccountNo之一为准）
func (dao *AccountDao) Update(account *po.Account) (affectd int64, error error) {
	rs, err := dao.runner.Update(account)
	if err != nil {
		logrus.Error(
			fmt.Sprintf("更新账户失败, err = %+v", err))
		return -1, err
	}
	return rs.RowsAffected()
}

// 根据UserId获取账户
func (dao *AccountDao) GetAccountByUserId(userId string) (*po.Account, error) {
	account := &po.Account{}
	sql := "select * from account where user_id = ?"
	ok, err := dao.runner.Get(account, sql, userId)
	if err != nil || !ok {
		logrus.Error(
			fmt.Sprintf("根据UserId查询账户数据失败, err = %+v", err))
		return nil, err
	}
	return account, nil
}

// 根据AccountNo获取账户
func (dao *AccountDao) GetAccountByAccountNo(accountNo string) (*po.Account, error) {
	account := &po.Account{}
	sql := "select * from account where account_no = ?"
	ok, err := dao.runner.Get(account, sql, accountNo)
	if err != nil || !ok {
		logrus.Error(
			fmt.Sprintf("根据AccountNo查询账户数据失败, err = %+v", err))
		return nil, err
	}
	return account, nil
}

// 根据Id获取账户
func (dao *AccountDao) GetAccountById(id int64) (*po.Account, error) {
	account := &po.Account{}
	sql := "select * from account where id = ?"
	ok, err := dao.runner.Get(account, sql, id)
	if err != nil || !ok {
		logrus.Error(
			fmt.Sprintf("根据id查询账户数据失败, err = %+v", err))
		return nil, err
	}
	return account, nil
}
