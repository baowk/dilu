package tools

// import (
// 	"errors"

// 	"github.com/baowk/dilu-core/core"
// 	"github.com/gin-gonic/gin"

// 	"dilu/modules/tools/models/tools"
// 	"dilu/modules/tools/service/dto"
// )

// // GetDBTableList 分页列表数据
// // @Summary 分页列表数据 / page list data
// // @Description 数据库表分页列表 / database table page list
// // @Tags 工具 / 生成工具
// // @Param data body dto.DBReq true "body"
// // @Success 200 {object} base.Resp "{"code": 200, "data": [...]}"
// // @Router /api/tools/db/tables/page [post]
// func (e *Gen) GetDBTableList(c *gin.Context) {
// 	//var res base.Resp
// 	var data tools.DBTables
// 	var err error
// 	var req dto.DBReq
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}

// 	if core.Cfg.DBCfg.Driver == "sqlite3" || core.Cfg.DBCfg.Driver == "postgres" {
// 		err = errors.New("对不起，sqlite3 或 postgres 不支持代码生成！")
// 		e.Error(c, err)
// 		return
// 	}
// 	if req.TableName != "" {
// 		data.TableName = req.TableName
// 	}

// 	var dbname = req.DBName
// 	db, mdbn, sdbn := GetDb(dbname)

// 	result, total, err := data.GetPage(db, req.GetSize(), req.GetPage(), sdbn, mdbn)
// 	if err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Page(c, result, total, req.Page, req.PageSize)
// }
