package model

import (
	"confuse/common/entity"
)

var User = &userModel{
	DbBase: &DbBase{
		readDbName:  "slave",
		writeDbName: "master",
	},
}

type userModel struct {
	*DbBase
}

func (u *userModel) BatchInsertUsers() (err error) {
	db, err := u.getWriteDB()

	if err != nil {
		return
	}

	dataUser := &entity.DataUser{
		Name:       "test1",
		CreateTime: 111,
		UpdateTime: 222,
	}

	dataUser2 := &entity.DataUser{
		Name:       "test2",
		CreateTime: 333,
		UpdateTime: 444,
	}

	users := []*entity.DataUser{dataUser, dataUser2}

	//err = db.Create(users).Error
	err = db.CreateInBatches(users, 1000).Error

	return
}

func (u *userModel) AddByMap() (err error) {
	db, err := u.getWriteDB()

	if err != nil {
		return
	}

	userMap := map[string]interface{}{
		"name":        "map",
		"create_time": 200,
	}

	err = db.Model(&entity.DataUser{}).Create(userMap).Error

	return
}

//func (u *userModel) AddWithAssociation() (err error) {
//	db, err := u.getWriteDB()
//
//	if err != nil {
//		return
//	}
//
//	dataUser := &entity.DataUser{
//		Name:       "test1",
//		CreateTime: 111,
//		UpdateTime: 222,
//		Ex: &entity.DataUserExtend{
//			Remark: "remark",
//		},
//	}
//
//	err = db.Create(dataUser).Error
//	return
//}

func (u *userModel) QueryByName(name string) (list []*entity.DataUser, err error) {
	db, err := u.getWriteDB()

	if err != nil {
		return
	}

	list = make([]*entity.DataUser, 0, 8)

	err = db.Where("name = ?", name).Order("create_time desc, id").Find(&list).Error
	//err = db.Find(&list, "name = ? and id = ?", name, 19).Error

	return
}

func (u *userModel) QueryWithRowsScan(name string) (err error) {
	db, err := u.getWriteDB()

	if err != nil {
		return
	}

	rows, err := db.Model(&entity.DataUser{}).Where("name = ?", name).Rows()
	if err != nil {
		return
	}

	for rows.Next() {
		var id, createTime, updateTime int64
		var n string
		user := &entity.DataUser{}

		rows.Scan(&id, &n, &createTime, &updateTime)
		db.ScanRows(rows, user)

		continue
	}

	return
}
