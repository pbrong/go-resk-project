package mapper

import (
	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/tietang/dbx"
	"go-resk/src/entity/po"
	"go-resk/src/infra/base"
	"go-resk/src/utils"
	"reflect"
	"testing"
)

func TestAccountDao_DeleteByAccountNo(t *testing.T) {
	type fields struct {
		runner *dbx.TxRunner
	}
	type args struct {
		accountNo string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantAffectd int64
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountDao{
				runner: tt.fields.runner,
			}
			gotAffectd, err := dao.DeleteByAccountNo(tt.args.accountNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteByAccountNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAffectd != tt.wantAffectd {
				t.Errorf("DeleteByAccountNo() gotAffectd = %v, want %v", gotAffectd, tt.wantAffectd)
			}
		})
	}
}

func TestAccountDao_DeleteByUserId(t *testing.T) {
	type fields struct {
		runner *dbx.TxRunner
	}
	type args struct {
		userId string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantAffectd int64
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountDao{
				runner: tt.fields.runner,
			}
			gotAffectd, err := dao.DeleteByUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAffectd != tt.wantAffectd {
				t.Errorf("DeleteByUserId() gotAffectd = %v, want %v", gotAffectd, tt.wantAffectd)
			}
		})
	}
}

func TestAccountDao_GetAccountByAccountNo(t *testing.T) {
	type fields struct {
		runner *dbx.TxRunner
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *po.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountDao{
				runner: tt.fields.runner,
			}
			got, err := dao.GetAccountByAccountNo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountByAccountNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountByAccountNo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountDao_GetAccountByUserId(t *testing.T) {
	type fields struct {
		runner *dbx.TxRunner
	}
	type args struct {
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *po.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountDao{
				runner: tt.fields.runner,
			}
			got, err := dao.GetAccountByUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountByUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountDao_Insert(t *testing.T) {
	Convey("TestAccountDao_Insert", t, func() {
		base.DbxDatabase().Tx(func(runner *dbx.TxRunner) error {
			accountDao := NewAccountDao(runner)
			account := &po.Account{
				UserId:    utils.GetUUID(),
				AccountNo: utils.GetUUID(),
				Balance:   decimal.NewFromFloat(100),
			}
			id, err := accountDao.Insert(account)
			So(err, ShouldBeNil)
			So(id, ShouldBeGreaterThan, 0)
			// 回查测试
			accountRS, err := accountDao.GetAccountById(id)
			So(accountRS, ShouldNotBeNil)
			So(accountRS.UserId, ShouldEqual, account.UserId)
			So(accountRS.AccountNo, ShouldEqual, account.AccountNo)
			So(accountRS.Balance, ShouldEqual, account.Balance)
			runner.Rollback()
			return err
		})
	})
}

func TestAccountDao_Update(t *testing.T) {
	type fields struct {
		runner *dbx.TxRunner
	}
	type args struct {
		account *po.Account
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantAffectd int64
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := &AccountDao{
				runner: tt.fields.runner,
			}
			gotAffectd, err := dao.Update(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAffectd != tt.wantAffectd {
				t.Errorf("Update() gotAffectd = %v, want %v", gotAffectd, tt.wantAffectd)
			}
		})
	}
}

func TestNewAccountDao(t *testing.T) {
	Convey("TestNewAccountDao", t, func() {
		accountDao := NewAccountDao(nil)
		So(accountDao, ShouldNotBeNil)
		So(accountDao.runner, ShouldBeNil)
	})
}
