package hotels

type Hotel struct {
	ID              int64    `json:"id"`
	Name            string   `json:"name"`
	Address         string   `json:"address"`
	City            string   `json:"city"`
	State           string   `json:"state"`
	Rating          float64  `json:"rating"`
	Amenities       []string `json:"amenities"`
	Price           float64  `json:"price"`
	Available_rooms int64    `json:"available_rooms"`
}
