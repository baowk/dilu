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

// Get 获取接口列表
// @Summary 获取接口列表
// @Tags sys-SysMenu
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
// @Tags sys-SysMenu
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
// @Tags sys-SysMenu
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
// @Tags sys-SysMenu
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

// GetMenus 获取所有菜单
// @Summary 获取所有菜单
// @Tags sys-SysMenu
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=[]models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-menu/all [post]
// @Security Bearer
func (e *SysMenuApi) GetMenus(c *gin.Context) {
	list := make([]models.SysMenu, 0)
	if err := service.SerSysMenu.GetMenus(c, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// GetGrantMenus 获取授权菜单
// @Summary 获取授权菜单
// @Tags sys-SysMenu
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=[]models.SysMenu} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-menu/grant [post]
// @Security Bearer
func (e *SysMenuApi) GetGrantMenus(c *gin.Context) {
	list := make([]models.SysMenu, 0)
	if err := service.SerSysMenu.GetMenus(c, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// GetUserMenus 获取用户菜单
// @Summary 获取用户菜单
// @Tags sys-SysMenu
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Success 200 {object} base.Resp{data=[]dto.MenuVo} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-menu/userMenus [post]
// @Security Bearer
func (e *SysMenuApi) GetUserMenus(c *gin.Context) {
	var ms []dto.MenuVo
	if err := service.SerSysMenu.GetUserMenus(c, &ms); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, ms)
}

// GetUserMenus 获取用户菜单
// @Summary 获取用户菜单
// @Tags sys-SysMenu
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysMenuGetReq true "body"
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/canAccess [post]
func (e *SysMenuApi) CanAccess(c *gin.Context) {
	var req dto.SysMenuGetReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysMenu.CanAccess(c, req.Id); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// // GetUserPerms 获取用户权限
// // @Summary 获取用户权限
// // @Tags sys-SysMenu
// // @Accept application/json
// // @Product application/json
// // @Success 200 {object} base.Resp{data=[]string} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/sys-menu/perms [post]
// // @Security Bearer
// func (e *SysMenuApi) GetUserPerms(c *gin.Context) {
// 	role := utils.GetRoleId(c)
// 	if role < 1 {
// 		e.Code(c, codes.InvalidToken_401)
// 		return
// 	}
// 	var ms []string
// 	if err := service.SerSysMenu.GetUserPerms(role, &ms); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, ms)
// }
