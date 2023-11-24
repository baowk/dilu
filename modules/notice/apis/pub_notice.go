package apis

import (
	"dilu/modules/notice/models"
	"dilu/modules/notice/service"
	"dilu/modules/notice/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type PubNoticeApi struct {
	base.BaseApi
}

var ApiPubNotice = PubNoticeApi{}

// QueryPage 获取公用通知列表
// @Summary 获取公用通知列表
// @Tags notice-PubNotice
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.PubNoticeGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.PubNotice}} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/pub-notice/page [post]
// @Security Bearer
func (e *PubNoticeApi) QueryPage(c *gin.Context) {
	var req dto.PubNoticeGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.PubNotice, 10)
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerPubNotice.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取公用通知
// @Summary 获取公用通知
// @Tags notice-PubNotice
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.PubNotice} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/pub-notice/get [post]
// @Security Bearer
func (e *PubNoticeApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.PubNotice
	if err := service.SerPubNotice.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建公用通知
// @Summary 创建公用通知
// @Tags notice-PubNotice
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.PubNoticeDto true "body"
// @Success 200 {object} base.Resp{data=models.PubNotice} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/pub-notice/create [post]
// @Security Bearer
func (e *PubNoticeApi) Create(c *gin.Context) {
	var req dto.PubNoticeDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	var data models.PubNotice
	copier.Copy(&data, req)
	if err := service.SerPubNotice.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新公用通知
// @Summary 更新公用通知
// @Tags notice-PubNotice
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.PubNoticeDto true "body"
// @Success 200 {object} base.Resp{data=models.PubNotice} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/pub-notice/update [post]
// @Security Bearer
func (e *PubNoticeApi) Update(c *gin.Context) {
	var req dto.PubNoticeDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.PubNotice
	copier.Copy(&data, req)
	if err := service.SerPubNotice.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除公用通知
// @Summary 删除公用通知
// @Tags notice-PubNotice
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.PubNotice} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/pub-notice/del [post]
// @Security Bearer
func (e *PubNoticeApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerPubNotice.DelIds(&models.PubNotice{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
