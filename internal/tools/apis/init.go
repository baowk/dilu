package apis

import (
	"dilu/internal/tools/utils"
	"time"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

var (
	InitApi = Init{}
)

type Init struct {
	base.BaseApi
}

var last time.Time

var server utils.Server

// Monitor 监控
// @Summary 监控
// @Tags 工具 / 监控
// @Accept application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=utils.Server} "{"code": 200, "data": [...]}"
// @Router /api/v1/tools/monitor [post]
// @Security Bearer
func (e *Init) Monitor(c *gin.Context) {
	cur := time.Now()
	if cur.Sub(last) < time.Second {
		e.Ok(c, server)
		return
	}
	last = cur
	server.Os = utils.InitOS()
	cpu, err := utils.InitCPU()
	if err == nil {
		server.Cpu = cpu
	}
	d, err := utils.InitDisk()
	if err == nil {
		server.Disk = d
	}
	ram, err := utils.InitRAM()
	if err == nil {
		server.Ram = ram
	}
	e.Ok(c, server)
}
