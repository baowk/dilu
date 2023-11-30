package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	Id             int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"`       //主键
	Username       string    `json:"username" gorm:"size:32;comment:用户名"`                                   //用户名
	Phone          string    `json:"phone" gorm:"size:11;comment:手机号"`                                      //手机号
	Email          string    `json:"email" gorm:"size:128;comment:邮箱"`                                      //邮箱
	Password       string    `json:"password" gorm:"size:128;comment:密码"`                                   //密码
	Nickname       string    `json:"nickname" gorm:"size:128;comment:昵称"`                                   //昵称
	Name           string    `json:"name" gorm:"size:64;comment:姓名"`                                        //姓名
	Avatar         string    `json:"avatar" gorm:"size:255;comment:头像"`                                     //头像
	Bio            string    `json:"bio" gorm:"type:varchar(255);default:(-);comment:签名"`                   //签名
	Birthday       string    `json:"birthday" gorm:"type:date;default:(-);comment:生日 格式 yyyy-MM-dd"`        //生日
	Gender         int       `json:"gender" gorm:"type:tinyint;default:2;comment:性别 1男 2女 3未知"`             //性别 1男 2女 3未知
	PlatformRoleId int       `json:"platformRoleId" gorm:"type:int;size:20;comment:平台角色ID 大于0为平台账户,0为团队账户"` //平台角色ID 大于0为平台账户,0为团队账户
	Status         int       `json:"status" gorm:"type:tinyint;comment:状态 1正常 "`                            //状态 1正常
	Remark         string    `json:"remark" gorm:"size:255;comment:备注"`                                     //备注
	LockTime       time.Time `json:"lockTime" gorm:"type:datetime;default:(-);comment:锁定结束时间"`              //锁定结束时间
	CreatedAt      time.Time `json:"createdAt" gorm:"comment:创建时间"`                                         //创建时间
	UpdatedAt      time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`                                       //更新时间
	UpdateBy       int       `json:"updateBy" gorm:"type:int unsigned;index;comment:更新者"`                   //更新者id
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
