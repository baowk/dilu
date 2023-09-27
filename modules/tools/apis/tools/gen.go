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
	if err != nil {
		core.Log.Error("gen err", zap.Error(err))
	}
	var b2 bytes.Buffer
	err = t2.Execute(&b2, tab)
	if err != nil {
		core.Log.Error("gen err", zap.Error(err))
	}
	var b3 bytes.Buffer
	err = t3.Execute(&b3, tab)
	if err != nil {
		core.Log.Error("gen err", zap.Error(err))
	}
	var b4 bytes.Buffer
	err = t4.Execute(&b4, tab)
	if err != nil {
		core.Log.Error("gen err", zap.Error(err))
	}
	var b5 bytes.Buffer
	err = t5.Execute(&b5, tab)
	if err != nil {
		core.Log.Error("gen err", zap.Error(err))
	}
	var b6 bytes.Buffer
	err = t6.Execute(&b6, tab)
	if err != nil {
		core.Log.Error("gen err", zap.Error(err))
	}
	var b7 bytes.Buffer
	err = t7.Execute(&b7, tab)
	if err != nil {
		core.Log.Error("gen err", zap.Error(err))
	}

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
// @Param force path string false "force"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/tools/gen/code/{tableId}/{force} [get]
func (e *Gen) GenCode(c *gin.Context) {
	table := tools.GenTable{}
	id, err := strconv.Atoi(c.Param("tableId"))
	if err != nil {
		e.Error(c, err)
		return
	}
	force, err := strconv.ParseBool(c.Param("force"))
	if err != nil {
		force = false
	}
	fmt.Println(force)
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}

	table.TableId = id

	db, _, _ := GetDb(consts.DB_DEF)
	tab, _ := table.Get(db, false)
	tab.ApiRoot = cons.ApiRoot

	e.NOMethodsGen(c, tab, force)

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

