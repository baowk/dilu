package apis

import (
	cons "dilu/common/consts"
	"dilu/modules/tools/models"
	"dilu/modules/tools/models/tools"
	"dilu/modules/tools/service"
	"dilu/modules/tools/service/dto"

	sModels "dilu/modules/sys/models"
	sService "dilu/modules/sys/service"
	"errors"
	"fmt"
	"strings"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GenTablesApi struct {
	base.BaseApi
}

var ApiGenTables = GenTablesApi{}

// GetDBTableList 分页列表数据
// @Summary 分页列表数据 / page list data
// @Description 数据库表分页列表 / database table page list
// @Tags 工具 / 生成工具
// @Param data body dto.DBReq true "body"
// @Success 200 {object} base.Resp "{"code": 200, "data": [...]}"
// @Router /api/v1/tools/gen/db/tables [post]
func (e *GenTablesApi) GetDBTableList(c *gin.Context) {
	//var res base.Resp
	var data tools.DBTables
	var err error
	var req dto.DBReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	if core.Cfg.DBCfg.Driver == "sqlite3" || core.Cfg.DBCfg.Driver == "postgres" {
		err = errors.New("对不起，sqlite3 或 postgres 不支持代码生成！")
		e.Error(c, err)
		return
	}
	if req.TableName != "" {
		data.TableName = req.TableName
	}

	var dbname = req.DBName
	db, mdbn, sdbn := service.GetDb(dbname)

	result, total, err := data.GetPage(db, req.GetSize(), req.GetPage(), sdbn, mdbn)
	if err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, result, total, req.GetPage(), req.GetSize())
}

// QueryPage 获取GenTables列表
// @Summary 获取GenTables列表
// @Tags 工具 / 生成工具
// @Accept application/json
// @Product application/json
// @Param data body dto.GenTablesGetPageReq true "body"
// @Success 200 {object} base.Resp{data=base.PageResp{list=[]models.GenTables}} "{"code": 200, "data": [...]}"
// @Router /api/v1/tools/gen/page [post]
func (e *GenTablesApi) QueryPage(c *gin.Context) {
	var req dto.GenTablesGetPageReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	list := make([]models.GenTables, 10)
	var total int64
	if err := service.SerGenTables.Page(&req, &list, &total); err != nil {
		e.Error(c, err)
		return
	}
	e.Page(c, list, total, req.GetPage(), req.GetSize())
}

// Del 删除GenTables
// @Summary 删除GenTables
// @Tags 工具 / 生成工具
// @Accept application/json
// @Product application/json
// @Param data body base.ReqIds true "body"
// @Success 200 {object} base.Resp{data=models.GenTables} "{"code": 200, "data": [...]}"
// @Router /api/v1/tools/gen/del [post]
func (e *GenTablesApi) Del(c *gin.Context) {
	if !core.Cfg.Gen.Enable {
		e.Error(c, errors.New("当前生成表已关闭"))
		return
	}
	var req base.ReqIds
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	if err := service.SerGenTables.Del(req); err != nil {
		e.Error(c, err)
		return
	}
	e.Ok(c)
}

// GetDBS
// @Summary 获取配置的数据库
// @Description 获取配置的数据库
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Success 200 {object} base.Resp{data=[]dto.DbOption} "{"code": 200, "data": [...]}"
// @Router /api/v1/tools/gen/dbs [post]
func (e *GenTablesApi) GetDBS(c *gin.Context) {
	e.Ok(c, service.SerGenTables.GetDbs())
}

// Insert
// @Summary 添加表结构
// @Description 添加表结构
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param data body dto.ImpTablesReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/tools/gen/add [post]
func (e *GenTablesApi) Insert(c *gin.Context) {
	if !core.Cfg.Gen.Enable {
		e.Error(c, errors.New("添加表结构已关闭"))
		return
	}
	var req dto.ImpTablesReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	for _, tableName := range req.Tables {
		fmt.Println(tableName)
		data, err := service.SerGenTables.GenTableInit(req.DbName, tableName, false)
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
		err = service.SerGenTables.Create(&data)
		if err != nil {
			core.Log.Error("Gen", zap.Error(err))
			e.Error(c, err)
			return
		}
	}
	e.Ok(c)

}

