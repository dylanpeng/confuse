package model

import (
	"confuse/common/entity"
	"gorm.io/gorm"
	"time"
)

var User = &userModel{
	baseModel: createModel(
		"main-slave",
		"main-master",
		"main-slave",
		"main-master",
		"user",
		time.Minute*10,
	),
}

type userModel struct {
	*baseModel
}

func (u *userModel) BatchInsertUsers() (err error) {
	db, err := u.DB.getDB(true)

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
	db, err := u.DB.getDB(true)

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
//	db, err := u.DB.getDB()
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
	db, err := u.DB.getDB(false)

	if err != nil {
		return
	}

	list = make([]*entity.DataUser, 0, 8)

	err = db.Where("name = ?", name).Order("create_time desc, id").Find(&list).Error
	//err = db.Find(&list, "name = ? and id = ?", name, 19).Error

	return
}

func (u *userModel) QueryWithRowsScan(name string) (err error) {
	db, err := u.DB.getDB(false)

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

func (u *userModel) PageQuery(page int, pageSize int) (list []*entity.DataUserPart, err error) {
	db, err := u.DB.getDB(false)

	if err != nil {
		return
	}

	list = make([]*entity.DataUserPart, 0, 8)

	err = db.Session(&gorm.Session{QueryFields: true}).Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&list).Error
	//err = db.Find(&list, "name = ? and id = ?", name, 19).Error

	//db.Clauses()
	return
}

func (u *userModel) ConditionName(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name = ?", name)
	}
}

func (u *userModel) ConditionId(db *gorm.DB) *gorm.DB {
	return db.Where("id = ?", 13)
}

func (u *userModel) QueryWithScope() (list []*entity.DataUser, err error) {
	db, err := u.DB.getDB(false)

	if err != nil {
		return
	}

	list = make([]*entity.DataUser, 0, 8)

	err = db.Scopes(u.ConditionName("test1"), u.ConditionId).Find(&list).Error
	//err = db.Find(&list, "name = ? and id = ?", name, 19).Error

	return
}

func (u *userModel) QueryUserCount() (count int64, err error) {
	db, err := u.DB.getDB(false)

	if err != nil {
		return
	}

	err = db.Model(&entity.DataUser{}).Where("name = ?", "test1").Count(&count).Error
	return
}

func (u *userModel) QueryByRaw() (list []*entity.DataUser, err error) {
	db, err := u.DB.getDB(false)

	if err != nil {
		return
	}

	list = make([]*entity.DataUser, 0, 8)
	err = db.Raw("select * from data_user where name = ?", "test").Scan(&list).Error
	return
}

func (u *userModel) Trans() (err error) {
	db, err := u.DB.getDB(true)
	if err != nil {
		return
	}

	user := &entity.DataUser{
		Id: 1,
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}

		if err != nil {
			tx.Rollback()
		}
	}()

	err = tx.First(user).Error
	if err != nil {
		return
	}

	user.Name = "tans1"

	err = tx.Save(user).Error
	if err != nil {
		return
	}

	err = tx.Commit().Error

	return
}
