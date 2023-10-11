package apis

import (
	"dilu/common/codes"
	"dilu/common/utils"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysMemberApi struct {
	base.BaseApi
}

var ApiSysMember = SysMemberApi{}

// QueryPage 获取会员列表
// @Summary 获取会员列表
// @Tags sys-SysMember
// @Accept application/json
// @Product application/json
// @Param teamId header int false "teamId"
// @Param data body dto.SysMemberGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysMember}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-member/page [post]
// @Security Bearer
func (e *SysMemberApi) QueryPage(c *gin.Context) {
	var req dto.SysMemberGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysMember, 10)
	var total int64

	if err := service.SerSysMember.Query(req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// MyTeams 获取加入的团队
// @Summary 获取加入的团队
// @Tags sys-SysMember
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=[]dto.TeamMemberResp} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-member/myTeams [post]
// @Security Bearer
func (e *SysMemberApi) MyTeams(c *gin.Context) {
	uid := utils.GetUserId(c)
	if uid < 1 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	var list []dto.TeamMemberResp
	if err := service.SerSysMember.GetUserTeams(uid, &list); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, list)
}

// MyInfo 获取我的信息
// @Summary 获取我的信息
// @Tags sys-SysMember
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=[]dto.TeamMemberResp} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-member/myInfo [post]
// @Security Bearer
func (e *SysMemberApi) MyInfo(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	uid := utils.GetUserId(c)
	if uid < 1 {
		e.Code(c, codes.InvalidToken_401)
		return
	}
	var data dto.TeamMemberResp
	if err := service.SerSysMember.GetTeamUser(req.Id, uid, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Get 获取会员
// @Summary 获取会员
// @Tags sys-SysMember
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-member/get [post]
// @Security Bearer
func (e *SysMemberApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMember
	if err := service.SerSysMember.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建会员
// @Summary 创建会员
// @Tags sys-SysMember
// @Accept application/json
// @Product application/json
// @Param data body dto.SysMemberDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-member/create [post]
// @Security Bearer
func (e *SysMemberApi) Create(c *gin.Context) {
	var req dto.SysMemberDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMember
	copier.Copy(&data, req)
	if err := service.SerSysMember.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新会员
// @Summary 更新会员
// @Tags sys-SysMember
// @Accept application/json
// @Product application/json
// @Param data body dto.SysMemberDto true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-member/update [post]
// @Security Bearer
func (e *SysMemberApi) Update(c *gin.Context) {
	var req dto.SysMemberDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysMember
	copier.Copy(&data, req)
	if err := service.SerSysMember.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除会员
// @Summary 删除会员
// @Tags sys-SysMember
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysMember} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-member/del [post]
// @Security Bearer
func (e *SysMemberApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysMember.DelIds(&models.SysMember{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
