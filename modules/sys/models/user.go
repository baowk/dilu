package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/baowk/dilu-core/core/base"
)

type User struct {
	base.Model

	Username   string   `json:"username" gorm:"type:varchar(50);default:(-);comment:用户名"`                  //用户名
	Phone      string   `json:"phone" gorm:"type:varchar(18);default:(-);comment:手机号"`                     //用户手机号
	Email      string   `json:"email" gorm:"type:varchar(128);default:(-);comment:邮箱"`                     //邮箱
	Bind       UserBind `json:"bind" gorm:"type:json;default:(-);comment:绑定情况 exp email:1代表邮箱已验证，第三方登录一样"` //绑定情况 exp email:1代表邮箱已验证，第三方登录一样
	Password   string   `json:"password" gorm:"type:varchar(80);default:(-);comment:密码"`                   //密码
	FirstName  string   `json:"firstName" gorm:"type:varchar(64);default:(-);comment:名"`                   //名
	LastName   string   `json:"lastName" gorm:"type:varchar(64);default:(-);comment:姓"`                    //姓
	Nickname   string   `json:"nickname" gorm:"type:varchar(64);default:(-);comment:昵称"`                   //昵称
	Avatar     string   `json:"avatar" gorm:"type:varchar(255);default:(-);comment:头像"`                    //头像
	Bio        string   `json:"bio" gorm:"type:varchar(255);default:(-);comment:签名"`                       //签名
	Gender     string   `json:"gender" gorm:"type:char(1);default:'2';comment:性别 1男 2女 3未知 "`              //性别  1 男 2女  3 未知
	Birthday   int      `json:"birthday" gorm:"type:date;default:(-);comment:生日"`                          //生日
	LockEnd    int      `json:"lockEnd" gorm:"type:datetime;default:(-);comment:锁定截至时间"`                   //锁定截至时间
	Source     string   `json:"source" gorm:"type:varchar(32);default:(-);comment:来源"`                     //来源（最长32位）
	Ip         string   `json:"ip" gorm:"type:varchar(32);default:(-);comment:ip"`                         //ip
	IpLocation string   `json:"ipLocation" gorm:"type:varchar(255);default:(-);comment:ip位置"`              //ip位置
	Inviter    string   `json:"inviter" gorm:"type:varchar(32);default:(-);comment:邀请人"`                   //邀请人
	InviteType int      `json:"inviteType" gorm:"type:tinyint;comment:邀请类型"`                               //邀请类型
	base.ControlBy
	base.ModelIntTime
}

type UserBind struct {
	Wechat int `json:"wechat" comment:"微信"` //微信 1绑定 下同
	Ding   int `json:"ding" comment:"钉钉"`   //
}

func (c UserBind) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *UserBind) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

func (User) TableName() string {
	return "user"
}
