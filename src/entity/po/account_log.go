package po

import (
	"github.com/shopspring/decimal"
	"go-resk/src/entity/dto"
	"go-resk/src/entity/service_flag"
	"time"
)

// account_log表对应po
type AccountLog struct {
	Id              int64                   `db:"id,omitempty"`         //自增ID
	LogNo           string                  `db:"log_no,uni"`           //流水编号 全局不重复字符或数字，唯一性标识
	TradeNo         string                  `db:"trade_no"`             //交易单号 全局不重复字符或数字，唯一性标识
	AccountNo       string                  `db:"account_no"`           //账户编号 账户ID
	UserId          string                  `db:"user_id"`              //用户编号
	Username        string                  `db:"username"`             //用户名称
	TargetAccountNo string                  `db:"target_account_no"`    //账户编号 账户ID
	TargetUserId    string                  `db:"target_user_id"`       //目标用户编号
	TargetUsername  string                  `db:"target_username"`      //目标用户名称
	Amount          decimal.Decimal         `db:"amount"`               //交易金额,该交易涉及的金额
	Balance         decimal.Decimal         `db:"balance"`              //交易后余额,该交易后的余额
	ChangeType      service_flag.ChangeType `db:"change_type"`          //流水交易类型，0 创建账户，>0 为收入类型，<0 为支出类型，自定义
	ChangeFlag      service_flag.ChangeFlag `db:"change_flag"`          //交易变化标识：-1 出账 1为进账，枚举
	Status          int                     `db:"status"`               //交易状态：
	Decs            string                  `db:"decs"`                 //交易描述
	CreatedAt       time.Time               `db:"created_at,omitempty"` //创建时间
}

func (po *AccountLog) ToDTO() *dto.AccountLogDTO {
	accountLogDTO := &dto.AccountLogDTO{
		TradeNo:         po.TradeNo,
		LogNo:           po.LogNo,
		AccountNo:       po.AccountNo,
		TargetAccountNo: po.TargetAccountNo,
		UserId:          po.UserId,
		Username:        po.Username,
		TargetUserId:    po.TargetUserId,
		TargetUsername:  po.TargetUsername,
		Amount:          po.Amount,
		Balance:         po.Balance,
		ChangeType:      po.ChangeType,
		ChangeFlag:      po.ChangeFlag,
		Status:          po.Status,
		Decs:            po.Decs,
		CreatedAt:       po.CreatedAt,
	}
	return accountLogDTO
}

func (po *AccountLog) FromDTO(accountLogDTO *dto.AccountLogDTO) {

	po.TradeNo = accountLogDTO.TradeNo
	po.LogNo = accountLogDTO.LogNo
	po.AccountNo = accountLogDTO.AccountNo
	po.TargetAccountNo = accountLogDTO.TargetAccountNo
	po.UserId = accountLogDTO.UserId
	po.Username = accountLogDTO.Username
	po.TargetUserId = accountLogDTO.TargetUserId
	po.TargetUsername = accountLogDTO.TargetUsername
	po.Amount = accountLogDTO.Amount
	po.Balance = accountLogDTO.Balance
	po.ChangeType = accountLogDTO.ChangeType
	po.ChangeFlag = accountLogDTO.ChangeFlag
	po.Status = accountLogDTO.Status
	po.Decs = accountLogDTO.Decs
	po.CreatedAt = accountLogDTO.CreatedAt
}

func (po *AccountLog) FromTransferDTO(transferDTO *dto.AccountTransferDTO) {
	po.TradeNo = transferDTO.TradeNo
	po.AccountNo = transferDTO.TradeBody.AccountNo
	po.TargetAccountNo = transferDTO.TradeTarget.AccountNo
	po.UserId = transferDTO.TradeBody.UserId
	po.Username = transferDTO.TradeBody.Username
	po.TargetUserId = transferDTO.TradeTarget.UserId
	po.TargetUsername = transferDTO.TradeTarget.Username
	po.Amount = transferDTO.Amount
	po.ChangeType = transferDTO.ChangeType
	po.ChangeFlag = transferDTO.ChangeFlag
	po.Decs = transferDTO.Decs
}
