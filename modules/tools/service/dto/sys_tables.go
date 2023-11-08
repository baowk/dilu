package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysTableSearch struct {
	TBName       string `form:"tableName" search:"type:exact;column:table_name;table:table_name"`
	TableComment string `form:"tableComment" search:"type:icontains;column:table_comment;table:table_comment"`
}

type SysTableReq struct {
	base.ReqPage
	DBName string `json:"dbname" form:"dbname"`
}

type DBReq struct {
	base.ReqPage `search:"-"`
	DBName       string `json:"dbName" form:"dbName"`
	TableName    string `json:"tableName"`
}

type ImpTablesReq struct {
	DbName string   `json:"dbName"`
	Tables []string `json:"tables"`
}

type GenCodeReq struct {
	TableId int  `json:"tableId"`
	Force   bool `json:"force"`
}

type GenMenuReq struct {
	TableId int `json:"tableId"`
	MenuId  int `json:"menuId"`
}
