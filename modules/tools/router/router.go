package router

import (
	"dilu/common/consts"
	"dilu/docs"
	"fmt"

	"github.com/gin-gonic/gin"

	"dilu/modules/tools/apis"

	"github.com/baowk/dilu-core/core"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	//routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// InitRouter 路由初始化
func InitRouter() {
	r := core.GetGinEngine()
	if core.Cfg.Server.Mode != core.ModeProd.String() {
		fmt.Printf("%s %s  \r\n", docs.SwaggerInfo.Title, docs.SwaggerInfo.Version)
		//初始化swagger
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	if core.Cfg.Gen.Enable {
		//r.GET("/init", apis.InitApi.Init)
		r.POST("/doInit", apis.InitApi.DoInit)
	}
	noCheckRoleRouter(r)
}

// noCheckRoleRouter 无需认证的路由
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group(consts.ApiRoot + "/tools")

	for _, f := range routerNoCheckRole {
		f(v)
	}
}
