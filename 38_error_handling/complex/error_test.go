package complex

import "testing"

func TestHandleErrorWithCustom(t *testing.T) {
	type args struct {
		key string
	}
	//测试数组
	testArray := []struct {
		name string
		args args
	}{
		{name: "查询", args: args{"age"}},
		{name: "查询", args: args{"name"}},
		{name: "查询", args: args{"sex"}},
		{name: "查询", args: args{"addr"}},
	}
	//遍历数组
	for _, item := range testArray {
		t.Run(item.name, func(t *testing.T) {
			HandleErrorWithCustom(item.args.key)
		})
	}
}
