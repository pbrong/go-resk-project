package do

import (
	"errors"
	"fmt"
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
	var rdto *dto.AccountDTO
	err := base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(runner)
		accountLogDao := mapper.NewAccountLogDao(runner)
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

// 转账，扣减付款者的余额、增加收款者的余额，并记录流水付款及收款流水
func (domain *AccountDomain) Transfer(transferDTO dto.AccountTransferDTO) (status service_flag.TransferedStatus, err error) {
	if transferDTO.ChangeFlag != service_flag.CHANGE_FLAG_OUTCOME {
		// 变化标示不为支出类型
		return service_flag.TRANSFER_STATUS_FAILURE, errors.New(
			fmt.Sprintf("转账时ChangeFlag需要为CHANGE_FLAG_OUTCOME, From:%+v, To:%+v",
				transferDTO.TradeBody, transferDTO.TradeTarget))
	}
	if transferDTO.Amount.IsZero() {
		return service_flag.TRANSFER_STATUS_FAILURE, errors.New(
			fmt.Sprintf("转账时交易金额不可为0, From:%+v, To:%+v",
				transferDTO.TradeBody, transferDTO.TradeTarget))
	}
	amount := transferDTO.Amount
	// 修正交易金额为负数
	if amount.GreaterThan(decimal.NewFromFloat(0)) {
		amount = amount.Mul(decimal.NewFromFloat(-1))
	}
	// 检查余额是否足够和更新余额：通过乐观锁来验证，更新余额的同时来验证余额是否足够
	// 更新成功后，写入流水记录
	err = base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(runner)
		accountLogDao := mapper.NewAccountLogDao(runner)
		// 扣减付款账户余额并生成支出流水
		err = transferTo(accountDao, accountLogDao, domain, transferDTO, amount)
		if err != nil {
			status = service_flag.TRANSFER_STATUS_SUCCESS
			return err
		}
		// 新增收款账户余额并生成收入流水
		err = transferFrom(accountDao, accountLogDao, domain, transferDTO, amount)
		if err != nil {
			status = service_flag.TRANSFER_STATUS_SUCCESS
			return err
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

// 账户2进行收款，新增账户2金额，并记录账户2收款流水
func transferFrom(accountDao *mapper.AccountDao, accountLogDao *mapper.AccountLogDao, domain *AccountDomain, transferDTO dto.AccountTransferDTO, amount decimal.Decimal) error {
	// 在收款场景中，发起对象为收款者，状态为收入
	transferDTO.ChangeFlag = service_flag.CHANGE_FLAG_INCOME
	transferDTO.ChangeType = service_flag.CHANGE_TYPE_INCOME
	transferDTO.Decs = "收款"
	transferDTO.TradeBody, transferDTO.TradeTarget =
		transferDTO.TradeTarget, transferDTO.TradeBody
	// 判断账户是否存在
	account := accountDao.GetOne(transferDTO.TradeBody.AccountNo)
	if account == nil {
		return errors.New("账户不存在")
	}
	// 新增金额
	if amount.LessThan(decimal.NewFromFloat(0)) {
		amount = amount.Mul(decimal.NewFromFloat(-1))
	}
	rows, err := accountDao.UpdateBalance(transferDTO.TradeBody.AccountNo, amount)
	if err != nil {
		return err
	}
	if rows <= 0 {
		return errors.New("增加金额错误")
	}
	// 创建账户流水记录
	domain.accountLog = AccountLog{}
	domain.accountLog.FromTransferDTO(&transferDTO)
	domain.createAccountLogNo()
	domain.account = *account
	domain.accountLog.Balance = domain.account.Balance
	// 新增收款流水
	id, err := accountLogDao.Insert(&domain.accountLog)
	if err != nil || id <= 0 {
		return errors.New("账户流水创建失败")
	}
	return nil
}

// 账户1进行付款，减去账户1金额，并记录账户1付款流水
func transferTo(accountDao *mapper.AccountDao, accountLogDao *mapper.AccountLogDao, domain *AccountDomain, transferDTO dto.AccountTransferDTO, amount decimal.Decimal) error {
	// 判断账户是否存在
	account := accountDao.GetOne(transferDTO.TradeBody.AccountNo)
	if account == nil {
		return errors.New("账户不存在")
	}
	// 扣款
	rows, err := accountDao.UpdateBalance(transferDTO.TradeBody.AccountNo, amount)
	if err != nil {
		return err
	}
	if rows <= 0 && transferDTO.ChangeFlag == service_flag.CHANGE_FLAG_OUTCOME {
		return errors.New("余额不足")
	}
	// 创建账户流水记录
	domain.accountLog = AccountLog{}
	domain.accountLog.FromTransferDTO(&transferDTO)
	domain.createAccountLogNo()
	domain.account = *account
	domain.accountLog.Balance = domain.account.Balance
	// 新增扣款流水
	id, err := accountLogDao.Insert(&domain.accountLog)
	if err != nil || id <= 0 {
		return errors.New("账户流水创建失败")
	}
	return nil
}

// 根据账户编号来查询账户信息
func (domain *AccountDomain) GetAccount(accountNo string) *dto.AccountDTO {
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
func (domain *AccountDomain) GetEnvelopeAccountByUserId(userId string) *dto.AccountDTO {
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
func (domain *AccountDomain) GetAccountByUserIdAndType(userId string, accountType service_flag.AccountType) *dto.AccountDTO {
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
func (domain *AccountDomain) GetAccountLog(logNo string) *dto.AccountLogDTO {
	var log *AccountLog
	err := base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		dao := mapper.NewAccountLogDao(runner)
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
func (domain *AccountDomain) GetAccountLogByTradeNo(tradeNo string) *dto.AccountLogDTO {
	var log *AccountLog
	err := base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		dao := mapper.NewAccountLogDao(runner)
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
