package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type PubNoticeService struct {
	*base.BaseService
}

var SerPubNotice = PubNoticeService{
	base.NewService("notice"),
}

