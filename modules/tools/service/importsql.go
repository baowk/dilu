package service

import (
	"os"
	"strings"

	"github.com/baowk/dilu-core/core"
)

func ImportSql(sqlFile, dbName string) error {
	_, err := os.Stat(sqlFile)
	if os.IsNotExist(err) {
		core.GetApp().GetLogger().Error("Sql 文件不存在", "err", err)
		return err
	}
	db := core.GetApp().Db(dbName)
	sqls, _ := os.ReadFile(sqlFile)
	sqlArr := strings.Split(string(sqls), ";")
	for _, sql := range sqlArr {
		sql = strings.TrimSpace(sql)
		if sql == "" {
			continue
		}
		err := db.Exec(sql).Error
		if err != nil {
			core.GetApp().GetLogger().Error("数据库导入失败:", "err", err)
			return err
		} else {
			core.GetApp().GetLogger().Info("[success]", "sql", sql)
		}
	}
	return nil
}
