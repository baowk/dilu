package dto

import (
	"dilu/modules/notice/models"
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type PubNoticeGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int    `json:"status" query:"column:status"` //状态

}

func (PubNoticeGetPageReq) TableName() string {
	return models.TBPubNotice
}

// 公用通知
type PubNoticeDto struct {
	Id         int       `json:"id"`         //主键
	TeamId     int       `json:"teamId"`     //针对组消息
	Title      string    `json:"title"`      //标题
	Content    string    `json:"content"`    //内容
	NoticeType int       `json:"noticeType"` //消息类型
	Op         int       `json:"op"`         //操作类型
	OpId       int       `json:"opId"`       //操作id
	Status     int       `json:"status"`     //状态
	Expired    time.Time `json:"expired"`    //到期时间
}
