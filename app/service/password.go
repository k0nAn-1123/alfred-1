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
	password := model.Password{}
	companyID := a.IsCompanyExist(r)

	password.Account = r.Account
	password.Password = r.Password
	password.Level = r.Level
	password.WebName = r.WebName
	password.Company = companyID
	password.Remark = r.Remark

	if _, err := dao.Password.Save(password); err != nil {
		return false
	}
	return true
}

func (a *passService) Modify(ctx context.Context, r *model.PasswordReq) bool {
	//passWord, err := dao.Password.FindOne("ID=?", r.ID)
	//if err != nil {
	//	glog.Error("Service modify password error. ")
	//	return false
	//} else {
	//	passWord.Account = r.Account
	//	passWord.Company = r.Company
	//	passWord.Domain = r.Domain
	//	passWord.Password = r.Password
	//	passWord.Level = r.Level
	//	passWord.WebName = r.WebName
	//	passWord.Remark = r.Remark
	//
	//	dao.Password.Save(passWord)
	//}

	return true
}

func (a *passService) IsCompanyExist(r *model.PasswordReq) int {
	company, err := dao.PwdCompany.FindOne("domain=?", r.Domain)

	if err != nil {
		glog.Error("Service IsCompanyExist Error: ", err)
		return 0
	}

	if company != nil {
		return company.ID
	} else {
		newCompany := model.PwdCompany{}
		newCompany.CompanyName = r.Company
		newCompany.Domain = r.Domain
		id, _ := dao.PwdCompany.InsertAndGetId(newCompany)
		return int(id)
	}
}