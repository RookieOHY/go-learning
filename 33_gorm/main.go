package main

/*
	Gorm v2 学习
		基本使用：
			创建表、属性、索引、约束等。基本的数据库配置。
		模型的操作：
			模型概念：
				落地gorm标准的结构体或者接口
			标签：
				使用标签实现对表结构的定义以及约束
				用法：
					`gorm:"key:value;key:value"`
			注意点：
				使用自动的自动迁移创建表结构，多次运行时，不会更新表结构。（需要重新删除原有的表结构）
				有些结构体的属性为指针类型，使得在插入数据时，允许为空（如：日期类型，很多时候都不需要我们自己指定，如果不是指针类型，插入时会出现0000-00-00 00:00:00）
		CURD:
			Select("structFieldName"): 仅插入表某一个字段的值
			Create(对象的指针):插入数据
			Omit("structFieldName"): 忽略这个字段的插入


*/

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type UserInfo struct {
	Name string
	Age  uint8
}

type EmbedModel struct {
	UUID  uint       `gorm:"primaryKey;type:int(11) auto_increment;not null;unique;comment:id"`
	Time1 *time.Time `gorm:"type:datetime;column:create_time;comment:创建时间"`
	Time2 *time.Time `gorm:"type:datetime;column:update_time;comment:更新时间"`
}

type Model struct {
	EmbedModel EmbedModel `gorm:"embedded;embeddedPrefix:rookie_"`
	Name       string     `gorm:"type:varchar(255);comment:姓名"`
	Age        uint8      `gorm:"type:int(11);comment:年龄"`
	Birthday   *time.Time `gorm:"type:datetime;comment:生日"`
	Email      string     `gorm:"type:varchar(32);comment:邮箱地址"`
}

type User struct {
	gorm.Model
	Name string
	Age  int
	//add
	Sex int
}

