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

type BillApi struct {
	base.BaseApi
}

var ApiBill = BillApi{}

// QueryPage 获取Bill列表
// @Summary Page接口
// @Tags dental-Bill
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
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

// Get 获取Bill
// @Summary 获取Bill
// @Tags dental-Bill
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
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

// Create 创建Bill
// @Summary 创建Bill
// @Tags dental-Bill
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.IdentifyBillDto true "body"
// @Success 200 {object} base.Resp{data=models.Bill} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/bill/create [post]
// @Security Bearer
func (e *BillApi) Create(c *gin.Context) {
	var req dto.IdentifyBillDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	req.TeamId = utils.GetTeamId(c)
	var data models.Bill
	if err := service.SerBill.CreateBill(e.GetReqId(c), req, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新Bill
// @Summary 更新Bill
// @Tags dental-Bill
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
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

// Del 删除Bill
// @Summary 删除Bill
// @Tags dental-Bill
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
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

// Identify 智能识别
// @Summary 智能识别
// @Tags dental-Bill
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.BillTmplReq true "body"
// @Success 200 {object} base.Resp{data=dto.IdentifyBillDto} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/bill/identify [post]
// @Security Bearer
func (e *BillApi) Identify(c *gin.Context) {
	var req dto.BillTmplReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.TeamId == 0 {
		req.TeamId = utils.GetTeamId(c)
	}
	var ib dto.IdentifyBillDto
	if err := service.SerBill.Identify(req, &ib); err != nil {
		e.Err(c, err)
	}
	e.Ok(c, ib)
}
