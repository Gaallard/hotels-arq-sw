package hotels

import (
	"context"
	"fmt"
	"hotels-api/dao/hotels"
	hotelsDAO "hotels-api/dao/hotels"
	hotelsDomain "hotels-api/domain/hotels"
)

type Repository interface {
	GetHotelByID(ctx context.Context, id int64) (hotels.Hotel, error)
	InsertHotel(ctx context.Context, hotel hotelsDAO.Hotel) error
	DeleteHotel(ctx context.Context, id int64) error
	UpdateHotel(ctx context.Context, id int64, hotel hotelsDomain.Hotel) (hotelsDomain.Hotel, error)
}

type Service struct {
	mainRepository  Repository
	cacheRepository Repository
}

func NewService(mainRepository Repository, cacheRepository Repository) Service {
	return Service{
		mainRepository:  mainRepository,
		cacheRepository: cacheRepository,
	}
}

func (service Service) GetHotelByID(ctx context.Context, id int64) (hotelsDomain.Hotel, error) {
	hotelDAO, err := service.cacheRepository.GetHotelByID(ctx, id)
	if err != nil {
		// Get hotel from main repository
		hotelDAO, err = service.mainRepository.GetHotelByID(ctx, id)
		if err != nil {
			return hotelsDomain.Hotel{}, fmt.Errorf("error getting hotel from repository: %v", err)
		}

		// TODO: service.cacheRepository.CreateHotel
	}

	// Convert DAO to DTO
	return hotelsDomain.Hotel{
		ID:        hotelDAO.ID,
		Name:      hotelDAO.Name,
		Address:   hotelDAO.Address,
		City:      hotelDAO.City,
		State:     hotelDAO.State,
		Rating:    hotelDAO.Rating,
		Amenities: hotelDAO.Amenities,
	}, nil
}

func (service Service) InsertHotel(ctx context.Context, hotel hotelsDomain.Hotel) error {

	hotelDAO := hotelsDAO.Hotel{
		ID:        hotel.ID,
		Name:      hotel.Name,
		Address:   hotel.Address,
		City:      hotel.City,
		State:     hotel.State,
		Rating:    hotel.Rating,
		Amenities: hotel.Amenities,
	}

	if err := service.mainRepository.InsertHotel(ctx, hotelDAO); err != nil {
		return fmt.Errorf("Error inserting hotel into main repository: %v", err)
	}

	if err := service.cacheRepository.InsertHotel(ctx, hotelDAO); err != nil {
		return fmt.Errorf("Error inserting hotel into cache: %v", err)
	}

	return nil
}

func (service Service) DeleteHotel(ctx context.Context, id int64) error {
	// Intenta eliminar el hotel del main repository
	err := service.mainRepository.DeleteHotel(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting hotel from repository: %v", err)
	}

	// Intenta eliminar el hotel del cache repository (si existe)
	err = service.cacheRepository.DeleteHotel(ctx, id)
	if err != nil {
		fmt.Printf("Error deleting hotel from cache: %v\n", err)
	}

	return nil
}

func (service Service) UpdateHotel(ctx context.Context, id int64, hotel hotelsDomain.Hotel) (hotelsDomain.Hotel, error) {

	hotelDomain := hotelsDomain.Hotel{
		Name:      hotel.Name,
		Address:   hotel.Address,
		City:      hotel.City,
		State:     hotel.State,
		Rating:    hotel.Rating,
		Amenities: hotel.Amenities,
	}

	if err := service.mainRepository.UpdateHotel(ctx, id, hotelDomain); err != nil {
		return hotelsDomain.Hotel{}, fmt.Errorf("Error updating hotel into main repository: %v", err)
	}

	if err := service.cacheRepository.UpdateHotel(ctx, id, hotelDomain); err != nil {
		return hotelsDomain.Hotel{}, fmt.Errorf("Error updating hotel into cache: %v", err)
	}

	return hotelDomain, nil
}
