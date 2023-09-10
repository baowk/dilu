package tools

import (
	//common "go-admin/common/models"
	"strings"

	"github.com/baowk/dilu-core/core/base"
	"gorm.io/gorm"
)

// 代码生成表
type GenTable struct {
	TableId             int    `gorm:"primaryKey;autoIncrement" json:"tableId"`        //表编码
	DBName              string `gorm:"column:db_name;size:64;" json:"dbName"`          //库名
	TBName              string `gorm:"column:table_name;size:128;" json:"tableName"`   //表名称
	MLTBName            string `gorm:"-" json:"-"`                                     //表名称
	TableComment        string `gorm:"size:128;" json:"tableComment"`                  //表备注
	ClassName           string `gorm:"size:128;" json:"className"`                     //类名
	TplCategory         string `gorm:"size:128;" json:"tplCategory"`                   //
	PackageName         string `gorm:"size:128;" json:"packageName"`                   //包名
	ModuleName          string `gorm:"size:128;" json:"moduleName"`                    //go文件名
	ModuleFrontName     string `gorm:"size:255;comment:前端文件名;" json:"moduleFrontName"` //前端文件名
	BusinessName        string `gorm:"size:255;" json:"businessName"`                  //
	FunctionName        string `gorm:"size:255;" json:"functionName"`                  //功能名称
	FunctionAuthor      string `gorm:"size:255;" json:"functionAuthor"`                //功能作者
	PkColumn            string `gorm:"size:255;" json:"pkColumn"`
	PkGoField           string `gorm:"size:255;" json:"pkGoField"`
	PkJsonField         string `gorm:"size:255;" json:"pkJsonField"`
	Options             string `gorm:"size:255;" json:"options"`
	TreeCode            string `gorm:"size:255;" json:"treeCode"`
	TreeParentCode      string `gorm:"size:255;" json:"treeParentCode"`
	TreeName            string `gorm:"size:255;" json:"treeName"`
	Tree                bool   `gorm:"size:1;default:0;" json:"tree"`
	Crud                bool   `gorm:"size:1;default:1;" json:"crud"`
	Remark              string `gorm:"size:255;" json:"remark"`
	IsDataScope         int    `gorm:"size:1;" json:"isDataScope"`
	IsActions           int    `gorm:"size:1;" json:"isActions"`
	IsAuth              int    `gorm:"size:1;" json:"isAuth"`
	IsLogicalDelete     string `gorm:"size:1;" json:"isLogicalDelete"`
	LogicalDelete       bool   `gorm:"size:1;" json:"logicalDelete"`
	LogicalDeleteColumn string `gorm:"size:128;" json:"logicalDeleteColumn"`
	base.ModelTime
	base.ControlBy
	DataScope string      `gorm:"-" json:"dataScope"`
	Params    Params      `gorm:"-" json:"params"`
	Columns   []GenColumn `gorm:"-" json:"columns"`
	ApiRoot   string      `gorm:"-" json:"apiRoot"`

	//models.BaseModel
}

func (GenTable) TableName() string {
	return "gen_tables"
}

type Params struct {
	TreeCode       string `gorm:"-" json:"treeCode"`
	TreeParentCode string `gorm:"-" json:"treeParentCode"`
	TreeName       string `gorm:"-" json:"treeName"`
}

func (e *GenTable) GetPage(tx *gorm.DB, pageSize int, pageIndex int) ([]GenTable, int64, error) {
	var doc []GenTable

	table := tx.Table("gen_tables")

	if e.TBName != "" {
		table = table.Where("table_name = ?", e.TBName)
	}
	if e.TableComment != "" {
		table = table.Where("table_comment = ?", e.TableComment)
	}

	var count int64

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	//table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, count, nil
}

