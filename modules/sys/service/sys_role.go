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
	// var m models.SysMember
	// if err := SerSysMember.GetMember(teamId, userId, &m); err != nil {
	// 	return err
	// }
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
		// for _, mid := range mids {
		// 	if _, ok := m[mid]; !ok {
		// 		rms = append(rms, models.SysRoleMenu{RoleId: model.Id, MenuId: mid})
		// 		m[mid] = true
		// 	}
		// }
		rms = append(rms, models.SysRoleMenu{RoleId: model.Id, MenuId: mids})
	}
	if err := tx.Create(rms).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

func (s *SysRoleService) Update(userId, teamId int, req dto.SysRoleDto) error {
	// var m models.SysMember
	// if err := SerSysMember.GetMember(teamId, userId, &m); err != nil {
	// 	return err
	// }
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
		// for _, mid := range mids {
		// 	if _, ok := m[mid]; !ok {
		// 		rms = append(rms, models.SysRoleMenu{RoleId: model.Id, MenuId: mid})
		// 		m[mid] = true
		// 	}
		// }
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
	// if *&model.TeamId != teamId {
	// 	return codes.Err403(nil)
	// }
	copier.Copy(data, model)
	var menuIds []int
	if err := s.DB().Model(&models.SysRoleMenu{}).Select("menu_id").Where("role_id = ?", id).Find(&menuIds).Error; err != nil {
		return err
	}
	data.MenuIds = menuIds
	return nil
}

// func (s *SysRoleService) Query(teamId int) ([]*models.SysRole, error) {
// 	q, t := gplus.NewQuery[models.SysRole]()
// 	q.Eq(t.Status, 1).Eq(t.TeamId, teamId)
// 	list, rdb := gplus.SelectList[models.SysRole](q)
// 	if rdb.Error != nil {
// 		return nil, rdb.Error
// 	}
// 	return list, nil
// }
