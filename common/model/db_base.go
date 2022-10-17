package model

import (
	"confuse/common"
	"confuse/common/entity"
	"gorm.io/gorm"
)

type DbBase struct {
	readDbName  string
	writeDbName string
}

func (d *DbBase) getReadDB() (db *gorm.DB, err error) {
	db, err = common.GetDb(d.readDbName)

	if err != nil {
		return
	}

	return
}

func (d *DbBase) getWriteDB() (db *gorm.DB, err error) {
	db, err = common.GetDb(d.writeDbName)

	if err != nil {
		return
	}

	return
}

func (d *DbBase) Add(entity entity.IEntity) (err error) {
	db, err := common.GetDb(d.writeDbName)

	if err != nil {
		return
	}

	err = db.Create(entity).Error
	return
}

func (d *DbBase) Update() (err error) {
	return
}

func (d *DbBase) Get(entity entity.IEntity) (err error) {
	db, err := common.GetDb(d.readDbName)

	if err != nil {
		return
	}

	err = db.First(entity).Error

	return
}
