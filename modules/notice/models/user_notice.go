package models

import (
	"time"

	"gorm.io/gorm"
)

// 用户通知
type UserNotice struct {
	Id         int            `json:"id" gorm:"type:bigint unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	TeamId     int            `json:"teamId" gorm:"type:int;comment:团队id"`                                //团队id
	UserId     int            `json:"userId" gorm:"type:int;comment:用户id"`                                //用户id
	Title      string         `json:"title" gorm:"type:varchar(255);comment:标题"`                          //标题
	Content    string         `json:"content" gorm:"type:varchar(1024);comment:内容"`                       //内容
	NoticeType int            `json:"noticeType" gorm:"type:tinyint;comment:消息类型"`                        //消息类型
	Op         int            `json:"op" gorm:"type:tinyint;comment:操作类型"`                                //操作类型
	OpId       int            `json:"opId" gorm:"type:int;comment:操作对象id"`                                //操作对象id
	Status     int            `json:"status" gorm:"type:tinyint;comment:状态 1未读 2已读 -1回收站"`                //状态 1未读 2已读 -1回收站
	PubId      int            `json:"pubId" gorm:"type:int;comment:公共id"`                                 //公共id
	CreateBy   int            `json:"createBy" gorm:"type:int;comment:创建人"`                               //创建人
	CreatedAt  time.Time      `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                        //创建时间
	UpdatedAt  time.Time      `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                        //更新时间
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"type:datetime;comment:删除时间"`                                //删除时间
}

const TBUserNotice = "user_notice"

func (UserNotice) TableName() string {
	return TBUserNotice
}

func NewUserNotice() *UserNotice {
	return &UserNotice{}
}
