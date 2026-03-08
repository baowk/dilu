package service

import (
	"bytes"
	"dilu/internal/tools/models"
	"dilu/internal/tools/models/tools"
	"dilu/internal/tools/service/dto"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"dilu/common/config"
	"dilu/common/utils/files"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type GenTablesService struct {
	*base.BaseService
}

var SerGenTables = GenTablesService{
	base.NewService("sys"),
}

func (s *GenTablesService) Page(req *dto.GenTablesGetPageReq, list *[]models.GenTables, total *int64) error {
	db := s.DB().Order("table_id desc").Offset(req.GetOffset()).Limit(req.GetSize())
	if req.DbName != "" {
		db.Where("db_name = ?", req.DbName)
	}
	if req.TableName != "" {
		db.Where("table_name = ?", req.TableName)
	}
	return db.Find(list).Offset(-1).Limit(-1).Count(total).Error
}

func (s *GenTablesService) Del(req base.ReqIds) error {
	err := s.DB().Where("table_id in ?", req.Ids).Delete(&models.GenColumns{}).Error
	if err != nil {
		return err
	}
	return s.DB().Delete(&models.GenTables{}, req.Ids).Error
}

func (e *GenTablesService) Get(tx *gorm.DB, exclude bool, tableId int) (models.GenTables, error) {
	var doc models.GenTables
	var err error
	table := tx
	if tx == nil {
		table = e.DB()
	}

	if tableId != 0 {
		table = table.Where("table_id = ?", tableId)
	}

	if err := table.First(&doc).Error; err != nil {
		return doc, err
	}
	var col models.GenColumns
	col.TableId = doc.TableId
	if doc.Columns, err = SerGenColumns.GetList(tx, exclude, tableId); err != nil {
		return doc, err
	}

	return doc, nil
}

func (e *GenTablesService) Create(m *models.GenTables) error {
	m.CreateBy = 0
	err := e.DB().Create(m).Error
	if err != nil {
		return err
	}
	for _, v := range m.Columns {
		v.TableId = m.TableId
		SerGenColumns.Create(&v)
	}
	return nil
}

func (e *GenTablesService) GetDbs() []dto.DbOption {
	var dbs []dto.DbOption
	if config.Get().DBCfg.DSN != "" {
		db := dto.DbOption{
			Lable: consts.DB_DEF,
		}
		db.Value = ParseDsn(config.Get().DBCfg.DSN)
		dbs = append(dbs, db)
	}
	for key, dbc := range config.Get().DBCfg.DBS {
		if !dbc.Disable {
			db := dto.DbOption{
				Lable: key,
			}
			db.Value = ParseDsn(dbc.DSN)
			dbs = append(dbs, db)
		}
	}
	return dbs
}

var (
	frontPath = "../dilu-admin/src"
)

func (e *GenTablesService) GenTableInit(dbname string, tableName string, force bool) (models.GenTables, error) {
	var data models.GenTables
	var dbTable tools.DBTables
	var dbColumn tools.DBColumns
	data.CreateBy = 0

	dbTable.TableName = tableName
	dstdb, _, sdbn, driver := GetDb(dbname)
	dbTable.TableSchema = sdbn
	data.DbName = sdbn
	data.ConfDbName = dbname
	data.TBName = tableName

	var db *gorm.DB
	if force {
		db = dstdb
	} else {
		db = e.DB()
	}

	dbtable, err := dbTable.Get(db, sdbn, driver)
	if err != nil {
		return data, err
	}
	dbColumn.TableName = tableName
	tablenamelist := strings.Split(dbColumn.TableName, "_")
	for i := 0; i < len(tablenamelist); i++ {
		strStart := string([]byte(tablenamelist[i])[:1])
		strend := string([]byte(tablenamelist[i])[1:])
		// 大驼峰表名 结构体使用
		data.ClassName += strings.ToUpper(strStart) + strend
		// 小驼峰表名 js函数名和权限标识使用
		if i == 0 {
			data.BusinessName += strings.ToLower(strStart) + strend
		} else {
			data.BusinessName += strings.ToUpper(strStart) + strend
		}
	}
	if dbname == "master" {
		data.PackageName = consts.DB_DEF
	} else {
		data.PackageName = dbname
	}
	data.PackageName = resolveModuleName(data.PackageName)
	data.TplCategory = "crud"
	data.Crud = true
	// 中横线表名称，接口路径、前端文件夹名称和js名称使用
	data.ModuleName = strings.Replace(tableName, "_", "-", -1)
	dbcolumn, err := dbColumn.GetList(db, sdbn, driver)
	data.CreateBy = 0
	data.TableComment = dbtable.TableComment
	if dbtable.TableComment == "" {
		data.TableComment = data.ClassName
	}

	data.FunctionName = data.TableComment
	data.IsLogicalDelete = "1"
	data.LogicalDelete = true
	data.LogicalDeleteColumn = "is_del"
	data.IsActions = 2
	data.IsDataScope = 1
	data.IsAuth = 1

	data.FunctionAuthor = "baowk"
	for i := 0; i < len(dbcolumn); i++ {
		var column models.GenColumns
		column.ColumnComment = dbcolumn[i].ColumnComment
		column.ColumnName = dbcolumn[i].ColumnName
		column.ColumnType = dbcolumn[i].ColumnType
		column.ColumnDefault = dbcolumn[i].ColumnDefault

		if column.ColumnType == "" {
			column.ColumnType = dbcolumn[i].DataType
		}

		column.Sort = i + 1
		column.Insert = true
		column.IsInsert = "1"
		column.QueryType = "EQ"
		column.IsPk = "0"

		namelist := strings.Split(dbcolumn[i].ColumnName, "_")
		column.GoField = toGoField(namelist)
		column.JsonField = toJSONField(namelist)
		if strings.Contains(dbcolumn[i].ColumnKey, "PR") {
			column.IsPk = "1"
			column.Pk = true
			data.PkColumn = dbcolumn[i].ColumnName
			//column.GoField = strings.ToUpper(column.GoField)
			//column.JsonField = strings.ToUpper(column.JsonField)
			data.PkGoField = column.GoField
			data.PkJsonField = column.JsonField
			column.IsList = "1"
		}
		column.IsRequired = "0"
		if strings.Contains(dbcolumn[i].IsNullable, "NO") {
			column.IsRequired = "1"
			column.Required = true
		}

		if strings.Contains(dbcolumn[i].ColumnKey, "tinyint unsigned") {
			column.GoType = "uint8"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "tinyint") {
			column.GoType = "int8"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "float") {
			column.GoType = "float64"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "double") {
			column.GoType = "float64"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "bigint unsigned") {
			column.GoType = "uint64"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "bigint") {
			column.GoType = "int64"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "smallint unsigned") {
			column.GoType = "uint16"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "smallint") {
			column.GoType = "int16"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "int unsigned") {
			column.GoType = "uint32"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "int") {
			column.GoType = "int32"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "decimal") {
			column.GoType = "float64"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "numeric") {
			column.GoType = "float64"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "binary") {
			column.GoType = "[]byte"
			column.HtmlType = "input"
		} else if strings.Contains(column.ColumnType, "blob") {
			column.GoType = "[]byte"
			column.HtmlType = "input"
		} else if strings.Contains(column.ColumnType, "bool") {
			column.GoType = "bool"
			column.HtmlType = "input"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "timestamp") {
			column.GoType = "time.Time"
			column.HtmlType = "datetime"
			column.IsList = "1"
		} else if strings.Contains(column.ColumnType, "date") {
			column.GoType = "time.Time"
			column.HtmlType = "datetime"
			column.IsList = "1"
		} else {
			column.GoType = "string"
			column.HtmlType = "input"
			column.IsEdit = "1"
			column.IsList = "1"
		}

		if column.ColumnName == "update_by" || column.ColumnName == "create_by" {
			column.IsEdit = ""
		}

		if strings.Contains(column.ColumnName, "status") {
			column.IsQuery = "1"
		}

		data.Columns = append(data.Columns, column)
	}
	return data, err
}

const ROOT = "./internal/"

func (e *GenTablesService) NOMethodsGen(tab models.GenTables, force bool) error {

	tab.MLTBName = strings.Replace(tab.TBName, "_", "-", -1)
	genFront := config.Get().Gen.GenFront

	if genFront && config.Get().Gen.FrontPath != "" {
		frontPath = config.Get().Gen.FrontPath
	}

	_ = files.PathCreate(ROOT + tab.PackageName + "/apis/")
	_ = files.PathCreate(ROOT + tab.PackageName + "/repository/")
	_ = files.PathCreate(ROOT + tab.PackageName + "/router/")
	_ = files.PathCreate(ROOT + tab.PackageName + "/service/dto/")

	if err := e.generateRepository(tab); err != nil {
		core.GetApp().GetLogger().Error("Gen repository", "err", err)
		return err
	}

	if genFront {
		_ = files.PathCreate(frontPath + "/api/" + tab.PackageName + "/")
		err := files.PathCreate(frontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/utils")
		if err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
			return err
		}
	}

	m := map[string]string{}
	m["modName"] = tab.PackageName

	//路由
	cmdApi := "cmd/start/" + tab.PackageName + ".go"
	if files.CheckExist(cmdApi) || force {
		rt1, err := parseTemplate("go/router/cmd_api.template")
		if err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
			return err

		}
		var rb1 bytes.Buffer
		if err = rt1.Execute(&rb1, m); err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
			return err
		}
		files.FileCreate(rb1, cmdApi)
	}

	baseRouter := ROOT + tab.PackageName + "/router/router.go"
	if files.CheckExist(baseRouter) || force {
		rt2, err := parseTemplate("go/router/router.template")
		if err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
		}
		var rb2 bytes.Buffer
		err = rt2.Execute(&rb2, m)
		if err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
			return err
		}
		files.FileCreate(rb2, baseRouter)
	}

	//golang

	apigo := ROOT + tab.PackageName + "/apis/" + tab.TBName + ".go"
	if files.CheckExist(apigo) || force {
		t2, err := parseTemplate("go/service/apis.go.template")
		if err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
			return err
		}
		var b2 bytes.Buffer
		err = t2.Execute(&b2, tab)
		if err != nil {
			core.GetApp().GetLogger().Error("gen err", "err", err)
			return err
		}
		files.FileCreate(b2, apigo)
	}

	routergo := ROOT + tab.PackageName + "/router/" + tab.TBName + ".go"
	if files.CheckExist(routergo) || force {
		t3, err := parseTemplate("go/service/router_no_check_role.go.template")
		if err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
			return err
		}
		var b3 bytes.Buffer
		err = t3.Execute(&b3, tab)
		if err != nil {
			core.GetApp().GetLogger().Error("gen err", "err", err)
			return err
		}
		files.FileCreate(b3, routergo)
	}

	dto := ROOT + tab.PackageName + "/service/dto/" + tab.TBName + ".go"
	if files.CheckExist(dto) || force {
		t6, err := parseTemplate("go/service/dto.go.template")
		if err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
			return err
		}
		var b6 bytes.Buffer
		err = t6.Execute(&b6, tab)
		if err != nil {
			core.GetApp().GetLogger().Error("gen err", "err", err)
			return err
		}
		files.FileCreate(b6, dto)
	}

	service := ROOT + tab.PackageName + "/service/" + tab.TBName + ".go"
	if files.CheckExist(service) || force {
		t7, err := parseTemplate("go/service/service.go.template")
		if err != nil {
			core.GetApp().GetLogger().Error("Gen", "err", err)
			return err
		}
		var b7 bytes.Buffer
		err = t7.Execute(&b7, tab)
		if err != nil {
			core.GetApp().GetLogger().Error("gen err", "err", err)
			return err
		}
		files.FileCreate(b7, service)
	}

	if genFront {
		//前端部分
		js := frontPath + "/api/" + tab.PackageName + "/" + tab.MLTBName + ".ts"
		if files.CheckExist(js) || force {
			t4, err := parseTemplate("vue/api/api.ts.template")
			if err != nil {
				core.GetApp().GetLogger().Error("Gen", "err", err)
				return err
			}
			var b4 bytes.Buffer
			err = t4.Execute(&b4, tab)
			if err != nil {
				core.GetApp().GetLogger().Error("gen err", "err", err)
				return err
			}
			files.FileCreate(b4, js)
		}

		// types := FrontPath + "/api/" + tab.PackageName + "/" + tab.MLTBName + ".d.ts"
		// if files.CheckExist(types) || force {
		// 	t5, err := template.ParseFiles(basePath + "vue/api/types.ts.template")
		// 	if err != nil {
		// 		core.GetApp().GetLogger().Error("Gen", err)
		// 		e.Error(c, err)
		// 		return
		// 	}
		// 	var b5 bytes.Buffer
		// 	err = t5.Execute(&b5, tab)
		// 	if err != nil {
		// 		core.GetApp().GetLogger().Error("gen err", err)
		// 	}
		// 	files.FileCreate(b5, types)
		// }

		vue := frontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/index.vue"
		if files.CheckExist(vue) || force {
			t5, err := parseTemplate("vue/views/index.vue.template")
			if err != nil {
				core.GetApp().GetLogger().Error("Gen", "err", err)
				return err
			}
			var b5 bytes.Buffer
			err = t5.Execute(&b5, tab)
			if err != nil {
				core.GetApp().GetLogger().Error("gen err", "err", err)
				return err
			}
			files.FileCreate(b5, vue)
		}

		form := frontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/form.vue"
		if files.CheckExist(form) || force {
			t5, err := parseTemplate("vue/views/form.vue.template")
			if err != nil {
				core.GetApp().GetLogger().Error("Gen", "err", err)
				return err
			}
			var b5 bytes.Buffer
			err = t5.Execute(&b5, tab)
			if err != nil {
				core.GetApp().GetLogger().Error("gen err", "err", err)
				return err
			}
			files.FileCreate(b5, form)
		}

		hook := frontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/utils/hook.tsx"
		if files.CheckExist(hook) || force {
			t5, err := parseTemplate("vue/views/utils/hook.tsx.template")
			if err != nil {
				core.GetApp().GetLogger().Error("Gen", "err", err)
				return err
			}
			var b5 bytes.Buffer
			err = t5.Execute(&b5, tab)
			if err != nil {
				core.GetApp().GetLogger().Error("gen err", "err", err)
				return err
			}
			files.FileCreate(b5, hook)
		}
	}

	// rule := FrontPath + "/views/" + tab.PackageName + "/" + tab.MLTBName + "/utils/rule.ts"
	// if files.CheckExist(rule) || force {
	// 	t5, err := template.ParseFiles(basePath + "vue/views/utils/rule.ts.template")
	// 	if err != nil {
	// 		core.GetApp().GetLogger().Error("Gen", err)
	// 		e.Error(c, err)
	// 		return
	// 	}
	// 	var b5 bytes.Buffer
	// 	err = t5.Execute(&b5, tab)
	// 	if err != nil {
	// 		core.GetApp().GetLogger().Error("gen err", err)
	// 	}
	// 	files.FileCreate(b5, rule)
	// }
	return nil

}

