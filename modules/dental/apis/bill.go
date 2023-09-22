package apis

import (
	"dilu/modules/dental/models"
	"dilu/modules/dental/service"
	"dilu/modules/dental/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type BillApi struct {
	base.BaseApi
}

var ApiBill = BillApi{}

// QueryPage 获取账单列表
// @Summary Page接口
// @Tags Bill
// @Accept application/json
// @Product application/json
// @Param data body dto.BillGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.Bill}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/bill/page [post]
// @Security Bearer
func (e *BillApi) QueryPage(c *gin.Context) {
	var req dto.BillGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.Bill, 10)
	var total int64

	var model models.Bill
	if err := copier.Copy(&model, req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerBill.Page(model, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取账单
// @Summary 获取账单
// @Tags Bill
// @Accept application/json
// @Product application/json
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.Bill} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/bill/get [post]
// @Security Bearer
func (e *BillApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Bill
	if err := service.SerBill.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建账单
// @Summary 创建账单
// @Tags Bill
// @Accept application/json
// @Product application/json
// @Param data body dto.BillDto true "body"
// @Success 200 {object} base.Resp{data=models.Bill} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/bill/create [post]
// @Security Bearer
func (e *BillApi) Create(c *gin.Context) {
	var req dto.BillDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Bill
	copier.Copy(&data, req)
	if err := service.SerBill.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新账单
// @Summary 更新账单
// @Tags Bill
// @Accept application/json
// @Product application/json
// @Param data body dto.BillDto true "body"
// @Success 200 {object} base.Resp{data=models.Bill} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/bill/update [post]
// @Security Bearer
func (e *BillApi) Update(c *gin.Context) {
	var req dto.BillDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Bill
	copier.Copy(&data, req)
	if err := service.SerBill.Save(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除账单
// @Summary 删除账单
// @Tags Bill
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.Bill} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/bill/del [post]
// @Security Bearer
func (e *BillApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerBill.DelIds(&models.Bill{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}
