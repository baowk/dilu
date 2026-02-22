package start

import (
	"dilu/common/codes"
	"dilu/common/config"
	"dilu/common/middleware"
	"fmt"
	"log/slog"
	"time"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/i18n"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:     "start",
		Short:   "Get Application config info",
		Example: "dilu start -c resources/config.dev.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "resources/config.dev.yaml", "Start server with provided configuration file")
}

func run() {
	if configYml == "" {
		panic("找不到配置文件")
	}
	v := viper.New()
	v.SetConfigFile(configYml)
	//v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %v \n", err))
	}

	var cfg config.Config // 修改为使用Extend结构体

	if err = v.Unmarshal(&cfg); err != nil {
		fmt.Println(err)
	}

	if cfg.Server.RemoteEnable {
		rviper := viper.New()
		if cfg.Remote.SecretKeyring == "" {
			err = rviper.AddRemoteProvider(cfg.Remote.Provider, cfg.Remote.Endpoint, cfg.Remote.Path)
		} else {
			err = rviper.AddSecureRemoteProvider(cfg.Remote.Provider, cfg.Remote.Endpoint, cfg.Remote.Path, cfg.Remote.SecretKeyring)
		}
		if err != nil {
			panic(fmt.Sprintf("Fatal error remote config : %v \n", err))
		}
		rviper.SetConfigType(cfg.Remote.GetConfigType())
		err = rviper.ReadRemoteConfig()
		if err != nil {
			panic(fmt.Sprintf("Fatal error remote config : %v \n", err))
		}
		var remoteCfg config.Config // 修改为使用Extend结构体
		rviper.Unmarshal(&remoteCfg)

		config.SaveConfig(&cfg, &remoteCfg)
		go func() {
			for {
				time.Sleep(time.Second * 5) // delay after each request
				err := rviper.WatchRemoteConfig()
				if err != nil {
					fmt.Println(err)
					continue
				}
				rviper.Unmarshal(&remoteCfg)

				config.SaveConfig(&cfg, &remoteCfg)
			}
		}()
	} else {
		config.SaveConfig(&cfg, nil)
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.String())
			if err = v.Unmarshal(&cfg); err != nil {
				fmt.Println(err)
			}
			config.SaveConfig(&cfg, nil)
		})
	}

	if err := core.Init(&cfg); err != nil {
		panic(err)
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
	core.GetApp().Run()
	slog.Info("Server exited")
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
