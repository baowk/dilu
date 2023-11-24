package models

import (
	"time"
)

// 公用通知
type PubNotice struct {
	Id         int       `json:"id" gorm:"type:int;primaryKey;autoIncrement;comment:主键"` //主键
	TeamId     int       `json:"teamId" gorm:"type:int;comment:针对组消息"`                   //针对组消息
	Title      string    `json:"title" gorm:"type:varchar(255);comment:标题"`              //标题
	Content    string    `json:"content" gorm:"type:varchar(1024);comment:内容"`           //内容
	NoticeType int       `json:"noticeType" gorm:"type:tinyint;comment:消息类型"`            //消息类型
	Op         int       `json:"op" gorm:"type:tinyint;comment:操作类型"`                    //操作类型
	OpId       int       `json:"opId" gorm:"type:int;comment:操作id"`                      //操作id
	Status     int       `json:"status" gorm:"type:tinyint;comment:状态 1正常 2删除"`          //状态1正常 2删除
	CreateBy   int       `json:"createBy" gorm:"type:int;comment:创建人"`                   //创建人
	UpdateBy   int       `json:"updateBy" gorm:"type:int;comment:更新人"`                   //更新人
	Expired    time.Time `json:"expired" gorm:"type:datetime;comment:到期时间"`              //到期时间
	CreatedAt  time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`            //创建时间
	UpdatedAt  time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`            //更新时间
}

const TBPubNotice = "pub_notice"

func (PubNotice) TableName() string {
	return TBPubNotice
}

func NewPubNotice() *PubNotice {
	return &PubNotice{}
}
