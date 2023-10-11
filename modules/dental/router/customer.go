package router

import (
	"dilu/common/middleware"
	"dilu/modules/dental/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerCustomerRouter)
}

// 默认需登录认证的路由
func registerCustomerRouter(v1 *gin.RouterGroup) {
	r := v1.Group("customer").Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiCustomer.Get)
		r.POST("/create", apis.ApiCustomer.Create)
		r.POST("/update", apis.ApiCustomer.Update)
		r.POST("/page", apis.ApiCustomer.QueryPage)
		r.POST("/del", apis.ApiCustomer.Del)
	}
}
