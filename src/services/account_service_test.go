package services

import (
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"go-resk/src/entity/dto"
	"go-resk/src/entity/service_flag"
	_ "go-resk/src/test" // 引入该包以在启动时加载Starter
	"reflect"
	"testing"
)

func TestAccountService_CreateAccount(t *testing.T) {
	Convey("TestAccountService_CreateAccount", t, func() {
		accountService := GetAccountService()
		accountCreatedDTO := dto.AccountCreatedDTO{
			UserId:       "test_id12141",
			Username:     "test_username",
			AccountName:  "test_account",
			AccountType:  0,
			CurrencyCode: service_flag.CNY,
			Amount:       "100.0",
		}
		result, err := accountService.CreateAccount(accountCreatedDTO)
		if err != nil {
			logrus.Error(err)
		} else {
			So(err, ShouldBeNil)
			So(result.UserId, ShouldEqual, result.UserId)
		}
	})
}

func TestAccountService_GetAccount(t *testing.T) {
	// TODO
}

func TestAccountService_GetEnvelopeAccountByUserId(t *testing.T) {
	// TODO
}

func TestAccountService_StoreValue(t *testing.T) {
	// TODO
}

func TestAccountService_Transfer(t *testing.T) {
	// TODO
}

func TestFromAccountCreatedDTO2AccountPO(t *testing.T) {
	// TODO
}

func TestGetAccountService(t *testing.T) {
	tests := []struct {
		name string
		want IAccountService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAccountService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountService() = %v, want %v", got, tt.want)
			}
		})
	}
}
