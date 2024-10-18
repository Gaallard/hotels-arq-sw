package hotels

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotels-api/dao/hotels"
	hotelsDAO "hotels-api/dao/hotels"
	hotelsDomain "hotels-api/domain/hotels"
)

type Repository interface {
	GetHotelByID(ctx context.Context, id primitive.ObjectID) (hotels.Hotel, error)
	InsertHotel(ctx context.Context, hotel hotelsDAO.Hotel) (error, string)
	UpdateHotel(ctx context.Context, id primitive.ObjectID, hotel hotelsDomain.Hotel) (hotelsDomain.Hotel, error)
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

func (service Service) GetHotelByID(ctx context.Context, id primitive.ObjectID) (hotelsDomain.Hotel, error) {
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
		ID:              hotelDAO.ID,
		Name:            hotelDAO.Name,
		Address:         hotelDAO.Address,
		City:            hotelDAO.City,
		State:           hotelDAO.State,
		Rating:          hotelDAO.Rating,
		Amenities:       hotelDAO.Amenities,
		Price:           hotelDAO.Price,
		Available_rooms: hotelDAO.Available_rooms,
	}, nil
}

func (service Service) InsertHotel(ctx context.Context, hotel hotelsDomain.Hotel) (error, string) {

	hotelDAO := hotelsDAO.Hotel{
		Name:            hotel.Name,
		Address:         hotel.Address,
		City:            hotel.City,
		State:           hotel.State,
		Rating:          hotel.Rating,
		Amenities:       hotel.Amenities,
		Price:           hotel.Price,
		Available_rooms: hotel.Available_rooms,
	}

	err, insertedId := service.mainRepository.InsertHotel(ctx, hotelDAO)
	if err != nil {
		return fmt.Errorf("Error inserting hotel into main repository: %v", err), ""
	}

	err, insertedId = service.cacheRepository.InsertHotel(ctx, hotelDAO)
	if err != nil {
		return fmt.Errorf("Error inserting hotel into cache: %v", err), insertedId
	}

	return nil, insertedId
}

func (service Service) UpdateHotel(ctx context.Context, id primitive.ObjectID, hotel hotelsDomain.Hotel) (hotelsDomain.Hotel, error) {

	updatedHotel, err := service.mainRepository.UpdateHotel(ctx, id, hotel)
	if err != nil {
		return hotelsDomain.Hotel{}, fmt.Errorf("Error updating hotel into main repository: %v", err)
	}

	_, err = service.cacheRepository.UpdateHotel(ctx, id, hotel)
	if err != nil {
		return hotelsDomain.Hotel{}, fmt.Errorf("Error updating hotel into cache: %v", err)
	}

	return updatedHotel, nil
}
