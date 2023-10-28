package router

import (
	"dilu/modules/sys/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerGenColumnsRouter)
}

// 默认需登录认证的路由
func registerGenColumnsRouter(v1 *gin.RouterGroup) {
	r := v1.Group("gen-columns").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiGenColumns.Get)
		r.POST("/create", apis.ApiGenColumns.Create)
		r.POST("/update", apis.ApiGenColumns.Update)
		r.POST("/page", apis.ApiGenColumns.QueryPage)
		r.POST("/del", apis.ApiGenColumns.Del)
	}
}