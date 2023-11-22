package dto

import (
	"dilu/modules/notice/models"
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type UserNoticeGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int    `json:"status" query:"column:status"` //状态 1未读 2已读 -1回收站

}

func (UserNoticeGetPageReq) TableName() string {
	return models.TBUserNotice
}

// 用户通知
type UserNoticeDto struct {
	Id         int       `json:"id"`         //主键
	TeamId     int       `json:"teamId"`     //团队id
	UserId     int       `json:"userId"`     //用户id
	Title      string    `json:"title"`      //标题
	Content    string    `json:"content"`    //内容
	NoticeType int       `json:"noticeType"` //消息类型
	Op         int       `json:"op"`         //操作类型
	OpId       int       `json:"opId"`       //操作对象id
	Status     int       `json:"status"`     //状态 1未读 2已读 -1回收站
	DeleteAt   time.Time `json:"deleteAt"`   //删除时间
}
