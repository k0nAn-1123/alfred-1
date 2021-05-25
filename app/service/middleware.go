package service

import (
	"alfred/app/model"
	"github.com/gogf/gf/net/ghttp"
)

var MiddleWare = new (serviceMiddleware)

type serviceMiddleware struct{}

func (s *serviceMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r, customCtx)
	if user := Session.GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       		user.ID,
			Identification: user.Identification,
			Password: 		user.Password,
			Nickname: 		user.NickName,
			Permission:		user.Permission,
		}
	}
	r.Middleware.Next()
}
//// 鉴权中间件，只有登录成功之后才能通过
//func (s *serviceMiddleware) Auth(r *ghttp.Request) {
//	if User.IsSignedIn(r.Context()) {
//		r.Middleware.Next()
//	} else {
//		r.Response.WriteStatus(http.StatusForbidden)
//	}
//}

// 允许接口跨域请求
func (s *serviceMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}