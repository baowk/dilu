package router

import (
	"dilu/modules/sys/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysOperaLogRouter)
}

// 默认需登录认证的路由
func registerSysOperaLogRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-opera-log").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiSysOperaLog.Get)
		r.POST("/create", apis.ApiSysOperaLog.Create)
		r.POST("/update", apis.ApiSysOperaLog.Update)
		r.POST("/page", apis.ApiSysOperaLog.QueryPage)
		r.POST("/del", apis.ApiSysOperaLog.Del)
	}
}