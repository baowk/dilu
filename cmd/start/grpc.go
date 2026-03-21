// grpc开启服务与grpc初始化
package start

import (
	"dilu/internal/common/config"
	"dilu/internal/common/container"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/baowk/dilu-core/common/utils/text"
	"github.com/baowk/dilu-core/core/logger"
	"github.com/baowk/dilu-rd/grpc/pb/health"
	"google.golang.org/grpc"
)

func grpcInit() {
	grpcServer := grpc.NewServer()
	container.Global().GrpcServer = grpcServer
	//注册grpc服务
	health.RegisterHealthServer(grpcServer, &health.HealthServerImpl{}) //健康检测服务
	//service.RegisterGreeterServer(grpcServer, &impl.TempimplementedGreeterServer{})

	//注册服务完成
	grpcAddr := fmt.Sprintf("%s:%d", config.Get().GrpcServer.GetHost(), config.Get().GrpcServer.GetPort())
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Error().Err(err).Msg("failed to listen")
		log.Fatal("failed to listen:", err)
	}
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error().Err(err).Msg("failed to serve")
			log.Fatal("failed to serve:", err)
		}
	}()
	fmt.Println(text.Green("Dilu GRPC Server started ,Listen: ") + text.Red("[ "+grpcAddr+" ]"))
}

func closeGrpc() {
	s := container.Global().GrpcServer
	if s == nil {
		return
	}
	done := make(chan struct{})
	go func() {
		s.GracefulStop()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(10 * time.Second):
		logger.Warn().Msg("grpc graceful stop timeout, forcing stop")
		s.Stop()
	}
}
