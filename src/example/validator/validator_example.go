package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translator "github.com/go-playground/validator/v10/translations/zh"
)

type Person struct {
	Id   int    `validate:"required,gt=0"`
	Name string `validate:"required"`
	Age  int    `validate:"gte=0,lte=130"`
}

var trans ut.Translator

func main() {
	p := &Person{
		Id:   -1,
		Name: "",
		Age:  22,
	}
	//valid := validator.New()
	//err := valid.Struct(p)
	//if err != nil {
	//	fmt.Println(err)
	//}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zh_translator.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(p)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans)) //Age必须大于18
		}
	}
}