func (e *Gen) NOMethodsGen(c *gin.Context, tab tools.GenTable, force bool) {

	tab.MLTBName = strings.Replace(tab.TBName, "_", "-", -1)

	basePath := "resources/template/v4/"

	_ = files.PathCreate(ROOT + tab.PackageName + "/apis/")
	_ = files.PathCreate(ROOT + tab.PackageName + "/models/")
	_ = files.PathCreate(ROOT + tab.PackageName + "/router/")
	_ = files.PathCreate(ROOT + tab.PackageName + "/service/dto/")
	_ = files.PathCreate(FrontPath + "/api/" + tab.PackageName + "/")
	err := files.PathCreate(FrontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/utils")
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}

	m := map[string]string{}
	m["modName"] = tab.PackageName

	//路由
	cmdApi := "cmd/start/" + tab.PackageName + ".go"
	if files.CheckExist(cmdApi) || force {
		rt1, err := template.ParseFiles("resources/template/cmd_api.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
		}
		var rb1 bytes.Buffer
		if err = rt1.Execute(&rb1, m); err != nil {
			fmt.Println(err)
		}
		files.FileCreate(rb1, cmdApi)
	}

	baseRouter := ROOT + tab.PackageName + "/router/router.go"
	if files.CheckExist(baseRouter) || force {
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
		files.FileCreate(rb2, baseRouter)
	}

	//golang

	modelgo := ROOT + tab.PackageName + "/models/" + tab.TBName + ".go"
	if files.CheckExist(modelgo) || force {
		t1, err := template.ParseFiles(basePath + "model.go.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b1 bytes.Buffer
		err = t1.Execute(&b1, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b1, modelgo)
	}

	apigo := ROOT + tab.PackageName + "/apis/" + tab.TBName + ".go"
	if files.CheckExist(apigo) || force {
		t2, err := template.ParseFiles(basePath + "no_actions/apis.go.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b2 bytes.Buffer
		err = t2.Execute(&b2, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b2, apigo)
	}

	routergo := ROOT + tab.PackageName + "/router/" + tab.TBName + ".go"
	if files.CheckExist(routergo) || force {
		routerFile := basePath + "no_actions/router_no_check_role.go.template"
		t3, err := template.ParseFiles(routerFile)
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b3 bytes.Buffer
		err = t3.Execute(&b3, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b3, routergo)
	}

	dto := ROOT + tab.PackageName + "/service/dto/" + tab.TBName + ".go"
	if files.CheckExist(dto) || force {
		t6, err := template.ParseFiles(basePath + "dto.go.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b6 bytes.Buffer
		err = t6.Execute(&b6, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b6, dto)
	}

	service := ROOT + tab.PackageName + "/service/" + tab.TBName + ".go"
	if files.CheckExist(service) || force {
		t7, err := template.ParseFiles(basePath + "no_actions/service.go.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b7 bytes.Buffer
		err = t7.Execute(&b7, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b7, service)
	}

	//前端部分
	js := FrontPath + "/api/" + tab.PackageName + "/" + tab.MLTBName + ".ts"
	if files.CheckExist(js) || force {
		t4, err := template.ParseFiles(basePath + "vue/api.ts.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b4 bytes.Buffer
		err = t4.Execute(&b4, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b4, js)
	}

	vue := FrontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/index.vue"
	if files.CheckExist(vue) || force {
		t5, err := template.ParseFiles(basePath + "vue/index.vue.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b5 bytes.Buffer
		err = t5.Execute(&b5, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b5, vue)
	}

	form := FrontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/form.vue"
	if files.CheckExist(form) || force {
		t5, err := template.ParseFiles(basePath + "vue/form.vue.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b5 bytes.Buffer
		err = t5.Execute(&b5, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b5, form)
	}

	hook := FrontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/utils/hook.tsx"
	if files.CheckExist(hook) || force {
		t5, err := template.ParseFiles(basePath + "vue/utils/hook.tsx.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b5 bytes.Buffer
		err = t5.Execute(&b5, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b5, hook)
	}

	rule := FrontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/utils/rule.ts"
	if files.CheckExist(rule) || force {
		t5, err := template.ParseFiles(basePath + "vue/utils/rule.ts.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b5 bytes.Buffer
		err = t5.Execute(&b5, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b5, rule)
	}

	types := FrontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/utils/types.ts"
	if files.CheckExist(types) || force {
		t5, err := template.ParseFiles(basePath + "vue/utils/types.ts.template")
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		var b5 bytes.Buffer
		err = t5.Execute(&b5, tab)
		if err != nil {
			core.Log.Error("gen err", zap.Error(err))
		}
		files.FileCreate(b5, types)
	}

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
	if err != nil {
		core.Log.Error("gen err", zap.Error(err))
	}

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
		Mmenu := models.SysMenu{}
		Mmenu.Title = tab.TableComment
		Mmenu.Icon = "pass"
		Mmenu.Path = "/" + tab.PackageName
		Mmenu.MenuType = 1
		Mmenu.ParentId = 0
		Mmenu.NoCache = false
		Mmenu.Component = "Layout"
		Mmenu.Sort = 0
		Mmenu.Hidden = false
		Mmenu.CreateBy = 1
		service.SerSysMenu.Insert(&Mmenu)
		menuPid = Mmenu.Id
	}

	curPath := fmt.Sprintf("%s/%s/%s/page", tab.ApiRoot, tab.PackageName, tab.ModuleName)
	where := map[string]any{
		"path":   curPath,
		"method": "POST",
	}
	mApi := models.NewSysApi()
	service.SerSysApi.GetByWhere(where, mApi)
	if mApi.Id == 0 {
		mApi = mApi.SetMethod("POST").SetPermType("t").SetPath(curPath).
			SetStatus(3).SetTitle("分页获取" + tab.TableComment)
		service.SerSysApi.Create(mApi)
	}

	Cmenu := models.SysMenu{}
	Cmenu.MenuName = tab.ClassName + "Manage"
	Cmenu.Title = tab.TableComment + "管理"
	Cmenu.Icon = "pass"
	Cmenu.Path = "/" + tab.PackageName + "/" + tab.MLTBName
	Cmenu.MenuType = 2
	Cmenu.Permission = tab.PackageName + ":" + tab.BusinessName + ":list"
	Cmenu.ParentId = menuPid
	Cmenu.NoCache = false
	Cmenu.Component = "/" + tab.PackageName + "/" + tab.MLTBName + "/index"
	Cmenu.Sort = 0
	Cmenu.Hidden = false
	Cmenu.CreateBy = 1
	Cmenu.UpdateBy = 1
	Cmenu.SysApi = []models.SysApi{*mApi}
	service.SerSysMenu.Insert(&Cmenu)

	curPath = fmt.Sprintf("%s/%s/%s/get", tab.ApiRoot, tab.PackageName, tab.ModuleName)
	where["path"] = curPath

	gApi := models.NewSysApi()
	service.SerSysApi.GetByWhere(where, gApi)
	if gApi.Id == 0 {
		gApi := gApi.SetMethod("POST").SetPermType("t").SetPath(curPath).
			SetStatus(3).SetTitle("根据id获取" + tab.TableComment)
		service.SerSysApi.Create(gApi)
	}

	MList := models.SysMenu{}
	MList.MenuName = ""
	MList.Title = tab.TableComment + "详情"
	MList.Icon = ""
	MList.Path = tab.TBName + "_detail"
	MList.MenuType = 3
	MList.Permission = tab.PackageName + ":" + tab.BusinessName + ":query"
	MList.ParentId = Cmenu.Id
	MList.NoCache = false
	MList.Sort = 0
	MList.Hidden = false
	MList.CreateBy = 1
	MList.UpdateBy = 1
	MList.SysApi = []models.SysApi{*gApi}
	service.SerSysMenu.Insert(&MList)

	curPath = fmt.Sprintf("%s/%s/%s/create", tab.ApiRoot, tab.PackageName, tab.ModuleName)
	where["path"] = curPath
	cApi := models.NewSysApi()
	service.SerSysApi.GetByWhere(where, cApi)
	if cApi.Id == 0 {
		cApi := cApi.SetMethod("POST").SetPermType("t").SetPath(curPath).
			SetStatus(3).SetTitle("创建" + tab.TableComment)
		service.SerSysApi.Create(cApi)
	}

	MCreate := models.SysMenu{}
	MCreate.MenuName = ""
	MCreate.Title = tab.TableComment + "创建"
	MCreate.Icon = ""
	MCreate.Path = tab.TBName + "_create"
	MCreate.MenuType = 3
	MCreate.Permission = tab.PackageName + ":" + tab.BusinessName + ":add"
	MCreate.ParentId = Cmenu.Id
	MCreate.NoCache = false
	MCreate.Sort = 0
	MCreate.Hidden = false
	MCreate.CreateBy = 1
	MCreate.UpdateBy = 1
	MCreate.SysApi = []models.SysApi{*cApi}
	service.SerSysMenu.Insert(&MCreate)

	curPath = fmt.Sprintf("%s/%s/%s/update", tab.ApiRoot, tab.PackageName, tab.ModuleName)
	where["path"] = curPath
	uApi := models.NewSysApi()
	service.SerSysApi.GetByWhere(where, uApi)
	if uApi.Id == 0 {
		uApi = uApi.SetMethod("POST").SetPermType("t").SetPath(curPath).
			SetStatus(3).SetTitle("修改" + tab.TableComment)
		service.SerSysApi.Create(uApi)
	}

	MUpdate := models.SysMenu{}
	MUpdate.MenuName = ""
	MUpdate.Title = tab.TableComment + "修改"
	MUpdate.Icon = ""
	MUpdate.Path = tab.TBName + "_update"
	MUpdate.MenuType = 3
	MUpdate.Permission = tab.PackageName + ":" + tab.BusinessName + ":edit"
	MUpdate.ParentId = Cmenu.Id
	MUpdate.NoCache = false
	MUpdate.Sort = 0
	MUpdate.Hidden = false
	MUpdate.CreateBy = 1
	MUpdate.UpdateBy = 1
	MUpdate.SysApi = []models.SysApi{*uApi}
	service.SerSysMenu.Insert(&MUpdate)

	curPath = fmt.Sprintf("%s/%s/%s/del", tab.ApiRoot, tab.PackageName, tab.ModuleName)
	where["path"] = curPath
	dApi := models.NewSysApi()
	service.SerSysApi.GetByWhere(where, dApi)
	if dApi.Id == 0 {
		dApi = dApi.SetMethod("POST").SetPermType("t").SetPath(curPath).
			SetStatus(3).SetTitle("删除" + tab.TableComment)
		service.SerSysApi.Create(dApi)
	}

	MDelete := models.SysMenu{}
	MDelete.MenuName = ""
	MDelete.Title = tab.TableComment + "删除"
	MDelete.Icon = ""
	MDelete.Path = tab.TBName + "_del"
	MDelete.MenuType = 3
	MDelete.Permission = tab.PackageName + ":" + tab.BusinessName + ":remove"
	MDelete.ParentId = Cmenu.Id
	MDelete.NoCache = false
	MDelete.Sort = 0
	MDelete.Hidden = false
	MDelete.CreateBy = 1
	MDelete.UpdateBy = 1
	MDelete.SysApi = []models.SysApi{*dApi}
	service.SerSysMenu.Insert(&MDelete)

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
