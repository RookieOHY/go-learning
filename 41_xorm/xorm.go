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
	"xorm.io/xorm/names"
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

type AppUser struct {
	ID        int64      `xorm:"varchar(25) notnull  pk unique  comment('id')"`
	UserName  string     `xorm:"varchar(25) notnull  comment('姓名')"`
	UserAge   int8       `xorm:"int(11) null comment('年龄')"`
	UserBirth *time.Time `xorm:"datetime null comment('生日')"`
	Email     string     `xorm:"varchar(32) null comment('邮箱')"`
	//结构体变动
	Address string `xorm:"varchar(32) null comment('住址')"`
}

func SnakeMapper() {
	engine := GetEngine()

	//驼峰映射 (默认作用于：表名字+字段名字)
	engine.SetMapper(names.NewPrefixMapper(names.SnakeMapper{}, "rk_"))

	//这是数据库连接时的字符集 而不是 建表字符集 (仅mysql使用)
	//设置存储引擎(仅mysql使用)
	err := engine.StoreEngine("MyISAM").Charset("utf8mb4").CreateTable(&AppUser{})

	//值得一提的是：当结构体新增属性的时候 重新执行建表操作 并不会更新表结构 （CREATE TABLE IF NOT EXISTS ......）
	engine.Sync2(&AppUser{})

	if err != nil {
		panic(err)
	}
	//判断表是否为空（无任何记录）
	empty, err := engine.IsTableEmpty(&AppUser{})
	if err != nil {
		panic(err)
	}
	fmt.Println("不存在记录", empty)

	//判断表是否存在
	exist, err := engine.IsTableExist(&AppUser{})
	if err != nil {
		panic(err)
	}
	fmt.Println("存在表", exist)

	//删除表
	//err = engine.DropTables(&AppUser{})
	//if err != nil {
	//	panic(err)
	//}

	//获取数据库信息
	metas, err := engine.DBMetas()
	for _, meta := range metas {
		name := meta.Name
		fmt.Println(name)
	}

	//提取实际的表结构
	infos, err := engine.TableInfo(&AppUser{})
	fmt.Println(infos.Name)
}

func SameMapper() {
	engine := GetEngine()
	//表、字段名字 和 结构体声明
	engine.SetMapper(names.SameMapper{})
}

func GonicMapper() {

}
