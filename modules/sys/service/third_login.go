package service

import (
	"dilu/modules/sys/models"
	"time"

	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ThirdLogin struct {
	*base.BaseService
}

var SerThirdLogin = ThirdLogin{
	base.NewService(consts.DB_DEF),
}

// Insert 创建ThirdLogin对象
func (e *ThirdLogin) Create(data *models.ThirdLogin) error {
	err := core.DB().Create(data).Error
	if err != nil {
		core.Log.Error("ThirdLogin", zap.Error(err))
		return err
	}
	return nil
}

func (e *ThirdLogin) GetById(id int, model *models.ThirdLogin) error {
	var data models.ThirdLogin
	err := core.DB().Model(&data).
		First(model, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	if err != nil {
		core.Log.Error("ThirdLogin", zap.Error(err))
		return err
	}
	return nil
}

func (e *ThirdLogin) GetTL(platform int, openId, unionId string, model *models.ThirdLogin) error {
	var data models.ThirdLogin

	if openId == "" && unionId == "" {
		return errors.New("参数错误")
	}

	db := core.DB().Model(&data)

	unionFlag := false

	if unionId != "" {
		db.Where("platform = ? and union_id = ", platform, unionId)
		unionFlag = true
	} else {
		db.Where("platform = ? and  open_id = ?", platform, openId)
	}

	err := db.First(model).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if unionFlag {
			db.Where("platform = ? and  open_id = ?", platform, openId)
			err = db.First(model).Error
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				core.Log.Error("ThirdLogin", zap.Error(err))
				return err
			}
		}
	} else {
		core.Log.Error("ThirdLogin", zap.Error(err))
		return err
	}
	return nil
}

// Update 修改ThirdLogin对象
func (e *ThirdLogin) UpdateUserId(userId int, data models.ThirdLogin) error {
	var err error
	updates := models.ThirdLogin{
		UserId: userId,
	}
	updates.UpdatedAt = time.Now().Unix()

	db := core.DB().Model(data).Updates(updates)
	if err = db.Error; err != nil {
		core.Log.Error("ThirdLogin", zap.Error(err))
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}
