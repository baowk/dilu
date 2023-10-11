package router

import (
	"dilu/common/middleware"
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysJobRouter)
}

// 默认需登录认证的路由
func registerSysJobRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-job").Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiSysJob.Get)
		r.POST("/create", apis.ApiSysJob.Create)
		r.POST("/update", apis.ApiSysJob.Update)
		r.POST("/page", apis.ApiSysJob.QueryPage)
		r.POST("/del", apis.ApiSysJob.Del)
	}
}
