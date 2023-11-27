package router

import (
	"dilu/common/middleware"
	"dilu/modules/ai/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerBillRouter)
}

// 默认需登录认证的路由
func registerBillRouter(v1 *gin.RouterGroup) {
	r := v1.Group("").Use(middleware.JwtHandler())
	{
		r.POST("/chat", apis.ApiAi.Chat)
	}

}
