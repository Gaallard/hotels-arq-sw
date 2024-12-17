package reservas

import "time"

type Reserva struct {
	ID           int64     `gorm:"primaryKey"`
	User         int       `gorm:"type:int;not null"`
	Hotel        string    `gorm:"type:varchar(24);not null"` //varchar(24) porque el id de Mongo está en 24 caracteres hexadecimales
	Noches       int       `gorm:"type:int;not null"`
	FechaIngreso time.Time `gorm:"type:date;not null"` // Removed extra quotes and spaces
	FechaSalida  time.Time `gorm:"type:date;not null"` // Fixed comma placement
	Estado       int       `gorm:"type:int;default:1"`
}
