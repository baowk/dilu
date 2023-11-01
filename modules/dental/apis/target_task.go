package apis

import (
	"dilu/common/utils"
	"dilu/modules/dental/models"
	"dilu/modules/dental/service"
	"dilu/modules/dental/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type TargetTaskApi struct {
	base.BaseApi
}

var ApiTargetTask = TargetTaskApi{}

// QueryPage 获取TargetTask列表
// @Summary 获取TargetTask列表
// @Tags dental-TargetTask
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.TargetTaskGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.TargetTask}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/target-task/page [post]
// @Security Bearer
func (e *TargetTaskApi) QueryPage(c *gin.Context) {
	var req dto.TargetTaskGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.TargetTask, 10)
	var total int64

	teamId := utils.GetTeamId(c)
	userId := utils.GetUserId(c)
	if err := service.SerTargetTask.Page(req, teamId, userId, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取TargetTask
// @Summary 获取TargetTask
// @Tags dental-TargetTask
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.TargetTask} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/target-task/get [post]
// @Security Bearer
func (e *TargetTaskApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.TargetTask
	if err := service.SerTargetTask.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建TargetTask
// @Summary 创建TargetTask
// @Tags dental-TargetTask
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.TargetTaskDto true "body"
// @Success 200 {object} base.Resp{data=models.TargetTask} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/target-task/create [post]
// @Security Bearer
func (e *TargetTaskApi) Create(c *gin.Context) {
	var req dto.TargetTaskDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	teamId := utils.GetTeamId(c)
	userId := utils.GetUserId(c)
	var data models.TargetTask
	copier.Copy(&data, req)
	if err := service.SerTargetTask.Create(teamId, userId, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新TargetTask
// @Summary 更新TargetTask
// @Tags dental-TargetTask
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.TargetTaskDto true "body"
// @Success 200 {object} base.Resp{data=models.TargetTask} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/target-task/update [post]
// @Security Bearer
func (e *TargetTaskApi) Update(c *gin.Context) {
	var req dto.TargetTaskDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	teamId := utils.GetTeamId(c)
	userId := utils.GetUserId(c)
	var data models.TargetTask
	copier.Copy(&data, req)
	if err := service.SerTargetTask.Update(teamId, userId, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除TargetTask
// @Summary 删除TargetTask
// @Tags dental-TargetTask
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.TargetTask} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/target-task/del [post]
// @Security Bearer
func (e *TargetTaskApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerTargetTask.DelIds(&models.TargetTask{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
