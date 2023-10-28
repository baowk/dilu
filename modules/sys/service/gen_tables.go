package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type GenTablesService struct {
	*base.BaseService
}

var SerGenTables = GenTablesService{
	base.NewService("sys"),
}

