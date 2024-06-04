package services

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlDatabase struct {
	db gorm.DB
}

func (sqlDatabase *MySqlDatabase) Connect(connection_string string) {
	db, err := gorm.Open(mysql.Open(connection_string), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDatabase.db = *db
}
