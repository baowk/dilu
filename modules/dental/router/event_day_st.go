package router

import (
	"dilu/modules/dental/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerEventDayStRouter)
}

// 默认需登录认证的路由
func registerEventDayStRouter(v1 *gin.RouterGroup) {
	r := v1.Group("event-day-st").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiEventDaySt.Get)
		r.POST("/create", apis.ApiEventDaySt.Create)
		r.POST("/update", apis.ApiEventDaySt.Update)
		r.POST("/page", apis.ApiEventDaySt.QueryPage)
		r.POST("/del", apis.ApiEventDaySt.Del)
	}
}