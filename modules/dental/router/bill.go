package router

import (
	"dilu/common/middleware"
	"dilu/modules/dental/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerBillRouter)
}

// 默认需登录认证的路由
func registerBillRouter(v1 *gin.RouterGroup) {
	r := v1.Group("bill").Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiBill.Get)
		r.POST("/create", apis.ApiBill.Create)
		r.POST("/update", apis.ApiBill.Update)
		r.POST("/page", apis.ApiBill.QueryPage)
		r.POST("/del", apis.ApiBill.Del)
		r.POST("identify", apis.ApiBill.Identify)
		r.POST("export", apis.ApiBill.BillExport)
	}
	r1 := v1.Group("st").Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r1.POST("day", apis.ApiBill.StDay)
		r1.POST("month", apis.ApiBill.StMonth)
		r1.POST("query", apis.ApiBill.StQuery)
		r1.POST("export", apis.ApiBill.StExport)
	}
}
