package models

import (
	"github.com/baowk/dilu-core/core/base"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	Id       int    `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Username string `json:"username" gorm:"size:32;comment:用户名"`                             //用户名
	Phone    string `json:"phone" gorm:"size:11;comment:手机号"`                                //手机号
	Email    string `json:"email" gorm:"size:128;comment:邮箱"`                                //邮箱
	Password string `json:"password" gorm:"size:128;comment:密码"`                             //密码
	Nickname string `json:"nickname" gorm:"size:128;comment:昵称"`                             //昵称
	Name     string `json:"name" gorm:"size:64;comment:姓名"`                                  //姓名
	Avatar   string `json:"avatar" gorm:"size:255;comment:头像"`                               //头像
	Bio      string `json:"bio" gorm:"type:varchar(255);default:(-);comment:签名"`             //签名
	Birthday string `json:"birthday" gorm:"type:date;default:(-);comment:生日 格式 yyyy-MM-dd"`  //生日
	Gender   string `json:"gender" gorm:"type:char(1);default:'2';comment:性别 1男 2女 3未知"`     //性别 1男 2女 3未知
	RoleId   int    `json:"roleId" gorm:"type:int unsigned;size:20;comment:角色ID"`            //角色id
	Post     string `json:"post" gorm:"size:32;comment:岗位"`                                  //岗位
	Remark   string `json:"remark" gorm:"size:255;comment:备注"`                               //备注
	Status   int    `json:"status" gorm:"type:tinyint;comment:状态 1冻结 2正常 3默认"`               //状态 1冻结 2正常 3默认
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

func (e *SysUser) GenPwd(pwd string) (enPwd string, err error) {
	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost); err != nil {
		return
	} else {
		enPwd = string(hash)
	}
	return
}

func (e *SysUser) CompPwd(srcPwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(srcPwd)); err != nil {
		return false
	}
	return true
}
