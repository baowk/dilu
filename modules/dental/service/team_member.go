package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type TeamMemberService struct {
	*base.BaseService
}

var SerTeamMember = TeamMemberService{
	base.NewService("dental"),
}

