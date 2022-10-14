package entity

type DataUser struct {
	Id         int64 `gorm:"primaryKey"`
	Name       string
	CreateTime int64
	UpdateTime int64
}

func (*DataUser) TableName() string {
	return "data_user"
}
