package apis

import (
	"dilu/modules/ai/service"
	"dilu/modules/ai/service/dto"

	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type AiApi struct {
	base.BaseApi
}

var ApiAi = AiApi{}

// Chat
// @Summary ai接口
// @Tags Ai
// @Accept application/json
// @Product application/json
// @Param teamId header int false "团队id"
// @Param data body dto.AiMsg true "body"
// @Success 200 {object} base.Resp{data=string} "{"code": 200, "data": [...]}"
// @Router /api/v1/ai/chat [post]
// @Security Bearer
func (e *AiApi) Chat(c *gin.Context) {
	var req dto.AiMsg
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	var res string
	if err := service.SerAi.Chat(req, e.GetReqId(c), &res); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c, res)
}
