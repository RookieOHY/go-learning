package _1_xorm

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

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

func TestGonicMapper(t *testing.T) {
	GonicMapper()
}

func TestTableName(t *testing.T) {
	TableName()
}

func TestColumnTag(t *testing.T) {
	ColumnTag()
}

func TestDump(t *testing.T) {
	Dump()
}

func TestImport(t *testing.T) {
	Import()
}

func TestInsert(t *testing.T) {
	Insert()
}

func TestInsertBatch(t *testing.T) {
	InsertBatch()
}

func TestInsertSlice(t *testing.T) {
	InsertSlice()
}

func TestInsertMultiTable(t *testing.T) {
	InsertMultiTable()
}

func TestAlias(t *testing.T) {
	Alias()
}

func TestOrderBy(t *testing.T) {
	OrderBy()
}

func TestQPrimary(t *testing.T){
	QPrimary()
}

func TestOr(t *testing.T){
	Or()
}

func TestSelect(t *testing.T){
	Select()
}

func TesSQL(t *testing.T) {
	SQL()
}

func TestIn(t *testing.T){
	In()
}

func TestCols(t *testing.T) {
	Cols()
}