package apis

import (
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type SysDeptApi struct {
	base.BaseApi
}

var ApiSysDept = SysDeptApi{}

// QueryPage 获取SysDept列表
// @Summary Page接口
// @Tags sys-SysDept
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysDeptGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysDept}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-dept/page [post]
// @Security Bearer
func (e *SysDeptApi) QueryPage(c *gin.Context) {
	var req dto.SysDeptGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysDept, 10)
	var total int64
	req.TeamId = utils.GetReqTeamId(c, req.TeamId)
	if err := service.SerSysDept.Page(&req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// List 获取全部部门
// @Summary 获取全部部门
// @Tags sys-SysDept
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysDeptGetPageReq true "body"
// @Success 200 {object} base.Resp{data=[]models.SysDept} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-dept/all [post]
// @Security Bearer
func (e *SysDeptApi) List(c *gin.Context) {
	var req dto.SysDeptGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	req.TeamId = utils.GetReqTeamId(c, req.TeamId)
	list := make([]models.SysDept, 10)

	if err := service.SerSysDept.GetDepts(req.TeamId, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// Get 获取SysDept
// @Summary 获取SysDept
// @Tags sys-SysDept
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysDept} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-dept/get [post]
// @Security Bearer
func (e *SysDeptApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysDept
	if err := service.SerSysDept.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysDept
// @Summary 创建SysDept
// @Tags sys-SysDept
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysDeptDto true "body"
// @Success 200 {object} base.Resp{data=models.SysDept} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-dept/create [post]
// @Security Bearer
func (e *SysDeptApi) Create(c *gin.Context) {
	var req dto.SysDeptDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	teamId := utils.GetTeamId(c)
	var data models.SysDept
	if teamId > 0 {
		req.TeamId = teamId
	}
	adminId := utils.GetUserId(c)
	if err := service.SerSysDept.CreateDept(req, adminId, e.GetReqId(c)); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysDept
// @Summary 更新SysDept
// @Tags sys-SysDept
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysDeptDto true "body"
// @Success 200 {object} base.Resp{data=models.SysDept} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-dept/update [post]
// @Security Bearer
func (e *SysDeptApi) Update(c *gin.Context) {
	var req dto.SysDeptDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	teamId := utils.GetTeamId(c)
	var data models.SysDept
	if teamId > 0 {
		req.TeamId = teamId
	}
	adminId := utils.GetUserId(c)
	if err := service.SerSysDept.UpdateDept(req, adminId, e.GetReqId(c)); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysDept
// @Summary 删除SysDept
// @Tags sys-SysDept
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysDept} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-dept/del [post]
// @Security Bearer
func (e *SysDeptApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysDept.DelIds(&models.SysDept{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
