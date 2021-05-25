package api

import (
	"alfred/app/model"
	"alfred/app/service"
	"alfred/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Auth = authApi{}

type authApi struct {}

func (a *healthApi) Auth(r *ghttp.Request) {
	var (
		req *model.AuthReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 500, err.Error())
	}
	if err := service.AuthService.Login(r.Context(), req); err != nil {
		response.JsonExit(r, 500, err.Error())
	} else {
		response.JsonExit(r, 200, "Welcome, master.")
	}
}
