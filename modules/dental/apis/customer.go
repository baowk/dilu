package apis

import (
	"dilu/common/utils"
	"dilu/modules/dental/models"
	"dilu/modules/dental/service"
	"dilu/modules/dental/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type CustomerApi struct {
	base.BaseApi
}

var ApiCustomer = CustomerApi{}

// QueryPage 获取Customer列表
// @Summary Page接口
// @Tags dental-Customer
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.CustomerGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.Customer}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/customer/page [post]
// @Security Bearer
func (e *CustomerApi) QueryPage(c *gin.Context) {
	var req dto.CustomerGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.Customer, 0)
	var total int64
	teamId := utils.GetTeamId(c)
	userId := utils.GetUserId(c)

	if err := service.SerCustomer.Page(req, teamId, userId, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取Customer
// @Summary 获取Customer
// @Tags dental-Customer
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.Customer} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/customer/get [post]
// @Security Bearer
func (e *CustomerApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Customer
	if err := service.SerCustomer.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建Customer
// @Summary 创建Customer
// @Tags dental-Customer
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.CustomerDto true "body"
// @Success 200 {object} base.Resp{data=models.Customer} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/customer/create [post]
// @Security Bearer
func (e *CustomerApi) Create(c *gin.Context) {
	var req dto.CustomerDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Customer
	copier.Copy(&data, req)

	if err := service.SerCustomer.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新Customer
// @Summary 更新Customer
// @Tags dental-Customer
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.CustomerDto true "body"
// @Success 200 {object} base.Resp{data=models.Customer} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/customer/update [post]
// @Security Bearer
func (e *CustomerApi) Update(c *gin.Context) {
	var req dto.CustomerDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Customer
	copier.Copy(&data, req)

	if err := service.SerCustomer.Update(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除Customer
// @Summary 删除Customer
// @Tags dental-Customer
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.Customer} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/customer/del [post]
// @Security Bearer
func (e *CustomerApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerCustomer.DelIds(&models.Customer{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
