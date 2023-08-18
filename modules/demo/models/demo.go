package models

type Demo struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Name string `json:"name" gorm:"type:varchar(128);comment:名字"`
}

func (Demo) TableName() string {
	return "demo"
}
