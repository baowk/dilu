package router

import (
	"dilu/modules/sys/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysCfgRouter)
}

// 默认需登录认证的路由
func registerSysCfgRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-cfg").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.SysCfgA.Get)
		r.POST("/create", apis.SysCfgA.Create)
		r.POST("/update", apis.SysCfgA.Update)
		r.POST("/page", apis.SysCfgA.QueryPage)
		r.POST("/del", apis.SysCfgA.Del)
	}
}