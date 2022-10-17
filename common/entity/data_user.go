package entity

type DataUser struct {
	Id         int64  `gorm:"primaryKey"`
	Name       string `gorm:"column:name"`
	CreateTime int64
	UpdateTime int64
	//Ex         *DataUserExtend `gorm:"foreignKey:Id;references:Id;"`
}

func (*DataUser) TableName() string {
	return "data_user"
}

//type DataUserExtend struct {
//	Id     int64  `gorm:"primaryKey"`
//	Remark string `gorm:"column:remark" json:"remark"`
//}
//
//func (*DataUserExtend) TableName() string {
//	return "data_user_extend"
//}
