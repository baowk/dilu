package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// Bill
type Bill struct {
	Id             int             `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"`  //主键
	No             string          `json:"no" gorm:"type:varchar(20);comment:订单号"`                           //订单号
	CustomerId     int             `json:"customerId" gorm:"type:int unsigned;comment:顾客"`                   //顾客
	UserId         int             `json:"userId" gorm:"type:int unsigned;comment:用户id"`                     //用户id
	TeamId         int             `json:"teamId" gorm:"type:int unsigned;comment:团队id"`                     //团队id
	DeptPath       string          `json:"deptPath" gorm:"type:varchar(255);comment:部门路径"`                   //路径
	Amount         decimal.Decimal `json:"amount" gorm:"type:decimal(10,2);comment:金额"`                      //金额
	RealAmount     decimal.Decimal `json:"realAmount" gorm:"type:decimal(10,2);comment:折后金额"`                //折后金额
	PaidAmount     decimal.Decimal `json:"paidAmount" gorm:"type:decimal(10,2);comment:已支付金额"`               //已支付金额
	DebtAmount     decimal.Decimal `json:"debtAmount" gorm:"type:decimal(10,2);comment:回收上月欠款"`              //回收上月欠款
	RefundAmount   decimal.Decimal `json:"refundAmount" gorm:"type:decimal(10,2);comment:退款"`                //退款
	LinkId         int             `json:"linkId" gorm:"type:int unsigned;comment:关联订单"`                     //关联订单
	TradeAt        time.Time       `json:"tradeAt" gorm:"type:datetime;default:(-);comment:交易日期"`            //交易日期
	TradeType      int             `json:"tradeType" gorm:"type:tinyint;comment:交易类型1 成交 2补尾款  3补上月欠款 10退款"` //交易类型 1 成交 2补尾款  3补上月欠款 10退款
	DentalCount    int             `json:"dentalCount" gorm:"type:tinyint;comment:颗数"`                       //颗数
	Brand          int             `json:"brand" gorm:"type:tinyint;comment:品牌"`                             //品牌
	ImplantedCount int             `json:"implantedCount" gorm:"type:tinyint;comment:已种颗数"`                  //已种颗数
	Implant        int             `json:"implant" gorm:"type:tinyint;comment:种植状态：1 未种 2部分 3已种"`            //种植状态：1 未种 2部分 3已种
	ImplantDate    time.Time       `json:"implantDate" gorm:"type:datetime;default:(-);comment:植入日期"`        //植入日期
	Doctor         string          `json:"doctor" gorm:"type:varchar(32);comment:医生"`                        //医生
	Pack           int             `json:"pack" gorm:"type:tinyint;comment:1 普通 2 半口 3 全口"`                  //1 普通 2 半口 3 全口
	PaybackDate    time.Time       `json:"paybackDate" gorm:"type:datetime;default:(-);comment:预定回款日期"`      //预定回款日期
	Tags           string          `json:"tags" gorm:"type:varchar(255);comment:标签"`                         //标签
	PrjName        string          `json:"prjName" gorm:"type:varchar(255);comment:项目"`                      //种植项目
	OtherPrj       string          `json:"otherPrj" gorm:"type:varchar(255);comment:其他项目"`                   //其他项目
	Remark         string          `json:"remark" gorm:"type:varchar(255);comment:备注"`                       //备注
	CreatedAt      time.Time       `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                      //创建时间
	UpdatedAt      time.Time       `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                      //更新时间
	CreateBy       int             `json:"createBy" gorm:"type:int unsigned;index;comment:创建者"`              //创建者id
	UpdateBy       int             `json:"updateBy" gorm:"type:int unsigned;index;comment:更新者"`              //更新者id
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
func (e *Bill) SetAmount(amount decimal.Decimal) *Bill {
	e.Amount = amount
	return e
}
func (e *Bill) SetRealAmount(realAmount decimal.Decimal) *Bill {
	e.RealAmount = realAmount
	return e
}
func (e *Bill) SetPaidAmount(paidAmount decimal.Decimal) *Bill {
	e.PaidAmount = paidAmount
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
func (e *Bill) SetTradeType(tradeType int) *Bill {
	e.TradeType = tradeType
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
