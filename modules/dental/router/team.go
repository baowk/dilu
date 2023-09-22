package router

import (
	"dilu/modules/dental/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerTeamRouter)
}

// 默认需登录认证的路由
func registerTeamRouter(v1 *gin.RouterGroup) {
	r := v1.Group("team").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiTeam.Get)
		r.POST("/create", apis.ApiTeam.Create)
		r.POST("/update", apis.ApiTeam.Update)
		r.POST("/page", apis.ApiTeam.QueryPage)
		r.POST("/del", apis.ApiTeam.Del)
	}
}