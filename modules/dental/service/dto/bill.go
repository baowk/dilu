package dto

import (
	"time"

	"github.com/baowk/dilu-core/core/base"
	"github.com/shopspring/decimal"
)

type BillGetPageReq struct {
	base.ReqPage `search:"-"`
	TradeType    int       `json:"tradeType" form:"tradeType"` //交易类型 1 成交 2补尾款  3补上月欠款 10退款
	Begin        time.Time `json:"begin"`                      //开始成交日期
	End          time.Time `json:"end"`                        //结束成交日期
	UserId       int       `json:"userId"`                     //咨询师id
}

// Bill
type BillDto struct {
	Id             int             `json:"id"`             //主键
	No             string          `json:"no"`             //订单号
	CustomerId     int             `json:"customerId"`     //顾客
	CustomerName   string          `json:"customerName"`   //用户姓名
	UserId         int             `json:"userId"`         //用户id
	TeamId         int             `json:"teamId"`         //团队id
	DeptPath       string          `json:"deptPath"`       //路径
	Amount         decimal.Decimal `json:"amount"`         //金额
	RealAmount     decimal.Decimal `json:"realAmount"`     //折后金额
	PaidAmount     decimal.Decimal `json:"paidAmount"`     //已支付金额
	DebtAmount     decimal.Decimal `json:"debtAmount"`     //回收上月欠款
	RefundAmount   decimal.Decimal `json:"refundAmount"`   //退款
	LinkId         int             `json:"linkId"`         //关联订单
	TradeAt        time.Time       `json:"tradeAt"`        //交易日期
	TradeType      int             `json:"tradeType"`      //交易类型 1 成交 2补尾款  3补上月欠款 10退款
	DentalCount    int             `json:"dentalCount"`    //颗数
	Brand          int             `json:"brand"`          //品牌
	ImplantedCount int             `json:"implantedCount"` //已种颗数
	Implant        int             `json:"implant"`        //种植状态：1 未种 2部分 3已种
	ImplantDate    time.Time       `json:"implantDate"`    //植入日期
	Doctor         string          `json:"doctor" `        //医生
	Pack           int             `json:"pack" `          //1 普通 2 半口 3 全口
	PaybackDate    time.Time       `json:"paybackDate"`    //预定回款日期
	Tags           string          `json:"tags"`           //标签
	PrjName        string          `json:"prjName"`        //种植项目
	OtherPrj       string          `json:"otherPrj"`       //其他项目
	Remark         string          `json:"remark"`         //备注
	CreatedAt      time.Time       `json:"createdAt"`      //创建时间
	UpdatedAt      time.Time       `json:"updatedAt"`      //更新时间
	CreateBy       int             `json:"createBy"`       //创建者id
	UpdateBy       int             `json:"updateBy"`       //更新者id
}

type Option struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

type IdentifyBillDto struct {
	Id             int      `json:"id"`
	PrjName        string   `json:"prjName"`        //种植项目
	OtherPrj       string   `json:"otherPrj"`       //其他项目
	CustomerId     int      `json:"customerId"`     //顾客Id
	CustomerName   string   `json:"customerName"`   //用户姓名
	Customers      []Option `json:"customers"`      //顾客列表
	UserId         int      `json:"userId"`         //用户id
	Name           string   `json:"name"`           //用户名
	TeamId         int      `json:"teamId"`         //团队id
	Amount         string   `json:"amount"`         //金额
	RealAmount     string   `json:"realAmount"`     //折后金额
	PaidAmount     string   `json:"paidAmount"`     //已支付金额
	Debts          string   `json:"debts"`          //欠款
	LinkId         int      `json:"linkId"`         //关联订单
	TradeAt        string   `json:"tradeAt"`        //交易日期
	TradeType      int      `json:"tradeType"`      //交易类型 1 成交 2补尾款  3补上月欠款 10退款
	DentalCount    int      `json:"dentalCount"`    //颗数
	Brand          int      `json:"brand"`          //品牌
	BrandName      string   `json:"brandName"`      //品牌名
	ImplantedCount int      `json:"implantedCount"` //已种颗数
	Implant        int      `json:"implant"`        //是否已种
	ImplantDate    string   `json:"implantDate"`    //植入日期
	Doctor         string   `json:"doctor"`         //医生
	Pack           int      `json:"pack"`           //1 普通 2 半口 3 全口
	PaybackDate    string   `json:"paybackDate"`    //预定回款日期
	Extensions     string   `json:"extensions"`     //延期情况
	Tags           string   `json:"tags"`           //标签
	Remark         string   `json:"remark"`         //备注
	Inviter        int      `json:"inviter"`        //邀请人id
	InviterName    string   `json:"inviterName"`    //邀请人名
}

type LinkBillDto struct {
	RealAmount     string `json:"realAmount"`     //折后金额
	PaidAmount     string `json:"paidAmount"`     //已支付金额
	LinkId         int    `json:"linkId"`         //关联订单
	TradeAt        string `json:"tradeAt"`        //交易日期
	TradeType      int    `json:"tradeType"`      //交易类型 1 成交 2补尾款  3补上月欠款 10退款
	ImplantedCount int    `json:"implantedCount"` //已种颗数
	ImplantDate    string `json:"implantDate"`    //植入日期
	Remark         string `json:"remark"`         //备注
}

type BillTmplReq struct {
	Text   string `json:"text"`   //文本
	TeamId int    `json:"teamId"` //团队id
}

type StQueryReq struct {
	TeamId   int       `json:"teamId"`
	UserId   int       `json:"userId"`
	DeptPath string    `json:"deptPath"`
	Begin    time.Time `json:"begin"`
	End      time.Time `json:"end"`
}

type BillUserStDto struct {
	UserId           int             `json:"userId"`
	Name             string          `json:"name"`             //姓名
	Target           decimal.Decimal `json:"target"`           //目标
	TargetNew        int             `json:"targetNew"`        //留存任务
	TargetFirst      int             `json:"targetFirst"`      //出诊任务
	NewCustomerCnt   int             `json:"newCustomerCnt"`   //留存任务
	FirstDiagnosis   int             `json:"firstDiagnosis"`   //导诊任务
	DealCnt          int             `json:"dealCnt"`          //成交患者
	FurtherDiagnosis int             `json:"furtherDiagnosis"` //复诊患者
	Deal             decimal.Decimal `json:"deal"`             //成交金额
	Paid             decimal.Decimal `json:"paid"`             //实收
	Debt             decimal.Decimal `json:"debt"`             //补上月欠款
	CurDebt          decimal.Decimal `json:"curDebt"`          //当前欠款
	Refund           decimal.Decimal `json:"refund"`           //退款
	Total            decimal.Decimal `json:"total"`            //总实收
}

type DentalStDto struct {
	Total     int `json:"total"`     //总数
	Implanted int `json:"implanted"` //已种植
}
