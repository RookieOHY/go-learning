package simple

import "testing"

/*
	Go错误处理方式的最佳实践测试用例（简单版）
*/
func TestServiceName(t *testing.T) {
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
	}
	//遍历数组
	for _, item := range testArray {
		t.Run(item.name, func(t *testing.T) {
			ServiceName(item.args.key)
		})
	}
}
