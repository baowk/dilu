package gen

import (
	"dilu/common/config"
	cons "dilu/common/consts"
	"dilu/modules/tools/service"
	"fmt"
	"strings"
	"time"

	coreCfg "github.com/baowk/dilu-core/config"

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

	var cfg coreCfg.AppCfg

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
		var remoteCfg coreCfg.AppCfg
		rviper.Unmarshal(&remoteCfg)

		mergeCfg(&cfg, &remoteCfg)

		extend := rviper.Sub("extend")
		if extend != nil {
			extend.Unmarshal(config.Ext)
		}
		go func() {
			for {
				time.Sleep(time.Second * 5) // delay after each request
				err := rviper.WatchRemoteConfig()
				if err != nil {
					fmt.Println(err)
					continue
				}
				rviper.Unmarshal(&remoteCfg)

				mergeCfg(&cfg, &remoteCfg)

				extend := rviper.Sub("extend")
				if extend != nil {
					extend.Unmarshal(config.Ext)
				}
			}
		}()
	} else {
		mergeCfg(&cfg, nil)
		v.Sub("extend").Unmarshal(&config.Ext)
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.String())
			if err = v.Unmarshal(cfg); err != nil {
				fmt.Println(err)
			}
			mergeCfg(&cfg, nil)
			extend := v.Sub("extend")
			if extend != nil {
				extend.Unmarshal(config.Ext)
			}
		})
	}

	core.Init()

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
		tab.Columns[i].TsType = TypeGo2Ts(v.GoType)
	}
	service.SerGenTables.NOMethodsGen(tab, force)
}

func TypeGo2Ts(t string) string {
	if strings.Contains(t, "int") {
		return "number"
	} else if strings.Contains(t, "time") {
		return "Date"
	} else if strings.Contains(t, "bool") {
		return "boolean"
	} else {
		return t
	}
}

func mergeCfg(local, remote *coreCfg.AppCfg) {
	if remote != nil {
		core.Cfg = *local
		core.Cfg = *remote
		core.Cfg.Server.Mode = local.Server.Mode
		core.Cfg.Server.RemoteEnable = local.Server.RemoteEnable
		core.Cfg.Remote = local.Remote
		core.Cfg.Server.Name = local.Server.Name
		core.Cfg.Server.Port = local.Server.Port
		core.Cfg.Server.Host = local.Server.Host
	} else {
		core.Cfg = *local
	}
}
