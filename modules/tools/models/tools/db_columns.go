package tools

import (
	"errors"

	"github.com/baowk/dilu-core/core"
	"gorm.io/gorm"
)

type DBColumns struct {
	TableSchema            string `gorm:"column:TABLE_SCHEMA" json:"tableSchema"`
	TableName              string `gorm:"column:TABLE_NAME" json:"tableName"`
	ColumnName             string `gorm:"column:COLUMN_NAME" json:"columnName"`
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT" json:"columnDefault"`
	IsNullable             string `gorm:"column:IS_NULLABLE" json:"isNullable"`
	DataType               string `gorm:"column:DATA_TYPE" json:"dataType"`
	CharacterMaximumLength string `gorm:"column:CHARACTER_MAXIMUM_LENGTH" json:"characterMaximumLength"`
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME" json:"characterSetName"`
	ColumnType             string `gorm:"column:COLUMN_TYPE" json:"columnType"`
	ColumnKey              string `gorm:"column:COLUMN_KEY" json:"columnKey"`
	Extra                  string `gorm:"column:EXTRA" json:"extra"`
	ColumnComment          string `gorm:"column:COLUMN_COMMENT" json:"columnComment"`
}

func (e *DBColumns) GetPage(tx *gorm.DB, pageSize int, pageIndex int, dbname string) ([]DBColumns, int64, error) {
	var doc []DBColumns
	var count int64
	table := new(gorm.DB)

	if core.Cfg.DBCfg.Driver == "mysql" {
		table = tx.Table("information_schema.`COLUMNS`")
		table = table.Where("table_schema= ? ", dbname)

		if e.TableName == "" {
			return nil, 0, errors.New("table name cannot be empty！")
		}

		table = table.Where("TABLE_NAME = ?", e.TableName)
	}

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	//table.Count(&count)
	return doc, count, nil

}

func (e *DBColumns) GetList(tx *gorm.DB, dbname string) ([]DBColumns, error) {
	var doc []DBColumns
	table := new(gorm.DB)

	if e.TableName == "" {
		return nil, errors.New("table name cannot be empty！")
	}

	if core.Cfg.DBCfg.Driver == "mysql" {
		table = tx.Table("information_schema.columns")
		table = table.Where("table_schema= ? ", dbname)

		table = table.Where("TABLE_NAME = ?", e.TableName).Order("ORDINAL_POSITION asc")
	} else {
		return doc, errors.New("只支持mysql")
	}
	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
