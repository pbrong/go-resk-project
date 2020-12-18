package do

import (
	"errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"go-resk/src/dao/mapper"
	"go-resk/src/entity/dto"
	. "go-resk/src/entity/po"
	"go-resk/src/entity/service_flag"
	"go-resk/src/infra/base"
	"go-resk/src/utils"
)

// 有状态的，每次使用时都要实例化
type AccountDomain struct {
	account    Account
	accountLog AccountLog
}

// 创建logNo的逻辑
func (domain *AccountDomain) createAccountLogNo() {
	//暂时采用ksuid的ID生成策略来创建No
	//后期会优化成可读性比较好的，分布式ID
	//全局唯一的ID
	domain.accountLog.LogNo = utils.GetUUID()
}

// 生成accountNo的逻辑
func (domain *AccountDomain) createAccountNo() {
	domain.account.AccountNo = utils.GetUUID()
}

// 创建流水的记录
func (domain *AccountDomain) createAccountLog() {
	//通过account来创建流水，创建账户逻辑在前
	domain.accountLog = AccountLog{}
	domain.createAccountLogNo()
	domain.accountLog.TradeNo = domain.accountLog.LogNo
	//流水中的交易主体信息
	domain.accountLog.AccountNo = domain.account.AccountNo
	domain.accountLog.UserId = domain.account.UserId
	domain.accountLog.Username = domain.account.Username.String
	//交易对象信息
	domain.accountLog.TargetAccountNo = domain.account.AccountNo
	domain.accountLog.TargetUserId = domain.account.UserId
	domain.accountLog.TargetUsername = domain.account.Username.String

	//交易金额
	domain.accountLog.Amount = domain.account.Balance
	domain.accountLog.Balance = domain.account.Balance
	//交易变化属性
	domain.accountLog.Decs = "账户创建"
	domain.accountLog.ChangeType = service_flag.CHANGE_TYPE_CREATE_ACCOUNT
	domain.accountLog.ChangeFlag = service_flag.CHANGE_FLAG_CREATE_ACCOUNT
}

// 账户创建的业务逻辑
func (domain *AccountDomain) Create(
	accountDTO dto.AccountDTO) (*dto.AccountDTO, error) {
	//创建账户持久化对象
	domain.account = Account{}
	domain.account.FromDTO(&accountDTO)
	domain.createAccountNo()
	domain.account.Username.Valid = true
	//创建账户流水持久化对象
	domain.createAccountLog()
	accountDao := mapper.AccountDao{}
	accountLogDao := mapper.AccountLogDao{}
	var rdto *dto.AccountDTO
	err := base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(runner)
		//插入账户数据
		id, err := accountDao.Insert(&domain.account)
		if err != nil {
			return err
		}
		if id <= 0 {
			return errors.New("创建账户失败")
		}
		//如果插入成功，就插入流水数据
		id, err = accountLogDao.Insert(&domain.accountLog)
		if err != nil {
			return err
		}
		if id <= 0 {
			return errors.New("创建账户流水失败")
		}
		domain.account = *accountDao.GetOne(domain.account.AccountNo)
		return nil
	})
	rdto = domain.account.ToDTO()
	return rdto, err

}

func (a *AccountDomain) Transfer(transferDTO dto.AccountTransferDTO) (status service_flag.TransferedStatus, err error) {
	// 如果交易变化是支出，修正amount
	amount := transferDTO.Amount
	if transferDTO.ChangeFlag == service_flag.CHANGE_FLAG_OUTCOME {
		amount = amount.Mul(decimal.NewFromFloat(-1))
	}
	// 创建账户流水记录
	a.accountLog = AccountLog{}
	a.accountLog.FromTransferDTO(&transferDTO)
	a.createAccountLogNo()
	// 检查余额是否足够和更新余额：通过乐观锁来验证，更新余额的同时来验证余额是否足够
	// 更新成功后，写入流水记录
	err = base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(runner)
		accountLogDao := mapper.NewAccountLogDao(runner)
		rows, err := accountDao.UpdateBalance(transferDTO.TradeBody.AccountNo, amount)
		if err != nil {
			status = service_flag.TRANSFER_STATUS_FAILURE
			return err
		}
		if rows <= 0 && transferDTO.ChangeFlag == service_flag.CHANGE_FLAG_OUTCOME {
			status = service_flag.TRANSFER_STATUS_FAILURE
			return errors.New("余额不足")
		}
		account := accountDao.GetOne(transferDTO.TradeBody.AccountNo)
		if account == nil {
			return errors.New("账户出错")
		}
		a.account = *account
		a.accountLog.Balance = a.account.Balance
		id, err := accountLogDao.Insert(&a.accountLog)
		if err != nil || id <= 0 {
			status = service_flag.TRANSFER_STATUS_FAILURE
			return errors.New("账户流水创建失败")
		}
		return nil
	})
	if err != nil {
		logrus.Error(err)
	} else {
		status = service_flag.TRANSFER_STATUS_SUCCESS
	}

	return status, err
}

// 根据账户编号来查询账户信息
func (a *AccountDomain) GetAccount(accountNo string) *dto.AccountDTO {
	var account *Account
	err := base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(runner)
		account = accountDao.GetOne(accountNo)
		return nil
	})
	if err != nil {
		return nil
	}
	if account == nil {
		return nil
	}
	return account.ToDTO()
}

// 根据用户ID来查询红包账户信息
func (a *AccountDomain) GetEnvelopeAccountByUserId(userId string) *dto.AccountDTO {
	var account *Account
	err := base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(runner)
		account = accountDao.GetByUserId(userId, int(service_flag.CHANGE_TYPE_OUTCOME))
		return nil
	})
	if err != nil {
		return nil
	}
	if account == nil {
		return nil
	}
	return account.ToDTO()

}

// 根据用户ID和账户类型来查询账户信息
func (a *AccountDomain) GetAccountByUserIdAndType(userId string, accountType services.AccountType) *services.AccountDTO {
	var account *Account
	err := base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(runner)
		account = accountDao.GetByUserId(userId, int(accountType))
		return nil
	})
	if err != nil {
		return nil
	}
	if account == nil {
		return nil
	}
	return account.ToDTO()

}

// 根据流水ID来查询账户流水
func (a *AccountDomain) GetAccountLog(logNo string) *dto.AccountLogDTO {
	dao := AccountLogDao{}
	var log *AccountLog
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao.runner = runner
		log = dao.GetOne(logNo)
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if log == nil {
		return nil
	}
	return log.ToDTO()
}

// 根据交易编号来查询账户流水
func (a *AccountDomain) GetAccountLogByTradeNo(tradeNo string) *services.AccountLogDTO {
	dao := AccountLogDao{}
	var log *AccountLog
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao.runner = runner
		log = dao.GetByTradeNo(tradeNo)
		return nil
	})
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if log == nil {
		return nil
	}
	return log.ToDTO()
}
