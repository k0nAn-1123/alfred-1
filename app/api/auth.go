package api

import (
	"alfred/app/model"
	"alfred/app/service"
	"alfred/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Auth = authApi{}

type authApi struct {}

func (a *authApi) Login(r *ghttp.Request) {
	var (
		req *model.AuthReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 500, err.Error())
	}
	if id, err := service.AuthService.Login(r.Context(), req); err != nil {
		response.JsonExit(r, 500, err.Error())
	} else {
		response.JsonExit(r, 200, "Welcome, master.", id)
	}
}

func (a *authApi) Logout(r *ghttp.Request) {
	var (
		req *model.AuthReq
	)
	if err := service.AuthService.Logout(r.Context(), req); err != nil {
		response.JsonExit(r, 500, err.Error())
	} else {
		response.JsonExit(r, 200, "Bye, master.")
	}

}