func (e *GenTablesService) Update(tab *models.GenTables) (err error) {

	//参数1:是要修改的数据
	//参数2:是修改的数据
	tab.UpdateBy = 0
	if err = e.DB().Where("table_id = ?", tab.TableId).Updates(tab).Error; err != nil {
		return
	}

	tableNames := make([]string, 0)
	for _, v := range tab.Columns {
		if v.FkTableName != "" {
			tableNames = append(tableNames, v.FkTableName)
		}
	}

	tables := make([]models.GenTables, 0)
	tableMap := make(map[string]*models.GenTables)
	if len(tableNames) > 0 {
		if err = e.DB().Where("table_name in (?)", tableNames).Find(&tables).Error; err != nil {
			return
		}
		for i := range tables {
			tableMap[tables[i].TBName] = &tables[i]
		}
	}

	for _, v := range tab.Columns {
		if v.FkTableName != "" {
			t, ok := tableMap[v.FkTableName]
			if ok {
				v.FkTableNameClass = t.ClassName
				t.MLTBName = strings.Replace(t.TBName, "_", "-", -1)
				v.FkTableNamePackage = t.MLTBName
			} else {
				tableNameList := strings.Split(v.FkTableName, "_")
				v.FkTableNameClass = ""
				//v.FkTableNamePackage = ""
				for a := 0; a < len(tableNameList); a++ {
					strStart := string([]byte(tableNameList[a])[:1])
					strEnd := string([]byte(tableNameList[a])[1:])
					v.FkTableNameClass += strings.ToUpper(strStart) + strEnd
					//v.FkTableNamePackage += strings.ToLower(strStart) + strings.ToLower(strEnd)
				}
			}
		}
		SerGenColumns.UpdateById(v)
	}
	return
}

