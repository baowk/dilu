package tools

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/baowk/dilu-core/common/consts"

	"github.com/baowk/dilu-core/common/utils/files"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	cons "dilu/common/consts"

	"dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"
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
// @Router /api/tools/gen/preview/{tableId} [get]
func (e *Gen) Preview(c *gin.Context) {
	table := tools.GenTable{}
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
// @Router /api/tools/gen/code/{tableId} [get]
func (e *Gen) GenCode(c *gin.Context) {
	table := tools.GenTable{}
	id, err := strconv.Atoi(c.Param("tableId"))
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}

	table.TableId = id

	db, _, _ := GetDb(consts.DB_DEF)
	tab, _ := table.Get(db, false)
	tab.ApiRoot = cons.ApiRoot

	e.NOMethodsGen(c, tab)

	e.Ok(c, "Code generated successfully！")
}

func (e *Gen) GenApiToFile(c *gin.Context) {
	var dbname string
	table := tools.GenTable{}
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

func (e *Gen) NOMethodsGen(c *gin.Context, tab tools.GenTable) {

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

	// fmt.Println(1)
	// flag, err := files.PathExists(ROOT + tab.PackageName)
	// if err != nil {
	// 	e.Error(c, err)
	// 	return
	// }
	// if !flag {
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

	rt1, err := template.ParseFiles("resources/template/cmd_api.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
	}
	m := map[string]string{}
	m["modName"] = tab.PackageName
	var rb1 bytes.Buffer
	if err = rt1.Execute(&rb1, m); err != nil {
		fmt.Println(err)
	}

	files.FileCreate(rb1, "cmd/start/"+tab.PackageName+".go")
	rt2, err := template.ParseFiles("resources/template/router.template")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
	}
	var rb2 bytes.Buffer
	err = rt2.Execute(&rb2, m)
	if err != nil {
		fmt.Println(err)
	}

	files.FileCreate(rb2, ROOT+tab.PackageName+"/router/router.go")
	//}

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

func (e *Gen) genApiToFile(c *gin.Context, tab tools.GenTable) {
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
		tools.GenTable
		GenerateTime string
	}{tab, i})

	files.FileCreate(b1, "./cmd/migrate/migration/version-local/"+i+"_migrate.go")

}

