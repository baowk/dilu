package router

import (
	"dilu/common/middleware"
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysDeptRouter)
}

// 默认需登录认证的路由
func registerSysDeptRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-dept").Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiSysDept.Get)
		r.POST("/create", apis.ApiSysDept.Create)
		r.POST("/update", apis.ApiSysDept.Update)
		r.POST("/page", apis.ApiSysDept.QueryPage)
		r.POST("/del", apis.ApiSysDept.Del)
		r.POST("/all", apis.ApiSysDept.List)
	}
}
