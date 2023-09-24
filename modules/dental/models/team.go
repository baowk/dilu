package models

import (
	"time"
)

// Team
type Team struct {
	Id        int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	ParentId  int       `json:"parentId" gorm:"type:int unsigned;index;comment:上级团队"`            //上级团队
	Path      string    `json:"path" gorm:"type:varchar(255);comment:团队路径"`                      //团队路径
	Name      string    `json:"name" gorm:"type:varchar(32);comment:团队名"`                        //团队名
	Owner     int       `json:"owner" gorm:"type:int unsigned;comment:团队拥有者"`                    //团队拥有者
	Status    int       `json:"status" gorm:"type:tinyint;comment:状态"`                           //状态
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                     //创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                     //更新时间
}

func (Team) TableName() string {
	return "team"
}

func NewTeam() *Team {
	return &Team{}
}

func (e *Team) SetId(id int) *Team {
	e.Id = id
	return e
}
func (e *Team) SetParentId(parentId int) *Team {
	e.ParentId = parentId
	return e
}
func (e *Team) SetPath(path string) *Team {
	e.Path = path
	return e
}
func (e *Team) SetName(name string) *Team {
	e.Name = name
	return e
}
func (e *Team) SetOwner(owner int) *Team {
	e.Owner = owner
	return e
}
func (e *Team) SetStatus(status int) *Team {
	e.Status = status
	return e
}
func (e *Team) SetCreatedAt(createdAt time.Time) *Team {
	e.CreatedAt = createdAt
	return e
}
func (e *Team) SetUpdatedAt(updatedAt time.Time) *Team {
	e.UpdatedAt = updatedAt
	return e
}
