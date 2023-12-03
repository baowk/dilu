package service

import (
	"dilu/common"
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"
	"encoding/json"
	"strings"
	"time"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
)

type SysMemberService struct {
	*base.BaseService
}

var SerSysMember = SysMemberService{
	base.NewService(consts.DB_DEF),
}

func (e *SysMemberService) Create(m *models.SysMember) error {
	if m.Name != "" {
		m.PY = utils.GetPinyin(m.Name)
	}
	if m.UserId == 0 {
		username := strings.ReplaceAll(m.PY, "-", "")
		user := models.SysUser{
			Username: username,
			Phone:    m.Phone,
			Nickname: m.Nickname,
			Name:     m.Name,
			Gender:   m.Gender,
			Status:   1,
		}
		if err := SerSysUser.Create(&user); err != nil {
			return err
		}
		m.UserId = user.Id
	}
	if m.DeptId > 0 {
		var dept models.SysDept
		if err := SerSysDept.Get(m.DeptId, &dept); err == nil {
			m.DeptPath = dept.DeptPath
			m.TeamId = dept.TeamId
		}
	} else {
		m.DeptPath = ""
	}
	return e.BaseService.Create(m)
}

func (e *SysMemberService) Update(m *models.SysMember) error {
	if m.Name != "" {
		m.PY = utils.GetPinyin(m.Name)
	}
	if m.DeptId > 0 {
		var dept models.SysDept
		if err := SerSysDept.Get(m.DeptId, &dept); err == nil {
			m.DeptPath = dept.DeptPath
		}
	}
	return e.UpdateById(m)
}

func (e *SysMemberService) Query(req dto.SysMemberGetPageReq, list *[]models.SysMember, total *int64) error {
	db := e.DB().Limit(req.GetSize()).Offset(req.GetOffset())
	if req.TeamId != 0 {
		db.Where("team_id = ?", req.TeamId)
	}
	if req.Status != 0 {
		db.Where("status = ?", req.Status)
	}
	if req.DeptId > 0 {
		db.Where("dept_id = ?", req.DeptId)
	} else if req.DeptPath != "" {
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
	return e.DB().Raw(sql, uid).Find(list).Error
}

func (e *SysMemberService) GetTeamMember(teamId, uid int, teamMember *dto.TeamMemberResp) error {
	sql := `Select t.name as team_name, t.owner ,user_id, team_id,nickname, m.name, phone, 
			dept_id, dept_path, post_id, roles, entry_time, gender,birthday
			From sys_team t,sys_member m 
			Where team_id = ? and  user_id = ? and m.status = 1 and t.status = 2 
			and m.team_id = t.id order by m.updated_at desc`
	return e.DB().Raw(sql, teamId, uid).Find(teamMember).Error
}

func (e *SysMemberService) GetMember(teamId, userId int, member *models.SysMember) error {
	cStr, err := e.Cache().Get(common.TeamMemberKey(teamId, userId))
	if err == nil && cStr != "" {
		err = json.Unmarshal([]byte(cStr), member)
		if err == nil {
			return nil
		}
	}
	err = e.DB().Where("team_id = ?", teamId).Where("user_id = ?", userId).First(member).Error
	if err != nil {
		return err
	}
	e.Cache().Set(common.TeamMemberKey(teamId, userId), *member, time.Hour)
	return nil
}

func (e *SysMemberService) GetMembersByUids(teamId int, uids []int, members *[]models.SysMember) error {
	return e.DB().Where("team_id = ?", teamId).Where("user_id in ?", uids).Find(members).Error
}

func (e *SysMemberService) GetMembers(teamId, userId int, deptPath string, name string, status int, members *[]models.SysMember) error {
	db := e.DB().Where("team_id = ?", teamId)
	if userId != 0 {
		db.Where("user_id = ?", userId)
	} else if deptPath != "" {
		db.Where("dept_path like ?", deptPath+"%")
	}
	if name != "" {
		db.Where("name like ?", "%"+name+"%")
	}
	if status != 0 {
		db.Where("status =?", status)
	}
	return db.Order("id").Find(members).Error
}
