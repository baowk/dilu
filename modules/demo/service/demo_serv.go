package service

import (
	"errors"
	"strconv"

	"dilu/modules/demo/models"
	"dilu/modules/demo/service/dto"

	"dilu/common/consts"

	"dilu/common/codes"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/errs"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type DemoService struct {
	//base.BaseService
}

func (e *DemoService) Page(req dto.DemePageReq, list *[]models.Demo, total *int64, reqId string) errs.IError {
	db := core.Db(consts.DB_DEMO)
	//TODO WHERE
	if err := db.Limit(req.GetSize()).Offset(req.GetOffset()).Find(list).Count(total).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*DemoService) Create(data *models.Demo, reqId string) errs.IError {
	if err := core.Db(consts.DB_DEMO).Create(data).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*DemoService) Update(data *models.Demo, reqId string) errs.IError {
	if err := core.Db(consts.DB_DEMO).Save(data).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*DemoService) Del(ids []int, reqId string) errs.IError {
	if err := core.Db(consts.DB_DEMO).Delete(&models.Demo{}, ids).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*DemoService) Get(id int, data *models.Demo, reqId string) errs.IError {
	if err := core.Db(consts.DB_DEMO).First(data, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			berr := codes.ErrNotFound(strconv.Itoa(id), "demo", reqId, err)
			core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
			return berr
		}
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (e *DemoService) GetIds(ids []int, list *[]models.Demo, reqId string) errs.IError {
	if len(ids) == 1 {
		var data models.Demo
		if err := e.Get(ids[0], &data, reqId); err != nil {
			return err
		}
		*list = append(*list, data)
	} else {
		if err := core.Db(consts.DB_DEMO).Where("id in ?", ids).Find(list).Error; err != nil {
			berr := errs.Err(codes.FAILURE, reqId, err)
			core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
			return berr
		}
	}
	return nil
}
