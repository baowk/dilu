package router

import (
	"dilu/common/middleware"
	"dilu/modules/notice/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerTaskRouter)
}

// 默认需登录认证的路由
func registerTaskRouter(v1 *gin.RouterGroup) {
	r := v1.Group("task").Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiTask.Get)
		r.POST("/create", apis.ApiTask.Create)
		r.POST("/update", apis.ApiTask.Update)
		r.POST("/page", apis.ApiTask.QueryPage)
		r.POST("/del", apis.ApiTask.Del)
	}

	r2 := v1.Group("task").Use(middleware.JwtHandler())
	{
		r2.POST("my", apis.ApiTask.UserTasks)
	}
}
