package common

import (
	"confuse/common/config"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DbMaps map[string]*gorm.DB

func InitDB(dbConfigs []*config.DBConfig) (err error) {
	DbMaps = make(map[string]*gorm.DB, 4)

	for _, c := range dbConfigs {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.UserName, c.Password, c.SourceUrl, c.Port, c.DataBaseName)

		db, e := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if e != nil {
			err = e
			return
		}

		DbMaps[c.Name] = db
	}

	return
}

func GetDb(name string) (db *gorm.DB, err error) {
	db, exist := DbMaps[name]

	if !exist {
		err = errors.New("db not exists")
	}

	return
}
