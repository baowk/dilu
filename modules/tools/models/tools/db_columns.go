package tools

import (
	"errors"

	"github.com/baowk/dilu-core/core"
	"gorm.io/gorm"
)

type DBColumns struct {
	TableSchema            string `gorm:"column:table_schema" json:"tableSchema"`
	TableName              string `gorm:"column:table_name" json:"tableName"`
	ColumnName             string `gorm:"column:column_name" json:"columnName"`
	ColumnDefault          string `gorm:"column:column_default" json:"columnDefault"`
	IsNullable             string `gorm:"column:is_nullable" json:"isNullable"`
	DataType               string `gorm:"column:data_type" json:"dataType"`
	CharacterMaximumLength string `gorm:"column:character_maximum_length" json:"characterMaximumLength"`
	CharacterSetName       string `gorm:"column:character_set_name" json:"characterSetName"`
	ColumnType             string `gorm:"column:column_type" json:"columnType"`
	ColumnKey              string `gorm:"column:column_key" json:"columnKey"`
	Extra                  string `gorm:"column:extra" json:"extra"`
	ColumnComment          string `gorm:"column:column_comment" json:"columnComment"`
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

func (e *DBColumns) GetList(tx *gorm.DB, dbname, driver string) ([]DBColumns, error) {
	var doc []DBColumns
	table := new(gorm.DB)

	if e.TableName == "" {
		return nil, errors.New("table name cannot be empty！")
	}

	if driver == "mysql" {
		table = tx.Table("information_schema.columns")
		table = table.Where("table_schema= ? ", dbname)

		table = table.Where("TABLE_NAME = ?", e.TableName).Order("ORDINAL_POSITION asc")
	} else if driver == "pgsql" {
		table = tx.Table("information_schema.columns AS col").Joins("LEFT JOIN pg_catalog.pg_description AS pgd ON (col.table_name::regclass = pgd.objoid AND col.ordinal_position = pgd.objsubid)").Select("col.*, pgd.description AS column_comment")
		table = table.Where("col.table_schema= ? ", "public") // 使用默认 public，将来可配置

		table = table.Where("col.table_name = ?", e.TableName).Order("col.ORDINAL_POSITION asc")
	} else {
		return doc, errors.New("只支持mysql、postgresql")
	}
	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}
