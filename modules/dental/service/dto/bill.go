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

type IdentifyBillDto struct {
	PrjName        string `json:"prjName"`        //种植项目
	OtherPrj       string `json:"otherPrj"`       //其他项目
	CustomerId     int    `json:"customerId"`     //顾客
	CustomerName   string `json:"customerName"`   //顾客姓名
	UserId         int    `json:"userId"`         //用户id
	Name           string `json:"name"`           //用户名
	TeamId         int    `json:"teamId"`         //团队id
	Total          string `json:"total"`          //金额
	RealTotal      string `json:"realTotal"`      //折后金额
	PaidTotal      string `json:"paidTotal"`      //已支付金额
	Debts          string `json:"debts"`          //欠款
	LinkId         int    `json:"linkId"`         //关联订单
	TradeAt        string `json:"tradeAt"`        //交易日期
	TradeStatus    int    `json:"tradeStatus"`    //交易类型 1 成交 2补尾款  3补上月欠款 10退款
	DentalCount    int    `json:"dentalCount"`    //颗数
	Brand          int    `json:"brand"`          //品牌
	BrandName      string `json:"brandName"`      //品牌名
	ImplantedCount int    `json:"implantedCount"` //已种颗数
	Implant        int    `json:"implant"`        //是否已种
	ImplantDate    string `json:"implantDate"`    //植入日期
	Doctor         string `json:"doctor"`         //医生
	Pack           int    `json:"pack"`           //1 普通 2 半口 3 全口
	PaybackDate    string `json:"paybackDate"`    //预定回款日期
	Extensions     string `json:"extensions"`     //延期情况
	Tags           string `json:"tags"`           //标签
	Remark         string `json:"remark"`         //备注
	Inviter        int    `json:"inviter"`        //邀请人id
	InviterName    string `json:"inviterName"`    //邀请人名
}

type BillTmplReq struct {
	Text   string `json:"text"`   //文本
	TeamId int    `json:"teamId"` //团队id
}

type StDayReq struct {
	TeamId   int       `json:"teamId"`
	UserId   int       `json:"userId"`
	DeptPath string    `json:"deptPath"`
	Day      time.Time `json:"day"`
}
