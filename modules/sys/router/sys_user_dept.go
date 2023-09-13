package router

import (
	"dilu/modules/sys/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysUserDeptRouter)
}

// 默认需登录认证的路由
func registerSysUserDeptRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-user-dept").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiSysUserDept.Get)
		r.POST("/create", apis.ApiSysUserDept.Create)
		r.POST("/update", apis.ApiSysUserDept.Update)
		r.POST("/page", apis.ApiSysUserDept.QueryPage)
		r.POST("/del", apis.ApiSysUserDept.Del)
	}
}