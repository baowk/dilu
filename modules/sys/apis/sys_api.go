package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysApiApi struct {
	base.BaseApi
}

var SysApiA = SysApiApi{}

// QueryPage 获取接口列表列表
// @Summary Page接口
// @Tags sys-SysApi
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysApi}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-api/page [post]
// @Security Bearer
func (e *SysApiApi) QueryPage(c *gin.Context) {
	var req dto.SysApiGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysApi, 10)
	var total int64

	var model models.SysApi
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysApi.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// QueryPage 获取接口列表列表
// @Summary Page接口
// @Tags sys-SysApi
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiGetPageReq true "body"
// @Success 200 {object} base.Resp{data=[]models.SysApi} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-api/all [post]
func (e *SysApiApi) Query(c *gin.Context) {
	list := make([]models.SysApi, 10)
	if err := service.SerSysApi.GetByType(3, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// Get 获取接口列表
// @Summary 获取接口列表
// @Tags sys-SysApi
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysApi} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-api/get [post]
// @Security Bearer
func (e *SysApiApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysApi
	if err := service.SerSysApi.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建接口列表
// @Summary 创建接口列表
// @Tags sys-SysApi
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiDto true "body"
// @Success 200 {object} base.Resp{data=models.SysApi} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-api/create [post]
// @Security Bearer
func (e *SysApiApi) Create(c *gin.Context) {
	var req dto.SysApiDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysApi
	copier.Copy(&data, req)
	if err := service.SerSysApi.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新接口列表
// @Summary 更新接口列表
// @Tags sys-SysApi
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiDto true "body"
// @Success 200 {object} base.Resp{data=models.SysApi} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-api/update [post]
// @Security Bearer
func (e *SysApiApi) Update(c *gin.Context) {
	var req dto.SysApiDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysApi
	copier.Copy(&data, req)
	if err := service.SerSysApi.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除接口列表
// @Summary 删除接口列表
// @Tags sys-SysApi
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysApi} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-api/del [post]
// @Security Bearer
func (e *SysApiApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysApi.DelIds(&models.SysApi{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
