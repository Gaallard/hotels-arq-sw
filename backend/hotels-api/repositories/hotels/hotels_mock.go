package hotels

import (
	"context"
	"fmt"
	hotelsDAO "hotels-api/dao/hotels"

	"github.com/google/uuid"
)

type Mock struct {
	docs map[string]hotelsDAO.Hotel
}

func NewMock() Mock {
	return Mock{
		docs: make(map[string]hotelsDAO.Hotel),
	}
}

func (repository Mock) GetHotelByID(ctx context.Context, id string) (hotelsDAO.Hotel, error) {
	return repository.docs[id], nil
}

func (repository Mock) InsertHotel(ctx context.Context, hotel hotelsDAO.Hotel) (string, error) {
	id := uuid.New().String()
	hotel.Id = id
	repository.docs[id] = hotel
	return id, nil
}

func (repository Mock) UpdateHotel(ctx context.Context, id string, hotel hotelsDAO.Hotel) (hotelsDAO.Hotel, error) {
	currentHotel, exists := repository.docs[id]
	if !exists {
		return hotelsDAO.Hotel{}, fmt.Errorf("hotel with ID %s not found", id)
	}

	if hotel.Name != "" {
		currentHotel.Name = hotel.Name
	}
	if hotel.Address != "" {
		currentHotel.Address = hotel.Address
	}
	if hotel.City != "" {
		currentHotel.City = hotel.City
	}
	if hotel.State != "" {
		currentHotel.State = hotel.State
	}
	if hotel.Rating != 0 {
		currentHotel.Rating = hotel.Rating
	}
	if len(hotel.Amenities) > 0 {
		currentHotel.Amenities = hotel.Amenities
	}
	if hotel.Price != 0 {
		currentHotel.Price = hotel.Price
	}
	if hotel.Available_rooms != 0 {
		currentHotel.Available_rooms = hotel.Available_rooms
	}

	repository.docs[id] = currentHotel
	return currentHotel, nil
}
