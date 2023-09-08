package start

import (
	"dilu/common/codes"
	"dilu/common/config"
	"fmt"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/i18n"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	//AppRouters = make([]func(), 0)
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
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "resources/config.dev.yml", "Start server with provided configuration file")
}

func run() {
	if configYml == "" {
		panic("找不到配置文件")
	}
	v := viper.New()
	v.SetConfigFile(configYml)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("Fatal error config file: %v \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&core.Cfg); err != nil {
			fmt.Println(err)
		}
		v.Sub("extend").Unmarshal(&config.Ext)
	})
	if err = v.Unmarshal(&core.Cfg); err != nil {
		fmt.Println(err)
	}

	v.Sub("extend").Unmarshal(&config.Ext)

	core.Init()

	i18n.Register(&codes.Code{
		EnableI18N: core.Cfg.Server.I18n,
		Lang:       core.Cfg.Server.Lang,
	})

	core.Run(&AppRouters)
}
