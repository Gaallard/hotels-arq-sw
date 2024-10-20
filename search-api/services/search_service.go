package search

import (
	"context"
	"fmt"
	"search-api/dao"
	"search-api/domain"
)

type Repository interface {
	Search(ctx context.Context, query string, offset int, limit int) ([]dao.Hotel, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return Service{
		repository: repository,
	}
}

func (service Service) Search(ctx context.Context, query string, offset int, limit int) ([]domain.Hotel, error) {
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
