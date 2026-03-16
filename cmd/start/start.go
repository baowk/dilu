package start

import (
	"dilu/internal/bootstrap"
	"dilu/internal/common/codes"
	"dilu/internal/common/config"
	"dilu/internal/common/middleware"
	"log/slog"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/i18n"
	"github.com/spf13/cobra"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "start",
		Short:   "Get Application config info",
		Example: "dilu start -c resources/config.dev.yml",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "resources/config.dev.yaml", "Start server with provided configuration file")
}

func run() error {
	cfg, err := bootstrap.LoadConfig(configYml)
	if err != nil {
		return err
	}

	if err := core.Init(cfg); err != nil {
		return err
	}

	i18n.Register(&codes.Code{
		EnableI18N: config.Get().Server.I18n,
		Lang:       config.Get().Server.Lang,
	})

	//初始化gin
	r := core.GetApp().GetGinEngine()
	middleware.InitMiddleware(r, config.Get())
	//初始化路由
	for _, f := range AppRouters {
		f()
	}
	go func() { //主服务启动后回调
		<-core.GetApp().WaitForStart()
		startedInit()
	}()

	go func() { //服务关闭释放资源
		<-core.GetApp().WaitForClose()
		toClose()

	}()
	if err := core.GetApp().Run(); err != nil {
		return err
	}
	slog.Info("Server exited")
	return nil
}

// 服务启动后要初始化的资源
func startedInit() {
	if config.Get().GrpcServer.Enable {
		grpcInit()
	}
	rdInit()
	slog.Debug("服务启动，初始化执行完成")
}

// 服务关闭要释放的资源
func toClose() {
	if config.Get().GrpcServer.Enable {
		closeGrpc()
	}
	rdRelease()
	slog.Debug("服务关闭需要释放的资源")
}
