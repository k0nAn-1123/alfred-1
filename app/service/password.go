package service

import (
	"alfred/app/dao"
	"alfred/app/model"
	"context"
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
