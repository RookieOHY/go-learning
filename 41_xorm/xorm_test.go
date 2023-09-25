package _1_xorm

import "testing"

func TestCreateEngine(t *testing.T) {
	CreateEngine()
}

func TestPing(t *testing.T) {
	Ping()
}

func TestPingContext(t *testing.T) {
	PingContext()
}

func TestPingTimer(t *testing.T) {
	PingTimer()
}

func TestGetEngine(t *testing.T) {
	GetEngine()
}

func TestNewEngineWithParams(t *testing.T) {
	NewEngineWithParams()
}

func TestNewEngineWithDB(t *testing.T) {
	NewEngineWithDB()
}

func TestSnakeMapper(t *testing.T) {
	SnakeMapper()
}
