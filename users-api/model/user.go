package models

type User struct {
	Id_user  int    `gorm:"primaryKey;autoIncrement;not null"`
	Username string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(256);not null"`
	Role     int    `gorm:"type:int;not null"`
}

type Users []User
