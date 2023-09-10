package apis

import (
	"dilu/modules/sys/models"
	"dilu/modules/tools/models/tools"
	"fmt"
	"text/template"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
)

type Init struct {
	base.BaseApi
}

func (Init) Init(c *gin.Context) {
	t1, err := template.ParseFiles("modules/tools/apis/tmpls/index.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(c.Writer, "")
}

func (Init) DoInit(c *gin.Context) {
	fmt.Println("开始运行初始化")
	err := core.DB().AutoMigrate(
		&models.EmailLog{},
		&models.SmsLog{},
		&models.SysApi{},
		&models.SysCfg{},
		&models.SysDept{},
		&models.SysJob{},
		&models.SysMenu{},
		&models.SysOperaLog{},
		&models.SysRole{},
		&models.SysUser{},
		&tools.GenColumn{},
		&tools.GenColumn{},
		//&models.SysPost{},
	)
	result := "执行成功"
	if err != nil {
		result = "执行失败"
	}
	t1, err := template.ParseFiles("modules/tools/apis/tmpls/result.html")
	if err != nil {
		panic(err)
	}

	t1.Execute(c.Writer, result)
}
