package service

import (
	"dilu/common/codes"
	"dilu/common/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"dilu/modules/sys/enums"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"
)

type SysMenu struct {
	*base.BaseService
}

var SerSysMenu = SysMenu{
	base.NewService(consts.DB_DEF),
}

// GetPage 获取SysMenu列表
// func (e *SysMenu) GetPage(c *dto.SysMenuGetPageReq, menus *[]models.SysMenu) (*SysMenu, errs.IError) {
// 	var menu = make([]models.SysMenu, 0)
// 	db := core.DB()
// 	//TODO WHERE
// 	if err := db.Limit(c.GetSize()).Offset(c.GetOffset()).Find(menu).Error; err != nil {
// 		berr := errs.Err(codes.FAILURE, "", err)
// 		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
// 		return e, berr
// 	}
// 	for i := 0; i < len(menu); i++ {
// 		if menu[i].ParentId != 0 {
// 			continue
// 		}
// 		menusInfo := menuCall(&menu, menu[i])
// 		*menus = append(*menus, menusInfo)
// 	}
// 	return e, nil
// }

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
	return e, nil
}

// Insert 创建SysMenu对象
func (e *SysMenu) Insert(data *models.SysMenu) errs.IError {
	var err error
	tx := core.DB().Debug().Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = tx.Create(&data).Error
	if err != nil {
		tx.Rollback()
		core.Log.Error("sys_menu", zap.Error(err))
		berr := errs.Err(codes.FAILURE, "", err)
		return berr
	}
	tx.Commit()
	return nil
}

// func (e *SysMenu) initPaths(tx *gorm.DB, menu *models.SysMenu) error {
// 	var err error
// 	var data models.SysMenu
// 	parentMenu := new(models.SysMenu)
// 	if menu.ParentId != 0 {
// 		err = tx.Model(&data).First(parentMenu, menu.ParentId).Error
// 		if err != nil {
// 			return err
// 		}
// 		if parentMenu.Paths == "" {
// 			err = errors.New("父级paths异常，请尝试对当前节点父级菜单进行更新操作！")
// 			return err
// 		}
// 		menu.Paths = parentMenu.Paths + "/" + strconv.Itoa(menu.Id)
// 	} else {
// 		menu.Paths = "/0/" + strconv.Itoa(menu.Id)
// 	}
// 	err = tx.Model(&data).Where("id = ?", menu.Id).Update("paths", menu.Paths).Error
// 	return err
// }

