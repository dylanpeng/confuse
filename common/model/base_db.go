package model

import (
	"confuse/common"
	"confuse/common/entity"
	"errors"
	"gorm.io/gorm"
)

type baseDBModel struct {
	readInstance  string
	writeInstance string
}

var ErrPrimaryAttrEmpty = errors.New("primary attribute is empty")

func createDBModel(readInstance, writeInstance string) *baseDBModel {
	return &baseDBModel{
		readInstance:  readInstance,
		writeInstance: writeInstance,
	}
}

func (d *baseDBModel) getReadDB() (db *gorm.DB, err error) {
	db, err = common.GetDb(d.readInstance)

	if err != nil {
		return
	}

	return
}

func (d *baseDBModel) getWriteDB() (db *gorm.DB, err error) {
	db, err = common.GetDb(d.writeInstance)

	if err != nil {
		return
	}

	return
}

func (d *baseDBModel) Add(entity entity.IEntity) (err error) {
	db, err := common.GetDb(d.writeInstance)

	if err != nil {
		return
	}

	err = db.Create(entity).Error
	return
}

func (d *baseDBModel) Update(entity entity.IEntity, params map[string]interface{}) (err error) {
	db, err := common.GetDb(d.writeInstance)

	if err != nil {
		return
	}

	if params == nil {
		err = db.Save(entity).Error
	} else {
		err = db.Model(entity).Updates(params).Error
	}

	return
}

func (d *baseDBModel) Get(entity entity.IEntity) (err error) {
	db, err := common.GetDb(d.writeInstance)

	if err != nil {
		return
	}

	err = db.First(entity).Error

	return
}
