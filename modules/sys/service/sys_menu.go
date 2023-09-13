package service

import (
	"dilu/common/codes"
	"strconv"
	"strings"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"
)

type SysMenu struct {
	*base.BaseService
}

var SerSysMenu = SysMenu{
	base.NewService("sys"),
}

// GetPage 获取SysMenu列表
func (e *SysMenu) GetPage(c *dto.SysMenuGetPageReq, menus *[]models.SysMenu) (*SysMenu, errs.IError) {
	var menu = make([]models.SysMenu, 0)
	db := core.DB()
	//TODO WHERE
	if err := db.Limit(c.GetSize()).Offset(c.GetOffset()).Find(menu).Error; err != nil {
		berr := errs.Err(codes.FAILURE, "", err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return e, berr
	}
	for i := 0; i < len(menu); i++ {
		if menu[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&menu, menu[i])
		*menus = append(*menus, menusInfo)
	}
	return e, nil
}

// Get 获取SysMenu对象
func (e *SysMenu) Get(d *dto.SysMenuGetReq, model *models.SysMenu) (*SysMenu, errs.IError) {
	var err error
	var data models.SysMenu

	db := core.DB().Model(&data).Preload("SysApi").
		First(model, d.GetId())
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		core.Log.Error("sys_menu", zap.Error(err))
		berr := errs.Err(codes.FAILURE, "", err)
		return e, berr
	}
	if err != nil {
		core.Log.Error("sys_menu", zap.Error(err))
		berr := errs.Err(codes.FAILURE, "", err)
		return e, berr
	}
	// apis := make([]int, 0)
	// for _, v := range model.SysApi {
	// 	apis = append(apis, v.MenuId)
	// }
	//model.SysApi = apis
	return e, nil
}

// Insert 创建SysMenu对象
func (e *SysMenu) Insert(c *dto.SysMenuInsertReq) (*SysMenu, errs.IError) {
	var err error
	var data models.SysMenu
	c.Generate(&data)
	tx := core.DB().Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	// err = tx.Where("id in ?", c.Apis).Find(&data.SysApi).Error
	// if err != nil {
	// 	tx.Rollback()
	// 	core.Log.Error("sys_menu", zap.Error(err))
	// 	berr := errs.Err(codes.FAILURE, "", err)
	// 	return e, berr
	// }
	err = tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		core.Log.Error("sys_menu", zap.Error(err))
		berr := errs.Err(codes.FAILURE, "", err)
		return e, berr
	}
	c.Id = data.MenuId
	err = e.initPaths(tx, &data)
	if err != nil {
		tx.Rollback()
		core.Log.Error("sys_menu", zap.Error(err))
		berr := errs.Err(codes.FAILURE, "", err)
		return e, berr
	}
	tx.Commit()
	return e, nil
}

func (e *SysMenu) initPaths(tx *gorm.DB, menu *models.SysMenu) error {
	var err error
	var data models.SysMenu
	parentMenu := new(models.SysMenu)
	if menu.ParentId != 0 {
		err = tx.Model(&data).First(parentMenu, menu.ParentId).Error
		if err != nil {
			return err
		}
		if parentMenu.Paths == "" {
			err = errors.New("父级paths异常，请尝试对当前节点父级菜单进行更新操作！")
			return err
		}
		menu.Paths = parentMenu.Paths + "/" + strconv.Itoa(menu.MenuId)
	} else {
		menu.Paths = "/0/" + strconv.Itoa(menu.MenuId)
	}
	err = tx.Model(&data).Where("id = ?", menu.MenuId).Update("paths", menu.Paths).Error
	return err
}

