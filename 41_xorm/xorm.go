package _1_xorm

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/core"
)

func CreateEngine() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@/02-gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("创建单引擎异常", err)
		return nil
	}
	log.Println("创建单引擎成功")
	return engine
}

func Ping() {
	engine := CreateEngine()
	err := engine.Ping()
	if err != nil {
		fmt.Println("数据库连接正常")
	}
}

func PingContext() {
	engine := CreateEngine()
	//创建根ctx (empty ctx)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	status := "DB Connection is UP"
	err := engine.PingContext(ctx)
	if err != nil {
		fmt.Println("[PingContext Method] DONE!")
		status = "DB Connection is DOWN"
	}
	log.Println(status)
}

func PingTimer() {
	engine := CreateEngine()
	//engine.Close()
	// 新建channel
	ch := time.Tick(10 * time.Second)
	for {
		select {
		case <-ch:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			if err := engine.PingContext(ctx); err != nil {
				fmt.Println("DB Connection is DOWN")
			}
		}
	}
}

type EngineX struct {
	*xorm.Engine
}

var instance *EngineX
var once sync.Once

func newEngine() {
	var err error
	instance = &EngineX{}
	instance.Engine, err = xorm.NewEngine("mysql", "root:123456@/02-gorm?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
func GetEngine() *EngineX {
	once.Do(newEngine)
	return instance
}

func NewEngineWithParams() *xorm.Engine {
	// 参数 驱动名字 数据源 额外参数（map）
	engine, err := xorm.NewEngineWithParams("mysql", "root:123456@/02-gorm?charset=utf8", map[string]string{"loc": "local"})
	if err != nil {
		return nil
	}
	return engine
}

func NewEngineWithDB() *xorm.Engine {
	//已存在的数据库连接
	db, err := sql.Open("mysql", "root:123456@/02-gorm?charset=utf8")
	if err != nil {
		panic(err)
	}
	// 值得一提的是：coreDB引用了 *sql.DB
	coreDB := &core.DB{
		DB: db,
	}
	engine, err := xorm.NewEngineWithDB("mysql", "root:123456@/02-gorm?charset=utf8", coreDB)
	if err != nil {
		panic(err)
	}
	return engine
}
