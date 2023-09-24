package models

import (
	"time"
)

// TeamMember
type TeamMember struct {
	Id        int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"`                 //主键
	UserId    int       `json:"userId" gorm:"type:int unsigned;uniqueIndex:idx_tid_uid,priority:2;comment:用户id"` //用户id
	TeamId    int       `json:"teamId" gorm:"type:int unsigned;uniqueIndex:idx_tid_uid,priority:1;comment:团队id"` //团队id
	Name      string    `json:"name" gorm:"type:varchar(32);comment:姓名"`                                         //姓名
	Phone     string    `json:"phone" gorm:"type:varchar(11);comment:电话"`                                        //电话
	Gender    int       `json:"gender" gorm:"type:tinyint;comment:性别"`                                           //性别
	Status    int       `json:"status" gorm:"type:tinyint;comment:状态"`                                           //状态 1 在职 2离职
	Role      int       `json:"role" gorm:"type:tinyint;comment:角色 1主管 2副主管 4普通"`                                //角色 1主管 2副主管 4普通
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                                     //创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                                     //更新时间
}

func (TeamMember) TableName() string {
	return "team_member"
}

func NewTeamMember() *TeamMember {
	return &TeamMember{}
}

func (e *TeamMember) SetId(id int) *TeamMember {
	e.Id = id
	return e
}
func (e *TeamMember) SetUserId(userId int) *TeamMember {
	e.UserId = userId
	return e
}
func (e *TeamMember) SetTeamId(teamId int) *TeamMember {
	e.TeamId = teamId
	return e
}
func (e *TeamMember) SetName(name string) *TeamMember {
	e.Name = name
	return e
}
func (e *TeamMember) SetPhone(phone string) *TeamMember {
	e.Phone = phone
	return e
}
func (e *TeamMember) SetGender(gender int) *TeamMember {
	e.Gender = gender
	return e
}
func (e *TeamMember) SetStatus(status int) *TeamMember {
	e.Status = status
	return e
}
func (e *TeamMember) SetRole(role int) *TeamMember {
	e.Role = role
	return e
}
func (e *TeamMember) SetCreatedAt(createdAt time.Time) *TeamMember {
	e.CreatedAt = createdAt
	return e
}
func (e *TeamMember) SetUpdatedAt(updatedAt time.Time) *TeamMember {
	e.UpdatedAt = updatedAt
	return e
}
