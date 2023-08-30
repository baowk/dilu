package tools

import (
	"strings"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core"
	"gorm.io/gorm"
)

func GetDb(dbname string) (db *gorm.DB, mdb string, sdb string) {
	mdsn := core.Cfg.DBCfg.DSN
	mdb = ParseDsn(mdsn)
	if dbname != consts.DB_DEF {
		gdsn, ok := core.Cfg.DBCfg.DBS[dbname]
		if !ok {
			return
		}
		sdb = ParseDsn(gdsn.DSN)
	} else {
		sdb = mdb
	}
	db = core.Db(dbname)
	return
}

func ParseDsn(dsn string) string {
	idx := strings.LastIndex(dsn, ")/")
	end := strings.LastIndex(dsn, "?")
	if end < 0 {
		end = len(dsn)
	}
	return dsn[idx+2 : end]
}
