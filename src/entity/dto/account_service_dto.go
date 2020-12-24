package dto

import "github.com/shopspring/decimal"

type AccountDTO struct {
	Id        int64
	UserId    string
	AccountNo string
	Balance   decimal.Decimal
}
