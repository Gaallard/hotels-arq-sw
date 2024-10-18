package hotels

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	hotelsDAO "hotels-api/dao/hotels"
	hotelsDomain "hotels-api/domain/hotels"
	"time"

	"github.com/karlseguin/ccache"
	_ "github.com/karlseguin/ccache"
)

const (
	keyFormat = "hotel:%d"
)

type CacheConfig struct {
	MaxSize      int64
	ItemsToPrune uint32
}

type Cache struct {
	client *ccache.Cache
}

func NewCache(config CacheConfig) Cache {
	client := ccache.New(ccache.Configure().
		MaxSize(config.MaxSize).
		ItemsToPrune(config.ItemsToPrune))
	return Cache{
		client: client,
	}
}

func (repo Cache) GetHotelByID(ctx context.Context, id primitive.ObjectID) (hotelsDAO.Hotel, error) {
	key := id.Hex()
	item := repo.client.Get(key) //obtiene el id del hotel

	if item == nil {
		return hotelsDAO.Hotel{}, fmt.Errorf("hotel with id %d not found in cache", id)
	}

	hotelDAO, ok := item.Value().(hotelsDAO.Hotel)
	if !ok {
		return hotelsDAO.Hotel{}, fmt.Errorf("error converting item with key %s", key)
	}
	return hotelDAO, nil
}

func (repo Cache) InsertHotel(ctx context.Context, hotel hotelsDAO.Hotel) (error, string) {
	key := hotel.IdMongo.Hex()

	expiration := 5 * time.Minute
	repo.client.Set(key, hotel, expiration) //setea el id del hotel

	return nil, key
}

func (repo Cache) UpdateHotel(ctx context.Context, id primitive.ObjectID, hotel hotelsDomain.Hotel) (hotelsDomain.Hotel, error) {
	key := fmt.Sprintf(keyFormat, id.Hex())

	hotelJSON, err := json.Marshal(hotel)
	if err != nil {
		return hotelsDomain.Hotel{}, fmt.Errorf("error serializing hotel to JSON: %v", err)
	}

	expiration := 5 * time.Minute
	repo.client.Set(key, hotelJSON, expiration)

	return hotel, nil
}
