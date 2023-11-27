package dto

import (
	"dilu/modules/notice/models"

	"github.com/baowk/dilu-core/core/base"
)

type UserNoticeGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:created_at"`
	Status       int    `json:"status" query:"column:status"` //状态 1未读 2已读 -1回收站
	TeamId       int    `json:"teamId" query:""`
	UserId       int    `json:"userId" query:""`
}

func (UserNoticeGetPageReq) TableName() string {
	return models.TBUserNotice
}

// 用户通知
type UserNoticeDto struct {
	Id         int    `json:"id"`         //主键
	TeamId     int    `json:"teamId"`     //团队id
	UserId     int    `json:"userId"`     //用户id
	Title      string `json:"title"`      //标题
	Content    string `json:"content"`    //内容
	NoticeType int    `json:"noticeType"` //消息类型
	Op         int    `json:"op"`         //操作类型
	OpId       int    `json:"opId"`       //操作对象id
	Status     int    `json:"status"`     //状态 1未读 2已读 -1回收站
}

type NoticeDto struct {
	Key   string       `json:"key"`
	Name  string       `json:"name"`
	Count int64        `json:"count"`
	Total int64        `json:"total"`
	List  []NoticeItem `json:"list"`
}

type NoticeItem struct {
	Id             int    `json:"id"`
	Title          string `json:"title"`
	Content        string `json:"content"`
	Type           int    `json:"type"`
	NoticeType     int    `json:"noticeType"`
	Op             int    `json:"op"`
	OpId           int    `json:"opId"`
	Status         int    `json:"status"`
	CreatedAt      int64  `json:"createdAt,omitempty"`
	BeginAt        int64  `json:"beginAt,omitempty"`
	EndAt          int64  `json:"endAt,omitempty"`
	ReminderTime   int64  `json:"reminderTime,omitempty"`
	ReminderStatus int    `json:"reminderStatus"`
}

type ReadNoticeDto struct {
	Key int `json:"key"`
	Id  int `json:"id"`
}
