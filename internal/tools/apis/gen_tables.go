package apis

import (
	"dilu/common/config"
	cons "dilu/common/consts"
	"dilu/internal/tools/repository/model"
	"dilu/internal/tools/service"
	"dilu/internal/tools/service/dto"

	"errors"
	"fmt"
	"strings"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type GenTablesApi struct {
	base.BaseApi
}

var ApiGenTables = GenTablesApi{}

// GetDBTableList 分页列表数据
// @Summary 分页列表数据 / page list data
// @Description 数据库表分页列表 / database table page list
// @Tags 工具 / 生成工具
// @Param data body dto.DBReq true "body"
// @Success 200 {object} base.Resp "{"code": 200, "data": [...]}"
// @Router /api/v1/tools/gen/db/tables [post]
func (e *GenTablesApi) GetDBTableList(c *gin.Context) {
	//var res base.Resp
	var data model.DBTables
	var err error
	var req dto.DBReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	if config.Get().DBCfg.Driver == "postgres" || config.Get().DBCfg.Driver == "pgsql" {
		err = errors.New("对不起，postgres 暂不支持数据库表分页读取！")
		e.Error(c, err)
		return
	}
	if req.TableName != "" {
		data.TableName = req.TableName
	}

	var dbname = req.DBName
	db, mdbn, sdbn, _ := service.GetDb(dbname)

	result, total, err := data.GetPage(db, req.GetSize(), req.GetPage(), sdbn, mdbn)
	if err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, result, total, req.GetPage(), req.GetSize())
}

// QueryPage 获取GenTables列表
// @Summary 获取GenTables列表
// @Tags 工具 / 生成工具
// @Accept application/json
// @Product application/json
// @Param data body dto.GenTablesGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]model.GenTables}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tools/gen/page [post]
func (e *GenTablesApi) QueryPage(c *gin.Context) {
	var req dto.GenTablesGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]model.GenTables, 10)
	var total int64
	if err := service.SerGenTables.Page(&req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Del 删除GenTables
// @Summary 删除GenTables
// @Tags 工具 / 生成工具
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=model.GenTables} "{"code": 200, "data": [...]}"
// @Router /api/v1/tools/gen/del [post]
func (e *GenTablesApi) Del(c *gin.Context) {
	if !config.Get().Gen.Enable {
		e.Error(c, errors.New("当前生成表已关闭"))
		return
	}
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerGenTables.Del(req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// GetDBS
// @Summary 获取配置的数据库
// @Description 获取配置的数据库
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=[]dto.DbOption} "{"code": 200, "data": [...]}"
// @Router /api/v1/tools/gen/dbs [post]
func (e *GenTablesApi) GetDBS(c *gin.Context) {
	e.Ok(c, service.SerGenTables.GetDbs())
}

// Insert
// @Summary 添加表结构
// @Description 添加表结构
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param data body dto.ImpTablesReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/tools/gen/add [post]
func (e *GenTablesApi) Insert(c *gin.Context) {
	if !config.Get().Gen.Enable {
		e.Error(c, errors.New("添加表结构已关闭"))
		return
	}
	var req dto.ImpTablesReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	for _, tableName := range req.Tables {
		fmt.Println(tableName)
		data, err := service.SerGenTables.GenTableInit(req.DbName, tableName, false)
		if err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
			e.Error(c, err)
			return
		}
		err = service.SerGenTables.Create(&data)
		if err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
			e.Error(c, err)
			return
		}
	}
	e.Ok(c)

}

// Update
// @Summary 修改表结构
// @Description 修改表结构
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param data body model.GenTables true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/tools/gen/update [POST]
func (e *GenTablesApi) Update(c *gin.Context) {
	if !config.Get().Gen.Enable {
		e.Error(c, errors.New("修改表结构已关闭"))
		return
	}
	var data model.GenTables
	if err := c.ShouldBind(&data); err != nil {
		e.Error(c, err)
		return
	}
	data.UpdateBy = 0
	err := service.SerGenTables.Update(&data)
	if err != nil {
		core.GetApp().GetLogger().Error("Gen", "err", err)
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// GenCode
// @Summary 生成代码
// @Description 生成代码
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param data body dto.GenCodeReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tools/gen/code [post]
func (e *GenTablesApi) GenCode(c *gin.Context) {
	if !config.Get().Gen.Enable {
		e.Error(c, errors.New("生成代码已关闭"))
		return
	}
	table := model.GenTables{}

	var req dto.GenCodeReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	table.TableId = req.TableId

	tab, _ := service.SerGenTables.Get(nil, false, req.TableId)
	tab.ApiRoot = cons.ApiRoot

	for i, v := range tab.Columns {
		tab.Columns[i].TsType = TypeGo2Ts(v.GoType)
	}
	service.SerGenTables.NOMethodsGen(tab, req.Force)
	e.Ok(c, "Code generated successfully！")
}

func TypeGo2Ts(t string) string {
	if strings.Contains(t, "int") {
		return "number"
	} else if strings.Contains(t, "float") {
		return "number"
	} else if strings.Contains(t, "time") {
		return "Date"
	} else if strings.Contains(t, "bool") {
		return "boolean"
	} else {
		return t
	}
}
