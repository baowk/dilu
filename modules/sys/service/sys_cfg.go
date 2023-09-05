package service

import (
	"errors"
	"strconv"

	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"

	"dilu/common/codes"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/errs"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SysCfgService struct {
	//base.BaseService
}

func (e *SysCfgService) Page(req dto.SysCfgGetPageReq, list *[]models.SysCfg, total *int64, reqId string) errs.IError {
	db := core.Db("sys")
	//TODO WHERE
	if err := db.Limit(req.GetSize()).Offset(req.GetOffset()).Find(list).Count(total).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*SysCfgService) Create(data *models.SysCfg, reqId string) errs.IError {
	if err := core.Db("sys").Create(data).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*SysCfgService) Update(data *models.SysCfg, reqId string) errs.IError {
	if err := core.Db("sys").Save(data).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*SysCfgService) Del(ids []int, reqId string) errs.IError {
	if err := core.Db("sys").Delete(&models.SysCfg{}, ids).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (*SysCfgService) Get(id int, data *models.SysCfg, reqId string) errs.IError {
	if err := core.Db("sys").First(data, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			berr := codes.ErrNotFound(strconv.Itoa(id), "SysCfg", reqId, err)
			core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
			return berr
		}
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}

func (e *SysCfgService) GetIds(ids []int, list *[]models.SysCfg, reqId string) errs.IError {
	if len(ids) == 1 {
		var data models.SysCfg
		if err := e.Get(ids[0], &data, reqId); err != nil {
			return err
		}
		*list = append(*list, data)
	} else {
		if err := core.Db("sys").Where("id in ?", ids).Find(list).Error; err != nil {
			berr := errs.Err(codes.FAILURE, reqId, err)
			core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
			return berr
		}
	}
	return nil
}
