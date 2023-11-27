package apis

import (
	"dilu/common/utils"
	"dilu/modules/notice/models"
	"dilu/modules/notice/service"
	"dilu/modules/notice/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type TaskApi struct {
	base.BaseApi
}

var ApiTask = TaskApi{}

// QueryPage 获取Task列表
// @Summary 获取Task列表
// @Tags notice-Task
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.TaskGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.Task}} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/task/page [post]
// @Security Bearer
func (e *TaskApi) QueryPage(c *gin.Context) {
	var req dto.TaskGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.Task, 10)
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}

	if err := service.SerTask.QueryPage(req, &list, &total, req.GetSize(), req.GetOffset()); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Get 获取Task
// @Summary 获取Task
// @Tags notice-Task
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqId true "body"
// @Success 200 {object} base.Resp{data=models.Task} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/task/get [post]
// @Security Bearer
func (e *TaskApi) Get(c *gin.Context) {
	var req base.ReqId
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Task
	if err := service.SerTask.Get(req.Id, &data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Create 创建Task
// @Summary 创建Task
// @Tags notice-Task
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.TaskDto true "body"
// @Success 200 {object} base.Resp{data=models.Task} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/task/create [post]
// @Security Bearer
func (e *TaskApi) Create(c *gin.Context) {
	var req dto.TaskDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Task
	copier.Copy(&data, req)
	if err := service.SerTask.Create(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Update 更新Task
// @Summary 更新Task
// @Tags notice-Task
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.TaskDto true "body"
// @Success 200 {object} base.Resp{data=models.Task} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/task/update [post]
// @Security Bearer
func (e *TaskApi) Update(c *gin.Context) {
	var req dto.TaskDto
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var data models.Task
	copier.Copy(&data, req)
	if err := service.SerTask.UpdateById(&data); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// Del 删除Task
// @Summary 删除Task
// @Tags notice-Task
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.Task} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/task/del [post]
// @Security Bearer
func (e *TaskApi) Del(c *gin.Context) {
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerTask.DelIds(&models.Task{}, req.Ids); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// QueryPage 获取Task列表
// @Summary 获取Task列表
// @Tags notice-Task
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.TaskGetPageReq true "body"
// @Success 200 {object} base.Resp{data=dto.NoticeDto} "{"code": 200, "data": [...]}"
// @Router /api/v1/notice/task/my [post]
// @Security Bearer
func (e *TaskApi) UserTasks(c *gin.Context) {
	var req dto.TaskGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.Task, 10)
	var total int64

	if req.SortOrder == "" {
		req.SortOrder = "desc"
	}
	req.TeamId = utils.GetReqTeamId(c, req.TeamId)
	req.UserId = utils.GetUserId(c)
	var unReadCnt int64

	if err := service.SerTask.UserTasks(&req, &list, &total, &unReadCnt); err != nil {
		e.Error(c, err)
		return
	}

	res := dto.NoticeDto{
		Key:   "2",
		Name:  "任务",
		Total: total,
		Count: unReadCnt,
	}
	for _, v := range list {
		var item dto.NoticeItem
		copier.Copy(&item, v)
		item.Type = 2
		item.NoticeType = v.TaskType
		item.CreatedAt = v.CreatedAt.Unix()
		item.BeginAt = v.BeginAt.Unix()
		item.EndAt = v.EndAt.Unix()
		item.ReminderTime = v.ReminderTime.Unix()
		res.List = append(res.List, item)
	}
	e.Ok(c, res)
}
