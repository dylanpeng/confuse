package main

import (
	"confuse/api/config"
	"confuse/api/router"
	"confuse/common"
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var (
	configFile = flag.String("c", "config.toml", "config file path")
)

func main() {
	// parse flag
	flag.Parse()

	// set max cpu core
	runtime.GOMAXPROCS(runtime.NumCPU())

	// parse config file
	if err := config.Init(*configFile); err != nil {
		log.Fatalf("Fatal Error: can't parse config file!!!\n%s", err)
	}

	// init log
	if err := common.InitLogger(); err != nil {
		log.Fatalf("Fatal Error: can't initialize logger!!!\n%s", err)
	}

	defer func() {
		_ = common.Logger.Sync()
		_ = common.Logger.Close()
	}()

	// init cache clients
	common.InitCache()

	// init db
	if err := common.InitDB(); err != nil {
		log.Fatalf("Fatal Error: can't initialize db clients!!!\n%s", err)
	}

	// start http server
	common.InitHttpServer(router.Router)
	common.Logger.Infof("http server start at <%s>", config.GetConfig().Server.Http.GetAddr())

	// waitting for exit signal
	exit := make(chan os.Signal, 1)
	stopSigs := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGABRT,
		syscall.SIGKILL,
		syscall.SIGTERM,
	}
	signal.Notify(exit, stopSigs...)

	// catch exit signal
	sign := <-exit
	common.Logger.Infof("stop by exit signal '%s'", sign)

	// stop http server
	common.HttpServer.Stop()
	common.Logger.Info("http server stoped")

	//dataUser := &entity.DataUser{
	//	Name:       "test",
	//	CreateTime: 10009,
	//	UpdateTime: 20009,
	//}
	//
	//err := model.User.Add(dataUser)
	//if err != nil {
	//	fmt.Printf("Create err:%s\n", err)
	//	return
	//}
	//
	//dataUser2 := &entity.DataUser{
	//	Id: 6,
	//}
	//
	//_, err = model.User.Get(dataUser2)
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
