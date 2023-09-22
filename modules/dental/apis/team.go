package apis

import (
	"dilu/modules/dental/models"
	"dilu/modules/dental/service"
	"dilu/modules/dental/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type TeamApi struct {
	base.BaseApi
}

var ApiTeam = TeamApi{}

// QueryPage 获取团队列表
// @Summary Page接口
// @Tags Team
// @Accept application/json
// @Product application/json
// @Param data body dto.TeamGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.Team}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/team/page [post]
// @Security Bearer
func (e *TeamApi) QueryPage(c *gin.Context) {
	var req dto.TeamGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.Team, 10)
	var total int64

	var model models.Team
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerTeam.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取团队
// @Summary 获取团队
// @Tags Team
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.Team} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/team/get [post]
// @Security Bearer
func (e *TeamApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Team
	if err := service.SerTeam.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建团队
// @Summary 创建团队
// @Tags Team
// @Accept application/json
// @Product application/json
// @Param data body dto.TeamDto true "body"
// @Success 200 {object} base.Resp{data=models.Team} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/team/create [post]
// @Security Bearer
func (e *TeamApi) Create(c *gin.Context) {
	var req dto.TeamDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Team
	copier.Copy(&data, req)
	if err := service.SerTeam.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新团队
// @Summary 更新团队
// @Tags Team
// @Accept application/json
// @Product application/json
// @Param data body dto.TeamDto true "body"
// @Success 200 {object} base.Resp{data=models.Team} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/team/update [post]
// @Security Bearer
func (e *TeamApi) Update(c *gin.Context) {
	var req dto.TeamDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Team
	copier.Copy(&data, req)
	if err := service.SerTeam.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除团队
// @Summary 删除团队
// @Tags Team
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.Team} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/team/del [post]
// @Security Bearer
func (e *TeamApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerTeam.DelIds(&models.Team{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