// GenMenuAndApi
// @Summary 生成菜单
// @Description 生成菜单
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param tableId path int true "tableId"
// @Param menuPid path uint false "menuPid"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/tools/gen/menu/{tableId}/{menuPid} [get]
func (e *Gen) GenMenuAndApi(c *gin.Context) {
	table := tools.GenTable{}
	id, err := strconv.Atoi(c.Param("tableId"))
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	menuPid, err := strconv.Atoi(c.Param("menuPid"))
	if err != nil {
		menuPid = 0
	}

	table.TableId = id

	db, _, _ := GetDb(consts.DB_DEF)

	tab, _ := table.Get(db, true)
	tab.MLTBName = strings.Replace(tab.TBName, "_", "-", -1)

	tab.ApiRoot = cons.ApiRoot

	if menuPid == 0 {
		Mmenu := dto.SysMenuInsertReq{}
		Mmenu.Title = tab.TableComment
		Mmenu.Icon = "pass"
		Mmenu.Path = "/" + tab.MLTBName
		Mmenu.MenuType = "M"
		Mmenu.ParentId = 0
		Mmenu.NoCache = false
		Mmenu.Component = "Layout"
		Mmenu.Sort = 0
		Mmenu.Hidden = false
		Mmenu.CreateBy = 1
		service.SysMenuS.Insert(&Mmenu)
		menuPid = Mmenu.Id
	}

	Cmenu := dto.SysMenuInsertReq{}
	Cmenu.MenuName = tab.ClassName + "Manage"
	Cmenu.Title = tab.TableComment
	Cmenu.Icon = "pass"
	Cmenu.Path = "/" + tab.PackageName + "/" + tab.MLTBName
	Cmenu.MenuType = "C"
	Cmenu.Permission = tab.PackageName + ":" + tab.BusinessName + ":list"
	Cmenu.ParentId = menuPid
	Cmenu.NoCache = false
	Cmenu.Component = "/" + tab.PackageName + "/" + tab.MLTBName + "/index"
	Cmenu.Sort = 0
	Cmenu.Hidden = false
	Cmenu.CreateBy = 1
	Cmenu.UpdateBy = 1
	service.SysMenuS.Insert(&Cmenu)

	mApi := models.NewSysApi().SetMethod("POST").SetPermType("t").
		SetPath(fmt.Sprintf("%s/%s/%s/page", tab.ApiRoot, tab.PackageName, tab.ModuleName)).
		SetStatus(3).SetTitle("分页获取" + tab.TableComment)
	service.SysApiS.Create(mApi)

	gApi := models.NewSysApi().SetMethod("POST").SetPermType("t").
		SetPath(fmt.Sprintf("%s/%s/%s/get", tab.ApiRoot, tab.PackageName, tab.ModuleName)).
		SetStatus(3).SetTitle("根据id获取" + tab.TableComment)
	service.SysApiS.Create(gApi)

	MList := dto.SysMenuInsertReq{}
	MList.MenuName = ""
	MList.Title = "分页获取" + tab.TableComment
	MList.Icon = ""
	MList.Path = tab.TBName
	MList.MenuType = "F"
	MList.Permission = tab.PackageName + ":" + tab.BusinessName + ":query"
	MList.ParentId = Cmenu.Id
	MList.NoCache = false
	MList.Sort = 0
	MList.Hidden = false
	MList.CreateBy = 1
	MList.UpdateBy = 1
	MList.SysApi = []models.SysApi{*mApi, *gApi}
	service.SysMenuS.Insert(&MList)

	cApi := models.NewSysApi().SetMethod("POST").SetPermType("t").
		SetPath(fmt.Sprintf("%s/%s/%s/create", tab.ApiRoot, tab.PackageName, tab.ModuleName)).
		SetStatus(3).SetTitle("创建" + tab.TableComment)
	service.SysApiS.Create(cApi)

	MCreate := dto.SysMenuInsertReq{}
	MCreate.MenuName = ""
	MCreate.Title = "创建" + tab.TableComment
	MCreate.Icon = ""
	MCreate.Path = tab.TBName
	MCreate.MenuType = "F"
	MCreate.Permission = tab.PackageName + ":" + tab.BusinessName + ":add"
	MCreate.ParentId = Cmenu.Id
	MCreate.NoCache = false
	MCreate.Sort = 0
	MCreate.Hidden = false
	MCreate.CreateBy = 1
	MCreate.UpdateBy = 1
	MCreate.SysApi = []models.SysApi{*cApi}
	service.SysMenuS.Insert(&MCreate)

	uApi := models.NewSysApi().SetMethod("POST").SetPermType("t").
		SetPath(fmt.Sprintf("%s/%s/%s/update", tab.ApiRoot, tab.PackageName, tab.ModuleName)).
		SetStatus(3).SetTitle("修改" + tab.TableComment)
	service.SysApiS.Create(uApi)

	MUpdate := dto.SysMenuInsertReq{}
	MUpdate.MenuName = ""
	MUpdate.Title = "修改" + tab.TableComment
	MUpdate.Icon = ""
	MUpdate.Path = tab.TBName
	MUpdate.MenuType = "F"
	MUpdate.Permission = tab.PackageName + ":" + tab.BusinessName + ":edit"
	MUpdate.ParentId = Cmenu.Id
	MUpdate.NoCache = false
	MUpdate.Sort = 0
	MUpdate.Hidden = false
	MUpdate.CreateBy = 1
	MUpdate.UpdateBy = 1
	MUpdate.SysApi = []models.SysApi{*uApi}
	service.SysMenuS.Insert(&MUpdate)

	dApi := models.NewSysApi().SetMethod("POST").SetPermType("t").
		SetPath(fmt.Sprintf("%s/%s/%s/del", tab.ApiRoot, tab.PackageName, tab.ModuleName)).
		SetStatus(3).SetTitle("删除" + tab.TableComment)
	service.SysApiS.Create(dApi)

	MDelete := dto.SysMenuInsertReq{}
	MDelete.MenuName = ""
	MDelete.Title = "删除" + tab.TableComment
	MDelete.Icon = ""
	MDelete.Path = tab.TBName
	MDelete.MenuType = "F"
	MDelete.Permission = tab.PackageName + ":" + tab.BusinessName + ":remove"
	MDelete.ParentId = Cmenu.Id
	MDelete.NoCache = false
	MDelete.Sort = 0
	MDelete.Hidden = false
	MDelete.CreateBy = 1
	MDelete.UpdateBy = 1
	MUpdate.SysApi = []models.SysApi{*dApi}
	service.SysMenuS.Insert(&MDelete)

	e.Ok(c, "数据生成成功！")
}

// SaveSysApi
// @Summary 生成Api
// @Description 生成Api
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/tools/gen/api [get]
func (e *Gen) GenApis(c *gin.Context) {
	data, err := os.ReadFile("docs/swagger.json")
	if err != nil {
		e.Error(c, err)
		return
	}
	basePath, err := jsonparser.GetString(data, "basePath")
	if err != nil {
		fmt.Println(err)
		basePath = ""
	}
	db := core.DB()
	jsonparser.ObjectEach(data, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		path := basePath + string(key)
		if !strings.HasPrefix(path, "/api/tools/") {
			if reg.MatchString(path) {
				path = reg.ReplaceAllString(path, "${1}/{${2}}") // 把:id换成{id}
			}
			jsonparser.ObjectEach(value, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
				method := strings.ToUpper(string(key))
				apiTitle, _ := jsonparser.GetString(value, "summary")
				if apiTitle == "" {
					apiTitle, _ = jsonparser.GetString(value, "description")
				}

				pt := "n"
				if token, _, _, err := jsonparser.Get(value, "security"); err == nil {
					if len(token) > 0 {
						pt = "t"
					}
				}

				err := db.Debug().Where(models.SysApi{Path: path, Method: method}).
					Attrs(models.SysApi{Title: apiTitle, Status: 3, PermType: pt}).
					FirstOrCreate(&models.SysApi{}).Error
				if err != nil {
					e.Error(c, err)
					return err
				}
				return nil
			})
		}
		return nil
	}, "paths")
	e.Ok(c)
}

var idPatten = "(.*)/:(\\w+)" // 正则替换，把:id换成{id}
var reg, _ = regexp.Compile(idPatten)
