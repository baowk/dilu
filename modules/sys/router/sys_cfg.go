package router

import (
	"dilu/modules/sys/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysCfgRouter)
}

// 无需认证的路由示例
func registerSysCfgRouter(v1 *gin.RouterGroup) {
	api := apis.SysCfgApi{}
	r := v1.Group("sys-cfg").Use(middleware.JwtHandler())
	{
		r.POST("/get", api.Get)
		r.POST("/create", api.Create)
		r.POST("/update", api.Update)
		r.POST("/page", api.QueryPage)
		r.POST("/del", api.Del)
	}
}