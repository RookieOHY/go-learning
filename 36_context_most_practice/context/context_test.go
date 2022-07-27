package context

import "testing"

//测试
func TestContextWithCancel(t *testing.T) {
	for i := 0; i < 5; i++ {
		ContextWithCancel()
	}
}
