package utils

import (
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"go-resk/src/infra/base"
	"math/rand"
	"strings"
	"time"
)

// 红包最少为1分钱
var min int64 = 1

// 二倍均值，红包算法
// count 红包剩余个数
// amount 红包剩余金额(以分为单位)
func DoubleAverage(count, amount int64) int64 {
	// 只剩最后一个红包，直接返回
	if count == 1 {
		return amount
	}
	// 最大可用金额
	max := amount - count*min
	// 可用金额平均数
	avg := max / count
	// 两倍平均数即为红包上限
	avg2 := 2 * avg
	// 随机取 [0, avg2] + min 作为红包金额
	rand.Seed(time.Now().UnixNano())
	x := rand.Int63n(avg2) + min
	return x
}

func StructValidate(object interface{}) error {
	err := base.Validate().Struct(object)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			logrus.Error("验证错误", err)
		}
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, e := range errs {
				logrus.Error(e.Translate(base.Transtate()))
			}
		}
	}
	return err
}

func GetUUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
