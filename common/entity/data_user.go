package entity

type DataUser struct {
	Id         int64  `gorm:"primaryKey"`
	Name       string `gorm:"column:name"`
	CreateTime int64
	UpdateTime int64
}

func (*DataUser) TableName() string {
	return "data_user"
}
