package tools

import (
	"dilu/modules/tools/models/tools"
	"strconv"

	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetDBColumnList 分页列表数据
// @Summary 分页列表数据 / page list data
// @Description 数据库表列分页列表 / database table column page list
// @Tags 工具 / 生成工具
// @Param dbName query string true "dbname / 数据库"
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {object} base.Resp "{"code": 200, "data": [...]}"
// @Router /api/tools/db/columns/page [get]
func (e *Gen) GetDBColumnList(c *gin.Context) {

	var data tools.DBColumns
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, err = strconv.Atoi(size)
		if err != nil {
			core.Log.Error("", zap.Error(err))
		}
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, err = strconv.Atoi(index)
		if err != nil {
			core.Log.Error("", zap.Error(err))
		}
	}

	var dbname = c.Request.FormValue("dbName")
	db, _, sdbn := GetDb(dbname)

	data.TableName = c.Request.FormValue("tableName")
	result, count, err := data.GetPage(db, pageSize, pageIndex, sdbn)
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	e.Page(c, result, count, pageIndex, pageSize)
}