// Update
// @Summary 修改表结构
// @Description 修改表结构
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param data body models.GenTables true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/tools/gen/update [POST]
func (e *GenTablesApi) Update(c *gin.Context) {
	if !core.Cfg.Gen.Enable {
		e.Error(c, errors.New("修改表结构已关闭"))
		return
	}
	var data models.GenTables
	if err := c.ShouldBind(&data); err != nil {
		e.Error(c, err)
		return
	}
	data.UpdateBy = 0
	err := service.SerGenTables.Update(&data)
	if err != nil {
		core.Log.Error("Gen", zap.Error(err))
		e.Error(c, err)
		return
	}
	e.Ok(c, data)
}

// GenMenuAndApi
// @Summary 生成菜单
// @Description 生成菜单
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param data body dto.GenMenuReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/tools/gen/menu [post]
func (e *GenTablesApi) GenMenuAndApi(c *gin.Context) {
	if !core.Cfg.Gen.Enable {
		e.Error(c, errors.New("api和菜单已关闭"))
		return
	}
	table := models.GenTables{}
	var req dto.GenMenuReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}
	menuPid := req.MenuId
	table.TableId = req.TableId
	db, _, _ := service.GetDb(consts.DB_DEF)
	tab, _ := service.SerGenTables.Get(db, true, req.TableId)
	tab.MLTBName = strings.Replace(tab.TBName, "_", "-", -1)
	tab.ApiRoot = cons.ApiRoot
	if menuPid == 0 {
		Mmenu := sModels.SysMenu{}
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
		Mmenu.PlatformType = 2
		sService.SerSysMenu.Insert(&Mmenu)
		menuPid = Mmenu.Id
	}

	curPath := fmt.Sprintf("%s/%s/%s/page", tab.ApiRoot, tab.PackageName, tab.ModuleName)
	where := map[string]any{
		"path":   curPath,
		"method": "POST",
	}
	mApi := sModels.NewSysApi()
	sService.SerSysApi.GetByWhere(where, mApi)
	if mApi.Id == 0 {
		mApi = mApi.SetMethod("POST").SetPermType(3).SetPath(curPath).
			SetStatus(3).SetTitle("分页获取" + tab.TableComment)
		sService.SerSysApi.Create(mApi)
	}
	Cmenu := sModels.SysMenu{}
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
	Cmenu.PlatformType = 2
	Cmenu.SysApi = []sModels.SysApi{*mApi}
	sService.SerSysMenu.Insert(&Cmenu)

	curPath = fmt.Sprintf("%s/%s/%s/get", tab.ApiRoot, tab.PackageName, tab.ModuleName)
	where["path"] = curPath

	gApi := sModels.NewSysApi()
	sService.SerSysApi.GetByWhere(where, gApi)
	if gApi.Id == 0 {
		gApi := gApi.SetMethod("POST").SetPermType(3).SetPath(curPath).
			SetStatus(3).SetTitle("根据id获取" + tab.TableComment)
		sService.SerSysApi.Create(gApi)
	}

	MList := sModels.SysMenu{}
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
	MList.PlatformType = 2
	MList.SysApi = []sModels.SysApi{*gApi}
	sService.SerSysMenu.Insert(&MList)

	curPath = fmt.Sprintf("%s/%s/%s/create", tab.ApiRoot, tab.PackageName, tab.ModuleName)
	where["path"] = curPath
	cApi := sModels.NewSysApi()
	sService.SerSysApi.GetByWhere(where, cApi)
	if cApi.Id == 0 {
		cApi := cApi.SetMethod("POST").SetPermType(3).SetPath(curPath).
			SetStatus(3).SetTitle("创建" + tab.TableComment)
		sService.SerSysApi.Create(cApi)
	}

	MCreate := sModels.SysMenu{}
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
	MCreate.PlatformType = 2
	MCreate.SysApi = []sModels.SysApi{*cApi}
	sService.SerSysMenu.Insert(&MCreate)

	curPath = fmt.Sprintf("%s/%s/%s/update", tab.ApiRoot, tab.PackageName, tab.ModuleName)
	where["path"] = curPath
	uApi := sModels.NewSysApi()
	sService.SerSysApi.GetByWhere(where, uApi)
	if uApi.Id == 0 {
		uApi = uApi.SetMethod("POST").SetPermType(3).SetPath(curPath).
			SetStatus(3).SetTitle("修改" + tab.TableComment)
		sService.SerSysApi.Create(uApi)
	}

	MUpdate := sModels.SysMenu{}
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
	MUpdate.PlatformType = 2
	MUpdate.SysApi = []sModels.SysApi{*uApi}
	sService.SerSysMenu.Insert(&MUpdate)

	curPath = fmt.Sprintf("%s/%s/%s/del", tab.ApiRoot, tab.PackageName, tab.ModuleName)
	where["path"] = curPath
	dApi := sModels.NewSysApi()
	sService.SerSysApi.GetByWhere(where, dApi)
	if dApi.Id == 0 {
		dApi = dApi.SetMethod("POST").SetPermType(3).SetPath(curPath).
			SetStatus(3).SetTitle("删除" + tab.TableComment)
		sService.SerSysApi.Create(dApi)
	}

	MDelete := sModels.SysMenu{}
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
	MDelete.PlatformType = 2
	MDelete.SysApi = []sModels.SysApi{*dApi}
	sService.SerSysMenu.Insert(&MDelete)

	e.Ok(c, "数据生成成功！")
}

