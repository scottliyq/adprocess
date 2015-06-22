package main

import (
	"adexchange/lib"
	m "adexchange/models"
	"adexchange/tools"
	//"adprocess/tasks"
	"adprocess/handlers"
	//adpm "adprocess/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//func main() {
//	flag.Parse()
//	pool = newPool(*redisServer, *redisPassword)

//	c := pool.Get()
//	for {
//		rstMap, _ := redis.StringMap(c.Do("brpop", "ADMUX_IMP", "0"))
//		fmt.Printf(rstMap["ADMUX_IMP"])

//	}

//	defer c.Close()

//}

//func test() {
//	b, err := msgpack.Marshal("test")
//	beego.Debug(b)
//	if err == nil {
//		c := pool.Get()
//		c.Do("lpush", "ADMUX_IMP", b)

//		defer c.Close()
//	}
//}

func main() {

	beego.SetLogger("file", `{"filename":"logs/adprocess.log"}`)
	beego.SetLogFuncCall(true)
	orm.Debug, _ = beego.AppConfig.Bool("orm_debug")

	lib.Pool = lib.NewPool(beego.AppConfig.String("redis_server"), "")
	tools.Init("ip.dat")
	m.Connect()

	//tasks.DailyReportInit(30)
	go handlers.HandleImp()
}

func test() {
	beego.Debug("test")

}
