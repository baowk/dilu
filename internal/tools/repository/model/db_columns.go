package model

import (
	"dilu/common/config"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/jinzhu/copier"
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

type PgDBColumns struct {
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

type SqliteDBColumn struct {
	CID       int    `gorm:"column:cid"`
	Name      string `gorm:"column:name"`
	Type      string `gorm:"column:type"`
	NotNull   int    `gorm:"column:notnull"`
	DfltValue string `gorm:"column:dflt_value"`
	PK        int    `gorm:"column:pk"`
}

var tableNamePattern = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

func (e *DBColumns) GetPage(tx *gorm.DB, pageSize int, pageIndex int, dbname string) ([]DBColumns, int64, error) {
	var doc []DBColumns
	var count int64
	table := new(gorm.DB)

	if config.Get().DBCfg.Driver == "mysql" {
		table = tx.Table("information_schema.`COLUMNS`")
		table = table.Where("table_schema= ? ", dbname)

		if e.TableName == "" {
			return nil, 0, errors.New("table name cannot be empty！")
		}

		table = table.Where("TABLE_NAME = ?", e.TableName)
	} else if config.Get().DBCfg.Driver == "sqlite" {
		cols, err := e.GetList(tx, dbname, "sqlite")
		if err != nil {
			return nil, 0, err
		}
		count = int64(len(cols))
		start := (pageIndex - 1) * pageSize
		if start >= len(cols) {
			return []DBColumns{}, count, nil
		}
		end := start + pageSize
		if end > len(cols) {
			end = len(cols)
		}
		return cols[start:end], count, nil
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
		if err := table.Find(&doc).Error; err != nil {
			return doc, err
		}
	} else if driver == "pgsql" || driver == "postgres" {
		table = tx.Table("information_schema.columns AS col").Joins("LEFT JOIN pg_catalog.pg_description AS pgd ON (col.table_name::regclass = pgd.objoid AND col.ordinal_position = pgd.objsubid)").Select("col.*, pgd.description AS column_comment")
		table = table.Where("col.table_schema= ? ", "public") // 使用默认 public，将来可配置

		table = table.Where("col.table_name = ?", e.TableName).Order("col.ORDINAL_POSITION asc")

		var pgdoc []PgDBColumns
		if err := table.Find(&pgdoc).Error; err != nil {
			return doc, err
		}
		copier.Copy(&doc, pgdoc)
	} else if driver == "sqlite" {
		if !tableNamePattern.MatchString(e.TableName) {
			return doc, errors.New("invalid sqlite table name")
		}

		var sqliteCols []SqliteDBColumn
		sql := fmt.Sprintf("PRAGMA table_info(`%s`)", e.TableName)
		if err := tx.Raw(sql).Scan(&sqliteCols).Error; err != nil {
			return doc, err
		}

		doc = make([]DBColumns, 0, len(sqliteCols))
		for _, col := range sqliteCols {
			columnType := strings.ToLower(strings.TrimSpace(col.Type))
			dataType := columnType
			if idx := strings.IndexAny(dataType, "( "); idx > 0 {
				dataType = dataType[:idx]
			}

			isNullable := "YES"
			if col.NotNull == 1 || col.PK > 0 {
				isNullable = "NO"
			}

			columnKey := ""
			if col.PK > 0 {
				columnKey = "PRI"
			}

			doc = append(doc, DBColumns{
				TableSchema:   "main",
				TableName:     e.TableName,
				ColumnName:    col.Name,
				ColumnDefault: col.DfltValue,
				IsNullable:    isNullable,
				DataType:      dataType,
				ColumnType:    columnType,
				ColumnKey:     columnKey,
				Extra:         "",
				ColumnComment: "",
			})
		}
	} else {
		return doc, errors.New("只支持mysql、postgresql、sqlite")
	}
	return doc, nil
}
