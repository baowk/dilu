package models

import "time"

//Customer
type Customer struct {
	Id          int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Name        string    `json:"name" gorm:"type:varchar(32);comment:姓名"`                         //姓名
	PY          string    `json:"py" gorm:"type:varchar(32);comment:姓名拼音"`                         //姓名拼音
	Phone       string    `json:"phone" gorm:"type:varchar(11);comment:手机号"`                       //手机号
	Wechat      string    `json:"wechat" gorm:"type:varchar(64);comment:微信号"`                      //微信号
	Gender      int       `json:"gender" gorm:"type:tinyint;comment:性别"`                           //性别
	Age         int       `json:"age" gorm:"type:tinyint unsigned;comment:年龄"`                     //年龄
	Birthday    int       `json:"birthday" gorm:"type:int;comment:生日"`                             //生日
	Source      string    `json:"source" gorm:"type:varchar(64);comment:来源"`                       //来源
	Address     string    `json:"address" gorm:"type:varchar(255);comment:地址"`                     //地址
	Remark      string    `json:"remark" gorm:"type:varchar(255);comment:描述"`                      //描述
	UserId      int       `json:"userId" gorm:"type:int unsigned;index;comment:用户id"`              //用户id
	TeamId      int       `json:"teamId" gorm:"type:int unsigned;index;comment:团队id"`              //团队id
	DeptPath    string    `json:"deptPath" gorm:"type:varchar(255);comment:部门路径"`                  //路径
	Inviter     int       `json:"inviter" gorm:"type:int unsigned;comment:邀请人"`                    //邀请人
	InviterName string    `json:"inviterName" gorm:"type:varchar(32);comment:邀请人名"`                //邀请人名
	CreatedAt   time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                     //创建时间
	UpdatedAt   time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                     //更新时间
	CreateBy    int       `json:"createBy" gorm:"type:int unsigned;index;comment:创建者"`             //创建者id
	UpdateBy    int       `json:"updateBy" gorm:"type:int unsigned;index;comment:更新者"`             //更新者id
}

func (Customer) TableName() string {
	return "customer"
}

func NewCustomer() *Customer {
	return &Customer{}
}

func (e *Customer) SetId(id int) *Customer {
	e.Id = id
	return e
}
func (e *Customer) SetName(name string) *Customer {
	e.Name = name
	return e
}
func (e *Customer) SetBirthday(birthday int) *Customer {
	e.Birthday = birthday
	return e
}
func (e *Customer) SetPhone(phone string) *Customer {
	e.Phone = phone
	return e
}
func (e *Customer) SetWechat(wechat string) *Customer {
	e.Wechat = wechat
	return e
}
func (e *Customer) SetGender(gender int) *Customer {
	e.Gender = gender
	return e
}
func (e *Customer) SetAddress(address string) *Customer {
	e.Address = address
	return e
}
func (e *Customer) SetRemark(remark string) *Customer {
	e.Remark = remark
	return e
}
