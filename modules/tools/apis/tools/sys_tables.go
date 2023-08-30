package tools

import (
	"strconv"
	"strings"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"dilu/modules/tools/models/tools"
)

type SysTable struct {
	base.BaseApi
}

// GetPage 分页列表数据
// @Summary 分页列表数据
// @Description 生成表分页列表
// @Tags 工具 / 生成工具
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/tables/page [get]
func (e SysTable) GetPage(c *gin.Context) {
	var req base.ReqPage

	c.ShouldBind(&req)

	var dbname string

	var data tools.SysTables
	var err error
	var pageSize = 10
	var pageIndex = 1

	data.TBName = c.Request.FormValue("tableName")
	data.TableComment = c.Request.FormValue("tableComment")
	db, _, _ := GetDb(dbname)
	result, count, err := data.GetPage(db, req.Size, req.Page)
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	e.Page(c, result, count, pageIndex, pageSize)
}

// Get
// @Summary 获取配置
// @Description 获取JSON
// @Tags 工具 / 生成工具
// @Param configKey path int true "configKey"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/tables/info/{tableId} [get]
// @Security Bearer
func (e SysTable) Get(c *gin.Context) {
	var dbname string
	var data tools.SysTables
	data.TableId, _ = strconv.Atoi(c.Param("tableId"))
	db, _, _ := GetDb(dbname)
	result, err := data.Get(db, true)
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}

	mp := make(map[string]interface{})
	mp["list"] = result.Columns
	mp["info"] = result
	e.Ok(c, mp)
}

func (e SysTable) GetSysTablesInfo(c *gin.Context) {
	var dbname string
	var data tools.SysTables
	if c.Request.FormValue("tableName") != "" {
		data.TBName = c.Request.FormValue("tableName")
	}
	db, _, _ := GetDb(dbname)
	result, err := data.Get(db, true)
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}

	mp := make(map[string]interface{})
	mp["list"] = result.Columns
	mp["info"] = result
	e.Ok(c, mp)
}

func (e SysTable) GetSysTablesTree(c *gin.Context) {
	var dbname string
	var data tools.SysTables
	db, _, _ := GetDb(dbname)
	result, err := data.GetTree(db)
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}

	e.Ok(c, result)
}

// Insert
// @Summary 添加表结构
// @Description 添加表结构
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param tables query string false "tableName / 数据表名称"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/sys/tables/info [post]
// @Security Bearer
func (e SysTable) Insert(c *gin.Context) {
	var dbname string
	tablesList := strings.Split(c.Request.FormValue("tables"), ",")
	db, _, _ := GetDb(dbname)
	for i := 0; i < len(tablesList); i++ {

		data, err := genTableInit(db, tablesList, i, c)
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}

		_, err = data.Create(db)
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
	}
	e.Ok(c)

}

