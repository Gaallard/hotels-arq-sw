package services

import (
	"context"
	"testing"

	hotelsDAO "hotels-api/dao/hotels"
	mock "hotels-api/repositories/hotels"

	"github.com/stretchr/testify/assert"
)

func NewMock() mock.Mock {
	return mock.NewMock()
}

func TestHotelService(t *testing.T) {
	mockRepo := NewMock()
	ctx := context.Background()

	// 1. Inserta un hotel
	hotel := hotelsDAO.Hotel{
		Name:            "Test Hotel",
		Address:         "123 Test St",
		City:            "Test City",
		State:           "Test State",
		Rating:          4.5,
		Amenities:       []string{"Pool", "WiFi"},
		Price:           150,
		Available_rooms: 10,
	}
	id, err := mockRepo.InsertHotel(ctx, hotel)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	// 2. Obtiene el hotel por ID
	retrievedHotel, err := mockRepo.GetHotelByID(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, hotel.Name, retrievedHotel.Name)
	assert.Equal(t, hotel.Price, retrievedHotel.Price)

	// 3. Actualiza el hotel
	update := hotelsDAO.Hotel{
		Name:            "Updated Hotel",
		Price:           200,
		Available_rooms: 5,
	}
	updatedHotel, err := mockRepo.UpdateHotel(ctx, id, update)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Hotel", updatedHotel.Name)
	assert.Equal(t, 200.0, updatedHotel.Price)
	assert.Equal(t, int(update.Available_rooms), int(updatedHotel.Available_rooms)) // Conversión explícita

	// 4. Maneja un ID inexistente en Update
	_, err = mockRepo.UpdateHotel(ctx, "non-existent-id", update)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}
