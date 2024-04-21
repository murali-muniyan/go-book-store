package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "book_store"
)

type DBConf struct {
	db *gorm.DB
}

var dbConf *DBConf

func GetDB() *gorm.DB {
	return dbConf.db
}

func InitDB() {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil {
		log.Fatalf("error while connecting to DB: %v", err)
	}

	dbConf = &DBConf{db: db}
}
