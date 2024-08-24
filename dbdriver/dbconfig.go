package dbdriver

import (
	"fmt"
)

type DBType int

const (
	DBTy_MySql DBType = iota
	DBTy_Sqlite
)

// DBConfig is the interface for database config
type DBConfig interface {
	GetDSN() string
	GetDBType() DBType
}

// MySqlConfig is the config for MySql
type MySqlConfig struct {
	Host   string
	Port   int
	User   string
	Passwd string
	DBName string
}

func NewMySqlConfig(host string, port int, user, passwd, dbName string) *MySqlConfig {
	return &MySqlConfig{
		Host:   host,
		Port:   port,
		User:   user,
		Passwd: passwd,
		DBName: dbName,
	}
}

func (m *MySqlConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", m.User, m.Passwd, m.Host, m.Port, m.DBName)
}

func (m *MySqlConfig) GetDBType() DBType {
	return DBTy_MySql
}

// SqliteConfig is the config for Sqlite
type SqliteConfig struct {
	DBPath string
}

func NewSqliteConfig(dbPath string) *SqliteConfig {
	return &SqliteConfig{
		DBPath: dbPath,
	}
}

func (s *SqliteConfig) GetDSN() string {
	return s.DBPath
}

func (s *SqliteConfig) GetDBType() DBType {
	return DBTy_Sqlite
}
