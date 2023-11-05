package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type GenTablesApi struct {
	base.BaseApi
}

var ApiGenTables = GenTablesApi{}

// QueryPage 获取GenTables列表
// @Summary 获取GenTables列表
// @Tags sys-GenTables
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.GenTablesGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.GenTables}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/gen-tables/page [post]
// @Security Bearer
func (e *GenTablesApi) QueryPage(c *gin.Context) {
	var req dto.GenTablesGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.GenTables, 10)
	var total int64
	if err := service.SerGenTables.Page(&req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取GenTables
// @Summary 获取GenTables
// @Tags sys-GenTables
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.GenTables} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/gen-tables/get [post]
// @Security Bearer
func (e *GenTablesApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.GenTables
	if err := service.SerGenTables.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建GenTables
// @Summary 创建GenTables
// @Tags sys-GenTables
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.GenTablesDto true "body"
// @Success 200 {object} base.Resp{data=models.GenTables} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/gen-tables/create [post]
// @Security Bearer
func (e *GenTablesApi) Create(c *gin.Context) {
	var req dto.GenTablesDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.GenTables
	copier.Copy(&data, req)
	if err := service.SerGenTables.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新GenTables
// @Summary 更新GenTables
// @Tags sys-GenTables
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.GenTablesDto true "body"
// @Success 200 {object} base.Resp{data=models.GenTables} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/gen-tables/update [post]
// @Security Bearer
func (e *GenTablesApi) Update(c *gin.Context) {
	var req dto.GenTablesDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.GenTables
	copier.Copy(&data, req)
	if err := service.SerGenTables.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除GenTables
// @Summary 删除GenTables
// @Tags sys-GenTables
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.GenTables} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/gen-tables/del [post]
// @Security Bearer
func (e *GenTablesApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerGenTables.DelIds(&models.GenTables{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
