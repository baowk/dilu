package models

//客户
type Customer struct {
	Id       int    `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Name     string `json:"name" gorm:"type:varchar(32);comment:姓名"`                         //姓名
	Birthday string `json:"birthday" gorm:"type:date;comment:生日"`                            //生日
	Phone    string `json:"phone" gorm:"type:varchar(11);comment:手机号"`                       //手机号
	Wechat   string `json:"wechat" gorm:"type:varchar(64);comment:微信号"`                      //微信号
	Gender   int    `json:"gender" gorm:"type:tinyint;comment:性别"`                           //性别
	Address  string `json:"address" gorm:"type:varchar(255);comment:地址"`                     //地址
	Remark   string `json:"remark" gorm:"type:varchar(255);comment:描述"`                      //描述
	SalesId  int    `json:"salesId" gorm:"type:int;comment:销售人员"`                            //销售人员
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
func (e *Customer) SetBirthday(birthday string) *Customer {
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
func (e *Customer) SetSalesId(salesId int) *Customer {
	e.SalesId = salesId
	return e
}