func genTableInit(tx *gorm.DB, tablesList []string, i int, c *gin.Context) (tools.SysTables, error) {
	var dbname string
	var data tools.SysTables
	var dbTable tools.DBTables
	var dbColumn tools.DBColumns
	data.TBName = tablesList[i]
	data.CreateBy = 0

	dbTable.TableName = data.TBName
	tx, _, sdbn := GetDb(dbname)
	dbtable, err := dbTable.Get(tx, sdbn)
	if err != nil {
		return data, err
	}

	dbColumn.TableName = data.TBName
	tablenamelist := strings.Split(dbColumn.TableName, "_")
	for i := 0; i < len(tablenamelist); i++ {
		strStart := string([]byte(tablenamelist[i])[:1])
		strend := string([]byte(tablenamelist[i])[1:])
		// 大驼峰表名 结构体使用
		data.ClassName += strings.ToUpper(strStart) + strend
		// 小驼峰表名 js函数名和权限标识使用
		if i == 0 {
			data.BusinessName += strings.ToLower(strStart) + strend
		} else {
			data.BusinessName += strings.ToUpper(strStart) + strend
		}
		//data.PackageName += strings.ToLower(strStart) + strings.ToLower(strend)
		//data.ModuleName += strings.ToLower(strStart) + strings.ToLower(strend)
	}
	//data.ModuleFrontName = strings.ReplaceAll(data.ModuleName, "_", "-")
	data.PackageName = "admin"
	data.TplCategory = "crud"
	data.Crud = true
	// 中横线表名称，接口路径、前端文件夹名称和js名称使用
	data.ModuleName = strings.Replace(data.TBName, "_", "-", -1)
	dbcolumn, err := dbColumn.GetList(tx, sdbn)
	data.CreateBy = 0
	data.TableComment = dbtable.TableComment
	if dbtable.TableComment == "" {
		data.TableComment = data.ClassName
	}

	data.FunctionName = data.TableComment
	//data.BusinessName = data.ModuleName
	data.IsLogicalDelete = "1"
	data.LogicalDelete = true
	data.LogicalDeleteColumn = "is_del"
	data.IsActions = 2
	data.IsDataScope = 1
	data.IsAuth = 1

	data.FunctionAuthor = "wenjianzhang"
	for i := 0; i < len(dbcolumn); i++ {
		var column tools.SysColumns
		column.ColumnComment = dbcolumn[i].ColumnComment
		column.ColumnName = dbcolumn[i].ColumnName
		column.ColumnType = dbcolumn[i].ColumnType
		column.Sort = i + 1
		column.Insert = true
		column.IsInsert = "1"
		column.QueryType = "EQ"
		column.IsPk = "0"

		namelist := strings.Split(dbcolumn[i].ColumnName, "_")
		for i := 0; i < len(namelist); i++ {
			strStart := string([]byte(namelist[i])[:1])
			strend := string([]byte(namelist[i])[1:])
			column.GoField += strings.ToUpper(strStart) + strend
			if i == 0 {
				column.JsonField = strings.ToLower(strStart) + strend
			} else {
				column.JsonField += strings.ToUpper(strStart) + strend
			}
		}
		if strings.Contains(dbcolumn[i].ColumnKey, "PR") {
			column.IsPk = "1"
			column.Pk = true
			data.PkColumn = dbcolumn[i].ColumnName
			//column.GoField = strings.ToUpper(column.GoField)
			//column.JsonField = strings.ToUpper(column.JsonField)
			data.PkGoField = column.GoField
			data.PkJsonField = column.JsonField
		}
		column.IsRequired = "0"
		if strings.Contains(dbcolumn[i].IsNullable, "NO") {
			column.IsRequired = "1"
			column.Required = true
		}

		if strings.Contains(dbcolumn[i].ColumnType, "int") {
			if strings.Contains(dbcolumn[i].ColumnKey, "PR") {
				column.GoType = "int"
			} else {
				column.GoType = "string"
			}
			column.HtmlType = "input"
		} else if strings.Contains(dbcolumn[i].ColumnType, "timestamp") {
			column.GoType = "time.Time"
			column.HtmlType = "datetime"
		} else if strings.Contains(dbcolumn[i].ColumnType, "datetime") {
			column.GoType = "time.Time"
			column.HtmlType = "datetime"
		} else {
			column.GoType = "string"
			column.HtmlType = "input"
		}

		data.Columns = append(data.Columns, column)
	}
	return data, err
}

// Update
// @Summary 修改表结构
// @Description 修改表结构
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param data body tools.SysTables true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/sys/tables/info [put]
// @Security Bearer
func (e SysTable) Update(c *gin.Context) {
	var dbname string
	var data tools.SysTables
	err := c.ShouldBind(&data)

	data.UpdateBy = 0
	db, _, _ := GetDb(dbname)
	result, err := data.Update(db)
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	e.Ok(c, result)
}

// Delete
// @Summary 删除表结构
// @Description 删除表结构
// @Tags 工具 / 生成工具
// @Param tableId path int true "tableId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/sys/tables/info/{tableId} [delete]
func (e SysTable) Delete(c *gin.Context) {
	var dbname string
	var req base.ReqIds
	c.ShouldBind(&req)
	var data tools.SysTables
	db, _, _ := GetDb(dbname)
	_, err := data.BatchDelete(db, req.Ids)
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
