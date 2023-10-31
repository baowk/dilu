package apis

import (
	"dilu/common/utils"
	"dilu/modules/dental/models"
	"dilu/modules/dental/service"
	"dilu/modules/dental/service/dto"
	"time"

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
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]dto.BillDto}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/bill/page [post]
// @Security Bearer
func (e *BillApi) QueryPage(c *gin.Context) {
	var req dto.BillGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]dto.BillDto, 0)
	var total int64

	if err := service.SerBill.Page(utils.GetTeamId(c), req, &list, &total); err != nil {
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
	if err := service.SerBill.CreateBill(e.GetReqId(c), req, &data, utils.GetUserId(c)); err != nil {
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
	var req dto.IdentifyBillDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Bill
	copier.Copy(&data, req)
	data.UpdateBy = utils.GetUserId(c)
	if err := service.SerBill.UpdateBill(e.GetReqId(c), req, &data, utils.GetUserId(c)); err != nil {
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

// StDay 日统计
// @Summary 日统计
// @Tags dental-Bill
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.StQueryReq true "body"
// @Success 200 {object} base.Resp{data=[]string} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/st/day [post]
// @Security Bearer
func (e *BillApi) StDay(c *gin.Context) {
	var req dto.StQueryReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.Begin.IsZero() {
		req.Begin = time.Now()
	}
	teamId := utils.GetTeamId(c)
	if teamId > 0 {
		req.TeamId = teamId
	}
	if req.UserId == 0 {
		req.UserId = utils.GetUserId(c)
	}
	text, err := service.SerBill.StDay(req.TeamId, req.UserId, req.DeptPath, req.Begin, e.GetReqId(c))
	if err != nil {
		e.Error(c, err)
	} else {
		e.PureOk(c, text)
	}
}

// StMonth 月统计
// @Summary 月统计
// @Tags dental-Bill
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.StQueryReq true "body"
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/st/month [post]
// @Security Bearer
func (e *BillApi) StMonth(c *gin.Context) {
	var req dto.StQueryReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if req.Begin.IsZero() {
		req.Begin = time.Now()
	}
	teamId := utils.GetTeamId(c)
	if teamId > 0 {
		req.TeamId = teamId
	}
	if req.UserId == 0 {
		req.UserId = utils.GetUserId(c)
	}
	text, err := service.SerBill.StMonth(req.TeamId, req.UserId, req.DeptPath, req.Begin, e.GetReqId(c))
	if err != nil {
		e.Error(c, err)
	} else {
		e.PureOk(c, text)
	}
}

// StQuery 查询统计
// @Summary 查询统计
// @Tags dental-Bill
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.StQueryReq true "body"
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/v1/dental/st/query [post]
// @Security Bearer
func (e *BillApi) StQuery(c *gin.Context) {
	var req dto.StQueryReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	teamId := utils.GetTeamId(c)
	if teamId > 0 {
		req.TeamId = teamId
	}
	// if req.UserId == 0 {
	// 	req.UserId = utils.GetUserId(c)
	// }
	text, err := service.SerBill.StQuery(req.TeamId, req.UserId, req.DeptPath, req.Begin, req.End, e.GetReqId(c))
	if err != nil {
		e.Error(c, err)
	} else {
		e.Ok(c, text)
	}
}
