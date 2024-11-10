package dao

type Hotel struct {
	ID              int64    `bson:"id"`
	IdMongo         string   `bson:"_id"`
	Name            string   `bson:"name"`
	Address         string   `bson:"address"`
	City            string   `bson:"city"`
	State           string   `bson:"state"`
	Rating          float64  `bson:"rating"`
	Amenities       []string `bson:"amenities"`
	Price           float64  `bson:"price"`
	Available_rooms int64    `bson:"available_rooms"`
}