// // Preview
// // @Summary 生成预览
// // @Description 生成预览
// // @Tags 工具 / 生成工具
// // @Accept  application/json
// // @Product application/json
// // @Param tableId path int true "tableId"
// // @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// // @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// // @Router /api/tools/gen/preview/{tableId} [get]
// func (e *GenTablesApi) Preview(c *gin.Context) {
// 	table := models.GenTables{}
// 	id, err := strconv.Atoi(c.Param("tableId"))
// 	if err != nil {
// 		core.Log.Error("Gen", zap.Error(err))
// 		e.Error(c, err)
// 		return
// 	}
// 	table.TableId = id
// 	t1, err := template.ParseFiles("resources/template/go/service/model.go.template")
// 	if err != nil {
// 		core.Log.Error("Gen", zap.Error(err))
// 		e.Error(c, err)
// 		return
// 	}
// 	t2, err := template.ParseFiles("resources/template/go/service/apis.go.template")
// 	if err != nil {
// 		core.Log.Error("Gen", zap.Error(err))
// 		e.Error(c, err)
// 		return
// 	}
// 	t3, err := template.ParseFiles("resources/template/vue/api/api.ts.template")
// 	if err != nil {
// 		core.Log.Error("Gen", zap.Error(err))
// 		e.Error(c, err)
// 		return
// 	}
// 	t4, err := template.ParseFiles("resources/template/vue/views/index.vue.template")
// 	if err != nil {
// 		core.Log.Error("Gen", zap.Error(err))
// 		e.Error(c, err)
// 		return
// 	}
// 	t5, err := template.ParseFiles("resources/template/go/service/router_no_check_role.go.template")
// 	if err != nil {
// 		core.Log.Error("Gen", zap.Error(err))
// 		e.Error(c, err)
// 		return
// 	}
// 	t6, err := template.ParseFiles("resources/template/go/service/dto.go.template")
// 	if err != nil {
// 		core.Log.Error("Gen", zap.Error(err))
// 		e.Error(c, err)
// 		return
// 	}
// 	t7, err := template.ParseFiles("resources/template/go/service/service.go.template")
// 	if err != nil {
// 		core.Log.Error("Gen", zap.Error(err))
// 		e.Error(c, err)
// 		return
// 	}

