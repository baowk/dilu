package models

import (
	"time"
)

// 团队
type SysTeam struct {
	Id        int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Name      string    `json:"name" gorm:"type:varchar(32);comment:团队名"`                        //团队名
	Owner     int       `json:"owner" gorm:"type:int unsigned;comment:团队拥有者"`                    //团队拥有者
	Status    int       `json:"status" gorm:"type:tinyint;comment:状态"`                           //状态
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                     //创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                     //更新时间
}

func (SysTeam) TableName() string {
	return "sys_team"
}

func NewSysTeam() *SysTeam {
	return &SysTeam{}
}

func (e *SysTeam) SetId(id int) *SysTeam {
	e.Id = id
	return e
}
func (e *SysTeam) SetName(name string) *SysTeam {
	e.Name = name
	return e
}
func (e *SysTeam) SetOwner(owner int) *SysTeam {
	e.Owner = owner
	return e
}
func (e *SysTeam) SetStatus(status int) *SysTeam {
	e.Status = status
	return e
}
func (e *SysTeam) SetCreatedAt(createdAt time.Time) *SysTeam {
	e.CreatedAt = createdAt
	return e
}
func (e *SysTeam) SetUpdatedAt(updatedAt time.Time) *SysTeam {
	e.UpdatedAt = updatedAt
	return e
}
