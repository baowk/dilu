package service

import (
	//"github.com/acmestack/gorm-plus/gplus"

	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
	"github.com/jinzhu/copier"
)

type SysRoleService struct {
	*base.BaseService
}

var SerSysRole = SysRoleService{
	base.NewService(consts.DB_DEF),
}

func (s *SysRoleService) Create(userId, teamId int, req dto.SysRoleDto) error {
	var model models.SysRole
	copier.Copy(&model, req)
	model.TeamId = teamId
	model.CreateBy = userId
	tx := s.DB().Begin()
	if err := tx.Create(&model).Error; err != nil {
		tx.Rollback()
		return err
	}
	//m := make(map[int]bool, 0)
	var rms []models.SysRoleMenu
	for _, mids := range req.MenuIds {
		rms = append(rms, models.SysRoleMenu{RoleId: model.Id, MenuId: mids})
	}
	if err := tx.Create(rms).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

func (s *SysRoleService) Update(userId, teamId int, req dto.SysRoleDto) error {
	var model models.SysRole
	copier.Copy(&model, req)
	model.TeamId = teamId
	model.UpdateBy = userId

	tx := s.DB().Begin()
	if err := tx.Updates(&model).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("role_id = ?", model.Id).Delete(&models.SysRoleMenu{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	//m := make(map[int]bool, 0)
	var rms []models.SysRoleMenu
	for _, mids := range req.MenuIds {
		rms = append(rms, models.SysRoleMenu{RoleId: model.Id, MenuId: mids})
	}
	if err := tx.Create(rms).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *SysRoleService) GetRole(id, userId, teamId int, data *dto.SysRoleDtoResp) error {
	var model models.SysRole
	if err := s.DB().First(&model, id).Error; err != nil {
		return err
	}
	copier.Copy(data, model)
	var menuIds []int
	if err := s.DB().Model(&models.SysRoleMenu{}).Select("menu_id").Where("role_id = ?", id).Find(&menuIds).Error; err != nil {
		return err
	}
	data.MenuIds = menuIds
	return nil
}

func (s *SysRoleService) Query(teamId, status int, list *[]models.SysRole) error {
	db := s.DB()
	if teamId != 0 {
		db = db.Where("team_id = ?", teamId)
	}
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	return db.Find(list).Error
}

func (s *SysRoleService) Page(req *dto.SysRoleGetPageReq, list *[]models.SysRole, total *int64) error {
	db := s.DB().Offset(req.GetOffset()).Limit(req.GetSize())
	if req.TeamId != 0 {
		db.Where("team_id = ?", req.TeamId)
	}
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	return db.Find(list).Offset(-1).Limit(-1).Count(total).Error
}
