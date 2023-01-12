package redis

import "testing"

func TestGetRedisConnection(t *testing.T) {
	GetRedisConnection()
}

func TestGetRedisConnectionWithPool(t *testing.T) {
	GetRedisConnectionWithPool()
}

func TestGetJsonWithUnmarshal(t *testing.T) {
	GetJsonWithUnmarshal()
}

func TestMultipartSetString(t *testing.T) {
	MultipartSetString()
}

func TestExpireKey(t *testing.T) {
	ExpireKey()
}

func TestListOperation(t *testing.T) {
	ListOperation()
}

func TestHashOperation(t *testing.T) {
	HashOperation()
}
