package router

import (
	"alfred/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		//group.Middleware(
		//	service.MiddleWare.Ctx,
		//	service.MiddleWare.CORS,
		//)
		group.POST("/login", api.Auth.Login)
		group.POST("/logout", api.Auth.Logout)

		group.GET("/health", api.Health.GetHealthInfo)
		group.POST("/health", api.Health.AddBodyInfo)

		group.POST("/pwd", api.Password.Add)
		group.PUT("/pwd", api.Password.Modify)
		group.GET("/pwd", api.Password.GetPwd)
		group.GET("/pwd/list", api.Password.GetList)
	})
}
