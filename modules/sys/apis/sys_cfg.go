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

// QueryPage 获取SysCfg列表
// @Summary Page接口
// @Tags SysCfg
// @Accept application/json
// @Product application/json
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
	list := make([]models.SysCfg, 10)
	var total int64
	if err := service.SysCfgS.Page(req, &list, &total, e.GetReqId(c)); err != nil {
		e.Err(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取SysCfg
// @Summary 获取SysCfg
// @Tags SysCfg
// @Accept application/json
// @Product application/json
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
	if err := service.SysCfgS.Get(req.Id, &data, e.GetReqId(c)); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建SysCfg
// @Summary 创建SysCfg
// @Tags SysCfg
// @Accept application/json
// @Product application/json
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
	if err := service.SysCfgS.Create(&data, e.GetReqId(c)); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新SysCfg
// @Summary 更新SysCfg
// @Tags SysCfg
// @Accept application/json
// @Product application/json
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
	if err := service.SysCfgS.Update(&data, e.GetReqId(c)); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除SysCfg
// @Summary 删除SysCfg
// @Tags SysCfg
// @Accept application/json
// @Product application/json
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
	if err := service.SysCfgS.Del(req.Ids, e.GetReqId(c)); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c)
}
