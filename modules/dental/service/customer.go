package service

import (
	"dilu/common/codes"
	"dilu/common/consts"
	"dilu/modules/dental/models"

	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
)

type CustomerService struct {
	*base.BaseService
}

var SerCustomer = CustomerService{
	base.NewService(consts.DB_CRM),
}

func (s *CustomerService) GetByUserIdAndName(userId, teamId int, name string, customer *[]models.Customer) errs.IError {
	where := models.Customer{
		UserId: userId,
		TeamId: teamId,
		Name:   name,
	}
	if err := s.GetByWhere(where, customer); err != nil {
		return codes.ErrSys(err)
	}
	return nil
}
