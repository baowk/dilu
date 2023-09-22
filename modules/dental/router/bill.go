package router

import (
	"dilu/modules/dental/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerBillRouter)
}

// 默认需登录认证的路由
func registerBillRouter(v1 *gin.RouterGroup) {
	r := v1.Group("bill").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiBill.Get)
		r.POST("/create", apis.ApiBill.Create)
		r.POST("/update", apis.ApiBill.Update)
		r.POST("/page", apis.ApiBill.QueryPage)
		r.POST("/del", apis.ApiBill.Del)
	}
}