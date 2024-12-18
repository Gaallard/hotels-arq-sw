package reservas

import "time"

type Reserva struct {
	ID           int64     `json:"id"`
	User         int64     `json:"user_id"`
	Hotel        string    `json:"hotel_id"`
	Noches       int64     `json:"noches"`
	FechaIngreso time.Time `json:"fecha_ingreso" time_format:"2006-01-02"`
	FechaSalida  time.Time `json:"fecha_salida" time_format:"2006-01-02"`
	Estado       int64     `json:"estado"`
}

type Hotel struct {
	Id              string    `json:"id"`
	Name            string    `json:"name"`
	Address         string    `json:"address"`
	City            string    `json:"city"`
	State           string    `json:"state"`
	Rating          float64   `json:"rating"`
	Amenities       []string  `json:"amenities"`
	Price           float64   `json:"price"`
	Available_rooms int64     `json:"available_rooms"`
	Noches          int64     `json:"noches"`
	FechaIngreso    time.Time `json:"fecha_ingreso" time_format:"2006-01-02"`
	FechaSalida     time.Time `json:"fecha_salida" time_format:"2006-01-02"`
}
