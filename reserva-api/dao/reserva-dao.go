package reservas

type Reserva struct {
	ID     int64  `gorm:"primaryKey"`
	User   int    `gorm:"type:int;not null"`
	Hotel  string `gorm:"type:varchar(600);not null"`
	Noches int    `gorm:"type:int;not null"`
}
