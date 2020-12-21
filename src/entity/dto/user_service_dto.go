package dto

type UserDTO struct {
	Id       int64
	UserId   string
	UserName string
}

type UserCreateDTO struct {
	UserId       string ``
	UserName     string `validate:"required"`
	UserPassword string `validate:"required"`
}

type UserUpdateDTO struct {
	UserId       string `validate:"required"`
	UserName     string `validate:"required"`
	UserPassword string `validate:"required"`
}
