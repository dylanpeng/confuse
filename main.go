package main

import (
	"confuse/common"
	"confuse/common/config"
	"confuse/common/entity"
	"confuse/common/model"
	"confuse/lib/gorm"
	"fmt"
)

func main() {
	conf := config.Config{}

	conf.DB = make(map[string]*gorm.Config)

	conf.DB["master"] = &gorm.Config{
		Host:         "127.0.0.1",
		Port:         3306,
		User:         "dev",
		Password:     "123!@#qweASD",
		Charset:      "utf8mb4",
		Database:     "test",
		Timeout:      3,
		MaxIdleConns: 20,
		MaxConnTtl:   300,
	}

	conf.DB["slave"] = &gorm.Config{
		Host:         "127.0.0.1",
		Port:         3306,
		User:         "dev",
		Password:     "123!@#qweASD",
		Charset:      "utf8mb4",
		Database:     "test",
		Timeout:      3,
		MaxIdleConns: 20,
		MaxConnTtl:   300,
	}

	_ = conf.Init()

	err := common.InitDB()
	if err != nil {
		fmt.Printf("Init Db failed. err: %s", err)
	}

	//dataUser := &entity.DataUser{
	//	Name:       "test",
	//	CreateTime: 10009,
	//	UpdateTime: 20009,
	//}
	//
	//err = model.User.Add(dataUser)
	//if err != nil {
	//	fmt.Printf("Create err:%s\n", err)
	//	return
	//}
	//
	dataUser2 := &entity.DataUser{
		Id: 6,
	}

	err = model.User.Get(dataUser2)
	if err != nil {
		fmt.Printf("Get err:%s\n", err)
		return
	}

	//err = model.User.BatchInsertUsers()
	//if err != nil {
	//	fmt.Printf("BatchInsertUsers Db failed. err: %s", err)
	//}

	//err = model.User.AddByMap()
	//if err != nil {
	//	fmt.Printf("AddByMap Db failed. err: %s", err)
	//}

	//err = model.User.AddWithAssociation()
	//if err != nil {
	//	fmt.Printf("AddWithAssociation Db failed. err: %s", err)
	//}

	//list, err := model.User.QueryByName("test1")
	//if err != nil {
	//	fmt.Printf("QueryByName Db failed. err: %s", err)
	//	return
	//}
	//fmt.Printf("list: %+v", list)

	//_, err = model.User.PageQuery(2, 2)
	//if err != nil {
	//	fmt.Printf("QueryByName Db failed. err: %s", err)
	//	return
	//}

	//_, err = model.User.QueryWithScope()
	//if err != nil {
	//	fmt.Printf("QueryWithScope Db failed. err: %s", err)
	//	return
	//}

	//_, err = model.User.QueryUserCount()
	//if err != nil {
	//	fmt.Printf("QueryUserCount Db failed. err: %s", err)
	//	return
	//}

	//user := &entity.DataUser{
	//	Id: 1,
	//}
	//model.User.Get(user)
	//params := map[string]interface{}{}
	//params["name"] = "aaa"
	//model.User.Update(user, params)

	//_, err = model.User.QueryByRaw()
	//if err != nil {
	//	fmt.Printf("QueryByRaw Db failed. err: %s", err)
	//	return
	//}

	//err = model.User.Trans()
	//if err != nil {
	//	fmt.Printf("Trans Db failed. err: %s", err)
	//	return
	//}

	return
}
