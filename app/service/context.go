package service

import (
	"alfred/app/model"
	"context"
	"github.com/gogf/gf/net/ghttp"
)

var Context = new (serviceContext)

type serviceContext struct {}

func (s *serviceContext) Init(r *ghttp.Request, customCtx *model.Context)  {
	r.SetCtxVar(model.ContextKey, customCtx)
}

func (s *serviceContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

func (s *serviceContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}