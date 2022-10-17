package model

import "confuse/common/entity"

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
