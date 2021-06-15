package api

import (
	"alfred/app/model"
	"alfred/app/service"
	"alfred/library/response"
	"github.com/gogf/gf/net/ghttp"
)
var Password = passApi {}

type passApi struct {}

func (a *passApi) Add(r *ghttp.Request) {
	var (
		req *model.PasswordReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 500, err.Error())
	}
	if result := service.PassService.Add(r.Context(), req); !result {
		response.JsonExit(r, 500, "添加失败")
	}
	response.JsonExit(r, 200, "添加成功")
}

func (a *passApi) Modify(r *ghttp.Request) {
	var (
		req *model.PasswordReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 500, err.Error())
	}
	if result := service.PassService.Modify(r.Context(), req); !result {
		response.JsonExit(r, 500, "添加失败")
	}
	response.JsonExit(r, 200, "添加成功")
}

func (a *passApi) GetList(r *ghttp.Request) {

}

func (a *passApi) Remove(r *ghttp.Request) {

}

func (a *passApi) Search(r *ghttp.Request) {

}

func (a *passApi) Get(r *ghttp.Request) {

}