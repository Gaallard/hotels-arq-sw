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

/*
func (service Service) HandleHotelNew(hotelNew hotelsDomain.HotelNew) {
	switch hotelNew.Operation {
	case "CREATE", "UPDATE":
		// Fetch hotel details from the local service
		hotel, err := service.hotelsAPI.GetHotelByID(context.Background(), hotelNew.HotelID)
		if err != nil {
			fmt.Printf("Error getting hotel (%s) from API: %v\n", hotelNew.HotelID, err)
			return
		}

		hotelDAO := hotelsDAO.Hotel{
			ID:        hotel.ID,
			Name:      hotel.Name,
			Address:   hotel.Address,
			City:      hotel.City,
			State:     hotel.State,
			Rating:    hotel.Rating,
			Amenities: hotel.Amenities,
		}

		// Handle Index operation
		if hotelNew.Operation == "CREATE" {
			if _, err := service.repository.Index(context.Background(), hotelDAO); err != nil {
				fmt.Printf("Error indexing hotel (%s): %v\n", hotelNew.HotelID, err)
			} else {
				fmt.Println("Hotel indexed successfully:", hotelNew.HotelID)
			}
		} else { // Handle Update operation
			if err := service.repository.Update(context.Background(), hotelDAO); err != nil {
				fmt.Printf("Error updating hotel (%s): %v\n", hotelNew.HotelID, err)
			} else {
				fmt.Println("Hotel updated successfully:", hotelNew.HotelID)
			}
		}

	case "DELETE":
		// Call Delete method directly since no hotel details are needed
		if err := service.repository.Delete(context.Background(), hotelNew.HotelID); err != nil {
			fmt.Printf("Error deleting hotel (%s): %v\n", hotelNew.HotelID, err)
		} else {
			fmt.Println("Hotel deleted successfully:", hotelNew.HotelID)
		}

	default:
		fmt.Printf("Unknown operation: %s\n", hotelNew.Operation)
	}
}
*/
