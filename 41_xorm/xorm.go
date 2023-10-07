package _1_xorm

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
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
	// 已存在的数据库连接
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


// 当结构体是自定义类型 默认会被存储为Text 且使用json序列化和反序列化
type RKString string
type RKMap map[RKString]RKString
type RKSlicp []RKString
type Extends struct{
	CityName string
	ProvinceName string
}
type Additional struct{
	FatherName string
	MothoreName string
}

// 当属性是 *time.Time 要求在xorm里面手动声明表类型 datetime（不然无法映射）
// 字段的类型和容纳大小推荐自己设置，否则就是xorm的默认映射类型（可能不合适）
type AppUsers struct{
	ID int64 `xorm:"'id'  notnull  pk unique autoincr  comment('id')"`
	UserName string `xorm:"'username' null varchar(32) comment('用户名')"`
	Age int8 `xorm:" 'age' null comment('年龄')"`
	Email string `xorm:" 'email' null comment('邮箱')"`
	CreateTime *time.Time `xorm:"datetime created null comment('创建时间')"`
	UpdateTime *time.Time `xorm:"datetime updated deleted null comment('更新时间')"`
	IsDeleted bool `xorm:"  default(0) comment('是否删除')"`
	NonMaping string `xorm:"-"` // 不映射的属性
	OnlyWrite string `xorm:"->"`
	OnlyRead string `xorm:"<-"`
	Data Additional `xorm:"data" json:"data"` //写入json
	Extends Extends `xorm:"extends"` //将这个接口的属性映射为字段
	RKString RKString
	RKMap RKMap
	RKSlicp RKSlicp
}

type AppUser struct {
	// xorm:"'column_name'" 可以设置被映射的属性对应的数据表字段名字
	ID        int64      `xorm:"'id' varchar(25) notnull  pk unique  comment('id')"`
	UserName  string     `xorm:"varchar(25) notnull  comment('姓名')"`
	UserAge   int8       `xorm:"int(11) null comment('年龄')"`
	UserBirth *time.Time `xorm:"datetime null comment('生日')"`
	Email     string     `xorm:"varchar(32) null comment('邮箱')"`
	// 结构体变动
	Address string `xorm:"varchar(32) null comment('住址')"`
}

func SnakeMapper() {
	engine := GetEngine()

	// 关于前缀（作用于表 和 字段名字）和后缀（作用于表 和 字段名字）、不包括id
	// 会创建2张表，互不影响
	engine.SetMapper(names.NewPrefixMapper(names.SnakeMapper{}, "rk_"))
	engine.SetMapper(names.NewSuffixMapper(names.SnakeMapper{},"_end"))

	// 这是数据库连接时的字符集 而不是 建表字符集 (仅mysql使用)
	// 设置存储引擎(仅mysql使用)
	err := engine.StoreEngine("MyISAM").Charset("utf8mb4").CreateTable(&AppUser{})

	// 值得一提的是：当结构体新增属性的时候 重新执行建表操作 并不会更新表结构 （CREATE TABLE IF NOT EXISTS ......）
	engine.Sync2(&AppUser{})

	if err != nil {
		panic(err)
	}
	// 判断表是否为空（无任何记录）
	empty, err := engine.IsTableEmpty(&AppUser{})
	if err != nil {
		panic(err)
	}
	fmt.Println("不存在记录", empty)

	// 判断表是否存在
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

	// 获取数据库信息
	metas, _ := engine.DBMetas()
	for _, meta := range metas {
		name := meta.Name
		fmt.Println(name)
	}

	// 提取实际的表结构
	infos, _ := engine.TableInfo(&AppUser{})
	fmt.Println(infos.Name)
}

func SameMapper() {
	engine := GetEngine()
	// 表、字段名字 和 结构体声明
	engine.SetMapper(names.SameMapper{})
}

func GonicMapper() {
	engine := GetEngine()
	engine.SetMapper(names.GonicMapper{})
}

func(appUser *AppUser) TableName() string{
	return "your_app_user"
}

func TableName(){
	// doc : 结构体拥有 TableName() string 的成员方法，那么此方法的返回值即是该结构体对应的数据库表名
	engine := GetEngine()
	err := engine.CreateTables(&AppUser{})
	// engine.Table("your_app_user_02")
	if err != nil{
		panic(err)
	}
	fmt.Println("create table success!")

	user := &AppUser{ID:2,UserName:"RookieOHY",UserAge:25}

	_ , err = engine.InsertOne(user)
	if err != nil{
		panic(err)
	}
	//result 看起来id从0开始的
	fmt.Println("insert row success!",user.ID)
}

// Column tag example

func ColumnTag(){
	ex := GetEngine()
	err := ex.CreateTables(&AppUsers{})
	if err != nil{
		panic(err)
	}
	fmt.Println("[测试行标签-创建表]成功")

	// 如果表已经存在 不再创建 此时更新
	err = ex.Sync2(&AppUsers{})
	if err != nil{
		panic(err)
	}
}

func saveToFile(buffer io.Writer,fileName string){
	file, err := os.Create(fileName)
	if err!=nil{
		panic(err)
		return
	}
	defer file.Close()
}

func Dump(){
	// 直接dump文件
	ex := GetEngine()
	err := ex.DumpAllToFile("db.sql")
	if err!=nil{
		panic(err)
	}
	fmt.Println("dump db to file success!")
	// dump到buffer
	var buffer bytes.Buffer
	err = ex.DumpAll(&buffer)
	if err!=nil{
		panic(err)
	}
	// buffer to file 
	saveToFile(&buffer,"dump.sql")
	fmt.Println("dump db to buffer success!")
}

func Import(){
	GetEngine()
	//ex := GetEngine()
}