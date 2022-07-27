package syncronized

import "testing"

func TestUnsafeAdd(t *testing.T) {
	UnsafeAdd()
}
func TestSafeAdd(t *testing.T) {
	SafeAdd()
}
func TestMutex(t *testing.T) {
	Mutex()
}
