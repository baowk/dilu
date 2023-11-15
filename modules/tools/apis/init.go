package apis

import (
	dm "dilu/modules/dental/models"
	"dilu/modules/sys/models"
	tm "dilu/modules/tools/models"

	"fmt"
	"text/template"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type Init struct {
	base.BaseApi
}

// func (Init) Init(c *gin.Context) {
// 	t1, err := template.ParseFiles("modules/tools/apis/tmpls/index.html")
// 	if err != nil {
// 		panic(err)
// 	}
// 	t1.Execute(c.Writer, "")
// }

func (Init) DoInit(c *gin.Context) {
	fmt.Println("开始运行初始化")

	// service.ImportSql("resources/dbs/dilu-db.sql", "sys")
	// service.ImportSql("resources/dbs/dental-db.sql", "dental")

	result := "执行成功"
	if err := core.DB().AutoMigrate(
		&models.SysEmail{},
		&models.SysSms{},
		&models.SysApi{},
		&models.SysCfg{},
		&models.SysDept{},
		&models.SysJob{},
		&models.SysMenu{},
		&models.SysOperaLog{},
		&models.SysRole{},
		&models.SysUser{},
		&models.ThirdLogin{},
		&models.SysMember{},
		&models.SysTeam{},
		&models.SysRoleMenu{},
		&tm.GenTables{},
		&tm.GenColumns{},
		// &tools.GenColumn{},
		// &tools.GenTable{},
	); err != nil {
		result = "sys执行失败"
	}
	if err := core.Db("dental").AutoMigrate(
		&dm.Bill{},
		&dm.Customer{},
		&dm.EventDaySt{},
		&dm.SummaryPlanDay{},
		&dm.TargetTask{},
	); err != nil {
		result = "dental执行失败"
	}
	t1, err := template.ParseFiles("modules/tools/apis/tmpls/result.html")
	if err != nil {
		panic(err)
	}

	t1.Execute(c.Writer, result)
}
