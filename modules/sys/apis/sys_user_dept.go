package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysUserDeptApi struct {
	base.BaseApi
}

var ApiSysUserDept = SysUserDeptApi{}

// QueryPage 获取SysUserDept列表
// @Summary Page接口
// @Tags SysUserDept
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserDeptGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysUserDept}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-user-dept/page [post]
// @Security Bearer
func (e *SysUserDeptApi) QueryPage(c *gin.Context) {
	var req dto.SysUserDeptGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysUserDept, 10)
	var total int64

	var model models.SysUserDept
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysUserDept.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysUserDept
// @Summary 获取SysUserDept
// @Tags SysUserDept
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysUserDept} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-user-dept/get [post]
// @Security Bearer
func (e *SysUserDeptApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUserDept
	if err := service.SerSysUserDept.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysUserDept
// @Summary 创建SysUserDept
// @Tags SysUserDept
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserDeptDto true "body"
// @Success 200 {object} base.Resp{data=models.SysUserDept} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-user-dept/create [post]
// @Security Bearer
func (e *SysUserDeptApi) Create(c *gin.Context) {
	var req dto.SysUserDeptDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUserDept
	copier.Copy(&data, req)
	if err := service.SerSysUserDept.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysUserDept
// @Summary 更新SysUserDept
// @Tags SysUserDept
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserDeptDto true "body"
// @Success 200 {object} base.Resp{data=models.SysUserDept} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-user-dept/update [post]
// @Security Bearer
func (e *SysUserDeptApi) Update(c *gin.Context) {
	var req dto.SysUserDeptDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUserDept
	copier.Copy(&data, req)
	if err := service.SerSysUserDept.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysUserDept
// @Summary 删除SysUserDept
// @Tags SysUserDept
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysUserDept} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-user-dept/del [post]
// @Security Bearer
func (e *SysUserDeptApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysUserDept.DelIds(&models.SysUserDept{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
