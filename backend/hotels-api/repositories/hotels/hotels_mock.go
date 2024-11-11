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
	hotel.Id = uuid.New().String()
	repository.docs[id] = hotel
	return id, nil
}

func (repository Mock) UpdateHotel(ctx context.Context, id string, hotel hotelsDAO.Hotel) (hotelsDAO.Hotel, error) {
	// Check if the hotel exists in the mock storage
	currentHotel, exists := repository.docs[hotel.Id]
	if !exists {
		return repository.docs[id], fmt.Errorf("hotel with ID %s not found", hotel.Id)
	}

	// Update only the fields that are non-zero or non-empty
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

	// Save the updated hotel back to the mock storage
	repository.docs[hotel.Id] = currentHotel
	return repository.docs[id], nil
}
