package tools

import (
	"errors"

	"github.com/baowk/dilu-core/core"
	"gorm.io/gorm"
)

type DBTables struct {
	TableName      string `gorm:"column:TABLE_NAME" json:"tableName"`
	TableSchema    string `gorm:"column:TABLE_SCHEMA" json:"tableSchema"`
	Engine         string `gorm:"column:ENGINE" json:"engine"`
	TableRows      string `gorm:"column:TABLE_ROWS" json:"tableRows"`
	TableCollation string `gorm:"column:TABLE_COLLATION" json:"tableCollation"`
	CreateTime     string `gorm:"column:CREATE_TIME" json:"createTime"`
	UpdateTime     string `gorm:"column:UPDATE_TIME" json:"updateTime"`
	TableComment   string `gorm:"column:TABLE_COMMENT" json:"tableComment"`
}

func (e *DBTables) GetPage(tx *gorm.DB, pageSize int, pageIndex int, dbname string, defDbName string) ([]DBTables, int64, error) {
	var doc []DBTables
	table := new(gorm.DB)
	var count int64

	if core.Cfg.DBCfg.Driver == "mysql" {

		table = tx.Table("information_schema.tables")
		table = table.Where("TABLE_NAME not in (select table_name from `" + defDbName + "`.gen_tables where db_name = ? )", dbname)
		table = table.Where("table_schema= ? ", dbname)

		if e.TableName != "" {
			table = table.Where("TABLE_NAME = ?", e.TableName)
		}
		if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
			return nil, 0, err
		}
	} else {
		return doc, 500, errors.New("只支持mysql")
	}

	//table.Count(&count)
	return doc, count, nil
}

func (e *DBTables) Get(tx *gorm.DB, dbname string) (DBTables, error) {
	var doc DBTables
	if core.Cfg.DBCfg.Driver == "mysql" {
		table := tx.Table("information_schema.tables")
		table = table.Where("table_schema= ? ", dbname)
		if e.TableName == "" {
			return doc, errors.New("table name cannot be empty！")
		}
		table = table.Where("TABLE_NAME = ?", e.TableName)
		if err := table.First(&doc).Error; err != nil {
			return doc, err
		}
	} else {
		return doc, errors.New("只支持mysql")
	}
	return doc, nil
}
