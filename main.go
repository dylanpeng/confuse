package main

import (
	"confuse/common"
	"confuse/common/config"
	"fmt"
)

func main() {
	conf := config.Config{}

	conf.DBConfigs = make([]*config.DBConfig, 0, 8)

	conf.DBConfigs = append(conf.DBConfigs, &config.DBConfig{
		Name:         "master",
		UserName:     "dev",
		Password:     "123!@#qweASD",
		SourceUrl:    "127.0.0.1",
		Port:         "3306",
		DataBaseName: "test",
	})

	conf.DBConfigs = append(conf.DBConfigs, &config.DBConfig{
		Name:         "slave",
		UserName:     "dev",
		Password:     "123!@#qweASD",
		SourceUrl:    "127.0.0.1",
		Port:         "3306",
		DataBaseName: "test2",
	})

	err := common.InitDB(conf.DBConfigs)
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
	//dataUser2 := &entity.DataUser{
	//	Id: 6,
	//}
	//
	//err = model.User.Get(dataUser2)
	//if err != nil {
	//	fmt.Printf("Get err:%s\n", err)
	//	return
	//}

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
}
