package container

import (
	"dilu/internal/common/third/file_store"
	"dilu/internal/common/third/sms"

	"github.com/baowk/dilu-rd/rd"
	"google.golang.org/grpc"
)

// Container 服务容器，集中管理原本分散在各包中的全局变量
type Container struct {
	RDClient   rd.RDClient
	GrpcServer *grpc.Server
	SMS        sms.SmsSend
	OSS        file_store.OSSHandler
}

var global = &Container{}

// Global 返回全局服务容器
func Global() *Container {
	return global
}