// Update 修改SysMenu对象
// func (e *SysMenu) Update(c *models.SysMenu) (*SysMenu, errs.IError) {
// 	var err error
// 	tx := core.DB().Debug().Begin()
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		} else {
// 			tx.Commit()
// 		}
// 	}()
// 	var alist = make([]models.SysApi, 0)
// 	var model = models.SysMenu{}
// 	tx.Preload("SysApi").First(&model, c.Id())
// 	oldPath := model.Paths
// 	tx.Where("id in ?", c.Apis).Find(&alist)

// 	var menuList []models.SysMenu
// 	tx.Where("paths like ?", oldPath+"%").Find(&menuList)
// 	for _, v := range menuList {
// 		v.Paths = strings.Replace(v.Paths, oldPath, model.Paths, 1)
// 		tx.Model(&v).Update("paths", v.Paths)
// 	}
// 	return e, nil
// }

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

func (e *SysMenu) GetMenus(c *gin.Context, mvs *[]models.SysMenu) errs.IError {
	role := utils.GetRoleId(c)
	var where string
	if role != 0 { //超管
		where = "platform_type <= ?"
	} else {
		where = "platform_type >= ?"
	}
	if err := e.DB().Where(where, enums.MenuPub).Find(mvs).Error; err != nil {
		return codes.ErrSys(err)
	}
	return nil
}

// func treeMenu(ms []models.SysMenu) []models.SysMenu {
// 	mvs := make([]models.SysMenu, 0)
// 	for _, menu := range ms {
// 		if menu.ParentId == 0 {
// 			menuCall(ms, &menu)
// 			mvs = append(mvs, menu)
// 		}
// 	}
// 	return mvs
// }

// // menuCall 构建菜单树
// func menuCall(ms []models.SysMenu, menu *models.SysMenu) {
// 	children := make([]models.SysMenu, 0)
// 	for _, m := range ms {
// 		if menu.Id != m.ParentId {
// 			continue
// 		}
// 		menuCall(ms, &m)
// 		children = append(children, m)

// 	}
// 	menu.Children = children
// }

func (e *SysMenu) GetRoles(c *gin.Context) (platform, teamId int, roles []int, ierr errs.IError) {
	role := utils.GetRoleId(c)
	platform = 2
	if role != 0 { //超管
		platform = 1
		if role > 0 {
			roles = append(roles, role)
		}
	} else { //团队菜单
		teamId = utils.GetTeamId(c)
		if teamId < 1 {
			ierr = codes.Err403(errors.New("团队id不存在"))
			return
		}
		var tu models.SysMember
		if err := SerSysMember.GetMember(teamId, utils.GetUserId(c), &tu); err != nil {
			ierr = codes.Err403(err)
			return
		}
		if tu.Roles == "" {
			ierr = codes.Err403(nil)
			return
		}
		if !strings.Contains(tu.Roles, "-1") {
			arr := strings.Split(tu.Roles, ",")
			for _, sid := range arr {
				id, err := strconv.Atoi(sid)
				if err != nil {
					return
				}
				roles = append(roles, id)
			}
		}
	}
	return
}

func (e *SysMenu) CanAccess(c *gin.Context, apiId int) error {
	platform, _, roles, err := e.GetRoles(c)
	if err != nil {
		return err
	}
	var ids []int

	if len(roles) > 0 {
		if err := e.DB().Raw("select sys_api_id from sys_menu_api_rule r,sys_role_menu rm  where  rm.role_id in ? and rm.menu_id = r.sys_menu_id", roles).
			Find(&ids).Error; err != nil {
			return err
		}
	} else {
		if err := e.DB().Raw("select sys_api_id from sys_menu_api_rule r,sys_menu m  where m.platform_type >= ? and m.id = r.sys_menu_id", platform).
			Find(&ids).Error; err != nil {
			return err
		}
	}

	for _, id := range ids {
		if id == apiId {
			return nil
		}
	}
	return codes.Err403(nil)
}

func (e *SysMenu) GetUserMenus(c *gin.Context, mvs *[]dto.MenuVo) errs.IError {
	platform, _, roles, err := e.GetRoles(c)
	if err != nil {
		return err
	}
	var where string
	if platform == 1 {
		where = "platform_type <= ?"
	} else {
		where = "platform_type >= ?"
	}
	db := e.DB().Where(where, enums.MenuPub)
	if len(roles) > 0 {
		db.Joins(" left join sys_role_menu on sys_role_menu.menu_id = sys_menu.id").
			Where("sys_role_menu.role_id in ?", roles)
	}
	var ms []models.SysMenu
	if err := db.Find(&ms).Error; err != nil {
		return codes.ErrSys(err)
	}
	*mvs = treeMenuVo(ms)
	return nil
}

func treeMenuVo(ms []models.SysMenu) []dto.MenuVo {
	mvs := make([]dto.MenuVo, 0)
	for _, menu := range ms {
		if menu.ParentId == 0 {
			vo := menuToVo(menu)
			menuCallVo(ms, &vo)
			mvs = append(mvs, vo)
		}
	}
	return mvs
}

// menuCall 构建菜单树
func menuCallVo(ms []models.SysMenu, menu *dto.MenuVo) {
	children := make([]dto.MenuVo, 0)
	for _, m := range ms {
		if menu.Id != m.ParentId {
			continue
		}
		if m.MenuType < 3 {
			vo := menuToVo(m)
			menuCallVo(ms, &vo)
			children = append(children, vo)
		} else {
			menu.Meta.Auths = append(menu.Meta.Auths, m.Permission)
		}
	}
	menu.Children = children
}

func menuToVo(menu models.SysMenu) dto.MenuVo {
	meta := dto.RouteMeta{
		Title: menu.Title,
		Icon:  menu.Icon,
	}
	if !menu.Hidden {
		meta.ShowLink = true
	}
	if !menu.NoCache {
		meta.KeepAlive = true
	}
	if menu.Sort > 0 {
		meta.Rank = menu.Sort
	}
	vo := dto.MenuVo{
		Name:      menu.MenuName,
		Meta:      meta,
		Path:      menu.Path,
		Component: menu.Component,
		Id:        menu.Id,
	}
	return vo
}

func (e *SysMenu) GetUserPerms(roleId int, mvs *[]string) errs.IError {
	var sql string
	if roleId == 1 {
		sql = "Select permission from sys_menu  where menu_type > 1 "
	} else {
		sql = fmt.Sprintf("Select permission from sys_menu m,sys_role_menu r where role_id = %d and menu_type > 1 and m.id = r.menu_id", roleId)
	}
	var ms []string
	if err := core.DB().Raw(sql).Find(&ms).Error; err != nil {
		return codes.ErrSys(err)
	}
	*mvs = ms
	return nil
}
