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

type EmbedModel struct {
	UUID  uint       `gorm:"primaryKey;type:int(11) auto_increment;not null;unique;;comment:自增id"`
	Time1 *time.Time `gorm:"type:datetime;column:create_time;comment:创建时间"`
	Time2 *time.Time `gorm:"type:datetime;column:update_time;comment:更新时间"`
}

type Model struct {
	//继承
	EmbedModel EmbedModel `gorm:"embedded;embeddedPrefix:rookie_"`
	//自己的属性
	Name     string     `gorm:"type:varchar(255);comment:姓名"`
	Age      uint8      `gorm:"type:int(11);comment:年龄"`
	Birthday *time.Time `gorm:"type:datetime;comment:生日"`
	Email    string     `gorm:"type:varchar(32);comment:邮箱地址"`
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
		DSN: "root:root12#$@tcp(127.0.0.1:3306)/02-gorm?charset=utf8&parseTime=True&loc=Local",
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

/*CRUD*/
func QueryWithCondition() {
	db, err := initConnection()
	if err == nil {
		//var md Model
		var mds []Model
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
		dbptr := db.Find(&mds, Model{Age: 0})

		//Select(获取结果指定字段)

		if dbptr.Error != nil {
			//打印错误
			fmt.Println("发生的错误为", dbptr.Error) //record not found
			//判断错误的类型是否已存在的错误类型
			fmt.Println(errors.Is(dbptr.Error, gorm.ErrRecordNotFound))
		} else {
			fmt.Println(mds)
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
		DSN: "root:root12#$@tcp(127.0.0.1:3306)/02-gorm?charset=utf8&parseTime=True&loc=Local",
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

/*模型:落地gorm标准的结构体或者接口*/
func main() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root12#$@tcp(127.0.0.1:3306)/02-gorm?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction: false, //跳过gorm默认带上事物的设置
		NamingStrategy: schema.NamingStrategy{ //表、行的命名策略
			TablePrefix:   "t_",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	db.AutoMigrate(&Model{})
}

/*基本用法*/
func main01() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root12#$@tcp(127.0.0.1:3306)/02-gorm?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction: false, //跳过gorm默认带上事物的设置
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
	//err := db.AutoMigrate(&User{})
	//m := db.Migrator()

	//获取数据库
	//database := m.CurrentDatabase()
	//fmt.Println(database)

	//表操作
	//if m.HasTable(&User{}) {
	//	//删除表
	//	//m.DropTable(&User{})
	//}else {
	//	//建立表且重命名
	//	m.CreateTable(&User{})
	//	//m.RenameTable(&User{},"rename_user")
	//}

	//列操作
	//db.Migrator().AddColumn(&User{},"Sex")
	//db.Migrator().DropColumn(&User{},"Sex")
	//db.Migrator().RenameColumn(&User{},"Sex","NewSex")
}
