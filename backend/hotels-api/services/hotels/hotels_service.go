package hotels

import (
	"context"
	"fmt"

	"hotels-api/dao/hotels"
	hotelsDAO "hotels-api/dao/hotels"
	hotelsDomain "hotels-api/domain/hotels"
	"sync"
)

type Repository interface {
	GetHotelByID(ctx context.Context, id string) (hotels.Hotel, error)
	InsertHotel(ctx context.Context, hotel hotels.Hotel) (string, error)
	UpdateHotel(ctx context.Context, id string, hotel hotels.Hotel) error
	GetAllHotels(ctx context.Context) ([]hotels.Hotel, error)
}

type RabbitMQ interface {
	Publish(hotelNew hotelsDomain.HotelNew) error
}

type Service struct {
	mainRepository  Repository
	cacheRepository Repository
	rabbitRpo       RabbitMQ
}

type RoomAvailability struct {
	HotelID        string
	AvailableRooms int64
}

func NewService(mainRepository Repository, cacheRepository Repository, rabbitRepo RabbitMQ) Service {
	return Service{
		mainRepository:  mainRepository,
		cacheRepository: cacheRepository,
		rabbitRpo:       rabbitRepo,
	}
}

func (service Service) GetHotelByID(ctx context.Context, id string) (hotelsDomain.Hotel, error) {
	hotelDAO, err := service.cacheRepository.GetHotelByID(ctx, id)
	if err != nil {
		// Get hotel from main repository
		hotelDAO, err = service.mainRepository.GetHotelByID(ctx, id)
		if err != nil {
			return hotelsDomain.Hotel{}, fmt.Errorf("error getting hotel from repository: %v", err)
		}

		if _, err := service.cacheRepository.InsertHotel(ctx, hotelDAO); err != nil {
			return hotelsDomain.Hotel{}, fmt.Errorf("error creating hotel in cache: %w", err)
		}
	}
	// prueba que envia mensaje, cambiar dsp en funciones utiles
	//service.rabbitRpo.Publish(id)
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

func (service Service) GetAllHotels(ctx context.Context) error {
	//ctx := context.Background()
	hotelDAO, err := service.mainRepository.GetAllHotels(ctx)
	if err != nil {
		return fmt.Errorf("error getting hotel from mainrepo: %w", err)
	}

	// prueba que envia mensaje, cambiar dsp en funciones utiles
	//service.rabbitRpo.Publish(id)
	// Convert DAO to DTO

	for _, hotels := range hotelDAO {
		if err := service.rabbitRpo.Publish(hotelsDomain.HotelNew{
			Operation: "CREATE",
			HotelID:   hotels.IdMongo,
		}); err != nil {
			return fmt.Errorf("error publishing hotel new: %w", err)
		}
		println("Se encontro este hotel: ", hotels.IdMongo)
	}
	return nil
}

func (service Service) GetHotelsAvailability(ctx context.Context) (map[string]int64, error) {

	hotelDAO, err := service.mainRepository.GetAllHotels(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting hotels: %w", err)
	}

	result := make(map[string]int64)

	var waitGroup sync.WaitGroup
	waitGroup.Add(len(hotelDAO))

	ch := make(chan RoomAvailability)

	go func() {
		for roomAvailability := range ch {
			if roomAvailability.AvailableRooms > 0 {
				result[roomAvailability.HotelID] = roomAvailability.AvailableRooms
			}
		}
	}()

	for _, hotel := range hotelDAO {
		go service.GetHotelRooms(ctx, hotel.IdMongo, &waitGroup, ch)
	}

	waitGroup.Wait()
	close(ch)

	return result, nil
}

func (service Service) GetHotelRooms(ctx context.Context, hotelID string, group *sync.WaitGroup, ch chan RoomAvailability) {
	defer group.Done()

	hotel, err := service.mainRepository.GetHotelByID(ctx, hotelID)
	if err != nil {
		return
	}

	ch <- RoomAvailability{
		HotelID:        hotel.IdMongo,
		AvailableRooms: hotel.Available_rooms,
	}
}

func (service Service) InsertHotel(ctx context.Context, hotel hotelsDomain.Hotel) (string, error) {

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

	id, err := service.mainRepository.InsertHotel(ctx, hotelDAO)
	if err != nil {
		return "", fmt.Errorf("Error inserting hotel into main repository: %v", err)
	}

	hotelDAO.IdMongo = id
	_, err = service.cacheRepository.InsertHotel(ctx, hotelDAO)
	if err != nil {
		return "", fmt.Errorf("Error inserting hotel into cache: %v", err)
	}

	if err := service.rabbitRpo.Publish(hotelsDomain.HotelNew{
		Operation: "CREATE",
		HotelID:   id,
	}); err != nil {
		return "", fmt.Errorf("error publishing hotel new: %w", err)
	}

	return id, nil
}

func (service Service) UpdateHotel(ctx context.Context, id string, hotel hotelsDomain.Hotel) error {

	hotelDAO := hotelsDAO.Hotel{
		IdMongo:         hotel.IdMongo,
		Name:            hotel.Name,
		Address:         hotel.Address,
		City:            hotel.City,
		State:           hotel.State,
		Rating:          hotel.Rating,
		Amenities:       hotel.Amenities,
		Price:           hotel.Price,
		Available_rooms: hotel.Available_rooms,
	}

	err := service.cacheRepository.UpdateHotel(ctx, id, hotelDAO)
	if err != nil {
		println("No esta en cache")
		err = service.mainRepository.UpdateHotel(ctx, id, hotelDAO)
		if err != nil {
			return fmt.Errorf("error updating hotel in main repository: %w", err)
		}

	}

	if err := service.rabbitRpo.Publish(hotelsDomain.HotelNew{
		Operation: "UPDATE",
		HotelID:   id,
	}); err != nil {
		return fmt.Errorf("error publishing hotel update: %w", err)
	}

	return nil
}
