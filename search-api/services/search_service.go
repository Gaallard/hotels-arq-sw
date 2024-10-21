package search

import (
	"context"
	"fmt"
	"log"
	"search-api/dao"
	"search-api/domain"
)

type Repository interface {
	Search(ctx context.Context, query string, offset int, limit int) ([]dao.Hotel, error)
}

type RabbitMQ2 interface {
	//Receive(hotelNew hotels.HotelNew) error
	ConsumeCola()
}

type Service struct {
	repository Repository
	rabbitRepo RabbitMQ2
}

func NewService(repository Repository, rabbitRepo RabbitMQ2) Service {
	return Service{
		repository: repository,
		rabbitRepo: rabbitRepo,
	}
}

func (service Service) Search(ctx context.Context, query string, offset int, limit int) ([]domain.Hotel, error) {
	//prueba de conexion ue llega mensaje
	log.Println("Connecion correcta")
	service.rabbitRepo.ConsumeCola()

	hotels, err := service.repository.Search(ctx, query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("error searching hotelsDAO: %s", err.Error())
	}

	result := make([]domain.Hotel, 0)
	for _, hotel := range hotels {
		result = append(result, domain.Hotel{
			IdMongo:         hotel.IdMongo,
			Name:            hotel.Name,
			Address:         hotel.Address,
			City:            hotel.City,
			State:           hotel.State,
			Rating:          hotel.Rating,
			Amenities:       hotel.Amenities,
			Price:           hotel.Price,
			Available_rooms: hotel.Available_rooms,
		})
	}

	return result, nil
}
