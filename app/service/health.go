package service

import (
	"alfred/app/dao"
	"alfred/app/model"
	"context"
	"github.com/gogf/gf/os/gtime"
)

type healthService struct {}

var HealthService = healthService{}

func (a *healthService) AddBodyInfo(ctx context.Context, r *model.HealthReq) bool {
	var (
		info *model.Bodyinfo
	)
	info.FatRatio = r.FatRatio
	info.MuscleRatio = r.MuscleRatio
	info.Height = r.Height
	info.Weight = r.Weight
	info.RecordDate = gtime.Now()
	info.UserID = r.UserID

	if _, err := dao.Bodyinfo.Save(info); err != nil {
		return false
	}
	return true
}