// Update 修改SysMenu对象
func (e *SysMenu) Update(c *dto.SysMenuUpdateReq) (*SysMenu, errs.IError) {
	var err error
	tx := core.DB().Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	var alist = make([]models.SysApi, 0)
	var model = models.SysMenu{}
	tx.Preload("SysApi").First(&model, c.GetId())
	oldPath := model.Paths
	tx.Where("id in ?", c.Apis).Find(&alist)
	// err = tx.Model(&model).Association("SysApi").Delete(model.SysApi)
	// if err != nil {
	// 	core.Log.Error("sys_menu", zap.Error(err))
	// 	berr := errs.Err(codes.FAILURE, "", err)
	// 	return e, berr
	// }
	c.Generate(&model)
	// model.SysApi = alist
	// db := tx.Model(&model).Session(&gorm.Session{FullSaveAssociations: true}).Debug().Save(&model)
	// if err = db.Error; err != nil {
	// 	core.Log.Error("sys_menu", zap.Error(err))
	// 	berr := errs.Err(codes.FAILURE, "", err)
	// 	return e, berr
	// }
	// if db.RowsAffected == 0 {
	// 	berr := errs.Err(codes.FAILURE, "", err)
	// 	return e, berr
	// }
	var menuList []models.SysMenu
	tx.Where("paths like ?", oldPath+"%").Find(&menuList)
	for _, v := range menuList {
		v.Paths = strings.Replace(v.Paths, oldPath, model.Paths, 1)
		tx.Model(&v).Update("paths", v.Paths)
	}
	return e, nil
}

// Remove 删除SysMenu
func (e *SysMenu) Remove(d *dto.SysMenuDeleteReq) (*SysMenu, errs.IError) {
	var err error
	var data models.SysMenu

	db := core.DB().Model(&data).Delete(&data, d.Ids)
	if err = db.Error; err != nil {
		err = db.Error
		core.Log.Error("sys_menu", zap.Error(err))
		berr := errs.Err(codes.FAILURE, "", err)
		return e, berr
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		berr := errs.Err(codes.FAILURE, "", err)
		return e, berr
	}
	return e, nil
}

// GetList 获取菜单数据
func (e *SysMenu) GetList(c *dto.SysMenuGetPageReq, list *[]models.SysMenu) error {
	var err error
	var data models.SysMenu

	err = core.DB().Model(&data).
		Find(list).Error
	if err != nil {
		core.Log.Error("sys_menu", zap.Error(err))
		return err
	}
	return nil
}

// SetLabel 修改角色中 设置菜单基础数据
func (e *SysMenu) SetLabel() (m []dto.MenuLabel, err error) {
	var list []models.SysMenu
	err = e.GetList(&dto.SysMenuGetPageReq{}, &list)

	m = make([]dto.MenuLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.MenuLabel{}
		e.Id = list[i].MenuId
		e.Label = list[i].Title
		deptsInfo := menuLabelCall(&list, e)

		m = append(m, deptsInfo)
	}
	return
}

// GetSysMenuByRoleName 左侧菜单
func (e *SysMenu) GetSysMenuByRoleName(roleName ...string) ([]models.SysMenu, error) {
	var MenuList []models.SysMenu
	// var role models.SysRole
	// var err error
	// admin := false
	// for _, s := range roleName {
	// 	if s == "admin" {
	// 		admin = true
	// 	}
	// }

	// if len(roleName) > 0 && admin {
	// 	var data []models.SysMenu
	// 	err = core.DB().Where(" menu_type in ('M','C')").
	// 		Order("sort").
	// 		Find(&data).
	// 		Error
	// 	MenuList = data
	// } else {
	// 	err = core.DB().Model(&role).Preload("SysMenu", func(db *gorm.DB) *gorm.DB {
	// 		return db.Where(" menu_type in ('M','C')").Order("sort")
	// 	}).Where("role_name in ?", roleName).Find(&role).
	// 		Error
	// 	MenuList = *role.SysMenu
	// }

	// if err != nil {
	// 	core.Log.Error("sys_menu", zap.Error(err))
	// }
	return MenuList, nil
}

// menuLabelCall 递归构造组织数据
func menuLabelCall(eList *[]models.SysMenu, dept dto.MenuLabel) dto.MenuLabel {
	list := *eList

	min := make([]dto.MenuLabel, 0)
	for j := 0; j < len(list); j++ {

		if dept.Id != list[j].ParentId {
			continue
		}
		mi := dto.MenuLabel{}
		mi.Id = list[j].MenuId
		mi.Label = list[j].Title
		mi.Children = []dto.MenuLabel{}
		if list[j].MenuType != "F" {
			ms := menuLabelCall(eList, mi)
			min = append(min, ms)
		} else {
			min = append(min, mi)
		}
	}
	if len(min) > 0 {
		dept.Children = min
	} else {
		dept.Children = nil
	}
	return dept
}

