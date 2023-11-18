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

type SysTeamApi struct {
	base.BaseApi
}

var ApiSysTeam = SysTeamApi{}

// QueryPage 获取团队列表
// @Summary 获取团队列表
// @Tags sys-SysTeam
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysTeamGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysTeam}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-team/page [post]
// @Security Bearer
func (e *SysTeamApi) QueryPage(c *gin.Context) {
	var req dto.SysTeamGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysTeam, 10)
	var total int64

	var model models.SysTeam
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysTeam.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取团队
// @Summary 获取团队
// @Tags sys-SysTeam
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysTeam} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-team/get [post]
// @Security Bearer
func (e *SysTeamApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysTeam
	if err := service.SerSysTeam.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建团队
// @Summary 创建团队
// @Tags sys-SysTeam
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysTeamDto true "body"
// @Success 200 {object} base.Resp{data=models.SysTeam} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-team/create [post]
// @Security Bearer
func (e *SysTeamApi) Create(c *gin.Context) {
	var req dto.SysTeamDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysTeam
	copier.Copy(&data, req)
	if err := service.SerSysTeam.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新团队
// @Summary 更新团队
// @Tags sys-SysTeam
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysTeamDto true "body"
// @Success 200 {object} base.Resp{data=models.SysTeam} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-team/update [post]
// @Security Bearer
func (e *SysTeamApi) Update(c *gin.Context) {
	var req dto.SysTeamDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	req.Id = utils.GetReqTeamId(c, req.Id)
	var data models.SysTeam
	copier.Copy(&data, req)
	if err := service.SerSysTeam.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新团队
// @Summary 更新团队
// @Tags sys-SysTeam
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysTeamDto true "body"
// @Success 200 {object} base.Resp{data=models.SysTeam} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-team/change [post]
// @Security Bearer
func (e *SysTeamApi) ChangeName(c *gin.Context) {
	var req dto.SysTeamDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	req.Id = utils.GetTeamId(c)
	var data models.SysTeam
	copier.Copy(&data, req)
	if err := service.SerSysTeam.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除团队
// @Summary 删除团队
// @Tags sys-SysTeam
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysTeam} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-team/del [post]
// @Security Bearer
func (e *SysTeamApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysTeam.DelIds(&models.SysTeam{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
