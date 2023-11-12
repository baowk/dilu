package apis

import (
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysUserApi struct {
	base.BaseApi
}

var ApiSysUser = SysUserApi{}

// QueryPage 获取用户列表
// @Summary Page接口
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysUser}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-user/page [post]
// @Security Bearer
func (e *SysUserApi) QueryPage(c *gin.Context) {
	var req dto.SysUserGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysUser, 10)
	var total int64

	var model models.SysUser
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysUser.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取用户
// @Summary 获取用户
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-user/get [post]
// @Security Bearer
func (e *SysUserApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUser
	if err := service.SerSysUser.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysUser
// @Summary 创建SysUser
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserDto true "body"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-user/create [post]
// @Security Bearer
func (e *SysUserApi) Create(c *gin.Context) {
	var req dto.SysUserDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysUser
	copier.Copy(&data, req)
	if err := service.SerSysUser.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysUser
// @Summary 更新SysUser
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param data body dto.SysUserDto true "body"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-user/update [post]
// @Security Bearer
func (e *SysUserApi) Update(c *gin.Context) {
	var req dto.SysUserDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	var data models.SysUser
	copier.Copy(&data, req)
	if err := service.SerSysUser.ChangeUserinfo(utils.GetUserId(c), data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysUser
// @Summary 删除SysUser
// @Tags sys-SysUser
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysUser} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-user/del [post]
// @Security Bearer
func (e *SysUserApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysUser.DelIds(&models.SysUser{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
