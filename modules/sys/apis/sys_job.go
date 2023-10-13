package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysJobApi struct {
	base.BaseApi
}

var ApiSysJob = SysJobApi{}

// QueryPage 获取SysJob列表
// @Summary Page接口
// @Tags sys-SysJob
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysJobGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysJob}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-job/page [post]
// @Security Bearer
func (e *SysJobApi) QueryPage(c *gin.Context) {
	var req dto.SysJobGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysJob, 10)
	var total int64

	var model models.SysJob
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysJob.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysJob
// @Summary 获取SysJob
// @Tags sys-SysJob
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysJob} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-job/get [post]
// @Security Bearer
func (e *SysJobApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysJob
	if err := service.SerSysJob.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysJob
// @Summary 创建SysJob
// @Tags sys-SysJob
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysJobDto true "body"
// @Success 200 {object} base.Resp{data=models.SysJob} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-job/create [post]
// @Security Bearer
func (e *SysJobApi) Create(c *gin.Context) {
	var req dto.SysJobDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysJob
	copier.Copy(&data, req)
	if err := service.SerSysJob.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysJob
// @Summary 更新SysJob
// @Tags sys-SysJob
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysJobDto true "body"
// @Success 200 {object} base.Resp{data=models.SysJob} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-job/update [post]
// @Security Bearer
func (e *SysJobApi) Update(c *gin.Context) {
	var req dto.SysJobDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysJob
	copier.Copy(&data, req)
	if err := service.SerSysJob.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysJob
// @Summary 删除SysJob
// @Tags sys-SysJob
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysJob} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-job/del [post]
// @Security Bearer
func (e *SysJobApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysJob.DelIds(&models.SysJob{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
