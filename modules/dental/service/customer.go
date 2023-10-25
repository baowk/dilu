package service

import (
	"dilu/common/codes"
	"dilu/common/consts"
	"dilu/common/utils"
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

func (s *CustomerService) GetByIds(teamId int, ids []int, customers *[]models.Customer) error {
	db := s.DB().Where("id in ?", ids)
	if teamId > 0 {
		db.Where("team_id = ?", teamId)
	}
	return db.Find(customers).Error
}

func (s *CustomerService) Create(customer *models.Customer) errs.IError {
	if customer.Name != "" {
		customer.PY = utils.GetPinyin(customer.Name)
	}
	if err := s.BaseService.Create(customer); err != nil {
		return codes.ErrSys(err)
	}
	return nil
}

func (s *CustomerService) Update(customer *models.Customer) errs.IError {
	if customer.Name != "" {
		customer.PY = utils.GetPinyin(customer.Name)
	}
	if err := s.UpdateById(customer); err != nil {
		return codes.ErrSys(err)
	}
	return nil
}
