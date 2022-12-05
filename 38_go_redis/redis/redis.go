package redis

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

/*
	利用（官方推荐的库）redisgo 操作Redis服务端
*/
// 普通方式连接服务端、设置值、获取值
func GetRedisConnection() {
	// 普通连接
	cli, err := redis.Dial("tcp", "127.0.0.1:6379")
	// 账号密码方式连接
	//cli, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialUsername("ohy"), redis.DialPassword("123456"))
	if err != nil {
		fmt.Println("客户端连接异常", err)
	}
	_, err = cli.Do("Set", "authorName", "RookieOHY")
	if err != nil {
		fmt.Println("Set指令出现错误", err)
	}
	s, err := redis.String(cli.Do("Get", "authorName"))
	if err == nil {
		fmt.Println("Get Key [authorName] ==>", s)
	}
	//执行完毕 释放连接
	defer cli.Close()
}

// GetRedisConnectionWithPool 连接池方式连接
func GetRedisConnectionWithPool() {
	p := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	cli := p.Get()
	repl, err := cli.Do("Get", "authorName")
	str, err := redis.String(repl, err)
	if err != nil {
		fmt.Println("Get Key Error ==>", err)
		return
	}
	fmt.Println("Get Key [authorName] ==>", str)
	defer cli.Close()
}

type Student struct {
	Name string
	Age  int
}

var (
	pool        *redis.Pool
	serverParam = flag.String("127.0.0.1", ":6379", "")
)

func newPool(serverParam string) *redis.Pool {
	return &redis.Pool{
		//
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", serverParam) },
	}
}

// GetJsonWithUnmarshal redis设置对象的序列化和反序列化
func GetJsonWithUnmarshal() {
	flag.Parse()
	pool := newPool(*serverParam)
	cli := pool.Get()
	student := Student{
		Name: "RookieOHY",
		Age:  25,
	}
	jsonObj, err := json.Marshal(&student)
	if err != nil {
		fmt.Println("Data Marshal Error ==>", err)
		return
	}
	//set
	setResult, error2 := cli.Do("Set", "name", string(jsonObj))
	if error2 != nil {
		fmt.Println("set key Error ==>", error2)
		return
	}
	fmt.Println("set result ==> ", setResult)
	//get
	byteStr, err := redis.String(cli.Do("Get", "name"))
	if err != nil {
		fmt.Println("get key Error ==>", err)
		return
	}
	byteData := []byte(byteStr)
	std := Student{}
	json.Unmarshal(byteData, &std)
	fmt.Println("the unmarshal result ==>", std)
	defer cli.Close()
}
