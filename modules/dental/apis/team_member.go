package apis

import (
	"dilu/modules/dental/models"
	"dilu/modules/dental/service"
	"dilu/modules/dental/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type TeamMemberApi struct {
	base.BaseApi
}

var ApiTeamMember = TeamMemberApi{}

// QueryPage 获取TeamMember列表
// @Summary Page接口
// @Tags dental-TeamMember
// @Accept application/json
// @Product application/json
// @Param data body dto.TeamMemberGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.TeamMember}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/team-member/page [post]
// @Security Bearer
func (e *TeamMemberApi) QueryPage(c *gin.Context) {
	var req dto.TeamMemberGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.TeamMember, 10)
	var total int64

	var model models.TeamMember
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerTeamMember.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取TeamMember
// @Summary 获取TeamMember
// @Tags dental-TeamMember
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.TeamMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/team-member/get [post]
// @Security Bearer
func (e *TeamMemberApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.TeamMember
	if err := service.SerTeamMember.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建TeamMember
// @Summary 创建TeamMember
// @Tags dental-TeamMember
// @Accept application/json
// @Product application/json
// @Param data body dto.TeamMemberDto true "body"
// @Success 200 {object} base.Resp{data=models.TeamMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/team-member/create [post]
// @Security Bearer
func (e *TeamMemberApi) Create(c *gin.Context) {
	var req dto.TeamMemberDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.TeamMember
	copier.Copy(&data, req)
	if err := service.SerTeamMember.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新TeamMember
// @Summary 更新TeamMember
// @Tags dental-TeamMember
// @Accept application/json
// @Product application/json
// @Param data body dto.TeamMemberDto true "body"
// @Success 200 {object} base.Resp{data=models.TeamMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/team-member/update [post]
// @Security Bearer
func (e *TeamMemberApi) Update(c *gin.Context) {
	var req dto.TeamMemberDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.TeamMember
	copier.Copy(&data, req)
	if err := service.SerTeamMember.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除TeamMember
// @Summary 删除TeamMember
// @Tags dental-TeamMember
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.TeamMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/team-member/del [post]
// @Security Bearer
func (e *TeamMemberApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerTeamMember.DelIds(&models.TeamMember{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
