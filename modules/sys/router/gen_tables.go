package router

import (
	"dilu/modules/sys/apis"
	"dilu/common/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerGenTablesRouter)
}

// 默认需登录认证的路由
func registerGenTablesRouter(v1 *gin.RouterGroup) {
	r := v1.Group("gen-tables").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.ApiGenTables.Get)
		r.POST("/create", apis.ApiGenTables.Create)
		r.POST("/update", apis.ApiGenTables.Update)
		r.POST("/page", apis.ApiGenTables.QueryPage)
		r.POST("/del", apis.ApiGenTables.Del)
	}
}