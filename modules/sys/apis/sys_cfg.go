package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type SysCfgApi struct {
	base.BaseApi
}

var ApiSysCfg = SysCfgApi{}

// QueryPage 获取配置列表
// @Summary 获取配置列表
// @Tags sys-SysCfg
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysCfgGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.SysCfg}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-cfg/page [post]
// @Security Bearer
func (e *SysCfgApi) QueryPage(c *gin.Context) {
	var req dto.SysCfgGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.SysCfg, 0, req.GetSize())
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerSysCfg.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取配置
// @Summary 获取配置
// @Tags sys-SysCfg
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.SysCfg} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-cfg/get [post]
// @Security Bearer
func (e *SysCfgApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysCfg
	if err := service.SerSysCfg.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建配置
// @Summary 创建配置
// @Tags sys-SysCfg
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysCfgDto true "body"
// @Success 200 {object} base.Resp{data=models.SysCfg} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-cfg/create [post]
// @Security Bearer
func (e *SysCfgApi) Create(c *gin.Context) {
	var req dto.SysCfgDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysCfg
	copier.Copy(&data, req)
	if err := service.SerSysCfg.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新配置
// @Summary 更新配置
// @Tags sys-SysCfg
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.SysCfgDto true "body"
// @Success 200 {object} base.Resp{data=models.SysCfg} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-cfg/update [post]
// @Security Bearer
func (e *SysCfgApi) Update(c *gin.Context) {
	var req dto.SysCfgDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.SysCfg
	copier.Copy(&data, req)
	if err := service.SerSysCfg.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除配置
// @Summary 删除配置
// @Tags sys-SysCfg
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.SysCfg} "{"code": 200, "data": [...]}"
// @Router /api/v1/sys/sys-cfg/del [post]
// @Security Bearer
func (e *SysCfgApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerSysCfg.DelIds(&models.SysCfg{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
