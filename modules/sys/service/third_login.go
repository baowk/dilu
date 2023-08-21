package service

// import (
// 	"errors"
// 	"time"

// 	"github.com/go-admin-team/go-admin-core/sdk/service"
// 	"gorm.io/gorm"

// 	"go-admin/app/sso/models"
// 	"go-admin/app/sso/service/dto"
// 	"go-admin/common/actions"
// 	cDto "go-admin/common/dto"
// )

// type ThirdLogin struct {
// 	service.Service
// }

// // GetPage 获取ThirdLogin列表
// func (e *ThirdLogin) GetPage(c *dto.ThirdLoginGetPageReq, p *actions.DataPermission, list *[]models.ThirdLogin, count *int64) error {
// 	var err error
// 	var data models.ThirdLogin

// 	err = e.Orm.Model(&data).
// 		Scopes(
// 			cDto.MakeCondition(c.GetNeedSearch()),
// 			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
// 			actions.Permission(data.TableName(), p),
// 		).
// 		Find(list).Limit(-1).Offset(-1).
// 		Count(count).Error
// 	if err != nil {
// 		e.Log.Errorf("ThirdLoginService GetPage error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // Get 获取ThirdLogin对象
// func (e *ThirdLogin) Get(d *dto.ThirdLoginGetReq, p *actions.DataPermission, model *models.ThirdLogin) error {
// 	var data models.ThirdLogin

// 	err := e.Orm.Model(&data).
// 		Scopes(
// 			actions.Permission(data.TableName(), p),
// 		).
// 		First(model, d.GetId()).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		err = errors.New("查看对象不存在或无权查看")
// 		e.Log.Errorf("Service GetThirdLogin error:%s \r\n", err)
// 		return err
// 	}
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // Insert 创建ThirdLogin对象
// func (e *ThirdLogin) Insert(c *dto.ThirdLoginInsertReq) error {
// 	var err error
// 	var data models.ThirdLogin
// 	c.Generate(&data)
// 	err = e.Orm.Create(&data).Error
// 	if err != nil {
// 		e.Log.Errorf("ThirdLoginService Insert error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // Insert 创建ThirdLogin对象
// func (e *ThirdLogin) Create(data models.ThirdLogin) error {
// 	err := e.Orm.Create(&data).Error
// 	if err != nil {
// 		e.Log.Errorf("ThirdLoginService Insert error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // Update 修改ThirdLogin对象
// func (e *ThirdLogin) Update(c *dto.ThirdLoginUpdateReq, p *actions.DataPermission) error {
// 	var err error
// 	var data = models.ThirdLogin{}
// 	e.Orm.Scopes(
// 		actions.Permission(data.TableName(), p),
// 	).First(&data, c.GetId())
// 	c.Generate(&data)

// 	db := e.Orm.Save(&data)
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("ThirdLoginService Save error:%s \r\n", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权更新该数据")
// 	}
// 	return nil
// }

// // Remove 删除ThirdLogin
// func (e *ThirdLogin) Remove(d *dto.ThirdLoginDeleteReq, p *actions.DataPermission) error {
// 	var data models.ThirdLogin

// 	db := e.Orm.Model(&data).
// 		Scopes(
// 			actions.Permission(data.TableName(), p),
// 		).Delete(&data, d.GetId())
// 	if err := db.Error; err != nil {
// 		e.Log.Errorf("Service RemoveThirdLogin error:%s \r\n", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权删除该数据")
// 	}
// 	return nil
// }

// func (e *ThirdLogin) GetById(id int, model *models.ThirdLogin) error {
// 	var data models.ThirdLogin
// 	err := e.Orm.Model(&data).
// 		First(model, id).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil
// 	}
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// func (e *ThirdLogin) GetTL(platform int, openId, unionId string, model *models.ThirdLogin) error {
// 	var data models.ThirdLogin

// 	if openId == "" && unionId == "" {
// 		return errors.New("参数错误")
// 	}

// 	db := e.Orm.Model(&data)

// 	unionFlag := false

// 	if unionId != "" {
// 		db.Where("platform = ? and union_id = ", platform, unionId)
// 		unionFlag = true
// 	} else {
// 		db.Where("platform = ? and  open_id = ?", platform, openId)
// 	}

// 	err := db.First(model).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		if unionFlag {
// 			db.Where("platform = ? and  open_id = ?", platform, openId)
// 			err = db.First(model).Error
// 			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
// 				e.Log.Errorf("db error:%s", err)
// 				return err
// 			}
// 		}
// 	} else {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // Update 修改ThirdLogin对象
// func (e *ThirdLogin) UpdateUserId(userId string, data models.ThirdLogin) error {
// 	var err error
// 	updates := models.ThirdLogin{
// 		UserId:    userId,
// 		UpdatedAt: time.Now(),
// 	}

// 	db := e.Orm.Model(data).Updates(updates)
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("ThirdLoginService Save error:%s \r\n", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权更新该数据")
// 	}
// 	return nil
// }
