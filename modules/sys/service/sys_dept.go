package service

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"
	"fmt"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
	"github.com/jinzhu/copier"
)

type SysDeptService struct {
	*base.BaseService
}

var SerSysDept = SysDeptService{
	base.NewService(consts.DB_DEF),
}

func (s *SysDeptService) Page(req *dto.SysDeptGetPageReq, list *[]models.SysDept, total *int64) error {
	db := s.DB().Offset(req.GetOffset()).Limit(req.GetSize())
	if req.TeamId != 0 {
		db.Where("team_id = ?", req.TeamId)
	}
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	return db.Find(list).Offset(-1).Limit(-1).Count(total).Error
}

func (s *SysDeptService) GetDepts(teamId int, list *[]models.SysDept) error {
	if teamId != 0 {
		return s.DB().Where("team_id = ?", teamId).Find(list).Error
	}
	return s.DB().Find(list).Error
}

func (s *SysDeptService) CreateDept(req dto.SysDeptDto, adminId int, reqId string) error {
	var pdept models.SysDept
	if req.ParentId > 0 {
		if err := s.Get(req.ParentId, &pdept); err != nil {
			return err
		}
	}
	var data models.SysDept
	copier.Copy(&data, req)
	data.CreateBy = adminId
	if err := s.BaseService.Create(&data); err != nil {
		return err
	}
	if req.ParentId > 0 {
		data.DeptPath = fmt.Sprintf("%s%d/", pdept.DeptPath, data.Id)
	} else {
		data.DeptPath = fmt.Sprintf("/0/%d/", data.Id)
	}
	return s.UpdateById(data)
}

func (s *SysDeptService) UpdateDept(req dto.SysDeptDto, adminId int, reqId string) error {
	var old models.SysDept
	if err := s.Get(req.Id, &old); err != nil {
		return err
	}
	var pdept models.SysDept
	if req.ParentId > 0 && old.ParentId != req.ParentId {
		if err := s.Get(req.ParentId, &pdept); err != nil {
			return err
		}
	}
	copier.Copy(&old, req)
	old.UpdateBy = adminId

	if pdept.Id > 0 {
		old.DeptPath = fmt.Sprintf("%s%d/", pdept.DeptPath, old.Id)
	}
	return s.UpdateById(old)
}
