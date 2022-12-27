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

func (e *DataUser) TableName() string {
	return "data_user"
}

func (e *DataUser) PrimaryPairs() []interface{} {
	return []interface{}{"id", e.Id}
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

func (e *DataUserPart) TableName() string {
	return "data_user"
}

func (e *DataUserPart) PrimaryPairs() []interface{} {
	return []interface{}{"id", e.Id}
}

func (e *DataUserPart) PrimarySeted() bool {
	return e.Id > 0
}

func (e *DataUserPart) String() string {
	return fmt.Sprintf("%+v", *e)
}

func (e *DataUserPart) AfterFind(db *gorm.DB) (err error) {
	if e != nil {
		fmt.Printf("query success. DataUserPart: %+v", e)
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
