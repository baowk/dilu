package apis

import (
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

// QueryPage 获取月目标设定列表
// @Summary Page接口
// @Tags TargetTask
// @Accept application/json
// @Product application/json
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

	var model models.TargetTask
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerTargetTask.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取月目标设定
// @Summary 获取月目标设定
// @Tags TargetTask
// @Accept application/json
// @Product application/json
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

// Create 创建月目标设定
// @Summary 创建月目标设定
// @Tags TargetTask
// @Accept application/json
// @Product application/json
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
	var data models.TargetTask
	copier.Copy(&data, req)
	if err := service.SerTargetTask.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新月目标设定
// @Summary 更新月目标设定
// @Tags TargetTask
// @Accept application/json
// @Product application/json
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
	var data models.TargetTask
	copier.Copy(&data, req)
	if err := service.SerTargetTask.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除月目标设定
// @Summary 删除月目标设定
// @Tags TargetTask
// @Accept application/json
// @Product application/json
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
