package router

import (
	"dilu/docs"
	"dilu/modules/sys/apis"
	"fmt"

	"github.com/baowk/dilu-core/core"
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	//routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// InitRouter 路由初始化，不要怀疑，这里用到了
func InitRouter() {
	r := core.GetGinEngine()
	if core.Cfg.Server.Mode != core.ModeProd.String() {
		fmt.Printf("%s %s  \r\n", docs.SwaggerInfo.Title, docs.SwaggerInfo.Version)
		//初始化swagger
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		r.GET("/", apis.InitApi.Init)
		r.POST("/doInit", apis.InitApi.DoInit)
	}

	noCheckRoleRouter(r)
}

// func InitBusinessRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

// 	// 无需认证的路由
// 	noCheckRoleRouter(r)
// 	// 需要认证的路由
// 	checkRoleRouter(r, authMiddleware)

// 	return r
// }

// noCheckRoleRouter 无需认证的路由
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group("/api/v1")

	for _, f := range routerNoCheckRole {
		f(v)
	}
}
