package reservas

type Reserva struct {
	ID     int64  `json:"id"`
	User   int64  `json:"user_id"`
	Hotel  string `json:"hotel_id"`
	Noches int64  `json:"noches"`
	Estado int64  `json:"estado"`
}
