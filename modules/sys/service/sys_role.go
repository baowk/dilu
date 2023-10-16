package service

import (
	//"github.com/acmestack/gorm-plus/gplus"
	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
)

type SysRoleService struct {
	*base.BaseService
}

var SerSysRole = SysRoleService{
	base.NewService(consts.DB_DEF),
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
