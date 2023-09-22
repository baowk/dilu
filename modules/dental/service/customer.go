package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type CustomerService struct {
	*base.BaseService
}

var SerCustomer = CustomerService{
	base.NewService("dental"),
}