func resolveModuleName(defaultName string) string {
	m := config.Get().Gen.ModuleMap
	if len(m) == 0 {
		return defaultName
	}
	if mapped, ok := m[defaultName]; ok && strings.TrimSpace(mapped) != "" {
		return strings.TrimSpace(mapped)
	}
	return defaultName
}

func parseTemplate(relativePath string) (*template.Template, error) {
	for _, base := range templateBasePaths() {
		full := filepath.Join(base, filepath.FromSlash(relativePath))
		if _, err := os.Stat(full); err == nil {
			return template.ParseFiles(full)
		}
	}
	return nil, fmt.Errorf("template not found: %s", relativePath)
}

func templateBasePaths() []string {
	roots := make([]string, 0, 3)
	if custom := strings.TrimSpace(config.Get().Gen.TemplatePath); custom != "" {
		roots = append(roots, custom)
	}
	roots = append(roots, "templates", "resources/template")
	return roots
}

func (e *GenTablesService) generateRepository(tab models.GenTables) error {
	db := core.GetApp().Db(tab.ConfDbName)
	if db == nil {
		return fmt.Errorf("database not found: %s", tab.ConfDbName)
	}

	outPath := filepath.Join(ROOT, tab.PackageName, "repository", "query")
	g := gen.NewGenerator(gen.Config{
		OutPath:           outPath,
		Mode:              gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     false,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateModelAs(tab.TBName, tab.ClassName))
	g.Execute()
	return nil
}

func ParseDsn(dsn string) string {
	if len(dsn) < 3 {
		return ""
	}
	if strings.HasSuffix(strings.ToLower(dsn), ".db") || strings.HasSuffix(strings.ToLower(dsn), ".sqlite") || strings.HasSuffix(strings.ToLower(dsn), ".sqlite3") {
		base := filepath.Base(dsn)
		ext := filepath.Ext(base)
		return strings.TrimSuffix(base, ext)
	}
	idx := strings.LastIndex(dsn, ")/")
	end := strings.LastIndex(dsn, "?")
	if end < 0 {
		end = len(dsn)
	}
	if idx < 0 || idx+2 > end {
		return dsn
	}
	return dsn[idx+2 : end]
}

func ParsePgsqlDsn(dsn string) string {
	if strings.HasPrefix(dsn, "postgres://") || strings.HasPrefix(dsn, "postgresql://") {
		if len(dsn) < 3 {
			return ""
		}
		idx := strings.LastIndex(dsn, "/")
		end := strings.LastIndex(dsn, "?")
		if end < 0 {
			end = len(dsn)
		}
		return dsn[idx+1 : end]
	}

	re := regexp.MustCompile(`dbname=([^\s]*)`)

	match := re.FindStringSubmatch(dsn)
	if len(match) <= 1 {
		return ""
	}
	return match[1]

}

func GetDb(dbname string) (db *gorm.DB, mdb string, sdb, driver string) {
	mdsn := config.Get().DBCfg.DSN
	mdb = ParseDsn(mdsn)
	if dbname != consts.DB_DEF {
		gdsn, ok := config.Get().DBCfg.DBS[dbname]
		if !ok {
			return
		}
		core.GetApp().GetLogger().Debug("driver", "test", gdsn.Driver)
		if gdsn.Driver == "pgsql" || gdsn.Driver == "postgres" {
			sdb = ParsePgsqlDsn(gdsn.DSN)
		} else if gdsn.Driver == "sqlite" {
			sdb = ParseDsn(gdsn.DSN)
		} else {
			sdb = ParseDsn(gdsn.DSN)
		}

		if gdsn.Driver != "" {
			driver = gdsn.Driver
		} else {
			driver = config.Get().DBCfg.Driver
		}
	} else {
		driver = config.Get().DBCfg.Driver
		if driver == "sqlite" {
			sdb = ParseDsn(config.Get().DBCfg.DSN)
		} else {
			sdb = mdb
		}
	}
	if driver == "postgres" {
		driver = "pgsql"
	}
	db = core.GetApp().Db(dbname)
	return
}

var goInitialisms = map[string]string{
	"api":   "API",
	"id":    "ID",
	"ip":    "IP",
	"url":   "URL",
	"uri":   "URI",
	"uuid":  "UUID",
	"http":  "HTTP",
	"https": "HTTPS",
	"json":  "JSON",
	"sql":   "SQL",
}

func toGoField(parts []string) string {
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		lower := strings.ToLower(p)
		if v, ok := goInitialisms[lower]; ok {
			b.WriteString(v)
			continue
		}
		b.WriteString(strings.ToUpper(lower[:1]))
		if len(lower) > 1 {
			b.WriteString(lower[1:])
		}
	}
	return b.String()
}

func toJSONField(parts []string) string {
	var b strings.Builder
	for i, p := range parts {
		if p == "" {
			continue
		}
		lower := strings.ToLower(p)
		if i == 0 {
			b.WriteString(lower)
			continue
		}
		b.WriteString(strings.ToUpper(lower[:1]))
		if len(lower) > 1 {
			b.WriteString(lower[1:])
		}
	}
	return b.String()
}
