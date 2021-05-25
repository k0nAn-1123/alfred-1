package api

import (
	"alfred/app/model"
	"alfred/app/service"
	"alfred/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Health = healthApi{}

type healthApi struct {}

func (a *healthApi) GetHealthInfo(r *ghttp.Request) {
	r.Response.Writeln("Hello, master, here is your health information.")
}

func (a healthApi) AddBodyInfo(r *ghttp.Request) {
	var (
		req *model.HealthReq
	)
	if err := r.Parse(&req); err != nil {
		response.JsonExit(r, 500, err.Error())
	}
	if result := service.HealthService.AddBodyInfo(r.Context(), req); !result {
		response.JsonExit(r, 500, "添加失败")
	}
	response.JsonExit(r, 200, "添加成功")
}