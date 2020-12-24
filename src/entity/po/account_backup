package po

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"go-resk/src/entity/dto"
	"go-resk/src/entity/service_flag"
	"time"
)

// account表对应po
type Account struct {
	Id           int64                     `db:"id,omitempty"`         //账户ID
	AccountNo    string                    `db:"account_no,uni"`       //账户编号,账户唯一标识
	AccountName  string                    `db:"account_name"`         //账户名称,用来说明账户的简短描述,账户对应的名称或者命名，比如xxx积分、xxx零钱
	AccountType  int                       `db:"account_type"`         //账户类型，用来区分不同类型的账户：积分账户、会员卡账户、钱包账户、红包账户
	CurrencyCode service_flag.CurrencyCode `db:"currency_code"`        //货币类型编码：CNY人民币，EUR欧元，USD美元 。。。
	UserId       string                    `db:"user_id"`              //用户编号, 账户所属用户
	Username     sql.NullString            `db:"username"`             //用户名称
	Balance      decimal.Decimal           `db:"balance"`              //账户可用余额
	Status       int                       `db:"status"`               //账户状态，账户状态：0账户初始化，1启用，2停用
	CreatedAt    time.Time                 `db:"created_at,omitempty"` //创建时间
	UpdatedAt    time.Time                 `db:"updated_at,omitempty"` //更新时间
}

//,omitempty
func (po *Account) ToDTO() *dto.AccountDTO {
	accountDTO := &dto.AccountDTO{}
	accountDTO.AccountNo = po.AccountNo
	accountDTO.AccountName = po.AccountName
	accountDTO.AccountType = po.AccountType
	accountDTO.CurrencyCode = po.CurrencyCode
	accountDTO.UserId = po.UserId
	accountDTO.Username = po.Username.String
	accountDTO.Balance = po.Balance
	accountDTO.Status = po.Status
	accountDTO.CreatedAt = po.CreatedAt
	accountDTO.UpdatedAt = po.UpdatedAt
	return accountDTO
}

func (po *Account) FromDTO(accountDTO *dto.AccountDTO) {
	po.AccountNo = accountDTO.AccountNo
	po.AccountName = accountDTO.AccountName
	po.AccountType = accountDTO.AccountType
	po.CurrencyCode = accountDTO.CurrencyCode
	po.UserId = accountDTO.UserId
	po.Username = sql.NullString{
		String: accountDTO.Username,
		Valid:  true,
	}
	po.Balance = accountDTO.Balance
	po.Status = accountDTO.Status
	po.CreatedAt = accountDTO.CreatedAt
	po.UpdatedAt = accountDTO.UpdatedAt
}
