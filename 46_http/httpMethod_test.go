package _6_http

import (
	"fmt"
	"net/url"
	"testing"
)

// 测试

func TestReqWithKV(t *testing.T) {
	api, err := url.Parse("https://www.baidu.com/s")
	if err != nil {
		t.Fatal(err)
	}
	ReqWithKV(api, "wd", "golang")
	fmt.Println(api.String())
}

func TestReqWithPath(t *testing.T) {
	api, err := url.Parse("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}
	ReqWithPath(api, "/s")
	fmt.Println(api.String())
}
