package dto

import (
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type BillGetPageReq struct {
	base.ReqPage `search:"-"`
	TradeStatus  int `json:"tradeStatus" form:"tradeStatus"` //交易类型 1 成交 2补尾款 3退款
}

// Bill
type BillDto struct {
	Id             int       `json:"id"`             //主键
	No             string    `json:"no"`             //订单号
	CustomerId     int       `json:"customerId"`     //顾客
	UserId         int       `json:"userId"`         //用户id
	TeamId         int       `json:"teamId"`         //团队id
	Total          string    `json:"total"`          //金额
	RealTotal      string    `json:"realTotal"`      //折后金额
	PaidTotal      string    `json:"paidTotal"`      //已支付金额
	LinkId         int       `json:"linkId"`         //关联订单
	TradeAt        time.Time `json:"tradeAt"`        //交易日期
	TradeStatus    int       `json:"tradeStatus"`    //交易类型 1 成交 2补尾款 3退款
	DentalCount    int       `json:"dentalCount"`    //颗数
	Brand          int       `json:"brand"`          //品牌
	ImplantedCount int       `json:"implantedCount"` //已种颗数
	Implant        int       `json:"implant"`        //是否已种
	ImplantDate    time.Time `json:"implantDate"`    //植入日期
	Doctor         string    `json:"doctor"`         //医生
	Pack           int       `json:"pack"`           //1 普通 2 半口 3 全口
	PaybackDate    time.Time `json:"paybackDate"`    //预定回款日期
}
