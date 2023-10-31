package service

import (
	"dilu/common/codes"
	"dilu/common/consts"
	"dilu/common/utils"
	"dilu/modules/dental/models"
	"dilu/modules/dental/service/dto"
	"dilu/modules/sys/enums"
	smodels "dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"time"

	"github.com/baowk/dilu-core/core/base"
	"github.com/baowk/dilu-core/core/errs"
)

type CustomerService struct {
	*base.BaseService
}

var SerCustomer = CustomerService{
	base.NewService(consts.DB_CRM),
}

func (s *CustomerService) Page(req dto.CustomerGetPageReq, teamId, userId int, list *[]models.Customer, total *int64) error {
	if teamId == 0 {
		return codes.ErrSys(nil)
	}
	db := s.DB().Offset(req.GetOffset()).Limit(req.GetSize()).Where("team_id = ?", teamId)
	var tm smodels.SysMember
	if err := service.SerSysMember.GetMember(teamId, userId, &tm); err != nil {
		return err
	}
	if tm.PostId == enums.Staff.Id {
		req.UserId = userId
	} else if tm.PostId > enums.Admin.Id {
		req.DeptPath = tm.DeptPath
	}
	if req.UserId != 0 {
		db.Where("user_id = ?", req.UserId)
	} else if req.DeptPath != "" {
		db.Where("dept_path like ", req.DeptPath+"%")
	}
	return db.Order("id desc").Find(list).Offset(-1).Limit(-1).Count(total).Error
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
	if customer.Birthday > 0 {
		customer.Age = utils.CmpAge(time.UnixMilli(int64(customer.Birthday)))
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
	if customer.Birthday > 0 {
		customer.Age = utils.CmpAge(time.UnixMilli(int64(customer.Birthday)))
	}
	if err := s.UpdateById(customer); err != nil {
		return codes.ErrSys(err)
	}
	return nil
}
