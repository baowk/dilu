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
			// r.POST("/get", apis.ApiGenTables.Get)
			// r.POST("/create", apis.ApiGenTables.Create)
			r.POST("/update", apis.ApiGenTables.Update)
			r.POST("/page", apis.ApiGenTables.QueryPage)
			r.POST("/del", apis.ApiGenTables.Del)

			//r.GET("/tables/page", apis.SysTable.GetPage)
			//r.GET("tables/info/:tableId", apis.SysTable.Get)
			r.POST("/add", apis.ApiGenTables.Insert)
			//r.DELETE("/tables/info/:tableId", apis.SysTable.Delete)
			//r.GET("/db/columns/page", apis.GenApi.GetDBColumnList)
			r.POST("db/tables", apis.ApiGenTables.GetDBTableList)
			//r2.GET("gen/preview/:tableId", apis.GenApi.Preview)
			r.POST("/code", apis.ApiGenTables.GenCode)
			//r2.GET("gen/menu/:tableId/:menuPid", apis.GenApi.GenMenuAndApi)
			//r2.GET("gen/api", apis.GenApi.GenApis)
			r.POST("/dbs", apis.ApiGenTables.GetDBS)
		}
	}
}
