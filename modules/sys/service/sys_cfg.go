package service

import (
	"dilu/common/codes"
	"dilu/modules/sys/models"
	"dilu/modules/sys/service/dto"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"go.uber.org/zap"
)

type SysCfgService struct {
	base.BaseService
}

var SysCfgS = SysCfgService{
	base.BaseService{
		DbName: "sys",
	},
}

func (e *SysCfgService) PageT(req dto.SysCfgGetPageReq, list *[]models.SysCfg, total *int64, reqId string) errs.IError {
	if err := e.DB().Limit(req.GetSize()).Offset(req.GetOffset()).Find(list).Count(total).Error; err != nil {
		berr := errs.Err(codes.FAILURE, reqId, err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}
