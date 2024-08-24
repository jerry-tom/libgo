package dborm

import (
	"fmt"
	"libgo/dbdriver"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbOrm struct {
	db *gorm.DB
}

func NewDbOrm() *DbOrm {
	return &DbOrm{
		db: nil,
	}
}

func (r *DbOrm) Open(dbConf dbdriver.DBConfig) error {
	if dbConf == nil {
		return fmt.Errorf("db config is nil")
	}

	if r.db != nil {
		r.Close()
	}

	var err error
	switch dbConf.GetDBType() {
	case dbdriver.DBTy_MySql:
		r.db, err = gorm.Open(mysql.Open(dbConf.GetDSN()), &gorm.Config{})
	case dbdriver.DBTy_Sqlite:
		r.db, err = gorm.Open(sqlite.Open(dbConf.GetDSN()), &gorm.Config{})
	default:
		err = fmt.Errorf("unsupport db type")
	}

	return err
}

/*
func (r *DbOrm) FindWithWhere(recordData interface{}, query string, args ...string) error {
	ret := r.db.Where(query, args...).Find(recordData)
	return ret.Error
}
*/
func (r *DbOrm) Insert(recordData interface{}) error {
	ret := r.db.Create(recordData)
	return ret.Error
}

func (r *DbOrm) InsertIntoTable(tableName string, recordData interface{}) error {
	ret := r.db.Table(tableName).Create(recordData)
	return ret.Error
}

func (r *DbOrm) Close() {
	if r.db != nil {
		sqlDB, _ := r.db.DB()
		sqlDB.Close()
	}
}
