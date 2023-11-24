package router

import (
	"dilu/common/middleware"
	"dilu/modules/notice/apis"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerPubNoticeRouter)
}

// 默认需登录认证的路由
func registerPubNoticeRouter(v1 *gin.RouterGroup) {
	r := v1.Group("pub-notice").Use(middleware.JwtHandler()).Use(middleware.PermHandler())
	{
		r.POST("/get", apis.ApiPubNotice.Get)
		r.POST("/create", apis.ApiPubNotice.Create)
		r.POST("/update", apis.ApiPubNotice.Update)
		r.POST("/page", apis.ApiPubNotice.QueryPage)
		r.POST("/del", apis.ApiPubNotice.Del)
	}
}
