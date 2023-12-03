package router

import (
	"dilu/common/middleware"
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysMemberRouter)
}

// 默认需登录认证的路由
func registerSysMemberRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-member").Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiSysMember.Get)
		r.POST("/create", apis.ApiSysMember.Create)
		r.POST("/update", apis.ApiSysMember.Update)
		r.POST("/page", apis.ApiSysMember.QueryPage)
		r.POST("/del", apis.ApiSysMember.Del)
	}

	r2 := v1.Group("sys-member").Use(middleware.JwtHandler())
	{
		r2.POST("myTeams", apis.ApiSysMember.MyTeams)
		r2.POST("myInfo", apis.ApiSysMember.MyInfo)
		r2.POST("changeMyInfo", apis.ApiSysMember.ChangeMyInfo)
		r2.POST("members", apis.ApiSysMember.GetMembers)
	}
}
