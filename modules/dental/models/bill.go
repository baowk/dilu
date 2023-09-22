package models

import (
	"time"
)

// 账单
type Bill struct {
	Id             int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	No             string    `json:"no" gorm:"type:varchar(20);comment:订单号"`                          //订单号
	CustomerId     int       `json:"customerId" gorm:"type:int;comment:顾客"`                           //顾客
	UserId         int       `json:"userId" gorm:"type:int;comment:用户id"`                             //用户id
	TeamId         int       `json:"teamId" gorm:"type:int;comment:团队id"`                             //团队id
	Total          string    `json:"total" gorm:"type:decimal(10,2);comment:金额"`                      //金额
	RealTotal      string    `json:"realTotal" gorm:"type:decimal(10,2);comment:折后金额"`                //折后金额
	PaidTotal      string    `json:"paidTotal" gorm:"type:decimal(10,2);comment:已支付金额"`               //已支付金额
	LinkId         int       `json:"linkId" gorm:"type:int;comment:关联订单"`                             //关联订单
	TradeAt        time.Time `json:"tradeAt" gorm:"type:datetime;comment:交易日期"`                       //交易日期
	TradeStatus    int       `json:"tradeStatus" gorm:"type:tinyint;comment:交易类型 1 成交 2补尾款 3退款"`      //交易类型 1 成交 2补尾款 3退款
	DentalCount    int       `json:"dentalCount" gorm:"type:tinyint;comment:颗数"`                      //颗数
	Brand          int       `json:"brand" gorm:"type:tinyint;comment:品牌"`                            //品牌
	ImplantedCount int       `json:"implantedCount" gorm:"type:tinyint;comment:已种颗数"`                 //已种颗数
	Implant        int       `json:"implant" gorm:"type:tinyint;comment:是否已种"`                        //是否已种
	ImplantDate    time.Time `json:"implantDate" gorm:"type:datetime;comment:植入日期"`                   //植入日期
	Doctor         string    `json:"doctor" gorm:"type:varchar(32);comment:医生"`                       //医生
	Pack           int       `json:"pack" gorm:"type:tinyint;comment:1 普通 2 半口 3 全口"`                 //1 普通 2 半口 3 全口
	PaybackDate    time.Time `json:"paybackDate" gorm:"type:datetime;comment:预定回款日期"`                 //预定回款日期
	CreatedAt      time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                     //创建时间
	UpdatedAt      time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                     //更新时间
}

func (Bill) TableName() string {
	return "bill"
}

func NewBill() *Bill {
	return &Bill{}
}

func (e *Bill) SetId(id int) *Bill {
	e.Id = id
	return e
}
func (e *Bill) SetNo(no string) *Bill {
	e.No = no
	return e
}
func (e *Bill) SetCustomerId(customerId int) *Bill {
	e.CustomerId = customerId
	return e
}
func (e *Bill) SetUserId(userId int) *Bill {
	e.UserId = userId
	return e
}
func (e *Bill) SetTeamId(teamId int) *Bill {
	e.TeamId = teamId
	return e
}
func (e *Bill) SetTotal(total string) *Bill {
	e.Total = total
	return e
}
func (e *Bill) SetRealTotal(realTotal string) *Bill {
	e.RealTotal = realTotal
	return e
}
func (e *Bill) SetPaidTotal(paidTotal string) *Bill {
	e.PaidTotal = paidTotal
	return e
}
func (e *Bill) SetLinkId(linkId int) *Bill {
	e.LinkId = linkId
	return e
}
func (e *Bill) SetTradeAt(tradeAt time.Time) *Bill {
	e.TradeAt = tradeAt
	return e
}
func (e *Bill) SetTradeStatus(tradeStatus int) *Bill {
	e.TradeStatus = tradeStatus
	return e
}
func (e *Bill) SetDentalCount(dentalCount int) *Bill {
	e.DentalCount = dentalCount
	return e
}
func (e *Bill) SetBrand(brand int) *Bill {
	e.Brand = brand
	return e
}
func (e *Bill) SetImplantedCount(implantedCount int) *Bill {
	e.ImplantedCount = implantedCount
	return e
}
func (e *Bill) SetImplant(implant int) *Bill {
	e.Implant = implant
	return e
}
func (e *Bill) SetImplantDate(implantDate time.Time) *Bill {
	e.ImplantDate = implantDate
	return e
}
func (e *Bill) SetDoctor(doctor string) *Bill {
	e.Doctor = doctor
	return e
}
func (e *Bill) SetPack(pack int) *Bill {
	e.Pack = pack
	return e
}
func (e *Bill) SetPaybackDate(paybackDate time.Time) *Bill {
	e.PaybackDate = paybackDate
	return e
}
func (e *Bill) SetCreatedAt(createdAt time.Time) *Bill {
	e.CreatedAt = createdAt
	return e
}
func (e *Bill) SetUpdatedAt(updatedAt time.Time) *Bill {
	e.UpdatedAt = updatedAt
	return e
}
