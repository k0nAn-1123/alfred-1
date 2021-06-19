package service

import (
	"alfred/app/dao"
	"alfred/app/model"
	"context"
	"github.com/gogf/gf/os/glog"
)

type passService struct {}

var PassService = passService{}

func (a *passService) Add(ctx context.Context, r *model.PasswordReq) bool {
	var (
		password *model.Password
	)
	password.Account = r.Account
	password.Password = r.Password
	password.WebName = r.WebName
	password.Level = r.Level
	password.Domain = r.Domain
	password.Company = r.Company

	if _, err := dao.Bodyinfo.Save(password); err != nil {
		return false
	}
	return true
}

func (a *passService) Modify(ctx context.Context, r *model.PasswordReq) bool {
	if passWord, err := dao.Password.FindOne("ID=?", r.ID); err != nil {
		glog.Error("Service modify password error. ")
		return false
	} else {
		passWord.Account = r.Account
		passWord.Company = r.Company
		passWord.Domain = r.Domain
		passWord.Password = r.Password
		passWord.Level = r.Level
		passWord.WebName = r.WebName
		passWord.Remark = r.Remark

		dao.Password.Save(passWord)
	}

	return true
}