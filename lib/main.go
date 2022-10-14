package main

import (
	"confuse/common/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "dev:123!@#qweASD@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("err:%s\n", err)
		return
	}

	//dataUser := &DataUser{Id: 1}
	//db.First(dataUser)

	dataUser := &entity.DataUser{
		Name:       "test",
		CreateTime: 1000,
		UpdateTime: 2000,
	}

	db.Create(dataUser)

	fmt.Printf("%+v \n", dataUser)
}