// menuCall 构建菜单树
func menuCall(menuList *[]models.SysMenu, menu models.SysMenu) models.SysMenu {
	//list := *menuList

	// min := make([]models.SysMenu, 0)
	// for j := 0; j < len(list); j++ {

	// 	if menu.MenuId != list[j].ParentId {
	// 		continue
	// 	}
	// 	mi := models.SysMenu{}
	// 	mi.MenuId = list[j].MenuId
	// 	mi.MenuName = list[j].MenuName
	// 	mi.Title = list[j].Title
	// 	mi.Icon = list[j].Icon
	// 	mi.Path = list[j].Path
	// 	mi.MenuType = list[j].MenuType
	// 	mi.Action = list[j].Action
	// 	mi.Permission = list[j].Permission
	// 	mi.ParentId = list[j].ParentId
	// 	mi.NoCache = list[j].NoCache
	// 	mi.Breadcrumb = list[j].Breadcrumb
	// 	mi.Component = list[j].Component
	// 	mi.Sort = list[j].Sort
	// 	mi.Visible = list[j].Visible
	// 	mi.CreatedAt = list[j].CreatedAt
	// 	mi.SysApi = list[j].SysApi
	// 	mi.Children = []models.SysMenu{}

	// 	if mi.MenuType != cModels.Button {
	// 		ms := menuCall(menuList, mi)
	// 		min = append(min, ms)
	// 	} else {
	// 		min = append(min, mi)
	// 	}
	// }
	// menu.Children = min
	return menu
}

func menuDistinct(menuList []models.SysMenu) (result []models.SysMenu) {
	distinctMap := make(map[int]struct{}, len(menuList))
	for _, menu := range menuList {
		if _, ok := distinctMap[menu.MenuId]; !ok {
			distinctMap[menu.MenuId] = struct{}{}
			result = append(result, menu)
		}
	}
	return result
}

// func recursiveSetMenu(orm *gorm.DB, mIds []int, menus *[]models.SysMenu) error {
// 	if len(mIds) == 0 || menus == nil {
// 		return nil
// 	}
// 	var subMenus []models.SysMenu
// 	err := orm.Where(fmt.Sprintf(" menu_type in ('%s', '%s', '%s') and id in ?",
// 		cModels.Directory, cModels.Menu, cModels.Button), mIds).Order("sort").Find(&subMenus).Error
// 	if err != nil {
// 		return err
// 	}

// 	subIds := make([]int, 0)
// 	for _, menu := range subMenus {
// 		if menu.ParentId != 0 {
// 			subIds = append(subIds, menu.ParentId)
// 		}
// 		if menu.MenuType != cModels.Button {
// 			*menus = append(*menus, menu)
// 		}
// 	}
// 	return recursiveSetMenu(orm, subIds, menus)
// }

// SetMenuRole 获取左侧菜单树使用
func (e *SysMenu) SetMenuRole(roleName string) (m []models.SysMenu, err error) {
	menus, err := e.getByRoleName(roleName)
	m = make([]models.SysMenu, 0)
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentId != 0 {
			continue
		}
		menusInfo := menuCall(&menus, menus[i])
		m = append(m, menusInfo)
	}
	return
}

func (e *SysMenu) getByRoleName(roleName string) ([]models.SysMenu, error) {

	data := make([]models.SysMenu, 0)
	// var role models.SysRole
	// var err error
	// if roleName == "admin" {
	// 	err = core.DB().Where(" menu_type in ('M','C') and deleted_at is null").
	// 		Order("sort").
	// 		Find(&data).
	// 		Error
	// 	err = errors.WithStack(err)
	// } else {
	// 	role.RoleKey = roleName
	// 	err = core.DB().Model(&role).Where("role_key = ? ", roleName).Preload("SysMenu").First(&role).Error

	// 	if role.SysMenu != nil {
	// 		mIds := make([]int, 0)
	// 		for _, menu := range *role.SysMenu {
	// 			mIds = append(mIds, menu.MenuId)
	// 		}
	// 		if err := recursiveSetMenu(core.DB(), mIds, &data); err != nil {
	// 			return nil, err
	// 		}

	// 		data = menuDistinct(data)
	// 	}
	// }

	// sort.Sort(models.SysMenuSlice(data))
	return data, nil
}
