package service

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
)

type GenTablesService struct {
	*base.BaseService
}

var SerGenTables = GenTablesService{
	base.NewService("sys"),
}

func (s *GenTablesService) Page(req *dto.GenTablesGetPageReq, list *[]models.GenTables, total *int64) error {
	db := s.DB().Order("table_id desc").Offset(req.GetOffset()).Limit(req.GetSize())
	if req.DbName != "" {
		db.Where("db_name = ?", req.DbName)
	}
	if req.TableName != "" {
		db.Where("table_name = ?", req.TableName)
	}
	return db.Find(list).Offset(-1).Limit(-1).Count(total).Error
}
