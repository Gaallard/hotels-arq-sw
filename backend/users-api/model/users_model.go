package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	User     string `gorm:"type:varchar(600);not null"`
	Password string `gorm:"type:varchar(350);not null"`
	Admin    bool   `gorm:"not null"`
}

type Users []User
