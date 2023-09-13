package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysMenuApi struct {
	base.BaseApi
}

var SysMenuA = SysMenuApi{}

// QueryPage 获取接口列表列表
// @Summary Page接口
// @Tags SysMenu
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysMenu}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-menu/page [post]
// @Security Bearer
func (e *SysMenuApi) QueryPage(c *gin.Context) {
	var req dto.SysApiGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysMenu, 10)
	var total int64

	var model models.SysMenu
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysMenu.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取接口列表
// @Summary 获取接口列表
// @Tags SysMenu
// @Accept application/json
// @Product application/json
// @Param data body dto.SysMenuGetReq true "body"
// @Success 200 {object} base.Resp{data=models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-menu/get [post]
// @Security Bearer
func (e *SysMenuApi) Get(c *gin.Context) {
	var req dto.SysMenuGetReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMenu
	if _, err := service.SerSysMenu.Get(&req, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建接口列表
// @Summary 创建接口列表
// @Tags SysMenu
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-menu/create [post]
// @Security Bearer
func (e *SysMenuApi) Create(c *gin.Context) {
	var req dto.SysApiDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMenu
	copier.Copy(&data, req)
	if err := service.SerSysMenu.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新接口列表
// @Summary 更新接口列表
// @Tags SysMenu
// @Accept application/json
// @Product application/json
// @Param data body dto.SysApiDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-menu/update [post]
// @Security Bearer
func (e *SysMenuApi) Update(c *gin.Context) {
	var req dto.SysApiDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMenu
	copier.Copy(&data, req)
	if err := service.SerSysMenu.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除接口列表
// @Summary 删除接口列表
// @Tags SysMenu
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-menu/del [post]
// @Security Bearer
func (e *SysMenuApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysMenu.DelIds(&models.SysMenu{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
