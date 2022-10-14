package main

import (
	"confuse/common/entity"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "dev:123!@#qweASD@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn2 := "dev:123!@#qweASD@tcp(127.0.0.1:3306)/test2?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	sqlDB, err := sql.Open("mysql", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	sqlDB2, err := sql.Open("mysql", dsn2)
	db2, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB2,
	}), &gorm.Config{})

	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DriverName: "my_mysql_driver",
	//	DSN:        dsn, // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	//}), &gorm.Config{})

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

	err = db.Create(dataUser).Error

	if err != nil {
		fmt.Printf("Create err:%s\n", err)
		return
	}

	dataUser2 := &entity.DataUser{
		Name:       "test2",
		CreateTime: 1000,
		UpdateTime: 2000,
	}

	err = db2.Create(dataUser2).Error

	if err != nil {
		fmt.Printf("Create2 err:%s\n", err)
		return
	}

	fmt.Printf("%+v \n", dataUser)
}
