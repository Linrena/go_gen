package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MustNewMemDB() *gorm.DB {
	gormDB, err := gorm.Open(mysql.Open("file::memory:?parseTime=True"))
	if err != nil {
		panic(err)
	}
	return gormDB
}
