package service

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
)

type SysMemberService struct {
	*base.BaseService
}

var SerSysMember = SysMemberService{
	base.NewService("sys"),
}

func (e *SysMemberService) Query(req dto.SysMemberGetPageReq, list *[]models.SysMember, total *int64) error {
	db := core.DB().Limit(req.GetSize()).Offset(req.GetOffset())
	if req.TeamId != 0 {
		db.Where("team_id = ?", req.TeamId)
	}
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	if req.DeptPath != "" {
		db.Where("dept_path like ?", req.DeptPath+"%")
	}
	if req.Name != "" {
		db.Where("name like ?", "%"+req.Name+"%")
	}
	if req.Phone != "" {
		db.Where("phone = ?", req.Phone)
	}
	return db.Find(list).Limit(-1).Offset(-1).Count(total).Error
}
