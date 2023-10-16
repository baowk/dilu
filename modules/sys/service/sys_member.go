package service

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
)

type SysMemberService struct {
	*base.BaseService
}

var SerSysMember = SysMemberService{
	base.NewService(consts.DB_DEF),
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

func (e *SysMemberService) GetUserTeams(uid int, list *[]dto.TeamMemberResp) error {
	sql := `Select t.name as team_name, t.owner ,user_id, team_id,nickname, m.name, phone, 
			dept_id, dept_path, post_id, roles, entry_time, gender,birthday
			From sys_team t,sys_member m 
			Where user_id = ? and m.status = 1 and t.status = 2 and m.team_id = t.id 
			order by m.updated_at desc`
	return core.DB().Raw(sql, uid).Find(list).Error
}

func (e *SysMemberService) GetTeamUser(teamId, uid int, teamMember *dto.TeamMemberResp) error {
	sql := `Select t.name as team_name, t.owner ,user_id, team_id,nickname, m.name, phone, 
			dept_id, dept_path, post_id, roles, entry_time, gender,birthday
			From sys_team t,sys_member m 
			Where team_id = ? and  user_id = ? and m.status = 1 and t.status = 2 
			and m.team_id = t.id order by m.updated_at desc`
	return core.DB().Raw(sql, teamId, uid).Find(teamMember).Error
}

func (e *SysMemberService) GetMembers(teamId int, name string, members *[]models.SysMember) error {
	where := models.SysMember{
		TeamId: teamId,
		Name:   name,
	}
	return e.GetByWhere(where, members)
}
