package dto

import "github.com/baowk/dilu-core/core/base"

type SysTableSearch struct {
	TBName       string `form:"tableName" search:"type:exact;column:table_name;table:table_name"`
	TableComment string `form:"tableComment" search:"type:icontains;column:table_comment;table:table_comment"`
}

type SysTableReq struct {
	base.ReqPage
	DBName string `json:"dbname" form:"dbname"`
}
