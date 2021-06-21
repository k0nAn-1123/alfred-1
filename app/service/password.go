package service

import (
	"alfred/app/dao"
	"alfred/app/model"
	"context"
	"github.com/gogf/gf/os/glog"
)

type passService struct{}

var PassService = passService{}

func (a *passService) Add(ctx context.Context, r *model.PasswordReq) bool {
	passWord := model.Password{}
	companyID := a.IsCompanyExist(r)
	if companyID == 0 {
		glog.Error("Service add password check company err.")
		return false
	}

	passWord.Account = r.Account
	passWord.Password = r.Password
	passWord.Level = r.Level
	passWord.WebName = r.WebName
	passWord.Company = companyID
	passWord.Remark = r.Remark

	if _, err := dao.Password.Save(passWord); err != nil {
		glog.Error("Service add password database error. Message: ", err)
		return false
	}
	return true
}

func (a *passService) Modify(ctx context.Context, r *model.PasswordReq) bool {
	passWord, err := dao.Password.FindOne("ID=?", r.ID)
	if err != nil {
		glog.Error("Service modify password error. ")
		return false
	}
	if passWord == nil {
		glog.Error("Service modify password error. The data related to this ID doesn't exist. ID: ", r.ID)
		return false
	} else {
		companyID := a.IsCompanyExist(r)
		if companyID == 0 {
			glog.Error("Service modify password check company err.")
			return false
		}

		passWord.Account = r.Account
		passWord.Company = companyID
		passWord.Password = r.Password
		passWord.Level = r.Level
		passWord.WebName = r.WebName
		passWord.Remark = r.Remark

		if _, err := dao.Password.Save(passWord); err != nil {
			glog.Error("Service add password database error. Message: ", err)
			return false
		}
	}

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

func (a *passService) GetPwdList(r *model.PwdListReq) *model.PwdListRes {
	var (
		res *model.PwdListRes
	)

	if r.Content != "" {

	} else {

	}
	return res
}

func (a *passService) GetPwd(r *model.PwdReq) *model.PwdRes {
	var (
		res *model.PwdRes
	)

	pwd, err := dao.Password.FindOne(r.ID)
	if err != nil {
		glog.Error("Service Search password error. Message: ", err)
		return nil
	}

	company, _ := dao.PwdCompany.FindOne(pwd.Company)

	password := model.PwdRes{}
	password.Id = pwd.ID
	password.Company = company.CompanyName
	password.Domain = company.Domain
	password.Account = pwd.Account
	password.AccountInfo = pwd.WebName

	res = &password
	return res
}
