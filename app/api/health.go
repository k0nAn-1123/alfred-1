package api

import (
	"github.com/gogf/gf/net/ghttp"
)

var Health = healthApi{}

type healthApi struct {
}

func (a *healthApi) GetHealthInfo(r *ghttp.Request) {
	r.Response.Writeln("Hello, master, here is your health information.")
}
