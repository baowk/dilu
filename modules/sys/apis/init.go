package apis

import (
	"dilu/modules/sys/models"
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
	t1, err := template.ParseFiles("modules/sys/apis/tmpls/index.html")
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
		&models.ThirdLogin{},
		&models.User{},
	)
	result := "执行成功"
	if err != nil {
		result = "执行失败"
	}
	t1, err := template.ParseFiles("modules/sys/apis/tmpls/result.html")
	if err != nil {
		panic(err)
	}

	t1.Execute(c.Writer, result)
}
