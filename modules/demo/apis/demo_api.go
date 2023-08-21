package apis

import (
	"dilu/modules/demo/models"
	"dilu/modules/demo/service"
	"dilu/modules/demo/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type DemoApi struct {
	base.BaseApi
}

// QueryPage 获取列表
// @Summary Page接口
// @Tags Demo
// @Accept application/json
// @Product application/json
// @Param data body dto.DemePageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.Demo}} "{"code": 200, "data": [...]}"
// @Router /api/v1/demo/page [post]
func (e *DemoApi) QueryPage(c *gin.Context) {
	var req dto.DemePageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.Demo, 10)
	var total int64
	if err := service.DemoS.Page(req, &list, &total, e.GetReqId(c)); err != nil {
		e.Err(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获得
// @Summary 获得
// @Tags Demo
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.Demo} "{"code": 200, "data": [...]}"
// @Router /api/v1/demo/get [post]
func (e *DemoApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Demo
	if err := service.DemoS.Get(req.Id, &data, e.GetReqId(c)); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建
// @Summary 创建
// @Tags Demo
// @Accept application/json
// @Product application/json
// @Param data body dto.DemoDto true "body"
// @Success 200 {object} base.Resp{data=models.Demo} "{"code": 200, "data": [...]}"
// @Router /api/v1/demo/create [post]
func (e *DemoApi) Create(c *gin.Context) {
	var req dto.DemoDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Demo
	copier.Copy(&data, req)
	if err := service.DemoS.Create(&data, e.GetReqId(c)); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新
// @Summary 更新
// @Tags Demo
// @Accept application/json
// @Product application/json
// @Param data body dto.DemoDto true "body"
// @Success 200 {object} base.Resp{data=models.Demo} "{"code": 200, "data": [...]}"
// @Router /api/v1/demo/update [post]
func (e *DemoApi) Update(c *gin.Context) {
	var req dto.DemoDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Demo
	copier.Copy(&data, req)
	if err := service.DemoS.Update(&data, e.GetReqId(c)); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除
// @Summary 删除
// @Tags Demo
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.Demo} "{"code": 200, "data": [...]}"
// @Router /api/v1/demo/del [post]
func (e *DemoApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	//var data models.Demo
	if err := service.DemoS.Del(req.Ids, e.GetReqId(c)); err != nil {
		e.Err(c, err)
		return
	}
	e.Ok(c)
}
