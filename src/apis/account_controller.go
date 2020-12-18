package apis

import (
	"github.com/kataras/iris/context"
	"github.com/sirupsen/logrus"
	. "go-resk/src/entity"
	"go-resk/src/entity/dto"
	"go-resk/src/infra"
	"go-resk/src/infra/base"
	"go-resk/src/services"
)

// 将AccountApi注册到WebApiContainer中
func init() {
	infra.RegisterApi(&AccountApi{})
}

type AccountApi struct {
}

var _ infra.WebApi = new(AccountApi)

// 注册AccountController的API
func (a *AccountApi) Init() {
	irisServer := base.IrisServer()
	groupRouter := irisServer.Party("/v1/account")
	groupRouter.Post("/create", createHandler)
	//groupRouter.Post("/transfer", transferHandler)
	//groupRouter.Get("/envelope/get", getEnvelopeAccountHandler)
	//groupRouter.Get("/get", getAccountHandler)
}

//账户创建的接口: /v1/account/create
//POST body json
/*
{
	"UserId": "w123456",
	"Username": "测试用户1",
	"AccountName": "测试账户1",
	"AccountType": 0,
	"CurrencyCode": "CNY",
	"Amount": "100.11"
}*/
func createHandler(ctx context.Context) {
	//获取请求参数，
	account := dto.AccountCreatedDTO{}
	err := ctx.ReadJSON(&account)
	result := AjaxRes{
		Code: ResCodeOk,
	}
	if err != nil {
		result.Code = ResCodeRequestParamsError
		result.Message = err.Error()
		ctx.JSON(&result)
		logrus.Error(err)
		return
	}
	// 执行创建账户的代码
	service := services.GetAccountService()
	dto, err := service.CreateAccount(account)
	if err != nil {
		result.Code = ResCodeInnerServerError
		result.Message = err.Error()
		logrus.Error(err)
	}
	result.Data = dto
	ctx.JSON(&result)
}
