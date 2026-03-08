package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type GenColumnsGetPageReq struct {
	base.ReqPage `search:"-"`
}

// GenColumns
type GenColumnsDto struct {
	ColumnId           int    `json:"columnId"`           //主键
	TableId            int    `json:"tableId"`            //
	ColumnName         string `json:"columnName"`         //
	ColumnComment      string `json:"columnComment"`      //
	ColumnType         string `json:"columnType"`         //
	GoType             string `json:"goType"`             //
	GoField            string `json:"goField"`            //
	JsonField          string `json:"jsonField"`          //
	IsPk               string `json:"isPk"`               //
	IsIncrement        string `json:"isIncrement"`        //
	IsRequired         string `json:"isRequired"`         //
	IsInsert           string `json:"isInsert"`           //
	IsEdit             string `json:"isEdit"`             //
	IsList             string `json:"isList"`             //
	IsQuery            string `json:"isQuery"`            //
	QueryType          string `json:"queryType"`          //
	HtmlType           string `json:"htmlType"`           //
	DictType           string `json:"dictType"`           //
	Sort               int    `json:"sort"`               //
	List               string `json:"list"`               //
	Pk                 int    `json:"pk"`                 //
	Required           int    `json:"required"`           //
	SuperColumn        int    `json:"superColumn"`        //
	UsableColumn       int    `json:"usableColumn"`       //
	Increment          int    `json:"increment"`          //
	Insert             int    `json:"insert"`             //
	Edit               int    `json:"edit"`               //
	Query              int    `json:"query"`              //
	Remark             string `json:"remark"`             //
	FkTableName        string `json:"fkTableName"`        //
	FkTableNameClass   string `json:"fkTableNameClass"`   //
	FkTableNamePackage string `json:"fkTableNamePackage"` //
	FkLabelId          string `json:"fkLabelId"`          //
	FkLabelName        string `json:"fkLabelName"`        //
}
