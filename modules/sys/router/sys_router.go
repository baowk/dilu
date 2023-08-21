package router

import (
	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, sysNoCheckRoleRouter)
}

// 无需认证的路由示例
func sysNoCheckRoleRouter(v1 *gin.RouterGroup) {
	// api := apis.InitApi
	// r := v1.Group("")
	// {
	// 	r.GET("init", api.Init)
	// 	r.POST("doInit", api.DoInit)
	// }
}
