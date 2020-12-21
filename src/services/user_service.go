package services

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"go-resk/src/dao/mapper"
	"go-resk/src/entity/dto"
	"go-resk/src/entity/po"
	"go-resk/src/entity/service_flag"
	"go-resk/src/infra/base"
	"go-resk/src/utils"
)

type IUserService interface {
	// 用户创建注册
	CreateUser(userCreateDTO *dto.UserCreateDTO) (service_flag.UserStatus, error)
	// 修改用户信息
	UpdateUser(userUpdateDTO *dto.UserUpdateDTO) (service_flag.UserStatus, error)
	// 判断用户是否已注册
	CheckUserExist(userUpdateDTO *dto.UserCreateDTO) (bool, error)
	// 用户登陆校验
	UserLogin(userUpdateDTO *dto.UserCreateDTO) (*dto.UserDTO, error)
}

type userService struct {
}

var service IUserService

func GetUserService() IUserService {
	return service
}

var _ IUserService = new(userService)

// 创建用户
func (service userService) CreateUser(userCreateDTO *dto.UserCreateDTO) (status service_flag.UserStatus, err error) {
	// 参数校验
	err = utils.StructValidate(userCreateDTO)
	if err != nil {
		logrus.Error(fmt.Sprintf("CreateUser参数校验失败, userCreateDTO = %+v", userCreateDTO))
		return service_flag.USER_CREATE_FAILURE, err
	}
	err = base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		userDao := mapper.NewUserDao(runner)
		// 判断当前用户是否存在
		userRS, err := userDao.GetUserByUsername(userCreateDTO.UserName)
		if err != nil {
			logrus.Error(
				fmt.Sprintf("获取用户出现错误, err = %+v", err))
			status = service_flag.USER_CREATE_FAILURE
			return err
		}
		if userRS != nil {
			// 用户存在
			status = service_flag.USER_EXIST
			return nil
		}
		// 创建用户
		newUser := &po.User{
			UserId:       utils.GetUUID(),
			UserName:     userCreateDTO.UserName,
			UserPassword: userCreateDTO.UserPassword,
		}
		rows, err := userDao.Insert(newUser)
		if rows <= 0 {
			logrus.Error(
				fmt.Sprintf("获取用户出现错误, err = %+v", err))
			status = service_flag.USER_CREATE_FAILURE
			return nil
		}
		status = service_flag.USER_CREATE_SUCCESS
		return nil
	})
	return status, err
}

// 更新用户
func (service userService) UpdateUser(userUpdateDTO *dto.UserUpdateDTO) (status service_flag.UserStatus, err error) {
	err = utils.StructValidate(userUpdateDTO)
	if err != nil {
		logrus.Error(fmt.Sprintf("UpdateUser参数校验失败, userUpdateDTO = %+v", userUpdateDTO))
		return service_flag.USER_UPDATE_FAILURE, err
	}
	err = base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		userDao := mapper.NewUserDao(runner)
		user := &po.User{
			UserId:       userUpdateDTO.UserId,
			UserName:     userUpdateDTO.UserName,
			UserPassword: userUpdateDTO.UserPassword,
		}
		rows, err := userDao.UpdateUser(user)
		if err != nil || rows == 0 {
			logrus.Error(
				fmt.Sprintf("更新用户数据发生错误, user = %+v", user))
			status = service_flag.USER_UPDATE_FAILURE
			return err
		}
		status = service_flag.USER_UPDATE_SUCCESS
		return nil
	})
	return status, err
}

// 检查用户是否存在
func (service userService) CheckUserExist(userUpdateDTO *dto.UserCreateDTO) (status bool, err error) {
	username := userUpdateDTO.UserName
	if len(username) == 0 {
		logrus.Error("username为空")
		return false, errors.New("username为空")
	}
	err = base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		userDao := mapper.NewUserDao(runner)
		user, err := userDao.GetUserByUsername(username)
		if err != nil {
			logrus.Error(
				fmt.Sprintf("根据username获取用户发生错误, username = %+v", username))
			return err
		}
		if user != nil {
			// 用户存在
			status = true
			return nil
		}
		// 用户不存在
		status = false
		return nil
	})
	return status, err
}

// 用户登录校验
func (service userService) UserLogin(userUpdateDTO *dto.UserCreateDTO) (userDTO *dto.UserDTO, err error) {
	err = utils.StructValidate(userUpdateDTO)
	if err != nil {
		logrus.Error(fmt.Sprintf("UserLogin参数校验失败, userUpdateDTO = %+v", userUpdateDTO))
		return nil, err
	}
	err = base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
		userDao := mapper.NewUserDao(runner)
		userDTO = userDao.GetUserByUserNameAndPassword(userUpdateDTO.UserName, userUpdateDTO.UserPassword)
		return nil
	})
	return userDTO, err
}

func init() {
	service = new(userService)
}
