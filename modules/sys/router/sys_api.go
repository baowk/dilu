package router

import (
	"dilu/common/middleware"
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysApiRouter)
}

// 默认需登录认证的路由
func registerSysApiRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-api").Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.SysApiA.Get)
		r.POST("/create", apis.SysApiA.Create)
		r.POST("/update", apis.SysApiA.Update)
		r.POST("/page", apis.SysApiA.QueryPage)
		r.POST("/del", apis.SysApiA.Del)
	}
	r2 := v1.Group("sys-api")
	{
		r2.POST("/all", apis.SysApiA.Query)
	}
}
