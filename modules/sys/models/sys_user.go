package models

import (
	"github.com/baowk/dilu-core/core/base"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	UserId   int    `json:"userId" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:编码"`
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`
	Password string `json:"-" gorm:"size:128;comment:密码"`
	NickName string `json:"nickName" gorm:"size:128;comment:昵称"`
	Avatar   string `json:"avatar" gorm:"size:255;comment:头像"`
	Bio      string `json:"bio" gorm:"type:varchar(255);default:(-);comment:签名"`             //签名
	Gender   string `json:"gender" gorm:"type:char(1);default:'2';comment:性别  1女  2男  3未知 "` //性别  1 男 2女  3 未知
	RoleId   int    `json:"roleId" gorm:"type:int unsigned;size:20;comment:角色ID"`
	DeptId   int    `json:"deptId" gorm:"type:int unsigned;size:20;comment:部门"`
	Post     string `json:"post" gorm:"size:32;comment:岗位"`
	Remark   string `json:"remark" gorm:"size:255;comment:备注"`
	Status   string `json:"status" gorm:"type:tinyint;comment:状态"`
	base.ControlBy
	base.ModelTime
}

func (SysUser) TableName() string {
	return "sys_user"
}

// 加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *SysUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}
