package router

import (
	"dilu/modules/dental/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerTeamMemberRouter)
}

// 默认需登录认证的路由
func registerTeamMemberRouter(v1 *gin.RouterGroup) {
	r := v1.Group("team-member").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiTeamMember.Get)
		r.POST("/create", apis.ApiTeamMember.Create)
		r.POST("/update", apis.ApiTeamMember.Update)
		r.POST("/page", apis.ApiTeamMember.QueryPage)
		r.POST("/del", apis.ApiTeamMember.Del)
	}
}