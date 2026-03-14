package gen

import (
	cons "dilu/common/consts"
	"dilu/internal/bootstrap"
	"dilu/internal/tools/apis"
	"dilu/internal/tools/service"
	"fmt"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core"
	"github.com/spf13/cobra"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return gen()
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

func gen() error {
	cfg, err := bootstrap.LoadConfig(configYml)
	if err != nil {
		return err
	}

	if err := core.Init(cfg); err != nil {
		return err
	}

	// 生成
	// 如果未指定数据库名，使用默认值 sys
	if dbName == "" {
		dbName = consts.DB_DEF
	}
	tab, err := service.SerGenTables.GenTableInit(dbName, tableName, true)
	if err != nil {
		fmt.Println(err)
		return err
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
	return nil
}
