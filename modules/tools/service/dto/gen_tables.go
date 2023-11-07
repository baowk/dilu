package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type GenTablesGetPageReq struct {
	base.ReqPage `search:"-"`
	DbName       string `json:"dbName"`
	TableName    string `json:"tableName"`
}

// GenTables
type GenTablesDto struct {
	TableId             int    `json:"tableId"`             //主键
	DbName              string `json:"dbName"`              //
	TableName           string `json:"tableName"`           //
	TableComment        string `json:"tableComment"`        //
	ClassName           string `json:"className"`           //
	TplCategory         string `json:"tplCategory"`         //
	PackageName         string `json:"packageName"`         //
	ModuleName          string `json:"moduleName"`          //
	ModuleFrontName     string `json:"moduleFrontName"`     //前端文件名
	BusinessName        string `json:"businessName"`        //
	FunctionName        string `json:"functionName"`        //
	FunctionAuthor      string `json:"functionAuthor"`      //
	PkColumn            string `json:"pkColumn"`            //
	PkGoField           string `json:"pkGoField"`           //
	PkJsonField         string `json:"pkJsonField"`         //
	Options             string `json:"options"`             //
	TreeCode            string `json:"treeCode"`            //
	TreeParentCode      string `json:"treeParentCode"`      //
	TreeName            string `json:"treeName"`            //
	Tree                int    `json:"tree"`                //
	Crud                int    `json:"crud"`                //
	Remark              string `json:"remark"`              //
	IsDataScope         int    `json:"isDataScope"`         //
	IsActions           int    `json:"isActions"`           //
	IsAuth              int    `json:"isAuth"`              //
	IsLogicalDelete     string `json:"isLogicalDelete"`     //
	LogicalDelete       int    `json:"logicalDelete"`       //
	LogicalDeleteColumn string `json:"logicalDeleteColumn"` //
}

type DbOption struct {
	Lable string `json:"label"`
	Value string `json:"value"`
}
