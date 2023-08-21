package service

// import (
// 	"errors"
// 	"fmt"
// 	"time"

// 	"github.com/go-admin-team/go-admin-core/sdk/config"
// 	"github.com/go-admin-team/go-admin-core/sdk/service"
// 	"gorm.io/gorm"

// 	"go-admin/app/sso/models"
// 	"go-admin/app/sso/service/dto"
// 	"go-admin/common/actions"
// 	cDto "go-admin/common/dto"
// 	"go-admin/common/utils"
// 	"go-admin/common/utils/email"
// )

// type EmailLog struct {
// 	service.Service
// }

// // GetPage 获取EmailLog列表
// func (e *EmailLog) GetPage(c *dto.EmailLogGetPageReq, p *actions.DataPermission, list *[]models.EmailLog, count *int64) error {
// 	var err error
// 	var data models.EmailLog

// 	err = e.Orm.Model(&data).
// 		Scopes(
// 			cDto.MakeCondition(c.GetNeedSearch()),
// 			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
// 			actions.Permission(data.TableName(), p),
// 		).
// 		Find(list).Limit(-1).Offset(-1).
// 		Count(count).Error
// 	if err != nil {
// 		e.Log.Errorf("EmailLogService GetPage error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // Get 获取EmailLog对象
// func (e *EmailLog) Get(d *dto.EmailLogGetReq, p *actions.DataPermission, model *models.EmailLog) error {
// 	var data models.EmailLog

// 	err := e.Orm.Model(&data).
// 		Scopes(
// 			actions.Permission(data.TableName(), p),
// 		).
// 		First(model, d.GetId()).Error
// 	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
// 		err = errors.New("查看对象不存在或无权查看")
// 		e.Log.Errorf("Service GetEmailLog error:%s \r\n", err)
// 		return err
// 	}
// 	if err != nil {
// 		e.Log.Errorf("db error:%s", err)
// 		return err
// 	}
// 	return nil
// }

// // Insert 创建EmailLog对象
// func (e *EmailLog) Insert(c *dto.EmailLogInsertReq) error {
// 	var err error
// 	var data models.EmailLog
// 	c.Generate(&data)
// 	err = e.Orm.Create(&data).Error
// 	if err != nil {
// 		e.Log.Errorf("EmailLogService Insert error:%s \r\n", err)
// 		return err
// 	}
// 	return nil
// }

// // Update 修改EmailLog对象
// func (e *EmailLog) Update(c *dto.EmailLogUpdateReq, p *actions.DataPermission) error {
// 	var err error
// 	var data = models.EmailLog{}
// 	e.Orm.Scopes(
// 		actions.Permission(data.TableName(), p),
// 	).First(&data, c.GetId())
// 	c.Generate(&data)

// 	db := e.Orm.Save(&data)
// 	if err = db.Error; err != nil {
// 		e.Log.Errorf("EmailLogService Save error:%s \r\n", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权更新该数据")
// 	}
// 	return nil
// }

// // Remove 删除EmailLog
// func (e *EmailLog) Remove(d *dto.EmailLogDeleteReq, p *actions.DataPermission) error {
// 	var data models.EmailLog

// 	db := e.Orm.Model(&data).
// 		Scopes(
// 			actions.Permission(data.TableName(), p),
// 		).Delete(&data, d.GetId())
// 	if err := db.Error; err != nil {
// 		e.Log.Errorf("Service RemoveEmailLog error:%s \r\n", err)
// 		return err
// 	}
// 	if db.RowsAffected == 0 {
// 		return errors.New("无权删除该数据")
// 	}
// 	return nil
// }

// var mail_tmp = `
// 	你的验证码为：%s,10分钟内有效！
// `

// // 发送验证码
// func (e *EmailLog) Send(mail string) error {
// 	var err error
// 	code := utils.RandNumberByLen(6)
// 	data := models.EmailLog{
// 		Email:     mail,
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
// 	content := fmt.Sprintf(mail_tmp, code)
// 	email.Send("", mail, "验证码消息", content)
// 	return nil
// }

// // 验证
// func (e *EmailLog) Verify(mail, code string) bool {
// 	if config.ApplicationConfig.Mode == "dev" && code == "666666" {
// 		return true
// 	}
// 	var err error
// 	var data models.EmailLog

// 	err = e.Orm.Model(&data).Where(" email = ? ", mail).Order("id desc").First(&data).Error
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
