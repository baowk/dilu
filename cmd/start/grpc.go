// grpc开启服务与grpc初始化
package start

import (
	"fmt"
	"log"
	"net"

	"github.com/baowk/dilu-core/common/utils/text"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-rd/grpc/pb/health"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func grpcInit() {
	grpcServer = grpc.NewServer()
	//注册grpc服务
	health.RegisterHealthServer(grpcServer, &health.HealthServerImpl{}) //健康检测服务
	//service.RegisterGreeterServer(grpcServer, &impl.TempimplementedGreeterServer{})

	//注册服务完成
	grpcAddr := fmt.Sprintf("%s:%d", core.Cfg.GrpcServer.GetHost(), core.Cfg.GrpcServer.GetPort())
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		core.Log.Error("failed to listen", zap.Error(err))
		log.Fatal("failed to listen:", err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			core.Log.Error("failed to serve", zap.Error(err))
			log.Fatal("failed to serve:", err)
		}
	}()
	fmt.Println(text.Green("Dilu GRPC Server started ,Listen: ") + text.Red("[ "+grpcAddr+" ]"))
}

func closeGrpc() {
	grpcServer.GracefulStop()
}
