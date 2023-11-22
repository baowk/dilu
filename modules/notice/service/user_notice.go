package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type UserNoticeService struct {
	*base.BaseService
}

var SerUserNotice = UserNoticeService{
	base.NewService("notice"),
}

