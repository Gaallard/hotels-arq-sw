package hotels

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID              int64              `bson:"id"`
	IdMongo         primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name"`
	Address         string             `bson:"address"`
	City            string             `bson:"city"`
	State           string             `bson:"state"`
	Rating          float64            `bson:"rating"`
	Amenities       []string           `bson:"amenities"`
	Price           float64            `bson:"price"`
	Available_rooms int64              `bson:"available_rooms"`
}
