package service

import (
	"fmt"
	"time"

	"dilu/modules/sys/models"

	"dilu/common/codes"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
	"go.uber.org/zap"
)

type SysServ struct {
	base.BaseService
}

func (s *SysServ) Init() errs.IError {
	core.DB().AutoMigrate(
		&models.Sys{},
	)

	cstr, err := core.Cache.Get("test")
	fmt.Printf("Init %s ,%v \n", cstr, err)
	//if err != nil || cstr == "" {
	d := models.Sys{
		Name: "goods",
	}
	d.UpdatedAt = time.Now().Unix()
	d.CreatedAt = d.UpdatedAt
	d.Status = 3
	d.UpdateBy = 3
	d.CreateBy = d.UpdateBy

	if err := core.DB().Create(&d).Error; err != nil {
		berr := errs.Err(codes.FAILURE, "reqId", err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	if err := core.Cache.Set("test", d, time.Hour); err != nil {
		berr := errs.Err(codes.FAILURE, "reqId", err)
		core.Log.Error(errs.DB_ERR.String(), zap.Error(berr))
		return berr
	}
	return nil
}
