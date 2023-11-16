package apis

import (
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type SysRoleApi struct {
	base.BaseApi
}

var ApiSysRole = SysRoleApi{}

// QueryPage 获取SysRole列表
// @Summary Page接口
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysRoleGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysRole}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-role/page [post]
// @Security Bearer
func (e *SysRoleApi) QueryPage(c *gin.Context) {
	var req dto.SysRoleGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysRole, 0)
	var total int64
	req.TeamId = utils.GetReqTeamId(c, req.TeamId)

	if err := service.SerSysRole.Page(&req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// List 获取角色列表
// @Summary 获取角色列表
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysRoleGetPageReq true "body"
// @Success 200 {object} base.Resp{data=[]models.SysRole} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-role/list [post]
// @Security Bearer
func (e *SysRoleApi) List(c *gin.Context) {
	var req dto.SysRoleGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysRole, 0)

	if err := service.SerSysRole.Query(utils.GetReqTeamId(c, req.TeamId), req.Status, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// Get 获取SysRole
// @Summary 获取SysRole
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=dto.SysRoleDtoResp} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-role/get [post]
// @Security Bearer
func (e *SysRoleApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data dto.SysRoleDtoResp
	if err := service.SerSysRole.GetRole(req.Id, utils.GetUserId(c), utils.GetTeamId(c), &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysRole
// @Summary 创建SysRole
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysRoleDto true "body"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-role/create [post]
// @Security Bearer
func (e *SysRoleApi) Create(c *gin.Context) {
	var req dto.SysRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysRole.Create(utils.GetUserId(c), utils.GetTeamId(c), req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// Update 更新SysRole
// @Summary 更新SysRole
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysRoleDto true "body"
// @Success 200 {object} base.Resp{} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-role/update [post]
// @Security Bearer
func (e *SysRoleApi) Update(c *gin.Context) {
	var req dto.SysRoleDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysRole.Update(utils.GetUserId(c), utils.GetTeamId(c), req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// Del 删除SysRole
// @Summary 删除SysRole
// @Tags sys-SysRole
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysRole} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-role/del [post]
// @Security Bearer
func (e *SysRoleApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysRole.DelIds(&models.SysRole{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
