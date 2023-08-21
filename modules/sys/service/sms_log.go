package service

// import (
// 	"errors"
// 	"time"

// 	"github.com/go-admin-team/go-admin-core/sdk/config"
// 	"github.com/go-admin-team/go-admin-core/sdk/service"
// 	"gorm.io/gorm"

// 	"go-admin/app/sso/models"
// 	"go-admin/app/sso/service/dto"
// 	"go-admin/common/actions"
// 	cDto "go-admin/common/dto"
// 	"go-admin/common/utils"
// 	"go-admin/common/utils/sms"
// )

// type SmsLog struct {
// 	service.Service
// }

// // GetPage 获取SmsLog列表
// func (e *SmsLog) GetPage(c *dto.SmsLogGetPageReq, p *actions.DataPermission, list *[]models.SmsLog, count *int64) error {
// 	var err error
// 	var data models.SmsLog

// 	err = e.Orm.Model(&data).
// 		Scopes(
// 			cDto.MakeCondition(c.GetNeedSearch()),
// 			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
// 			actions.Permission(data.TableName(), p),
// 		).
// 		Find(list).Limit(-1).Offset(-1).
// 		Count(count).Error
// 	if err != nil {
// 		e.Log.Errorf("SmsLogService GetPage error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // Get 获取SmsLog对象
// func (e *SmsLog) Get(d *dto.SmsLogGetReq, p *actions.DataPermission, model *models.SmsLog) error {
// 	var data models.SmsLog

// 	err := e.Orm.Model(&data).
// 		Scopes(
// 			actions.Permission(data.TableName(), p),
// 		).
// 		First(model, d.GetId()).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		err = errors.New("查看对象不存在或无权查看")
// 		e.Log.Errorf("Service GetSmsLog error:%s \r\n", err)
// 		return err
// 	}
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // Insert 创建SmsLog对象
// func (e *SmsLog) Insert(c *dto.SmsLogInsertReq) error {
// 	var err error
// 	var data models.SmsLog
// 	c.Generate(&data)
// 	err = e.Orm.Create(&data).Error
// 	if err != nil {
// 		e.Log.Errorf("SmsLogService Insert error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // Update 修改SmsLog对象
// func (e *SmsLog) Update(c *dto.SmsLogUpdateReq, p *actions.DataPermission) error {
// 	var err error
// 	var data = models.SmsLog{}
// 	e.Orm.Scopes(
// 		actions.Permission(data.TableName(), p),
// 	).First(&data, c.GetId())
// 	c.Generate(&data)

// 	db := e.Orm.Save(&data)
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("SmsLogService Save error:%s \r\n", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权更新该数据")
// 	}
// 	return nil
// }

// // Remove 删除SmsLog
// func (e *SmsLog) Remove(d *dto.SmsLogDeleteReq, p *actions.DataPermission) error {
// 	var data models.SmsLog

// 	db := e.Orm.Model(&data).
// 		Scopes(
// 			actions.Permission(data.TableName(), p),
// 		).Delete(&data, d.GetId())
// 	if err := db.Error; err != nil {
// 		e.Log.Errorf("Service RemoveSmsLog error:%s \r\n", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权删除该数据")
// 	}
// 	return nil
// }

// func (e *SmsLog) Send(phone string) error {
// 	var err error
// 	code := utils.RandNumberByLen(6)
// 	data := models.SmsLog{
// 		Mobile:    phone,
// 		Code:      code,
// 		Type:      "Code",
// 		Status:    0,
// 		UseStatus: 0,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// 	err = e.Orm.Create(&data).Error
// 	if err != nil {
// 		e.Log.Errorf("EmailLogService Insert error:%s \r\n", err)
// 		return err
// 	}
// 	sms.Send(phone, code)
// 	return nil
// }

// // 验证
// func (e *SmsLog) Verify(phone, code string) bool {
// 	if config.ApplicationConfig.Mode == "dev" && code == "666666" {
// 		return true
// 	}
// 	var err error
// 	var data models.SmsLog
// 	err = e.Orm.Model(&data).Where(" mobile = ? ", phone).Order("id desc").First(&data).Error
// 	if err != nil {
// 		e.Log.Errorf("验证码错误:%s \r\n", err)
// 		return false
// 	}
// 	if data.UseStatus == 1 {
// 		return false
// 	}
// 	if data.CreatedAt.Add(time.Minute * 10).After(time.Now()) {
// 		if code == data.Code {
// 			data.UseStatus = 1
// 			data.UpdatedAt = time.Now()
// 			e.Orm.Save(&data)
// 			return true
// 		}
// 	}

// 	return false
// }
