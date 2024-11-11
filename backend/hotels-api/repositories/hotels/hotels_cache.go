package hotels

import (
	"context"
	"fmt"
	hotelsDAO "hotels-api/dao/hotels"
	"time"

	"github.com/karlseguin/ccache"
	_ "github.com/karlseguin/ccache"
)

const (
	keyFormat = "hotel:%s"
)

type CacheConfig struct {
	MaxSize      int64
	ItemsToPrune uint32
	Duration     time.Duration
}

type Cache struct {
	client   *ccache.Cache
	duration time.Duration
}

func NewCache(config CacheConfig) Cache {
	client := ccache.New(ccache.Configure().
		MaxSize(config.MaxSize).
		ItemsToPrune(config.ItemsToPrune))
	return Cache{
		client:   client,
		duration: config.Duration,
	}
}

func (repo Cache) GetAllHotels(ctx context.Context) ([]hotelsDAO.Hotel, error) {
	return []hotelsDAO.Hotel{}, nil
}

func (repo Cache) GetHotelByID(ctx context.Context, id string) (hotelsDAO.Hotel, error) {
	key := fmt.Sprintf(keyFormat, id)
	item := repo.client.Get(key)
	if item == nil {
		return hotelsDAO.Hotel{}, fmt.Errorf("not found item with key %s", key)
	}
	if item.Expired() {
		return hotelsDAO.Hotel{}, fmt.Errorf("item with key %s is expired", key)
	}
	hotelDAO, ok := item.Value().(hotelsDAO.Hotel)
	if !ok {
		return hotelsDAO.Hotel{}, fmt.Errorf("error converting item with key %s", key)
	}
	return hotelDAO, nil
}

func (repo Cache) InsertHotel(ctx context.Context, hotel hotelsDAO.Hotel) (string, error) {
	key := fmt.Sprintf(keyFormat, hotel.IdMongo)
	repo.client.Set(key, hotel, repo.duration)
	return hotel.IdMongo, nil
}

func (repository Cache) UpdateHotel(ctx context.Context, id string, hotel hotelsDAO.Hotel) error {
	key := fmt.Sprintf(keyFormat, hotel.ID)

	// Retrieve the current hotel data from the cache
	item := repository.client.Get(key)
	if item == nil {
		return fmt.Errorf("hotel with ID %s not found in cache", hotel.ID)
	}
	if item.Expired() {
		return fmt.Errorf("item with key %s is expired", key)
	}

	currentHotel, ok := item.Value().(hotelsDAO.Hotel)
	if !ok {
		return fmt.Errorf("error converting item with key %s", key)
	}

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

	repository.client.Set(key, currentHotel, repository.duration)

	return nil
}
