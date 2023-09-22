package router

import (
	"dilu/modules/dental/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSummaryPlanDayRouter)
}

// 默认需登录认证的路由
func registerSummaryPlanDayRouter(v1 *gin.RouterGroup) {
	r := v1.Group("summary-plan-day").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiSummaryPlanDay.Get)
		r.POST("/create", apis.ApiSummaryPlanDay.Create)
		r.POST("/update", apis.ApiSummaryPlanDay.Update)
		r.POST("/page", apis.ApiSummaryPlanDay.QueryPage)
		r.POST("/del", apis.ApiSummaryPlanDay.Del)
	}
}