package entity

import (
	"fmt"
	"gorm.io/gorm"
)

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

func (e *DataUser) PrimarySeted() bool {
	return e.Id > 0
}

func (e *DataUser) String() string {
	return fmt.Sprintf("%+v", *e)
}

type DataUserPart struct {
	Id   int64 `gorm:"primaryKey"`
	Name string
}

func (*DataUserPart) TableName() string {
	return "data_user"
}

func (e *DataUserPart) PrimarySeted() bool {
	return e.Id > 0
}

func (e *DataUserPart) String() string {
	return fmt.Sprintf("%+v", *e)
}

func (u *DataUserPart) AfterFind(db *gorm.DB) (err error) {
	if u != nil {
		fmt.Printf("query success. DataUserPart: %+v", u)
	}
	return
}

//type DataUserExtend struct {
//	Id     int64  `gorm:"primaryKey"`
//	Remark string `gorm:"column:remark" json:"remark"`
//}
//
//func (*DataUserExtend) TableName() string {
//	return "data_user_extend"
//}
