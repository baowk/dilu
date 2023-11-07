package models

import (
	"time"
)

// GenTables
type GenTables struct {
	TableId             int          `json:"tableId" gorm:"type:bigint;primaryKey;autoIncrement;comment:主键"`           //主键
	DbName              string       `json:"dbName" gorm:"type:varchar(64);comment:DbName"`                            //
	TBName              string       `json:"tableName" gorm:"column:table_name;type:varchar(128);comment:TableName"`   //
	TableComment        string       `json:"tableComment" gorm:"type:varchar(128);comment:TableComment"`               //
	ClassName           string       `json:"className" gorm:"type:varchar(128);comment:ClassName"`                     //
	TplCategory         string       `json:"tplCategory" gorm:"type:varchar(128);comment:TplCategory"`                 //
	PackageName         string       `json:"packageName" gorm:"type:varchar(128);comment:PackageName"`                 //
	ModuleName          string       `json:"moduleName" gorm:"type:varchar(128);comment:ModuleName"`                   //
	ModuleFrontName     string       `json:"moduleFrontName" gorm:"type:varchar(255);comment:前端文件名"`                   //前端文件名
	BusinessName        string       `json:"businessName" gorm:"type:varchar(255);comment:BusinessName"`               //
	FunctionName        string       `json:"functionName" gorm:"type:varchar(255);comment:FunctionName"`               //
	FunctionAuthor      string       `json:"functionAuthor" gorm:"type:varchar(255);comment:FunctionAuthor"`           //
	PkColumn            string       `json:"pkColumn" gorm:"type:varchar(255);comment:PkColumn"`                       //
	PkGoField           string       `json:"pkGoField" gorm:"type:varchar(255);comment:PkGoField"`                     //
	PkJsonField         string       `json:"pkJsonField" gorm:"type:varchar(255);comment:PkJsonField"`                 //
	Options             string       `json:"options" gorm:"type:varchar(255);comment:Options"`                         //
	TreeCode            string       `json:"treeCode" gorm:"type:varchar(255);comment:TreeCode"`                       //
	TreeParentCode      string       `json:"treeParentCode" gorm:"type:varchar(255);comment:TreeParentCode"`           //
	TreeName            string       `json:"treeName" gorm:"type:varchar(255);comment:TreeName"`                       //
	Tree                int          `json:"tree" gorm:"type:tinyint(1);comment:Tree"`                                 //
	Crud                bool         `json:"crud" gorm:"type:tinyint(1);comment:Crud"`                                 //
	Remark              string       `json:"remark" gorm:"type:varchar(255);comment:Remark"`                           //
	IsDataScope         int          `json:"isDataScope" gorm:"type:tinyint;comment:IsDataScope"`                      //
	IsActions           int          `json:"isActions" gorm:"type:tinyint;comment:IsActions"`                          //
	IsAuth              int          `json:"isAuth" gorm:"type:tinyint;comment:IsAuth"`                                //
	IsLogicalDelete     string       `json:"isLogicalDelete" gorm:"type:varchar(1);comment:IsLogicalDelete"`           //
	LogicalDelete       bool         `json:"logicalDelete" gorm:"type:tinyint(1);comment:LogicalDelete"`               //
	LogicalDeleteColumn string       `json:"logicalDeleteColumn" gorm:"type:varchar(128);comment:LogicalDeleteColumn"` //
	CreatedAt           time.Time    `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                           //创建时间
	UpdatedAt           time.Time    `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                         //最后更新时间
	CreateBy            int          `json:"createBy" gorm:"type:int unsigned;comment:创建者"`                            //创建者
	UpdateBy            int          `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                            //更新者
	Columns             []GenColumns `gorm:"-" json:"columns"`
	ApiRoot             string       `gorm:"-" json:"apiRoot"`
	MLTBName            string       `gorm:"-" json:"-"` //表名称
}

func (GenTables) TableName() string {
	return "gen_tables"
}
