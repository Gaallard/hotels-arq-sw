package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID              int64              `json:"id"`
	IdMongo         primitive.ObjectID `json:"_id,omitempty"`
	Name            string             `json:"name"`
	Address         string             `json:"address"`
	City            string             `json:"city"`
	State           string             `json:"state"`
	Rating          float64            `json:"rating"`
	Amenities       []string           `json:"amenities"`
	Price           float64            `json:"price"`
	Available_rooms int64              `json:"available_rooms"`
}
