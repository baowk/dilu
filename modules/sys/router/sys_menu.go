package router

import (
	"dilu/common/middleware"
	"dilu/modules/sys/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSysMenuRouter)
}

// 默认需登录认证的路由
func registerSysMenuRouter(v1 *gin.RouterGroup) {
	r := v1.Group("sys-menu").Use(middleware.JwtHandler())
	{
		r.POST("/get", apis.SysMenuA.Get)
		r.POST("/create", apis.SysMenuA.Create)
		r.POST("/update", apis.SysMenuA.Update)
		r.POST("/all", apis.SysMenuA.GetMenus)
		r.POST("/del", apis.SysMenuA.Del)
		r.POST("/userMenus", apis.SysMenuA.GetUserMenus)
		r.POST("/perms", apis.SysMenuA.GetUserPerms)
	}
}