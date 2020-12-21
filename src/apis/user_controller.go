package apis

import (
	"github.com/kataras/iris/context"
	. "go-resk/src/entity"
	"go-resk/src/entity/dto"
	"go-resk/src/entity/service_flag"
	"go-resk/src/infra"
	"go-resk/src/infra/base"
	"go-resk/src/services"
	"go-resk/src/utils"
)

// 将UserApi注册到WebApiContainer中
func init() {
	infra.RegisterApi(&UserApi{})
}

type UserApi struct {
}

var _ infra.WebApi = new(UserApi)

// 注册UserController的API
func (a *UserApi) Init() {
	irisServer := base.IrisServer()
	groupRouter := irisServer.Party("/v1/user")
	groupRouter.Post("/create", createUserHandler)
	groupRouter.Post("/login", loginUserHandler)
	groupRouter.Post("/exist", existUserHandler)
	groupRouter.Post("/update", updateUserHandler)
}

// 更新用户信息接口
// post body json
/**
{
	"UserId":"9f8db427c11e464c9a229c06ac8d2197",
	"UserName":"test",
	"UserPassword":"test"
}
*/
func updateUserHandler(ctx context.Context) {
	result := &AjaxRes{
		Code: ResCodeOk,
	}
	userUpdateDTO := &dto.UserUpdateDTO{}
	ctx.ReadJSON(userUpdateDTO)
	userService := services.GetUserService()
	status, err := userService.UpdateUser(userUpdateDTO)
	if err != nil {
		result.Code = ResCodeInnerServerError
		result.Message = "更新用户数据出错"
		ctx.JSON(result)
		return
	}
	if status != service_flag.USER_UPDATE_SUCCESS {
		result.Code = ResCodeOk
		result.Message = "更新用户失败"
		result.Data = nil
		ctx.JSON(result)
		return
	}
	result.Message = "更新用户数据成功"
	result.Data = userUpdateDTO
	ctx.JSON(result)
}

// 用户存在判断接口
// post body json
/**
{
	"UserName": "testUsername"
}
*/
func existUserHandler(ctx context.Context) {
	result := &AjaxRes{
		Code: ResCodeOk,
	}
	userExistDTO := &dto.UserCreateDTO{}
	ctx.ReadJSON(userExistDTO)
	userService := services.GetUserService()
	exist, err := userService.CheckUserExist(userExistDTO)
	if err != nil {
		result.Code = ResCodeInnerServerError
		result.Message = "用户校验失败"
		ctx.JSON(result)
		return
	}
	if exist {
		result.Code = ResCodeOk
		result.Message = "用户存在"
		result.Data = true
		ctx.JSON(result)
		return
	}
	result.Message = "用户不存在"
	result.Data = false
	ctx.JSON(result)
}

// 用户登陆校验接口
// post body json
/**
{
	"UserName": "testUsername",
	"UserPassword": "testUserPassword"
}
*/
func loginUserHandler(ctx context.Context) {
	result := &AjaxRes{
		Code: ResCodeOk,
	}
	userLoginDTO := &dto.UserCreateDTO{}
	ctx.ReadJSON(userLoginDTO)
	userService := services.GetUserService()
	userDTO, err := userService.UserLogin(userLoginDTO)
	if err != nil {
		result.Code = ResCodeInnerServerError
		result.Message = "用户登录失败"
		ctx.JSON(result)
		return
	}
	if userDTO == nil {
		result.Code = ResCodeOk
		result.Message = "用户不存在或密码错误"
		result.Data = nil
		ctx.JSON(result)
		return
	}
	result.Message = "用户登录成功"
	result.Data = userDTO
	ctx.JSON(result)
}

// 用户创建的接口: /v1/user/create
// POST body json
/*
{
	"UserName": "testUsername",
	"UserPassword": "testUserPassword"
}*/
func createUserHandler(ctx context.Context) {
	result := &AjaxRes{
		Code: ResCodeOk,
	}
	// 获取请求参数
	userCreateDTO := &dto.UserCreateDTO{}
	ctx.ReadJSON(userCreateDTO)
	err := utils.StructValidate(userCreateDTO)
	if err != nil {
		result.Code = ResCodeRequestParamsError
		result.Message = "参数校验错误"
		ctx.JSON(result)
		return
	}
	userService := services.GetUserService()
	status, err := userService.CreateUser(userCreateDTO)
	if err != nil || status == service_flag.USER_CREATE_FAILURE {
		result.Code = ResCodeInnerServerError
		result.Message = "业务处理异常"
		ctx.JSON(result)
		return
	}
	if status == service_flag.USER_EXIST {
		result.Code = ResCodeOk
		result.Message = "用户已存在"
		ctx.JSON(result)
		return
	}
	result.Message = "用户创建成功"
	result.Data = userCreateDTO
	ctx.JSON(result)
}