func (e *GenTable) Get(tx *gorm.DB, exclude bool) (GenTable, error) {
	var doc GenTable
	var err error
	table := tx.Table("gen_tables")

	if e.TBName != "" {
		table = table.Where("table_name = ?", e.TBName)
	}
	if e.TableId != 0 {
		table = table.Where("table_id = ?", e.TableId)
	}
	if e.TableComment != "" {
		table = table.Where("table_comment = ?", e.TableComment)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	var col GenColumn
	col.TableId = doc.TableId
	if doc.Columns, err = col.GetList(tx, exclude); err != nil {
		return doc, err
	}

	return doc, nil
}

func (e *GenTable) GetTree(tx *gorm.DB) ([]GenTable, error) {
	var doc []GenTable
	var err error
	table := tx.Table("gen_tables")

	if e.TBName != "" {
		table = table.Where("table_name = ?", e.TBName)
	}
	if e.TableId != 0 {
		table = table.Where("table_id = ?", e.TableId)
	}
	if e.TableComment != "" {
		table = table.Where("table_comment = ?", e.TableComment)
	}

	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	for i := 0; i < len(doc); i++ {
		var col GenColumn
		//col.FkCol = append(col.FkCol, GenColumn{ColumnId: 0, ColumnName: "请选择"})
		col.TableId = doc[i].TableId
		if doc[i].Columns, err = col.GetList(tx, false); err != nil {
			return doc, err
		}

	}

	return doc, nil
}

func (e *GenTable) Create(tx *gorm.DB) (GenTable, error) {
	var doc GenTable
	e.CreateBy = 0
	result := tx.Table("gen_tables").Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	for i := 0; i < len(e.Columns); i++ {
		e.Columns[i].TableId = doc.TableId

		_, _ = e.Columns[i].Create(tx)
	}

	return doc, nil
}

func (e *GenTable) Update(tx *gorm.DB) (update GenTable, err error) {
	//if err = orm.Eloquent.Table("gen_tables").First(&update, e.TableId).Error; err != nil {
	//	return
	//}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	e.UpdateBy = 0
	if err = tx.Table("gen_tables").Where("table_id = ?", e.TableId).Updates(&e).Error; err != nil {
		return
	}

	tableNames := make([]string, 0)
	for i := range e.Columns {
		if e.Columns[i].FkTableName != "" {
			tableNames = append(tableNames, e.Columns[i].FkTableName)
		}
	}

	tables := make([]GenTable, 0)
	tableMap := make(map[string]*GenTable)
	if len(tableNames) > 0 {
		if err = tx.Table("gen_tables").Where("table_name in (?)", tableNames).Find(&tables).Error; err != nil {
			return
		}
		for i := range tables {
			tableMap[tables[i].TBName] = &tables[i]
		}
	}

	for i := 0; i < len(e.Columns); i++ {
		if e.Columns[i].FkTableName != "" {
			t, ok := tableMap[e.Columns[i].FkTableName]
			if ok {
				e.Columns[i].FkTableNameClass = t.ClassName
				t.MLTBName = strings.Replace(t.TBName, "_", "-", -1)
				e.Columns[i].FkTableNamePackage = t.MLTBName
			} else {
				tableNameList := strings.Split(e.Columns[i].FkTableName, "_")
				e.Columns[i].FkTableNameClass = ""
				//e.Columns[i].FkTableNamePackage = ""
				for a := 0; a < len(tableNameList); a++ {
					strStart := string([]byte(tableNameList[a])[:1])
					strEnd := string([]byte(tableNameList[a])[1:])
					e.Columns[i].FkTableNameClass += strings.ToUpper(strStart) + strEnd
					//e.Columns[i].FkTableNamePackage += strings.ToLower(strStart) + strings.ToLower(strEnd)
				}
			}
		}
		_, _ = e.Columns[i].Update(tx)
	}
	return
}

func (e *GenTable) Delete(db *gorm.DB) (success bool, err error) {
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	if err = tx.Table("gen_tables").Delete(GenTable{}, "table_id = ?", e.TableId).Error; err != nil {
		success = false
		return
	}
	if err = tx.Table("sys_columns").Delete(GenColumn{}, "table_id = ?", e.TableId).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}

func (e *GenTable) BatchDelete(tx *gorm.DB, id []int) (Result bool, err error) {
	if err = tx.Unscoped().Table(e.TableName()).Where(" table_id in (?)", id).Delete(&GenColumn{}).Error; err != nil {
		return
	}
	Result = true
	return
}
