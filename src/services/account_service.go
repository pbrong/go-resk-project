package services

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/tietang/dbx"
	"go-resk/src/dao/mapper"
	. "go-resk/src/entity/dto"
	"go-resk/src/entity/po"
	. "go-resk/src/entity/service_flag"
	"go-resk/src/infra/base"
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

// 账户创建
func (service AccountService) CreateAccount(accountCreateDto AccountCreatedDTO) (*AccountDTO, error) {
	err := utils.StructValidate(accountCreateDto)
	if err != nil {
		return nil, err
	}
	accountPo, err := FromAccountCreatedDTO2AccountPO(accountCreateDto)
	if err != nil {
		return nil, err
	}
	err = base.DbxDatabase().Tx(func(tx *dbx.TxRunner) error {
		accountDao := mapper.NewAccountDao(tx)
		existAccount := accountDao.GetByUserId(accountCreateDto.UserId, 0)
		if existAccount != nil {
			return errors.New(
				fmt.Sprintf("can't create account because the account is existed：username=%s[%s],账户类型=%d",
					accountCreateDto.Username, accountCreateDto.UserId, accountCreateDto.AccountType))
		}
		id, err := accountDao.Insert(accountPo)
		if id <= 0 {
			return errors.New("Id small than zero")
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &AccountDTO{UserId: accountCreateDto.UserId}, nil
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
