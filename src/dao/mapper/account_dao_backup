package mapper

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	. "go-resk/src/entity/po"
)

type AccountDao struct {
	runner *dbx.TxRunner
}

func NewAccountDao(runner *dbx.TxRunner) *AccountDao {
	return &AccountDao{runner: runner}
}

func (dao *AccountDao) GetOne(accountNo string) *Account {
	account := &Account{}
	account.AccountNo = accountNo
	ok, err := dao.runner.GetOne(account)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return account
}

func (dao *AccountDao) GetByUserId(userId string, accountType int) *Account {
	account := &Account{}
	sqlString := "select * from account where user_id = ? and account_type = ?"
	ok, err := dao.runner.Get(account, sqlString, userId, accountType)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if !ok {
		return nil
	}
	return account
}

func (dao *AccountDao) Insert(account *Account) (id int64, err error) {
	rs, err := dao.runner.Insert(account)
	if err != nil {
		logrus.Error(err)
		return -1, err
	}
	return rs.LastInsertId()
}

func (dao *AccountDao) Update(account *Account) (id int64, err error) {
	if account.Id <= 0 {
		logrus.Warn("account id can't small than 0")
		return -1, err
	}
	rs, err := dao.runner.Update(account)
	if err != nil {
		logrus.Error(err)
		return -1, err
	}
	return rs.LastInsertId()
}

//账户余额的更新
//amount 如果是负数，就是扣减；如果是正数，就是增加
func (dao *AccountDao) UpdateBalance(
	accountNo string,
	amount decimal.Decimal) (rows int64, err error) {
	sql := "update account " +
		" set balance=balance+CAST(? AS DECIMAL(30,6))" +
		" where account_no=? " +
		" and balance>=-1*CAST(? AS DECIMAL(30,6)) "
	rs, err := dao.runner.Exec(sql,
		amount.String(),
		accountNo,
		amount.String())
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

//账户状态更新
func (dao *AccountDao) UpdateStatus(
	accountNo string,
	status int) (rows int64, err error) {
	sql := "update account set status=? " +
		" where account_no=? "
	rs, err := dao.runner.Exec(sql, status, accountNo)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()

}