// 	db, _, _ := GetDb(consts.DB_DEF)

// 	tab, _ := table.Get(db, false)
// 	var b1 bytes.Buffer
// 	err = t1.Execute(&b1, tab)
// 	if err != nil {
// 		core.Log.Error("gen err", zap.Error(err))
// 	}
// 	var b2 bytes.Buffer
// 	err = t2.Execute(&b2, tab)
// 	if err != nil {
// 		core.Log.Error("gen err", zap.Error(err))
// 	}
// 	var b3 bytes.Buffer
// 	err = t3.Execute(&b3, tab)
// 	if err != nil {
// 		core.Log.Error("gen err", zap.Error(err))
// 	}
// 	var b4 bytes.Buffer
// 	err = t4.Execute(&b4, tab)
// 	if err != nil {
// 		core.Log.Error("gen err", zap.Error(err))
// 	}
// 	var b5 bytes.Buffer
// 	err = t5.Execute(&b5, tab)
// 	if err != nil {
// 		core.Log.Error("gen err", zap.Error(err))
// 	}
// 	var b6 bytes.Buffer
// 	err = t6.Execute(&b6, tab)
// 	if err != nil {
// 		core.Log.Error("gen err", zap.Error(err))
// 	}
// 	var b7 bytes.Buffer
// 	err = t7.Execute(&b7, tab)
// 	if err != nil {
// 		core.Log.Error("gen err", zap.Error(err))
// 	}

// 	mp := make(map[string]interface{})
// 	mp["model.go.template"] = b1.String()
// 	mp["api.go.template"] = b2.String()
// 	mp["router.go.template"] = b5.String()
// 	mp["dto.go.template"] = b6.String()
// 	mp["service.go.template"] = b7.String()
// 	mp["js.go.template"] = b3.String()
// 	mp["vue.go.template"] = b4.String()
// 	e.Ok(c, mp)
// }

// GenCode
// @Summary 生成代码
// @Description 生成代码
// @Tags 工具 / 生成工具
// @Accept  application/json
// @Product application/json
// @Param data body dto.GenCodeReq true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/tools/gen/code [post]
func (e *GenTablesApi) GenCode(c *gin.Context) {
	if !core.Cfg.Gen.Enable {
		e.Error(c, errors.New("生成代码已关闭"))
		return
	}
	table := models.GenTables{}

	var req dto.GenCodeReq
	if err := c.ShouldBind(&req); err != nil {
		e.Error(c, err)
		return
	}

	table.TableId = req.TableId

	tab, _ := service.SerGenTables.Get(nil, false, req.TableId)
	tab.ApiRoot = cons.ApiRoot

	for i, v := range tab.Columns {
		tab.Columns[i].TsType = TypeGo2Ts(v.GoType)
	}
	service.SerGenTables.NOMethodsGen(tab, req.Force)
	e.Ok(c, "Code generated successfully！")
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

// // Get 获取GenTables
// // @Summary 获取GenTables
// // @Tags sys-GenTables
// // @Accept application/json
// // @Product application/json
// // @Param teamId header int false "团队id"
// // @Param data body base.ReqId true "body"
// // @Success 200 {object} base.Resp{data=models.GenTables} "{"code": 200, "data": [...]}"
// // @Router /api/v1/sys/gen-tables/get [post]
// // @Security Bearer
// func (e *GenTablesApi) Get(c *gin.Context) {
// 	var req base.ReqId
// 	if err := c.ShouldBind(&req); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	var data models.GenTables
// 	if err := service.SerGenTables.Get(req.Id, &data); err != nil {
// 		e.Error(c, err)
// 		return
// 	}
// 	e.Ok(c, data)
// }
