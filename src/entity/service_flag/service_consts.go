package service_flag

// 流水交易类型：0 创建账户，>0 为收入类型，<0 为支出类型
type ChangeType int

const (
	CHANGE_TYPE_CREATE_ACCOUNT ChangeType = 0
	CHANGE_TYPE_INCOME         ChangeType = 1
	CHANGE_TYPE_OUTCOME        ChangeType = -1
)

// 交易变化标识：-1 出账 1为进账 0账户创建
type ChangeFlag int

const (
	CHANGE_FLAG_CREATE_ACCOUNT ChangeFlag = 0
	CHANGE_FLAG_INCOME         ChangeFlag = 1
	CHANGE_FLAG_OUTCOME        ChangeFlag = -1
)

// 转账状态: -1 失败 1 成功
type TransferedStatus int

const (
	TRANSFER_STATUS_FAILURE TransferedStatus = -1
	TRANSFER_STATUS_SUCCESS TransferedStatus = 1
)

// 交易货币
type CurrencyCode string

const (
	CNY CurrencyCode = "CNY"
	EUR CurrencyCode = "EUR"
	USD CurrencyCode = "USD"
)

//账户类型
type AccountType int

const (
	EnvelopeAccountType       AccountType = 1
	SystemEnvelopeAccountType AccountType = 2
)
