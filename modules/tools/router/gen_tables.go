package router

import (
	"dilu/modules/tools/apis"

	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerGenTablesRouter)
}

// 默认需登录认证的路由
func registerGenTablesRouter(v1 *gin.RouterGroup) {
	r := v1.Group("gen")
	{
		if core.Cfg.Gen.Enable {
			r.POST("/menu", apis.ApiGenTables.GenMenuAndApi)
			r.POST("/update", apis.ApiGenTables.Update)
			r.POST("/page", apis.ApiGenTables.QueryPage)
			r.POST("/del", apis.ApiGenTables.Del)
			r.POST("/add", apis.ApiGenTables.Insert)
			r.POST("db/tables", apis.ApiGenTables.GetDBTableList)
			r.POST("/code", apis.ApiGenTables.GenCode)
			r.POST("/dbs", apis.ApiGenTables.GetDBS)
		}
	}
}
