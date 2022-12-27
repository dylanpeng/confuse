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

func (m *baseDBModel) getDB(write bool) (db *gorm.DB, err error) {
	if write {
		return common.GetDB(m.writeInstance)
	}

	return common.GetDB(m.readInstance)
}

func (m *baseDBModel) Add(e entity.IEntity) (err error) {
	db, err := common.GetDB(m.writeInstance)

	if err != nil {
		return
	}

	err = db.Create(e).Error
	return
}

func (m *baseDBModel) Update(e entity.IEntity, params map[string]interface{}) (err error) {
	if !e.PrimarySeted() {
		return ErrPrimaryAttrEmpty
	}

	db, err := common.GetDB(m.writeInstance)

	if err != nil {
		return
	}

	if params == nil {
		err = db.Save(e).Error
	} else {
		err = db.Model(e).Updates(params).Error
	}

	return
}

func (m *baseDBModel) Get(e entity.IEntity) (exist bool, err error) {
	if !e.PrimarySeted() {
		err = ErrPrimaryAttrEmpty
		return
	}

	db, err := m.getDB(false)

	if err != nil {
		return false, err
	}

	err = db.First(e).Error

	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
