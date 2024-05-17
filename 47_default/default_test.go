package _7_default

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

// 这是测试结构体
type Config struct {
	Name     string   `default:"defaultName"`
	Enabled  bool     `default:"true"`
	MaxCount int      `default:"100"`
	Timeout  float64  `default:"1.5"`
	Data     []byte   `default:"example"`
	Age      *int     `default:"25"`
	Price    *float64 `default:"19.99"`
}

func TestApply(t *testing.T) {
	config := &Config{}
	err := Apply(config)
	if err != nil {
		log.Fatalf("Apply error %v", err)
	}
	fmt.Println(config)
}

type ConfigNew struct {
	Host string `json:"Host" default:"0.0.0.0"`
	Port int    `json:"Port" default:"8080"`
}

func TestApply02(t *testing.T) {
	var cfgNew ConfigNew
	err := json.Unmarshal([]byte(`{"Host":"RookieOHY"}`), &cfgNew)
	if err != nil {
		log.Fatalf("Unmarshal error: %v", err)
	}
	err = Apply(&cfgNew)
	if err != nil {
		log.Fatalf("Apply error: %v", err)
	}
	fmt.Println(cfgNew.Port) // 输出端口号
	fmt.Println(cfgNew.Host) // 输出主机地址
}

type GroceryList struct {
	Fruit struct {
		Bananas int `default:"8"`
		Pears   int `default:"12"`
	}
	Vegetables *struct {
		Artichokes    int `default:"4"`
		SweetPotatoes int `default:"16"`
	}
}

func TestApply3(t *testing.T) {
	list := GroceryList{}
	err := Apply(&list)
	if err != nil {
		log.Fatalf("apply error %v", err)
	}
	fmt.Println(list)
	fmt.Println(list.Vegetables.Artichokes)

}

type ConfigNN struct {
	FInt  int32 `default:"12"`
	FRune rune  `default:"a"`
}

type ConfigNNN struct {
	FRune1 rune `default:"3"`  // this is a unicode ETX or ctrl+c
	FRune2 rune `default:"97"` // this is a unicode `a`
}

func TestBeyondRuneStruct(t *testing.T) {
	//n := ConfigNN{}
	n := ConfigNNN{}

	err := Apply(&n)
	if err != nil {
		log.Fatalf("apply error %v", err)
	}
	fmt.Println(n)
}
