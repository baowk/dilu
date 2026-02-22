package gen

import (
	"dilu/common/config"
	cons "dilu/common/consts"
	"dilu/modules/tools/apis"
	"dilu/modules/tools/service"
	"fmt"
	"time"

	"github.com/baowk/dilu-core/core"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var (
	configYml   string
	dbName      string
	tableName   string
	packageName string
	force       bool
	GenCmd      = &cobra.Command{
		Use:     "gen",
		Short:   "Generate code",
		Long:    "Generate code based on database tables",
		Example: "dilu gen -c resources/config.dev.yml -d sys -t sys_users",
		Run: func(cmd *cobra.Command, args []string) {
			gen()
		},
	}
)

func init() {
	GenCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "resources/config.dev.yml", "Start server with provided configuration file")
	GenCmd.PersistentFlags().StringVarP(&dbName, "db", "d", "", "database name")
	GenCmd.PersistentFlags().StringVarP(&tableName, "table", "t", "", "table name")
	GenCmd.PersistentFlags().StringVarP(&packageName, "package", "p", "", "package name")
	GenCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "if set to true, will overwrite existing files")
}

func gen() {
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

	var cfg config.Config

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
		var remoteCfg config.Config
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
			if err = v.Unmarshal(cfg); err != nil {
				fmt.Println(err)
			}
			config.SaveConfig(&cfg, nil)
		})
	}

	if err := core.Init(&cfg); err != nil {
		panic(err)
	}

	// 生成
	fmt.Printf("packageName %s db %s table %s\n", packageName, dbName, tableName)
	tab, err := service.SerGenTables.GenTableInit(dbName, tableName, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 自定义包名
	if packageName != "" {
		tab.PackageName = packageName
	}
	tab.ApiRoot = cons.ApiRoot

	for i, v := range tab.Columns {
		tab.Columns[i].TsType = apis.TypeGo2Ts(v.GoType)
	}
	service.SerGenTables.NOMethodsGen(tab, force)
}
