package router

import (
	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	r := core.GetGinEngine()

	noCheckRoleRouter(r)
}

// noCheckRoleRouter 无需认证的路由
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group("/api/v1")

	for _, f := range routerNoCheckRole {
		f(v)
	}
}
