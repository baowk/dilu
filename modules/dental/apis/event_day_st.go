package apis

import (
	"dilu/modules/dental/models"
	"dilu/modules/dental/service"
	"dilu/modules/dental/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type EventDayStApi struct {
	base.BaseApi
}

var ApiEventDaySt = EventDayStApi{}

// QueryPage 获取EventDaySt列表
// @Summary Page接口
// @Tags dental-EventDaySt
// @Accept application/json
// @Product application/json
// @Param data body dto.EventDayStGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.EventDaySt}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/event-day-st/page [post]
// @Security Bearer
func (e *EventDayStApi) QueryPage(c *gin.Context) {
	var req dto.EventDayStGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.EventDaySt, 10)
	var total int64

	var model models.EventDaySt
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerEventDaySt.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取EventDaySt
// @Summary 获取EventDaySt
// @Tags dental-EventDaySt
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.EventDaySt} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/event-day-st/get [post]
// @Security Bearer
func (e *EventDayStApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.EventDaySt
	if err := service.SerEventDaySt.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建EventDaySt
// @Summary 创建EventDaySt
// @Tags dental-EventDaySt
// @Accept application/json
// @Product application/json
// @Param data body dto.EventDayStDto true "body"
// @Success 200 {object} base.Resp{data=models.EventDaySt} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/event-day-st/create [post]
// @Security Bearer
func (e *EventDayStApi) Create(c *gin.Context) {
	var req dto.EventDayStDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.EventDaySt
	copier.Copy(&data, req)
	if err := service.SerEventDaySt.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新EventDaySt
// @Summary 更新EventDaySt
// @Tags dental-EventDaySt
// @Accept application/json
// @Product application/json
// @Param data body dto.EventDayStDto true "body"
// @Success 200 {object} base.Resp{data=models.EventDaySt} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/event-day-st/update [post]
// @Security Bearer
func (e *EventDayStApi) Update(c *gin.Context) {
	var req dto.EventDayStDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.EventDaySt
	copier.Copy(&data, req)
	if err := service.SerEventDaySt.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除EventDaySt
// @Summary 删除EventDaySt
// @Tags dental-EventDaySt
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.EventDaySt} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/event-day-st/del [post]
// @Security Bearer
func (e *EventDayStApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerEventDaySt.DelIds(&models.EventDaySt{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
