//客户端注册注册中心程序
//默认开启检测，http检测请求为 http://ip:port/api/health
//grpc检测为 ip:port/Health

package start

import (
	"dilu/common/config"
	"fmt"
	"net/http"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-rd/rd"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerHealthRouter)
	AppRouters = append(AppRouters, InitRouter)
}

func InitRouter() {
	r := core.GetGinEngine()
	noCheckRoleRouter(r)
}

func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v := r.Group("")

	for _, f := range routerNoCheckRole {
		f(v)
	}
}

func registerHealthRouter(v1 *gin.RouterGroup) {
	r := v1.Group("")
	{
		r.GET("/api/health", func(ctx *gin.Context) {
			ctx.AbortWithStatus(http.StatusOK)
		})
	}
}

var rdclient rd.RDClient

func rdInit() {
	//注册中心
	if config.Ext.RdConfig.Enable {
		rdcfg := config.Ext.RdConfig
		for _, v := range rdcfg.Registers {
			if v.Addr == "" {
				if v.Protocol == "http" {
					v.Addr = core.Cfg.Server.GetHost()
					v.Port = core.Cfg.Server.GetPort()
					v.HealthCheck = fmt.Sprintf("http://%s:%d/api/health", core.Cfg.Server.GetHost(), core.Cfg.Server.GetPort())
				} else {
					v.Addr = core.Cfg.GrpcServer.GetHost()
					v.Port = core.Cfg.GrpcServer.GetPort()
					v.HealthCheck = fmt.Sprintf("%s:%d/Health", core.Cfg.GrpcServer.GetHost(), core.Cfg.GrpcServer.GetPort())
				}
			}
			if len(v.Tags) == 0 {
				v.Tags = []string{core.Cfg.Server.Mode}
			}
		}
		core.Log.Debug("注册中心连接", zap.Any("rdcfg", rdcfg))
		var err error
		rdclient, err = rd.NewRDClient(&rdcfg, core.Log.Sugar())
		if err != nil {
			core.Log.Error("注册中心连接失败", zap.Error(err))
		}
	}
}

func rdRelease() {
	if rdclient != nil {
		rdclient.Deregister()
	}
}
