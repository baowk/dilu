//客户端注册注册中心程序
//默认开启检测，http检测请求为 http://ip:port/api/health
//grpc检测为 ip:port/Health

package start

import (
	"dilu/common/config"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/baowk/dilu-core/common/utils/ips"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-rd/rd"
	"github.com/gin-gonic/gin"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerHealthRouter)
	AppRouters = append(AppRouters, InitRouter)
}

func InitRouter() {
	r := core.GetApp().GetGinEngine()
	noCheckRoleRouter(r)
}

func noCheckRoleRouter(r *gin.Engine) {
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
	if config.Get().RdConfig.Enable {
		rdcfg := config.Get().RdConfig
		for _, v := range rdcfg.Registers {
			if v.Protocol != "grpc" && v.Protocol != "http" {
				slog.Error("rd register error", "protocol", v.Protocol)
				continue
			}
			if v.Protocol == "grpc" && !config.Get().GrpcServer.Enable {
				slog.Error("rd register error", "protocol", v.Protocol, "GrpcServer enable", false)
				continue
			}
			if v.Name == "" {
				if v.Protocol == "http" {
					v.Name = config.Get().Server.Name
				} else {
					if config.Get().GrpcServer.Name != "" {
						v.Name = config.Get().GrpcServer.Name
					} else {
						v.Name = config.Get().Server.Name + "_grpc"
					}
				}
			}
			if v.Addr == "" {
				if v.Protocol == "http" {
					if config.Get().Server.GetHost() != "0.0.0.0" {
						v.Addr = config.Get().Server.GetHost()
					} else {
						v.Addr = ips.GetLocalHost()
					}
					v.Port = config.Get().Server.GetPort()
					v.HealthCheck = fmt.Sprintf("http://%s:%d/api/health", v.Addr, config.Get().Server.GetPort())
				} else {
					if config.Get().GrpcServer.GetHost() != "0.0.0.0" {
						v.Addr = config.Get().GrpcServer.GetHost()
					} else {
						v.Addr = ips.GetLocalHost()
					}
					v.Port = config.Get().GrpcServer.GetPort()
					v.HealthCheck = fmt.Sprintf("%s:%d/Health", v.Addr, config.Get().GrpcServer.GetPort())
				}
			}
			if len(v.Tags) == 0 {
				v.Tags = []string{config.Get().Server.Mode}
			}
			if v.Id == "" {
				v.Id = fmt.Sprintf("%s:%d", v.Addr, v.Port)
			}
		}

		slog.Debug("注册中心连接", "rdcfg", rdcfg)
		var err error
		rdclient, err = rd.NewRDClient(&rdcfg)
		if err != nil {
			slog.Error("注册中心连接失败", err)
		}
	}
}

func rdRelease() {
	if rdclient != nil {
		rdclient.Deregister()
	}
}
