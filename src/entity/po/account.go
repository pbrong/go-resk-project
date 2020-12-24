package po

import (
	"github.com/shopspring/decimal"
	"go-resk/src/entity/dto"
	"time"
)

type Account struct {
	Id          int64           `db:"id,pk"`
	UserId      string          `db:"user_id,uni"`
	AccountNo   string          `db:"account_no,uni"`
	Balance     decimal.Decimal `db:"balance"`
	GmtCreated  time.Time       `db:"gmt_created,omitempty"`
	GmtModified time.Time       `db:"gmt_modified,omitempty"`
}

func (a *Account) ToDTO() *dto.AccountDTO {
	accountDTO := &dto.AccountDTO{
		Id:        a.Id,
		UserId:    a.UserId,
		AccountNo: a.AccountNo,
		Balance:   a.Balance,
	}
	return accountDTO
}
