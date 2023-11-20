package router

import (
	"dilu/common/middleware"
	"dilu/modules/tools/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerGenTablesRouter)
}

// 默认需登录认证的路由
func registerGenTablesRouter(v1 *gin.RouterGroup) {
	r := v1.Group("gen")
	{
		r.POST("/dbs", apis.ApiGenTables.GetDBS)
		r.POST("db/tables", apis.ApiGenTables.GetDBTableList)
		r.POST("/page", apis.ApiGenTables.QueryPage)
		r.POST("/menu", apis.ApiGenTables.GenMenuAndApi)
		r.POST("/update", apis.ApiGenTables.Update)
		r.POST("/del", apis.ApiGenTables.Del)
		r.POST("/add", apis.ApiGenTables.Insert)
		r.POST("/code", apis.ApiGenTables.GenCode)
	}
	r2 := v1.Use().Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r2.POST("/monitor", apis.InitApi.Monitor)
	}
}
