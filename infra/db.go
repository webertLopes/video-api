package infra

import (
	"github.com/jinzhu/gorm"
)

func ConnectToDb(connection string) *gorm.DB {

	db, err := gorm.Open("mysql", connection)

	if err != nil {
		panic(err)
	}

	return db
}
