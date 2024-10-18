package hotels

import (
	"context"
	"fmt"
	"hotels-api/dao/hotels"
	hotelsDAO "hotels-api/dao/hotels"
	hotelsDomain "hotels-api/domain/hotels"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	GetHotelByID(ctx context.Context, id primitive.ObjectID) (hotels.Hotel, error)
	InsertHotel(ctx context.Context, hotel hotels.Hotel) (primitive.ObjectID, error)
	UpdateHotel(ctx context.Context, id primitive.ObjectID, hotel hotels.Hotel) (hotels.Hotel, error)
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
		IdMongo:         hotelDAO.IdMongo,
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

func (service Service) InsertHotel(ctx context.Context, hotel hotelsDomain.Hotel) (primitive.ObjectID, error) {

	hotelDAO := hotelsDAO.Hotel{
		// IdMongo:         hotel.IdMongo,
		Name:            hotel.Name,
		Address:         hotel.Address,
		City:            hotel.City,
		State:           hotel.State,
		Rating:          hotel.Rating,
		Amenities:       hotel.Amenities,
		Price:           hotel.Price,
		Available_rooms: hotel.Available_rooms,
	}

	_, err := service.mainRepository.InsertHotel(ctx, hotelDAO)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("Error inserting hotel into main repository: %v", err)
	}

	_, err = service.cacheRepository.InsertHotel(ctx, hotelDAO)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("Error inserting hotel into cache: %v", err)
	}

	return primitive.NewObjectID(), nil
}

func (service Service) UpdateHotel(ctx context.Context, id primitive.ObjectID, hotel hotelsDomain.Hotel) (hotelsDomain.Hotel, error) {

	hotelDAO := hotelsDAO.Hotel{
		// IdMongo:         hotel.IdMongo,
		Name:            hotel.Name,
		Address:         hotel.Address,
		City:            hotel.City,
		State:           hotel.State,
		Rating:          hotel.Rating,
		Amenities:       hotel.Amenities,
		Price:           hotel.Price,
		Available_rooms: hotel.Available_rooms,
	}

	newHotelDAO, err := service.cacheRepository.UpdateHotel(ctx, id, hotelDAO)
	if err != nil {
		newHotelDAO, err = service.mainRepository.UpdateHotel(ctx, id, hotelDAO)
		if err != nil {
			return hotelsDomain.Hotel{}, fmt.Errorf("error getting hotel from repository: %v", err)
		}
	}

	return hotelsDomain.Hotel{
		ID:              newHotelDAO.ID,
		IdMongo:         newHotelDAO.IdMongo,
		Name:            newHotelDAO.Name,
		Address:         newHotelDAO.Address,
		City:            newHotelDAO.City,
		State:           newHotelDAO.State,
		Rating:          newHotelDAO.Rating,
		Amenities:       newHotelDAO.Amenities,
		Price:           newHotelDAO.Price,
		Available_rooms: newHotelDAO.Available_rooms,
	}, nil
}
