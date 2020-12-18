package services

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go-resk/src/entity/do"
	. "go-resk/src/entity/dto"
	"go-resk/src/entity/po"
	. "go-resk/src/entity/service_flag"
	"go-resk/src/utils"
	"sync"
	"time"
)

// IAccountService全局唯一暴露点
var accountService IAccountService

func GetAccountService() IAccountService {
	return accountService
}

type IAccountService interface {
	CreateAccount(accountCreateDto AccountCreatedDTO) (*AccountDTO, error)
	Transfer(accountTransferDTO AccountTransferDTO) (TransferedStatus, error)
	StoreValue(accountTransferDTO AccountTransferDTO) (TransferedStatus, error)
	GetEnvelopeAccountByUserId(userId string) *AccountDTO
	GetAccount(accountNo string) *AccountDTO
}

type AccountService struct{}

var _ IAccountService = new(AccountService)

// 账户创建
func (service AccountService) CreateAccount(accountCreateDto AccountCreatedDTO) (*AccountDTO, error) {
	err := utils.StructValidate(accountCreateDto)
	if err != nil {
		return nil, err
	}
	domain := new(do.AccountDomain)
	// 验证账户是否存在和幂等性
	acc := domain.GetAccountByUserIdAndType(accountCreateDto.UserId, AccountType(accountCreateDto.AccountType))
	if acc != nil {
		return acc, errors.New(
			fmt.Sprintf("用户的该类型账户已经存在：username=%s[%s],账户类型=%d",
				accountCreateDto.Username, accountCreateDto.UserId, accountCreateDto.AccountType))

	}
	// 执行账户创建的业务逻辑
	amount, err := decimal.NewFromString(accountCreateDto.Amount)
	if err != nil {
		return nil, err
	}
	accountDTO := AccountDTO{
		UserId:       accountCreateDto.UserId,
		Username:     accountCreateDto.Username,
		AccountType:  accountCreateDto.AccountType,
		AccountName:  accountCreateDto.AccountName,
		CurrencyCode: accountCreateDto.CurrencyCode,
		Status:       1,
		Balance:      amount,
	}
	dto, err := domain.Create(accountDTO)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return dto, err
}

// 将AccountCreatedDTO转化为AccountPO
func FromAccountCreatedDTO2AccountPO(dto AccountCreatedDTO) (*po.Account, error) {
	amount, err := decimal.NewFromString(dto.Amount)
	if err != nil {
		return nil, err
	}
	account := &po.Account{
		AccountNo:    utils.GetUUID(),
		AccountName:  dto.AccountName,
		AccountType:  0,
		CurrencyCode: CNY,
		UserId:       dto.UserId,
		Username:     sql.NullString{dto.Username, true},
		Balance:      amount,
		Status:       0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return account, nil
}

func (service AccountService) Transfer(accountTransferDTO AccountTransferDTO) (TransferedStatus, error) {
	panic("implement me")
}

func (service AccountService) StoreValue(accountTransferDTO AccountTransferDTO) (TransferedStatus, error) {
	panic("implement me")
}

func (service AccountService) GetEnvelopeAccountByUserId(userId string) *AccountDTO {
	panic("implement me")
}

func (service AccountService) GetAccount(accountNo string) *AccountDTO {
	panic("implement me")
}

var once sync.Once

// 初始化IAccountService
func init() {
	once.Do(func() {
		accountService = new(AccountService)
	})
}
