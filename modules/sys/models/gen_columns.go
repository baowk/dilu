package models

import ("gorm.io/gorm"
    "time"
    
)

//GenColumns
type GenColumns struct {
    
    ColumnId int `json:"columnId" gorm:"type:bigint;primaryKey;autoIncrement;comment:主键"` //主键
    TableId int `json:"tableId" gorm:"type:bigint;comment:TableId"` // 
    ColumnName string `json:"columnName" gorm:"type:varchar(128);comment:ColumnName"` // 
    ColumnComment string `json:"columnComment" gorm:"type:varchar(128);comment:ColumnComment"` // 
    ColumnType string `json:"columnType" gorm:"type:varchar(128);comment:ColumnType"` // 
    GoType string `json:"goType" gorm:"type:varchar(128);comment:GoType"` // 
    GoField string `json:"goField" gorm:"type:varchar(128);comment:GoField"` // 
    JsonField string `json:"jsonField" gorm:"type:varchar(128);comment:JsonField"` // 
    IsPk string `json:"isPk" gorm:"type:varchar(4);comment:IsPk"` // 
    IsIncrement string `json:"isIncrement" gorm:"type:varchar(4);comment:IsIncrement"` // 
    IsRequired string `json:"isRequired" gorm:"type:varchar(4);comment:IsRequired"` // 
    IsInsert string `json:"isInsert" gorm:"type:varchar(4);comment:IsInsert"` // 
    IsEdit string `json:"isEdit" gorm:"type:varchar(4);comment:IsEdit"` // 
    IsList string `json:"isList" gorm:"type:varchar(4);comment:IsList"` // 
    IsQuery string `json:"isQuery" gorm:"type:varchar(4);comment:IsQuery"` // 
    QueryType string `json:"queryType" gorm:"type:varchar(128);comment:QueryType"` // 
    HtmlType string `json:"htmlType" gorm:"type:varchar(128);comment:HtmlType"` // 
    DictType string `json:"dictType" gorm:"type:varchar(128);comment:DictType"` // 
    Sort int `json:"sort" gorm:"type:bigint;comment:Sort"` // 
    List string `json:"list" gorm:"type:varchar(1);comment:List"` // 
    Pk int `json:"pk" gorm:"type:tinyint(1);comment:Pk"` // 
    Required int `json:"required" gorm:"type:tinyint(1);comment:Required"` // 
    SuperColumn int `json:"superColumn" gorm:"type:tinyint(1);comment:SuperColumn"` // 
    UsableColumn int `json:"usableColumn" gorm:"type:tinyint(1);comment:UsableColumn"` // 
    Increment int `json:"increment" gorm:"type:tinyint(1);comment:Increment"` // 
    Insert int `json:"insert" gorm:"type:tinyint(1);comment:Insert"` // 
    Edit int `json:"edit" gorm:"type:tinyint(1);comment:Edit"` // 
    Query int `json:"query" gorm:"type:tinyint(1);comment:Query"` // 
    Remark string `json:"remark" gorm:"type:varchar(255);comment:Remark"` // 
    FkTableName string `json:"fkTableName" gorm:"type:longtext;comment:FkTableName"` // 
    FkTableNameClass string `json:"fkTableNameClass" gorm:"type:longtext;comment:FkTableNameClass"` // 
    FkTableNamePackage string `json:"fkTableNamePackage" gorm:"type:longtext;comment:FkTableNamePackage"` // 
    FkLabelId string `json:"fkLabelId" gorm:"type:longtext;comment:FkLabelId"` // 
    FkLabelName string `json:"fkLabelName" gorm:"type:varchar(255);comment:FkLabelName"` // 
    CreateBy int `json:"createBy" gorm:"type:mediumint;comment:CreateBy"` // 
    UpdateBy int `json:"updateBy" gorm:"type:mediumint;comment:UpdateBy"` // 
    CreatedAt time.Time `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"` //创建时间 
    UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"` //最后更新时间 
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`     //删除时间
}

func (GenColumns) TableName() string {
    return "gen_columns"
}

func NewGenColumns() *GenColumns{
    return &GenColumns{}
}

func (e *GenColumns) SetColumnId(columnId int) *GenColumns {
	e.ColumnId = columnId
	return e
}
func (e *GenColumns) SetTableId(tableId int) *GenColumns {
	e.TableId = tableId
	return e
}
func (e *GenColumns) SetColumnName(columnName string) *GenColumns {
	e.ColumnName = columnName
	return e
}
func (e *GenColumns) SetColumnComment(columnComment string) *GenColumns {
	e.ColumnComment = columnComment
	return e
}
func (e *GenColumns) SetColumnType(columnType string) *GenColumns {
	e.ColumnType = columnType
	return e
}
func (e *GenColumns) SetGoType(goType string) *GenColumns {
	e.GoType = goType
	return e
}
func (e *GenColumns) SetGoField(goField string) *GenColumns {
	e.GoField = goField
	return e
}
func (e *GenColumns) SetJsonField(jsonField string) *GenColumns {
	e.JsonField = jsonField
	return e
}
func (e *GenColumns) SetIsPk(isPk string) *GenColumns {
	e.IsPk = isPk
	return e
}
func (e *GenColumns) SetIsIncrement(isIncrement string) *GenColumns {
	e.IsIncrement = isIncrement
	return e
}
func (e *GenColumns) SetIsRequired(isRequired string) *GenColumns {
	e.IsRequired = isRequired
	return e
}
func (e *GenColumns) SetIsInsert(isInsert string) *GenColumns {
	e.IsInsert = isInsert
	return e
}
func (e *GenColumns) SetIsEdit(isEdit string) *GenColumns {
	e.IsEdit = isEdit
	return e
}
func (e *GenColumns) SetIsList(isList string) *GenColumns {
	e.IsList = isList
	return e
}
func (e *GenColumns) SetIsQuery(isQuery string) *GenColumns {
	e.IsQuery = isQuery
	return e
}
func (e *GenColumns) SetQueryType(queryType string) *GenColumns {
	e.QueryType = queryType
	return e
}
func (e *GenColumns) SetHtmlType(htmlType string) *GenColumns {
	e.HtmlType = htmlType
	return e
}
func (e *GenColumns) SetDictType(dictType string) *GenColumns {
	e.DictType = dictType
	return e
}
func (e *GenColumns) SetSort(sort int) *GenColumns {
	e.Sort = sort
	return e
}
func (e *GenColumns) SetList(list string) *GenColumns {
	e.List = list
	return e
}
func (e *GenColumns) SetPk(pk int) *GenColumns {
	e.Pk = pk
	return e
}
func (e *GenColumns) SetRequired(required int) *GenColumns {
	e.Required = required
	return e
}
func (e *GenColumns) SetSuperColumn(superColumn int) *GenColumns {
	e.SuperColumn = superColumn
	return e
}
func (e *GenColumns) SetUsableColumn(usableColumn int) *GenColumns {
	e.UsableColumn = usableColumn
	return e
}
func (e *GenColumns) SetIncrement(increment int) *GenColumns {
	e.Increment = increment
	return e
}
func (e *GenColumns) SetInsert(insert int) *GenColumns {
	e.Insert = insert
	return e
}
func (e *GenColumns) SetEdit(edit int) *GenColumns {
	e.Edit = edit
	return e
}
func (e *GenColumns) SetQuery(query int) *GenColumns {
	e.Query = query
	return e
}
func (e *GenColumns) SetRemark(remark string) *GenColumns {
	e.Remark = remark
	return e
}
func (e *GenColumns) SetFkTableName(fkTableName string) *GenColumns {
	e.FkTableName = fkTableName
	return e
}
func (e *GenColumns) SetFkTableNameClass(fkTableNameClass string) *GenColumns {
	e.FkTableNameClass = fkTableNameClass
	return e
}
func (e *GenColumns) SetFkTableNamePackage(fkTableNamePackage string) *GenColumns {
	e.FkTableNamePackage = fkTableNamePackage
	return e
}
func (e *GenColumns) SetFkLabelId(fkLabelId string) *GenColumns {
	e.FkLabelId = fkLabelId
	return e
}
func (e *GenColumns) SetFkLabelName(fkLabelName string) *GenColumns {
	e.FkLabelName = fkLabelName
	return e
}
func (e *GenColumns) SetCreateBy(createBy int) *GenColumns {
	e.CreateBy = createBy
	return e
}
func (e *GenColumns) SetUpdateBy(updateBy int) *GenColumns {
	e.UpdateBy = updateBy
	return e
}
func (e *GenColumns) SetCreatedAt(createdAt time.Time) *GenColumns {
	e.CreatedAt = createdAt
	return e
}
func (e *GenColumns) SetUpdatedAt(updatedAt time.Time) *GenColumns {
	e.UpdatedAt = updatedAt
	return e
}

