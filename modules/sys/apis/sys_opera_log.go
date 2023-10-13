package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysOperaLogApi struct {
	base.BaseApi
}

var ApiSysOperaLog = SysOperaLogApi{}

// QueryPage 获取SysOperaLog列表
// @Summary Page接口
// @Tags sys-SysOperaLog
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysOperaLogGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysOperaLog}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-opera-log/page [post]
// @Security Bearer
func (e *SysOperaLogApi) QueryPage(c *gin.Context) {
	var req dto.SysOperaLogGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysOperaLog, 10)
	var total int64

	var model models.SysOperaLog
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysOperaLog.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysOperaLog
// @Summary 获取SysOperaLog
// @Tags sys-SysOperaLog
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysOperaLog} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-opera-log/get [post]
// @Security Bearer
func (e *SysOperaLogApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysOperaLog
	if err := service.SerSysOperaLog.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysOperaLog
// @Summary 创建SysOperaLog
// @Tags sys-SysOperaLog
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysOperaLogDto true "body"
// @Success 200 {object} base.Resp{data=models.SysOperaLog} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-opera-log/create [post]
// @Security Bearer
func (e *SysOperaLogApi) Create(c *gin.Context) {
	var req dto.SysOperaLogDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysOperaLog
	copier.Copy(&data, req)
	if err := service.SerSysOperaLog.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysOperaLog
// @Summary 更新SysOperaLog
// @Tags sys-SysOperaLog
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysOperaLogDto true "body"
// @Success 200 {object} base.Resp{data=models.SysOperaLog} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-opera-log/update [post]
// @Security Bearer
func (e *SysOperaLogApi) Update(c *gin.Context) {
	var req dto.SysOperaLogDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysOperaLog
	copier.Copy(&data, req)
	if err := service.SerSysOperaLog.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysOperaLog
// @Summary 删除SysOperaLog
// @Tags sys-SysOperaLog
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysOperaLog} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-opera-log/del [post]
// @Security Bearer
func (e *SysOperaLogApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysOperaLog.DelIds(&models.SysOperaLog{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
