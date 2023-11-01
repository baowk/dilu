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

type SummaryPlanDayApi struct {
	base.BaseApi
}

var ApiSummaryPlanDay = SummaryPlanDayApi{}

// QueryPage 获取SummaryPlanDay列表
// @Summary Page接口
// @Tags dental-SummaryPlanDay
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SummaryPlanDayGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SummaryPlanDay}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/summary-plan-day/page [post]
// @Security Bearer
func (e *SummaryPlanDayApi) QueryPage(c *gin.Context) {
	var req dto.SummaryPlanDayGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SummaryPlanDay, 10)
	var total int64

	teamId := utils.GetTeamId(c)
	userId := utils.GetUserId(c)
	if err := service.SerSummaryPlanDay.Page(req, teamId, userId, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SummaryPlanDay
// @Summary 获取SummaryPlanDay
// @Tags dental-SummaryPlanDay
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SummaryPlanDay} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/summary-plan-day/get [post]
// @Security Bearer
func (e *SummaryPlanDayApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SummaryPlanDay
	if err := service.SerSummaryPlanDay.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SummaryPlanDay
// @Summary 创建SummaryPlanDay
// @Tags dental-SummaryPlanDay
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SummaryPlanDayDto true "body"
// @Success 200 {object} base.Resp{data=models.SummaryPlanDay} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/summary-plan-day/create [post]
// @Security Bearer
func (e *SummaryPlanDayApi) Create(c *gin.Context) {
	var req dto.SummaryPlanDayDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	teamId := utils.GetTeamId(c)
	userId := utils.GetUserId(c)
	var data models.SummaryPlanDay
	copier.Copy(&data, req)
	if err := service.SerSummaryPlanDay.Create(teamId, userId, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SummaryPlanDay
// @Summary 更新SummaryPlanDay
// @Tags dental-SummaryPlanDay
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SummaryPlanDayDto true "body"
// @Success 200 {object} base.Resp{data=models.SummaryPlanDay} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/summary-plan-day/update [post]
// @Security Bearer
func (e *SummaryPlanDayApi) Update(c *gin.Context) {
	var req dto.SummaryPlanDayDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	teamId := utils.GetTeamId(c)
	userId := utils.GetUserId(c)
	var data models.SummaryPlanDay
	copier.Copy(&data, req)
	if err := service.SerSummaryPlanDay.Update(teamId, userId, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SummaryPlanDay
// @Summary 删除SummaryPlanDay
// @Tags dental-SummaryPlanDay
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SummaryPlanDay} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/summary-plan-day/del [post]
// @Security Bearer
func (e *SummaryPlanDayApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSummaryPlanDay.DelIds(&models.SummaryPlanDay{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
