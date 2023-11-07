package models

import (
	"time"
)

// GenColumns
type GenColumns struct {
	ColumnId           int       `json:"columnId" gorm:"type:bigint;primaryKey;autoIncrement;comment:主键"`    //主键
	TableId            int       `json:"tableId" gorm:"type:bigint;comment:TableId"`                         //
	ColumnName         string    `json:"columnName" gorm:"type:varchar(128);comment:ColumnName"`             //
	ColumnComment      string    `json:"columnComment" gorm:"type:varchar(128);comment:ColumnComment"`       //
	ColumnType         string    `json:"columnType" gorm:"type:varchar(128);comment:ColumnType"`             //
	GoType             string    `json:"goType" gorm:"type:varchar(128);comment:GoType"`                     //
	GoField            string    `json:"goField" gorm:"type:varchar(128);comment:GoField"`                   //
	JsonField          string    `json:"jsonField" gorm:"type:varchar(128);comment:JsonField"`               //
	IsPk               string    `json:"isPk" gorm:"type:varchar(4);comment:IsPk"`                           //
	IsIncrement        string    `json:"isIncrement" gorm:"type:varchar(4);comment:IsIncrement"`             //
	IsRequired         string    `json:"isRequired" gorm:"type:varchar(4);comment:IsRequired"`               //
	IsInsert           string    `json:"isInsert" gorm:"type:varchar(4);comment:IsInsert"`                   //
	IsEdit             string    `json:"isEdit" gorm:"type:varchar(4);comment:IsEdit"`                       //
	IsList             string    `json:"isList" gorm:"type:varchar(4);comment:IsList"`                       //
	IsQuery            string    `json:"isQuery" gorm:"type:varchar(4);comment:IsQuery"`                     //
	QueryType          string    `json:"queryType" gorm:"type:varchar(128);comment:QueryType"`               //
	HtmlType           string    `json:"htmlType" gorm:"type:varchar(128);comment:HtmlType"`                 //
	DictType           string    `json:"dictType" gorm:"type:varchar(128);comment:DictType"`                 //
	Sort               int       `json:"sort" gorm:"type:bigint;comment:Sort"`                               //
	List               string    `json:"list" gorm:"type:varchar(1);comment:List"`                           //
	Pk                 bool      `json:"pk" gorm:"type:tinyint(1);comment:Pk"`                               //
	Required           bool      `json:"required" gorm:"type:tinyint(1);comment:Required"`                   //
	SuperColumn        int       `json:"superColumn" gorm:"type:tinyint(1);comment:SuperColumn"`             //
	UsableColumn       int       `json:"usableColumn" gorm:"type:tinyint(1);comment:UsableColumn"`           //
	Increment          int       `json:"increment" gorm:"type:tinyint(1);comment:Increment"`                 //
	Insert             bool      `json:"insert" gorm:"type:tinyint(1);comment:Insert"`                       //
	Edit               int       `json:"edit" gorm:"type:tinyint(1);comment:Edit"`                           //
	Query              int       `json:"query" gorm:"type:tinyint(1);comment:Query"`                         //
	Remark             string    `json:"remark" gorm:"type:varchar(255);comment:Remark"`                     //
	FkTableName        string    `json:"fkTableName" gorm:"type:longtext;comment:FkTableName"`               //
	FkTableNameClass   string    `json:"fkTableNameClass" gorm:"type:longtext;comment:FkTableNameClass"`     //
	FkTableNamePackage string    `json:"fkTableNamePackage" gorm:"type:longtext;comment:FkTableNamePackage"` //
	FkLabelId          string    `json:"fkLabelId" gorm:"type:longtext;comment:FkLabelId"`                   //
	FkLabelName        string    `json:"fkLabelName" gorm:"type:varchar(255);comment:FkLabelName"`           //
	CreateBy           int       `json:"createBy" gorm:"type:mediumint;comment:CreateBy"`                    //
	UpdateBy           int       `json:"updateBy" gorm:"type:mediumint;comment:UpdateBy"`                    //
	CreatedAt          time.Time `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                     //创建时间
	UpdatedAt          time.Time `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                   //最后更新时间
	TsType             string    `gorm:"-" json:"tsType"`
}

func (GenColumns) TableName() string {
	return "gen_columns"
}
