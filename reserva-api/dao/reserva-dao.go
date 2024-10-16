package reservas

type Reserva struct {
	ID     int64  `gorm:"primaryKey"`
	User   int    `gorm:"type:int;not null"`
	Hotel  string `gorm:"type:varchar(24);not null"` //varchar(24) porque el id de Mongo esta en 24 chars hexa
	Noches int    `gorm:"type:int;not null"`
	Estado int    `gorm:"type:int; default:1"`
}