func initConnection() (db *gorm.DB, err error) {
	d, e := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/02-gorm?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction: false, //跳过gorm默认带上事物的设置
		NamingStrategy: schema.NamingStrategy{ //表、行的命名策略
			TablePrefix:   "t_",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	return d, e
}

/* CRUD 增删改查*/
func Delete() {
	db, _ := initConnection()
	//实体类主键删除
	//model := Model{EmbedModel: EmbedModel{UUID: 13}}
	//dbptr := db.Delete(&model)

	//按照整形主键删除
	//dbptr:=db.Delete(&Model{},16)

	//批量删除（条件是非主属性，如果匹配，将会对所有匹配的记录做删除）
	//dbptr:=db.Where("name like ?","%o%").Delete(&Model{})
	dbptr := db.Delete(&Model{}, "name like ?", "%o%")
	if dbptr.Error == nil {
		fmt.Println("删除成功")
	} else {
		fmt.Println("发生的错误为", dbptr.Error)
	}
}

func Update() {
	db, err := initConnection()
	if err == nil {
		//var mds []Model
		var md Model
		//update 仅更新选择的字段
		//dbptr:=db.Model(&Model{}).Where("rookie_uuid = ?",14).Update("name","RookieOHY2")

		//save 无条件更新（默认更具主键来执行更新）
		//db.Model(&Model{}).Where("name like ?","%RookieOHY%").Find(&mds)
		//for k, _ := range mds {
		//	mds[k].Age = 99
		//}
		//dbptr:=db.Save(&mds)

		//updates (指定map更新、指定结构体更新（0值不参与更新）)
		//dbptr:=db.First(&md).Updates(&Model{Name:"xxx",Age: 0})
		dbptr := db.First(&md).Updates(map[string]interface{}{"name": "RookieOHY", "Age": 0})
		if dbptr.Error != nil {
			fmt.Println("发生的错误为", dbptr.Error) //record not found
			fmt.Println(errors.Is(dbptr.Error, gorm.ErrRecordNotFound))
		}
	}
}

func QueryWithCondition() {
	db, err := initConnection()
	if err == nil {
		//var md Model
		//var mds []Model
		var mfs []UserInfo
		//id主键查询（若查询不到返回一个nil model）
		//dbptr :=db.Model(&Model{}).Find(&md,13)
		//dbptr :=db.Model(&Model{}).First(&md,19)

		//where条件查询（按表字段、按结构体属性、按照map）
		//dbptr:=db.Where("name=? and age=?","RookieOHY03",0).Find(&md)
		//dbptr:=db.Where(Model{Name:"RookieOHY03",Age: 0}).Find(&md)
		//dbptr := db.Where(map[string]interface{}{
		//	"name": "RookieOHY03", "age": 0,
		//}).Find(&md)

		//or not order 等连接条件的使用
		//dbptr:=db.Where("name=?","RookieOHY03").Or("age",1).Find(&md)
		//dbptr:=db.Where("name=?","RookieOHY03").Not("age",1).Find(&md)
		//dbptr:=db.Order("rookie_uuid asc").Find(&mds)

		//内联（代替where）(按照表字段、按照map、按照结构体)
		//dbptr := db.Find(&mds, "age = ?", 0)
		//dbptr := db.Find(&mds, map[string]interface{}{"age":0})
		//dbptr := db.Find(&mds, Model{Age: 0})

		//Select(按照指定表字段、按照数组)
		//dbptr:=db.Select("name","age").Find(&mds)
		//dbptr:=db.Select([]string{"name","age"}).Find(&mds)

		dbptr := db.Model(&Model{}).Find(&mfs)

		if dbptr.Error != nil {
			//打印错误
			fmt.Println("发生的错误为", dbptr.Error) //record not found
			//判断错误的类型是否已存在的错误类型
			fmt.Println(errors.Is(dbptr.Error, gorm.ErrRecordNotFound))
		} else {
			//fmt.Println(md)
			//fmt.Println(mds)
			fmt.Println(mfs)
		}
	}
}

func Query() {
	db, err := initConnection()
	if err == nil {
		//使用map接收（没有被初始化map）
		//var resultMap map[string]interface{}
		//使用对应的结构体接收
		var md Model
		//按照主键排序 取第一条
		//db.Model(&Model{}).First(&md)
		//不按照主键排序
		//db.Model(&Model{}).Take(&md)
		//获取最后一条 按照主键排序
		db.Model(&Model{}).Last(&md)
		fmt.Println(md)
	}
}
func Main03() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/02-gorm?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction: false, //跳过gorm默认带上事物的设置
		NamingStrategy: schema.NamingStrategy{ //表、行的命名策略
			TablePrefix:   "t_",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	now := time.Now()
	//插入数据(单)
	db.Omit("Name").Create(&Model{
		EmbedModel: EmbedModel{
			//UUID: 1,
			Time1: &now,
			Time2: &now,
		},
		Name:     "RookieOHY",
		Age:      25,
		Birthday: &now,
		Email:    "204130199@qq.com",
	})

	//插入多条
	dbptr2 := db.Create(&[]Model{
		{Name: "RookieOHY02"},
		{Name: "RookieOHY03"},
	})

	err := dbptr2.Error
	rows := dbptr2.RowsAffected
	if err != nil {
		log.Println("新增数据错误", err)
	} else {
		log.Println("成功插入条数", rows)
	}
}

/*
Open
	作用：用于初始化数据库会话
	入参：2个接口对象。前者是数据库方言接口 Dialector,后者是 Option。
		- Dialector
			由具体的数据库方式驱动实现接口的全部方法（如引入的gorm mysql驱动包里的mysql.go便实现了该接口的全部方法，定义一个结构体Config且将其指针类型定义成新类型Dialector）
			可以使用mysql.go(驱动包)下的New函数，传入数据源初始化mysql驱动包的config结构体（实现了Dialector）返回对应Dialector对象
		- Option
			gorm的配置结构体gorm.go下的config实现了该接口，直接初始化
AutoMigrate
	作用：表的迁移（结构体建表）
	参数：一个或者多个结构体对象

*/
func main00() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/02-gorm?charset=utf8&parseTime=True&loc=Local", //数据源
	}), &gorm.Config{
		SkipDefaultTransaction: false, //跳过gorm默认带上事物的设置
		NamingStrategy: schema.NamingStrategy{ //表、行的命名策略
			TablePrefix:   "t_",  //创建表前缀
			SingularTable: false, //使用负数表
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	db.AutoMigrate(&Model{})
}

/*
	Migrator（接口）
		作用：可以操作表、列、索引、约束、视图等
*/
func main01() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(127.0.0.1:3306)/02-gorm?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction: false, //跳过gorm的默认事务
		NamingStrategy: schema.NamingStrategy{ //表、行的命名策略
			TablePrefix:   "t_",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	//设置数据库连接池
	sqlDB, _ := db.DB()
	//最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	//设置开启的最大打开连接数
	sqlDB.SetMaxOpenConns(100)
	//设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	//自动建表
	db.AutoMigrate(&User{})

	//获取Migrator对象
	m := db.Migrator()

	//获取数据库名字
	database := m.CurrentDatabase()
	fmt.Println(database)

	//表操作
	if m.HasTable(&User{}) {
		//删除表
		//m.DropTable(&User{})
	} else {
		//建立表且重命名
		m.CreateTable(&User{})
		//m.RenameTable(&User{},"rename_user")
	}

	//列操作
	//db.Migrator().AddColumn(&User{}, "sex")
	//db.Migrator().DropColumn(&User{}, "sex")
	db.Migrator().RenameColumn(&User{}, "sex", "new_sex")
}
