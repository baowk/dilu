package router

import (
	"dilu/modules/sys/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysApiRouter)
}

// 默认需登录认证的路由
func registerSysApiRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-api").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.SysApiA.Get)
		r.POST("/create", apis.SysApiA.Create)
		r.POST("/update", apis.SysApiA.Update)
		r.POST("/page", apis.SysApiA.QueryPage)
		r.POST("/del", apis.SysApiA.Del)
	}
}