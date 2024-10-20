package db

import (
	users "users-api/users-api/client"
	model "users-api/users-api/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	DBName := "users"
	DBUser := "root"
	DBPass := "root"
	DBHost := "localhost"

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	users.Db = db
}

func StartDbEngine() {
	db.AutoMigrate(&model.User{})

	log.Info("Finishing Migration Database Table")
}
