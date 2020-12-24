package do

import (
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
	"go-resk/src/dao/mapper"
	"go-resk/src/entity/dto"
	"go-resk/src/entity/service_flag"
	"go-resk/src/infra/base"
	_ "go-resk/src/test"
	"go-resk/src/utils"
	"testing"
)

func TestAccountDomain_Create(t *testing.T) {
	// TODO
}

func TestAccountDomain_GetAccount(t *testing.T) {
	// TODO
}

func TestAccountDomain_GetAccountByUserIdAndType(t *testing.T) {
	// TODO
}

func TestAccountDomain_GetAccountLog(t *testing.T) {
	// TODO
}

func TestAccountDomain_GetAccountLogByTradeNo(t *testing.T) {
	// TODO
}

func TestAccountDomain_GetEnvelopeAccountByUserId(t *testing.T) {
	// TODO
}

func TestAccountDomain_Transfer(t *testing.T) {

}

func TestAccountDomain_createAccountLog(t *testing.T) {
	// TODO
}

func TestAccountDomain_createAccountLogNo(t *testing.T) {
	// TODO
}

func TestAccountDomain_createAccountNo(t *testing.T) {
	// TODO
}

func Test_transferTo(t *testing.T) {
	Convey("Test_transferTo", t, func() {
		domain := new(AccountDomain)
		// 创建账户1
		accountDto1 := dto.AccountDTO{
			AccountName:  "转账测试账户1",
			AccountType:  0,
			CurrencyCode: service_flag.CNY,
			UserId:       utils.GetUUID(),
			Username:     "转账测试用户1",
			Balance:      decimal.NewFromFloat(10),
			Status:       0,
		}
		// 创建账户2
		accountDto2 := dto.AccountDTO{
			AccountName:  "转账测试账户2",
			AccountType:  0,
			CurrencyCode: service_flag.CNY,
			UserId:       utils.GetUUID(),
			Username:     "转账测试用户2",
			Balance:      decimal.NewFromFloat(10),
			Status:       0,
		}
		accountDtoRes1, err := domain.Create(accountDto1)
		accountDtoRes2, err := domain.Create(accountDto2)
		So(err, ShouldBeNil)
		// 开始转账测试
		transferDTO := dto.AccountTransferDTO{
			TradeNo: utils.GetUUID(),
			TradeBody: dto.TradeParticipator{
				AccountNo: accountDtoRes1.AccountNo,
				UserId:    accountDtoRes1.UserId,
				Username:  accountDtoRes1.Username,
			},
			TradeTarget: dto.TradeParticipator{
				AccountNo: accountDtoRes2.AccountNo,
				UserId:    accountDtoRes2.UserId,
				Username:  accountDtoRes2.Username,
			},
			AmountStr:  accountDto1.Balance.String(),
			Amount:     accountDto1.Balance,
			ChangeType: service_flag.CHANGE_TYPE_OUTCOME,
			ChangeFlag: service_flag.CHANGE_FLAG_OUTCOME,
			Decs:       "转账测试",
		}
		amount := accountDto1.Balance.Mul(decimal.NewFromFloat(-1))
		err = base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
			accountDao := mapper.NewAccountDao(runner)
			accountLogDao := mapper.NewAccountLogDao(runner)
			// 扣减付款账户余额并生成支出流水
			err := transferTo(accountDao, accountLogDao, domain, transferDTO, amount)
			if err != nil {
				return err
			}
			Convey("账户1余额校验", func() {
				account1 := accountDao.GetOne(transferDTO.TradeBody.AccountNo)
				So(account1, ShouldNotBeNil)
				So(account1.Balance, ShouldEqual, decimal.NewFromFloat(0))
			})
			Convey("账户1流水余额校验", func() {
				accountLog1 := accountLogDao.GetOne(domain.accountLog.LogNo)
				So(accountLog1, ShouldNotBeNil)
				So(accountLog1.Balance, ShouldEqual, accountDtoRes1.Balance)
				So(accountLog1.Amount, ShouldEqual, accountDtoRes1.Balance)
				So(accountLog1.AccountNo, ShouldEqual, accountDtoRes1.AccountNo)
				So(accountLog1.TargetAccountNo, ShouldEqual, accountDtoRes2.AccountNo)
				So(accountLog1.ChangeFlag, ShouldEqual, transferDTO.ChangeFlag)
			})
			return nil
		})
		So(err, ShouldBeNil)
	})
}

func Test_transferFrom(t *testing.T) {
	Convey("Test_transferTo", t, func() {
		domain := new(AccountDomain)
		// 创建账户1
		accountDto1 := dto.AccountDTO{
			AccountName:  "转账测试账户1",
			AccountType:  0,
			CurrencyCode: service_flag.CNY,
			UserId:       utils.GetUUID(),
			Username:     "转账测试用户1",
			Balance:      decimal.NewFromFloat(10),
			Status:       0,
		}
		// 创建账户2
		accountDto2 := dto.AccountDTO{
			AccountName:  "转账测试账户2",
			AccountType:  0,
			CurrencyCode: service_flag.CNY,
			UserId:       utils.GetUUID(),
			Username:     "转账测试用户2",
			Balance:      decimal.NewFromFloat(10),
			Status:       0,
		}
		accountDtoRes1, err := domain.Create(accountDto1)
		accountDtoRes2, err := domain.Create(accountDto2)
		So(err, ShouldBeNil)
		// 开始转账测试
		transferDTO := dto.AccountTransferDTO{
			TradeNo: utils.GetUUID(),
			TradeBody: dto.TradeParticipator{
				AccountNo: accountDtoRes1.AccountNo,
				UserId:    accountDtoRes1.UserId,
				Username:  accountDtoRes1.Username,
			},
			TradeTarget: dto.TradeParticipator{
				AccountNo: accountDtoRes2.AccountNo,
				UserId:    accountDtoRes2.UserId,
				Username:  accountDtoRes2.Username,
			},
			AmountStr:  accountDto1.Balance.String(),
			Amount:     accountDto1.Balance,
			ChangeType: service_flag.CHANGE_TYPE_OUTCOME,
			ChangeFlag: service_flag.CHANGE_FLAG_OUTCOME,
			Decs:       "转账测试",
		}
		amount := accountDto1.Balance.Mul(decimal.NewFromFloat(-1))
		err = base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
			accountDao := mapper.NewAccountDao(runner)
			accountLogDao := mapper.NewAccountLogDao(runner)
			// 扣减付款账户余额并生成支出流水
			err := transferFrom(accountDao, accountLogDao, domain, transferDTO, amount)
			if err != nil {
				return err
			}
			Convey("账户2余额校验", func() {
				account2 := accountDao.GetOne(accountDtoRes2.AccountNo)
				So(account2, ShouldNotBeNil)
				So(account2.Balance, ShouldEqual, accountDto2.Balance.Add(accountDto1.Balance))
			})
			Convey("账户2流水余额校验", func() {
				accountLog1 := accountLogDao.GetOne(domain.accountLog.LogNo)
				So(accountLog1, ShouldNotBeNil)
				So(accountLog1.Balance, ShouldEqual, accountDtoRes2.Balance)
				So(accountLog1.Amount, ShouldEqual, accountDtoRes2.Balance)
				So(accountLog1.AccountNo, ShouldEqual, accountDtoRes2.AccountNo)
				So(accountLog1.TargetAccountNo, ShouldEqual, accountDtoRes1.AccountNo)
				So(accountLog1.ChangeFlag, ShouldEqual, service_flag.CHANGE_FLAG_INCOME)
			})
			return nil
		})
		So(err, ShouldBeNil)
	})
}
