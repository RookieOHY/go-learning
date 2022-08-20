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
func TestOnce(t *testing.T) {
	for i := 0; i < 5; i++ {
		Once()
	}
}

func TestSyncCondAtConsumerProducer(t *testing.T) {
	ProduceAndConsumerSimulationWithSyncCond()
}
