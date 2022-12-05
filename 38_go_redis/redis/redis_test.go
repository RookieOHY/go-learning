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
