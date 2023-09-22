package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type TeamService struct {
	*base.BaseService
}

var SerTeam = TeamService{
	base.NewService("dental"),
}

