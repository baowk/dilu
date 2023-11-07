package service

import (
	"dilu/modules/tools/models"

	"github.com/baowk/dilu-core/core/base"
	"gorm.io/gorm"
)

type GenColumnsService struct {
	*base.BaseService
}

var SerGenColumns = GenColumnsService{
	base.NewService("sys"),
}

func (e *GenColumnsService) GetList(tx *gorm.DB, exclude bool, tableId int) ([]models.GenColumns, error) {
	var doc []models.GenColumns

	table := tx
	if table == nil {
		table = e.DB()
	}
	if err := table.Where("table_id = ?", tableId).Find(&doc).Error; err != nil {
		return nil, err
	}
	return doc, nil
}

// func (e *GenColumn) Create(tx *gorm.DB) (GenColumn, error) {
// 	var doc GenColumn
// 	e.CreateBy = 0
// 	result := tx.Table("gen_columns").Create(&e)
// 	if result.Error != nil {
// 		err := result.Error
// 		return doc, err
// 	}
// 	doc = *e
// 	return doc, nil
// }
