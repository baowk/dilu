package dto

import (
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type SysRoleGetPageReq struct {
	base.ReqPage `search:"-"`
	Status       int `json:"status" form:"status"` //状态
	TeamId       int `json:"teamId"`               //团队id
}

type SysRoleDto struct {
	Id        int       `json:"id" `       //主键
	Name      string    `json:"name"`      //角色名称
	RoleKey   string    `json:"roleKey"`   //角色代码
	RoleSort  int       `json:"roleSort"`  //排序
	Status    int       `json:"status"`    //状态
	TeamId    int       `json:"team_id"`   //团队
	Remark    string    `json:"remark" `   //备注
	MenuIds   []int     `json:"menuIds"`   //菜单id
	CreateBy  int       `json:"createBy" ` //创建者
	UpdateBy  int       `json:"updateBy" ` //更新者
	CreatedAt time.Time `json:"createdAt"` //创建时间
	UpdatedAt time.Time `json:"updatedAt"` //最后更新时间
}

// SysRole
type SysRoleDtoResp struct {
	Id        int       `json:"id" `       //主键
	Name      string    `json:"name"`      //角色名称
	RoleKey   string    `json:"roleKey"`   //角色代码
	RoleSort  int       `json:"roleSort"`  //排序
	Status    int       `json:"status"`    //状态
	TeamId    int       `json:"team_id"`   //团队
	Remark    string    `json:"remark" `   //备注
	MenuIds   []int     `json:"menuIds"`   //菜单id
	CreateBy  int       `json:"createBy" ` //创建者
	UpdateBy  int       `json:"updateBy" ` //更新者
	CreatedAt time.Time `json:"createdAt"` //创建时间
	UpdatedAt time.Time `json:"updatedAt"` //最后更新时间
}
