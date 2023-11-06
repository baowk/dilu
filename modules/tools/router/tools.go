package router

import (
	"dilu/modules/tools/apis"

	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckSsoRouter)
}

// 无需认证的路由示例
func sysNoCheckSsoRouter(v1 *gin.RouterGroup) {
	r := v1.Group("tools")
	{
		if core.Cfg.Gen {
			//r.GET("/tables/page", apis.SysTable.GetPage)
			//r.GET("tables/info/:tableId", apis.SysTable.Get)
			r.POST("/tables/info", apis.SysTable.Insert)
			r.PUT("/tables/info", apis.SysTable.Update)
			//r.DELETE("/tables/info/:tableId", apis.SysTable.Delete)
			//r.GET("/db/columns/page", apis.GenApi.GetDBColumnList)
			r.POST("db/tables/page", apis.GenApi.GetDBTableList)
			r.GET("gen/preview/:tableId", apis.GenApi.Preview)
			r.POST("gen/code", apis.GenApi.GenCode)
			r.GET("gen/menu/:tableId/:menuPid", apis.GenApi.GenMenuAndApi)
			r.GET("gen/api", apis.GenApi.GenApis)
			r.POST("/gen/dbs", apis.GenApi.GetDBS)
		}
	}
}
