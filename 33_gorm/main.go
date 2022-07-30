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
		CURD:


*/

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

type EmbedModel struct {
	UUID  uint      `gorm:"primaryKey;type:int(11) auto_increment;not null;unique;;comment:自增id"`
	Time1 time.Time `gorm:"type:datetime;column:create_time;comment:创建时间"`
	Time2 time.Time `gorm:"type:datetime;column:update_time;comment:更新时间"`
}

type Model struct {
	//继承
	EmbedModel EmbedModel `gorm:"embedded;embeddedPrefix:rookie_"`
	//自己的属性
	Name     string    `gorm:"type:varchar(255);comment:姓名"`
	Age      uint8     `gorm:"type:int(11);comment:年龄"`
	Birthday time.Time `gorm:"type:datetime;comment:生日"`
	Email    string    `gorm:"type:varchar(32);comment:邮箱地址"`
}

type User struct {
	gorm.Model
	Name string
	Age  int
	//add
	Sex int
}

/*CRUD*/
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
	//插入数据
	dbptr := db.Create(&Model{
		EmbedModel: EmbedModel{
			Time1: time.Now(),
			Time2: time.Now(),
		},
		Name:     "RookieOHY",
		Age:      25,
		Birthday: time.Now(),
		Email:    "204130199@qq.com",
	})
	err := dbptr.Error
	rows := dbptr.RowsAffected
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
