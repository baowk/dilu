package service

import (
	"dilu/internal/tools/repository/model"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
	"gorm.io/gorm"
)

type GenColumnsService struct {
	*base.BaseService
}

var SerGenColumns = GenColumnsService{
	base.NewService(consts.DB_DEF),
}

func (e *GenColumnsService) GetList(tx *gorm.DB, exclude bool, tableId int) ([]model.GenColumns, error) {
	var doc []model.GenColumns

	table := tx
	if table == nil {
		table = e.DB()
	}
	if err := table.Where("table_id = ?", tableId).Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}
