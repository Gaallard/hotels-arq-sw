package search

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	hotelsDAO "search-api/dao"
	hotelsDomain "search-api/domain"
	"sync"
)

type Repository interface {
	Index(ctx context.Context, hotel hotelsDAO.Hotel) (string, error)
	Update(ctx context.Context, hotel hotelsDAO.Hotel) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string, limit int, offset int) ([]hotelsDAO.Hotel, error) // Updated signature
}

type ExternalRepository interface {
	GetHotelByID(ctx context.Context, id string) (hotelsDomain.Hotel, error)
}

type Service struct {
	repository Repository
	hotelsAPI  ExternalRepository
}

func NewService(repository Repository, hotelsAPI ExternalRepository) Service {
	return Service{
		repository: repository,
		hotelsAPI:  hotelsAPI,
	}
}

func (service Service) Search(ctx context.Context, query string, offset int, limit int) ([]hotelsDomain.Hotel, error) {
	// Llamar al método Search del repositorio
	hotelsDAOList, err := service.repository.Search(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error searching hotels: %w", err)
	}

	result := make([]hotelsDomain.Hotel, 0)
	for _, hotels := range hotelsDAOList {
		println("Id buscado2: ", hotels.Id)
		result = append(result, hotelsDomain.Hotel{
			Id:              hotels.Id,
			Name:            hotels.Name,
			Address:         hotels.Address,
			City:            hotels.City,
			State:           hotels.State,
			Rating:          hotels.Rating,
			Amenities:       hotels.Amenities,
			Price:           hotels.Price,
			Available_rooms: hotels.Available_rooms,
		})
	}

	hotelsChannel := make(chan hotelsDomain.Hotel, len(result))
	var group sync.WaitGroup

	for _, hotel := range result {
		group.Add(1)
		go func(hotelID string) {
			defer group.Done()

			urlHotel := fmt.Sprintf("http://hotels-api:8081/hotels/%s", hotelID)
			response, err := http.Get(urlHotel)
			if err != nil {
				log.Printf("Error fetching hotel (%s): %v\n", hotelID, err)
				return
			}
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Printf("Error reading body for hotel (%s): %v\n", hotelID, err)
				return
			}

			var hotel hotelsDomain.Hotel
			err = json.Unmarshal(body, &hotel)
			if err != nil {
				log.Printf("Error unmarshalling hotel (%s): %v\n", hotelID, err)
				return
			}

			if hotel.Available_rooms > 0 {
				hotelsChannel <- hotel
			}
		}(hotel.Id)
	}

	go func() {
		group.Wait()
		close(hotelsChannel)
	}()

	hotelsDomainList := make([]hotelsDomain.Hotel, 0)
	for hotel := range hotelsChannel {
		hotelsDomainList = append(hotelsDomainList, hotel)
	}

	return hotelsDomainList, nil
}

/*
func (service Service) GetHotelRooms(ctx context.Context, hotelID string, group *sync.WaitGroup, ch chan hotelsDomain.Hotel) {
	defer group.Done()

	urlHotel := fmt.Sprintf("http://localhost:8081/hotels/%s ", hotelID)
	response, err := http.Get(urlHotel)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error al ller el body de hotel", err)
	}
	var hotel hotelsDomain.Hotel
	err = json.Unmarshal(body, &hotel)
	if err != nil {
		log.Fatal("error al cargar el hotel para reservas: ", err)
	}

	ch <- hotel
}*/

func (service Service) HandleHotelNew(hotelNew hotelsDomain.HotelNew) {
	println("op: ", hotelNew.Operation)
	println("id: ", hotelNew.HotelID)

	switch hotelNew.Operation {
	case "CREATE", "UPDATE":
		// Fetch hotel details from the local service
		hotel, err := service.hotelsAPI.GetHotelByID(context.Background(), hotelNew.HotelID)
		if err != nil {
			fmt.Printf("Error getting hotel (%s) from API: %v\n", hotelNew.HotelID, err)
			return
		}

		hotelDAO := hotelsDAO.Hotel{
			Id:              hotel.Id,
			Name:            hotel.Name,
			Address:         hotel.Address,
			City:            hotel.City,
			State:           hotel.State,
			Rating:          hotel.Rating,
			Amenities:       hotel.Amenities,
			Price:           hotel.Price,
			Available_rooms: hotel.Available_rooms,
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
