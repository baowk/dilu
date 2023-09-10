package service

import (
	"dilu/common/third/sms"
	"dilu/modules/sys/models"
	"time"

	"github.com/baowk/dilu-core/common/utils"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"go.uber.org/zap"
)

type SmsLog struct {
	*base.BaseService
}

var SerSms = SmsLog{
	base.NewService("sys"),
}

func (e *SmsLog) Send(phone string) error {
	var err error
	code := utils.RandNumberByLen(6)
	data := models.SmsLog{
		Phone:     phone,
		Code:      code,
		Type:      "Code",
		Status:    0,
		UseStatus: 0,
	}
	data.CreatedAt = time.Now().Unix()
	data.UpdatedAt = data.CreatedAt
	err = core.DB().Create(&data).Error
	if err != nil {
		core.Log.Error("EmailLogService Insert error", zap.Error(err))
		return err
	}
	sms.Send(phone, code)
	return nil
}

// 验证
func (e *SmsLog) Verify(phone, code string) bool {
	if core.Cfg.Server.Mode == "dev" && code == "666666" {
		return true
	}
	var err error
	var data models.SmsLog
	err = core.DB().Model(&data).Where(" mobile = ? ", phone).Order("id desc").First(&data).Error
	if err != nil {
		core.Log.Error("验证码错误", zap.Error(err))
		return false
	}
	if data.UseStatus == 1 {
		return false
	}
	if data.CreatedAt+int64(time.Minute.Seconds()*10) < time.Now().Unix() {
		if code == data.Code {
			data.UseStatus = 1
			data.UpdatedAt = time.Now().Unix()
			core.DB().Save(&data)
			return true
		}
	}
	return false
}
