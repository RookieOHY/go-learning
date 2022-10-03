package _9_panic_handling

import "testing"

func TestRunRoutines(t *testing.T) {
	testArray := []struct {
		testName string
	}{
		{testName: "测试异常处理"},
	}

	for _, tt := range testArray {
		t.Run(tt.testName, func(t *testing.T) {
			runRoutines()
		})
	}
}
