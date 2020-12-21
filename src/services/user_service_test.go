package services

import (
	. "github.com/smartystreets/goconvey/convey"
	"go-resk/src/entity/dto"
	"go-resk/src/entity/service_flag"
	_ "go-resk/src/test" // 引入该包以在启动时加载Starter
	"go-resk/src/utils"
	"testing"
)

func TestGetUserService(t *testing.T) {
	Convey("TestGetUserService", t, func() {
		userService1 := GetUserService()
		userService2 := GetUserService()
		So(userService1, ShouldNotBeNil)
		So(userService1, ShouldEqual, userService2)
	})
}

func Test_userService_CreateUser(t *testing.T) {
	Convey("Test_userService_CreateUser", t, func() {
		userService := GetUserService()
		createDTO := &dto.UserCreateDTO{
			UserId:       "",
			UserName:     utils.GetUUID(),
			UserPassword: "Test_userService_CreateUser",
		}
		status, err := userService.CreateUser(createDTO)
		So(err, ShouldBeNil)
		So(status, ShouldEqual, service_flag.USER_CREATE_SUCCESS)
	})
	Convey("Test_userService_CreateUser_Params_Lose", t, func() {
		userService := GetUserService()
		createDTO := &dto.UserCreateDTO{
			UserId:       "",
			UserName:     "",
			UserPassword: "Test_userService_CreateUser",
		}
		status, err := userService.CreateUser(createDTO)
		So(err, ShouldNotBeNil)
		So(status, ShouldEqual, service_flag.USER_CREATE_FAILURE)
	})
	Convey("Test_userService_CreateUser_User_Exist", t, func() {
		userService := GetUserService()
		createDTO := &dto.UserCreateDTO{
			UserId:       "",
			UserName:     utils.GetUUID(),
			UserPassword: "Test_userService_CreateUser",
		}
		status, err := userService.CreateUser(createDTO)
		So(err, ShouldBeNil)
		So(status, ShouldEqual, service_flag.USER_CREATE_SUCCESS)
		// 重复插入
		status, err = userService.CreateUser(createDTO)
		So(err, ShouldBeNil)
		So(status, ShouldEqual, service_flag.USER_EXIST)
	})
}

func Test_userService_UpdateUser(t *testing.T) {
	// TODO
}
