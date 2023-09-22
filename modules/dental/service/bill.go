package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type BillService struct {
	*base.BaseService
}

var SerBill = BillService{
	base.NewService("dental"),
}

