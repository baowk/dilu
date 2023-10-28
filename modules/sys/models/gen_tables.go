package models

import (
	"time"

	"gorm.io/gorm"
)

// GenTables
type GenTables struct {
	TableId             int            `json:"tableId" gorm:"type:bigint;primaryKey;autoIncrement;comment:主键"`           //主键
	DbName              string         `json:"dbName" gorm:"type:varchar(64);comment:DbName"`                            //
	ATableName          string         `json:"tableName" gorm:"type:varchar(128);comment:TableName"`                     //
	TableComment        string         `json:"tableComment" gorm:"type:varchar(128);comment:TableComment"`               //
	ClassName           string         `json:"className" gorm:"type:varchar(128);comment:ClassName"`                     //
	TplCategory         string         `json:"tplCategory" gorm:"type:varchar(128);comment:TplCategory"`                 //
	PackageName         string         `json:"packageName" gorm:"type:varchar(128);comment:PackageName"`                 //
	ModuleName          string         `json:"moduleName" gorm:"type:varchar(128);comment:ModuleName"`                   //
	ModuleFrontName     string         `json:"moduleFrontName" gorm:"type:varchar(255);comment:前端文件名"`                   //前端文件名
	BusinessName        string         `json:"businessName" gorm:"type:varchar(255);comment:BusinessName"`               //
	FunctionName        string         `json:"functionName" gorm:"type:varchar(255);comment:FunctionName"`               //
	FunctionAuthor      string         `json:"functionAuthor" gorm:"type:varchar(255);comment:FunctionAuthor"`           //
	PkColumn            string         `json:"pkColumn" gorm:"type:varchar(255);comment:PkColumn"`                       //
	PkGoField           string         `json:"pkGoField" gorm:"type:varchar(255);comment:PkGoField"`                     //
	PkJsonField         string         `json:"pkJsonField" gorm:"type:varchar(255);comment:PkJsonField"`                 //
	Options             string         `json:"options" gorm:"type:varchar(255);comment:Options"`                         //
	TreeCode            string         `json:"treeCode" gorm:"type:varchar(255);comment:TreeCode"`                       //
	TreeParentCode      string         `json:"treeParentCode" gorm:"type:varchar(255);comment:TreeParentCode"`           //
	TreeName            string         `json:"treeName" gorm:"type:varchar(255);comment:TreeName"`                       //
	Tree                int            `json:"tree" gorm:"type:tinyint(1);comment:Tree"`                                 //
	Crud                int            `json:"crud" gorm:"type:tinyint(1);comment:Crud"`                                 //
	Remark              string         `json:"remark" gorm:"type:varchar(255);comment:Remark"`                           //
	IsDataScope         int            `json:"isDataScope" gorm:"type:tinyint;comment:IsDataScope"`                      //
	IsActions           int            `json:"isActions" gorm:"type:tinyint;comment:IsActions"`                          //
	IsAuth              int            `json:"isAuth" gorm:"type:tinyint;comment:IsAuth"`                                //
	IsLogicalDelete     string         `json:"isLogicalDelete" gorm:"type:varchar(1);comment:IsLogicalDelete"`           //
	LogicalDelete       int            `json:"logicalDelete" gorm:"type:tinyint(1);comment:LogicalDelete"`               //
	LogicalDeleteColumn string         `json:"logicalDeleteColumn" gorm:"type:varchar(128);comment:LogicalDeleteColumn"` //
	CreatedAt           time.Time      `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                           //创建时间
	UpdatedAt           time.Time      `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                         //最后更新时间
	DeletedAt           gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                              //删除时间
	CreateBy            int            `json:"createBy" gorm:"type:int unsigned;comment:创建者"`                            //创建者
	UpdateBy            int            `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                            //更新者
}

func (GenTables) TableName() string {
	return "gen_tables"
}

func NewGenTables() *GenTables {
	return &GenTables{}
}

func (e *GenTables) SetTableId(tableId int) *GenTables {
	e.TableId = tableId
	return e
}
func (e *GenTables) SetDbName(dbName string) *GenTables {
	e.DbName = dbName
	return e
}
func (e *GenTables) SetTableName(tableName string) *GenTables {
	e.ATableName = tableName
	return e
}
func (e *GenTables) SetTableComment(tableComment string) *GenTables {
	e.TableComment = tableComment
	return e
}
func (e *GenTables) SetClassName(className string) *GenTables {
	e.ClassName = className
	return e
}
func (e *GenTables) SetTplCategory(tplCategory string) *GenTables {
	e.TplCategory = tplCategory
	return e
}
func (e *GenTables) SetPackageName(packageName string) *GenTables {
	e.PackageName = packageName
	return e
}
func (e *GenTables) SetModuleName(moduleName string) *GenTables {
	e.ModuleName = moduleName
	return e
}
func (e *GenTables) SetModuleFrontName(moduleFrontName string) *GenTables {
	e.ModuleFrontName = moduleFrontName
	return e
}
func (e *GenTables) SetBusinessName(businessName string) *GenTables {
	e.BusinessName = businessName
	return e
}
func (e *GenTables) SetFunctionName(functionName string) *GenTables {
	e.FunctionName = functionName
	return e
}
func (e *GenTables) SetFunctionAuthor(functionAuthor string) *GenTables {
	e.FunctionAuthor = functionAuthor
	return e
}
func (e *GenTables) SetPkColumn(pkColumn string) *GenTables {
	e.PkColumn = pkColumn
	return e
}
func (e *GenTables) SetPkGoField(pkGoField string) *GenTables {
	e.PkGoField = pkGoField
	return e
}
func (e *GenTables) SetPkJsonField(pkJsonField string) *GenTables {
	e.PkJsonField = pkJsonField
	return e
}
func (e *GenTables) SetOptions(options string) *GenTables {
	e.Options = options
	return e
}
func (e *GenTables) SetTreeCode(treeCode string) *GenTables {
	e.TreeCode = treeCode
	return e
}
func (e *GenTables) SetTreeParentCode(treeParentCode string) *GenTables {
	e.TreeParentCode = treeParentCode
	return e
}
func (e *GenTables) SetTreeName(treeName string) *GenTables {
	e.TreeName = treeName
	return e
}
func (e *GenTables) SetTree(tree int) *GenTables {
	e.Tree = tree
	return e
}
func (e *GenTables) SetCrud(crud int) *GenTables {
	e.Crud = crud
	return e
}
func (e *GenTables) SetRemark(remark string) *GenTables {
	e.Remark = remark
	return e
}
func (e *GenTables) SetIsDataScope(isDataScope int) *GenTables {
	e.IsDataScope = isDataScope
	return e
}
func (e *GenTables) SetIsActions(isActions int) *GenTables {
	e.IsActions = isActions
	return e
}
func (e *GenTables) SetIsAuth(isAuth int) *GenTables {
	e.IsAuth = isAuth
	return e
}
func (e *GenTables) SetIsLogicalDelete(isLogicalDelete string) *GenTables {
	e.IsLogicalDelete = isLogicalDelete
	return e
}
func (e *GenTables) SetLogicalDelete(logicalDelete int) *GenTables {
	e.LogicalDelete = logicalDelete
	return e
}
func (e *GenTables) SetLogicalDeleteColumn(logicalDeleteColumn string) *GenTables {
	e.LogicalDeleteColumn = logicalDeleteColumn
	return e
}
func (e *GenTables) SetCreatedAt(createdAt time.Time) *GenTables {
	e.CreatedAt = createdAt
	return e
}
func (e *GenTables) SetUpdatedAt(updatedAt time.Time) *GenTables {
	e.UpdatedAt = updatedAt
	return e
}
func (e *GenTables) SetCreateBy(createBy int) *GenTables {
	e.CreateBy = createBy
	return e
}
func (e *GenTables) SetUpdateBy(updateBy int) *GenTables {
	e.UpdateBy = updateBy
	return e
}
