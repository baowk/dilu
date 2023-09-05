package tools

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/common/utils/files"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"dilu/modules/tools/models/tools"
)

type Gen struct {
	base.BaseApi
}

var (
	FrontPath = "../vue"
)

// Preview
// @Summary 生成预览
// @Description 生成预览
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param tableId path int true "tableId"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/tools/gen/preview/{tableId} [get]
func (e Gen) Preview(c *gin.Context) {
	table := tools.SysTables{}
	id, err := strconv.Atoi(c.Param("tableId"))
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	table.TableId = id
	t1, err := template.ParseFiles("resources/template/v4/model.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t2, err := template.ParseFiles("resources/template/v4/no_actions/apis.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t3, err := template.ParseFiles("resources/template/v4/js.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t4, err := template.ParseFiles("resources/template/v4/vue.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t5, err := template.ParseFiles("resources/template/v4/no_actions/router_no_check_role.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t6, err := template.ParseFiles("resources/template/v4/dto.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t7, err := template.ParseFiles("resources/template/v4/no_actions/service.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}

	db, _, _ := GetDb(consts.DB_DEF)

	tab, _ := table.Get(db, false)
	var b1 bytes.Buffer
	err = t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	err = t3.Execute(&b3, tab)
	var b4 bytes.Buffer
	err = t4.Execute(&b4, tab)
	var b5 bytes.Buffer
	err = t5.Execute(&b5, tab)
	var b6 bytes.Buffer
	err = t6.Execute(&b6, tab)
	var b7 bytes.Buffer
	err = t7.Execute(&b7, tab)

	mp := make(map[string]interface{})
	mp["resources/template/model.go.template"] = b1.String()
	mp["resources/template/api.go.template"] = b2.String()
	mp["resources/template/js.go.template"] = b3.String()
	mp["resources/template/vue.go.template"] = b4.String()
	mp["resources/template/router.go.template"] = b5.String()
	mp["resources/template/dto.go.template"] = b6.String()
	mp["resources/template/service.go.template"] = b7.String()
	e.Ok(c, mp)
}

// GenCode
// @Summary 生成代码
// @Description 生成代码
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param tableId path int true "tableId"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/tools/gen/code/{tableId} [get]
func (e Gen) GenCode(c *gin.Context) {
	table := tools.SysTables{}
	id, err := strconv.Atoi(c.Param("tableId"))
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}

	table.TableId = id

	db, _, _ := GetDb(consts.DB_DEF)
	tab, _ := table.Get(db, false)

	e.NOActionsGen(c, tab)

	e.Ok(c, "Code generated successfully！")
}

func (e Gen) GenApiToFile(c *gin.Context) {
	var dbname string
	table := tools.SysTables{}
	id, err := strconv.Atoi(c.Param("tableId"))
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}

	table.TableId = id

	db, _, _ := GetDb(dbname)
	tab, _ := table.Get(db, false)
	e.genApiToFile(c, tab)

	e.Ok(c, "Code generated successfully！")
}

const ROOT = "./modules/"

func (e Gen) NOActionsGen(c *gin.Context, tab tools.SysTables) {

	tab.MLTBName = strings.Replace(tab.TBName, "_", "-", -1)

	basePath := "resources/template/v4/"
	// routerFile := basePath + "no_actions/router_check_role.go.template"

	// if tab.IsAuth == 2 {
	routerFile := basePath + "no_actions/router_no_check_role.go.template"
	//}

	t1, err := template.ParseFiles(basePath + "model.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t2, err := template.ParseFiles(basePath + "no_actions/apis.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t3, err := template.ParseFiles(routerFile)
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t4, err := template.ParseFiles(basePath + "js.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t5, err := template.ParseFiles(basePath + "vue.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t6, err := template.ParseFiles(basePath + "dto.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	t7, err := template.ParseFiles(basePath + "no_actions/service.go.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}

	fmt.Println(1)
	flag, err := files.PathExists(ROOT + tab.PackageName)
	if err != nil {
		e.Error(c, err)
		return
	}
	if !flag {
		_ = files.PathCreate(ROOT + tab.PackageName + "/apis/")
		_ = files.PathCreate(ROOT + tab.PackageName + "/models/")
		_ = files.PathCreate(ROOT + tab.PackageName + "/router/")
		_ = files.PathCreate(ROOT + tab.PackageName + "/service/dto/")
		_ = files.PathCreate(FrontPath + "/api/" + tab.PackageName + "/")
		err = files.PathCreate(FrontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}

		t1, err := template.ParseFiles("resources/template/cmd_api.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
		}
		m := map[string]string{}
		m["modName"] = tab.PackageName
		var b1 bytes.Buffer
		if err = t1.Execute(&b1, m); err != nil {
			fmt.Println(err)
		}

		files.FileCreate(b1, "cmd/start/"+tab.PackageName+".go")
		t2, err := template.ParseFiles("resources/template/router.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
		}
		var b2 bytes.Buffer
		err = t2.Execute(&b2, m)
		if err != nil {
			fmt.Println(err)
		}

		files.FileCreate(b2, ROOT+tab.PackageName+"/router/router.go")
	}

	var b1 bytes.Buffer
	err = t1.Execute(&b1, tab)
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)
	var b3 bytes.Buffer
	err = t3.Execute(&b3, tab)
	var b4 bytes.Buffer
	err = t4.Execute(&b4, tab)
	var b5 bytes.Buffer
	err = t5.Execute(&b5, tab)
	var b6 bytes.Buffer
	err = t6.Execute(&b6, tab)
	var b7 bytes.Buffer
	err = t7.Execute(&b7, tab)
	files.FileCreate(b1, ROOT+tab.PackageName+"/models/"+tab.TBName+".go")
	files.FileCreate(b2, ROOT+tab.PackageName+"/apis/"+tab.TBName+".go")
	files.FileCreate(b3, ROOT+tab.PackageName+"/router/"+tab.TBName+".go")
	files.FileCreate(b4, FrontPath+"/api/"+tab.PackageName+"/"+tab.MLTBName+".js")
	files.FileCreate(b5, FrontPath+"/views/"+tab.PackageName+"/"+tab.MLTBName+"/index.vue")
	files.FileCreate(b6, ROOT+tab.PackageName+"/service/dto/"+tab.TBName+".go")
	files.FileCreate(b7, ROOT+tab.PackageName+"/service/"+tab.TBName+".go")

}

func (e Gen) genApiToFile(c *gin.Context, tab tools.SysTables) {
	basePath := "resources/template/"
	t1, err := template.ParseFiles(basePath + "api_migrate.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	i := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	var b1 bytes.Buffer
	err = t1.Execute(&b1, struct {
		tools.SysTables
		GenerateTime string
	}{tab, i})

	files.FileCreate(b1, "./cmd/migrate/migration/version-local/"+i+"_migrate.go")

}

func (e Gen) GenMenuAndApi(c *gin.Context) {
	// var dbname string
	// table := tools.SysTables{}
	// id, err := strconv.Atoi(c.Param("tableId"))
	// if err != nil {
	// 	core.Log.Error("Gen", zap.Error(err))
	// 	e.Error(c, err)
	// 	return
	// }

	// table.TableId = id

	// db, _, _ := GetDb(dbname)

	// tab, _ := table.Get(db, true)
	// tab.MLTBName = strings.Replace(tab.TBName, "_", "-", -1)

	// Mmenu := dto.SysMenuInsertReq{}
	// Mmenu.Title = tab.TableComment
	// Mmenu.Icon = "pass"
	// Mmenu.Path = "/" + tab.MLTBName
	// Mmenu.MenuType = "M"
	// Mmenu.Action = "无"
	// Mmenu.ParentId = 0
	// Mmenu.NoCache = false
	// Mmenu.Component = "Layout"
	// Mmenu.Sort = 0
	// Mmenu.Visible = "0"
	// Mmenu.IsFrame = "0"
	// Mmenu.CreateBy = 1
	// s.Insert(&Mmenu)

	// Cmenu := dto.SysMenuInsertReq{}
	// Cmenu.MenuName = tab.ClassName + "Manage"
	// Cmenu.Title = tab.TableComment
	// Cmenu.Icon = "pass"
	// Cmenu.Path = "/" + tab.PackageName + "/" + tab.MLTBName
	// Cmenu.MenuType = "C"
	// Cmenu.Action = "无"
	// Cmenu.Permission = tab.PackageName + ":" + tab.BusinessName + ":list"
	// Cmenu.ParentId = Mmenu.MenuId
	// Cmenu.NoCache = false
	// Cmenu.Component = "/" + tab.PackageName + "/" + tab.MLTBName + "/index"
	// Cmenu.Sort = 0
	// Cmenu.Visible = "0"
	// Cmenu.IsFrame = "0"
	// Cmenu.CreateBy = 1
	// Cmenu.UpdateBy = 1
	// s.Insert(&Cmenu)

	// MList := dto.SysMenuInsertReq{}
	// MList.MenuName = ""
	// MList.Title = "分页获取" + tab.TableComment
	// MList.Icon = ""
	// MList.Path = tab.TBName
	// MList.MenuType = "F"
	// MList.Action = "无"
	// MList.Permission = tab.PackageName + ":" + tab.BusinessName + ":query"
	// MList.ParentId = Cmenu.MenuId
	// MList.NoCache = false
	// MList.Sort = 0
	// MList.Visible = "0"
	// MList.IsFrame = "0"
	// MList.CreateBy = 1
	// MList.UpdateBy = 1
	// s.Insert(&MList)

	// MCreate := dto.SysMenuInsertReq{}
	// MCreate.MenuName = ""
	// MCreate.Title = "创建" + tab.TableComment
	// MCreate.Icon = ""
	// MCreate.Path = tab.TBName
	// MCreate.MenuType = "F"
	// MCreate.Action = "无"
	// MCreate.Permission = tab.PackageName + ":" + tab.BusinessName + ":add"
	// MCreate.ParentId = Cmenu.MenuId
	// MCreate.NoCache = false
	// MCreate.Sort = 0
	// MCreate.Visible = "0"
	// MCreate.IsFrame = "0"
	// MCreate.CreateBy = 1
	// MCreate.UpdateBy = 1
	// s.Insert(&MCreate)

	// MUpdate := dto.SysMenuInsertReq{}
	// MUpdate.MenuName = ""
	// MUpdate.Title = "修改" + tab.TableComment
	// MUpdate.Icon = ""
	// MUpdate.Path = tab.TBName
	// MUpdate.MenuType = "F"
	// MUpdate.Action = "无"
	// MUpdate.Permission = tab.PackageName + ":" + tab.BusinessName + ":edit"
	// MUpdate.ParentId = Cmenu.MenuId
	// MUpdate.NoCache = false
	// MUpdate.Sort = 0
	// MUpdate.Visible = "0"
	// MUpdate.IsFrame = "0"
	// MUpdate.CreateBy = 1
	// MUpdate.UpdateBy = 1
	// s.Insert(&MUpdate)

	// MDelete := dto.SysMenuInsertReq{}
	// MDelete.MenuName = ""
	// MDelete.Title = "删除" + tab.TableComment
	// MDelete.Icon = ""
	// MDelete.Path = tab.TBName
	// MDelete.MenuType = "F"
	// MDelete.Action = "无"
	// MDelete.Permission = tab.PackageName + ":" + tab.BusinessName + ":remove"
	// MDelete.ParentId = Cmenu.MenuId
	// MDelete.NoCache = false
	// MDelete.Sort = 0
	// MDelete.Visible = "0"
	// MDelete.IsFrame = "0"
	// MDelete.CreateBy = 1
	// MDelete.UpdateBy = 1
	// s.Insert(&MDelete)

	// e.Ok(c, "数据生成成功！")
}
