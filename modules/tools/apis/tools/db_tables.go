package tools

import (
	"errors"
	"strconv"

	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"

	"dilu/modules/tools/models/tools"
)

// GetDBTableList 分页列表数据
// @Summary 分页列表数据 / page list data
// @Description 数据库表分页列表 / database table page list
// @Tags 工具 / 生成工具
// @Param tableName query string false "tableName / 数据表名称"
// @Param pageSize query int false "pageSize / 页条数"
// @Param pageIndex query int false "pageIndex / 页码"
// @Success 200 {object} base.Resp "{"code": 200, "data": [...]}"
// @Router /api/v1/db/tables/page [get]
func (e *Gen) GetDBTableList(c *gin.Context) {
	//var res base.Resp
	var data tools.DBTables
	var err error
	var pageSize = 10
	var pageIndex = 1

	if core.Cfg.DBCfg.Driver == "sqlite3" || core.Cfg.DBCfg.Driver == "postgres" {
		err = errors.New("对不起，sqlite3 或 postgres 不支持代码生成！")
		e.Error(c, err)
		return
	}

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, err = strconv.Atoi(size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, err = strconv.Atoi(index)
	}

	data.TableName = c.Request.FormValue("tableName")

	var dbname string
	db, mdbn, sdbn := GetDb(dbname)

	result, count, err := data.GetPage(db, pageSize, pageIndex, sdbn, mdbn)
	if err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, result, count, pageIndex, pageSize)
}
