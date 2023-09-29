package router

import (
	"dilu/modules/sys/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysTeamRouter)
}

// 默认需登录认证的路由
func registerSysTeamRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-team").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiSysTeam.Get)
		r.POST("/create", apis.ApiSysTeam.Create)
		r.POST("/update", apis.ApiSysTeam.Update)
		r.POST("/page", apis.ApiSysTeam.QueryPage)
		r.POST("/del", apis.ApiSysTeam.Del)
	}
}