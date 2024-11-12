package db

import (
	userClient "backend/clients/users"
	Model "backend/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {

	/*
		DBName := os.Getenv("DB_NAME")
		DBUser := os.Getenv("DB_USER")
		DBPass := os.Getenv("DB_PASSWORD")
		DBHost := os.Getenv("DB_HOST")
		DBPort := os.Getenv("DB_PORT")
		// ------------------------

		dsn := DBUser + ":" + DBPass + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8&parseTime=True"
		db, err = gorm.Open("mysql", dsn)

		if err != nil {
			log.Info("Connection Failed to Open")
			log.Fatal(err)
		} else {
			log.Info("Connection Established")
		}

		userClient.Db = db*/

	DBName := "users"
	DBUser := "root"
	DBPass := "Tomas1927"
	DBHost := "localhost"

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	userClient.Db = db
}

func StartDbEngine() {
	db.AutoMigrate(&Model.User{})
	log.Info("Finishing Migration Database Tables")
}
