package service

import (
	"alfred/app/dao"
	"alfred/app/model"
	"context"
	"errors"
)

type authService struct {}

var AuthService = authService{}

// 登录
func (a *authService) Login(ctx context.Context, r *model.AuthReq) (uint, error) {
	user, err := dao.User.FindOne("Identification=? and Password=?", r.Identification, r.Password)
	if err != nil {
		return 0, err
	}

	if user == nil {
		return 0, errors.New("There is no such a person in my database.")
	} else if user.Permission != 1 {
		err := "Sorry, " + user.NickName + " only my master can login to my system."
		return 0, errors.New(err)
	}

	Context.SetUser(ctx, &model.ContextUser{
		Id:       		user.ID,
		Identification: user.Identification,
		Password: 		user.Password,
		Nickname: 		user.NickName,
		Permission:		user.Permission,
	})
	return user.ID, nil
}

// 登出
func (a *authService) Logout(ctx context.Context, r *model.AuthReq) error {
	return Session.RemoveUser(ctx)
}
