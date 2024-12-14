package clientUsers

import (
	e "backend/errors"
	Model "backend/model"
	"fmt"
	"sync"
	"time"

	"github.com/karlseguin/ccache"
)

const (
	keyFormat = "user:%d"
)

type CacheConfig struct {
	MaxSize      int64
	ItemsToPrune uint32
	Duration     time.Duration
}

type Cache struct {
	client   *ccache.Cache
	duration time.Duration
	keys     sync.Map
}

func NewCache(config CacheConfig) Cache {
	client := ccache.New(ccache.Configure().
		MaxSize(config.MaxSize).
		ItemsToPrune(config.ItemsToPrune))
	return Cache{
		client:   client,
		duration: config.Duration,
		keys:     sync.Map{},
	}
}

func (repo Cache) GetUserByName(Usuario Model.User) (Model.User, e.ApiError) {
	println("name: ", Usuario.User)

	key := fmt.Sprintf("user:user:%s", Usuario.User)
	item := repo.client.Get(key)
	println("Key: ", key)

	if item == nil {
		println("error buscar user")
		return Model.User{}, e.NewBadRequestApiError("Error al buscar el usuario")
	}
	if item.Expired() {
		println("expire")

		return Model.User{}, e.NewBadRequestApiError("item expire")
	}
	userDAO, ok := item.Value().(Model.User)
	if !ok {
		return Model.User{}, e.NewBadRequestApiError("Error converting user")
	}

	println("se busco usuario en cache")

	return userDAO, nil

}

func (repo Cache) InsertUsuario(user Model.User) (Model.User, e.ApiError) {
	idkey := fmt.Sprintf("user:id:%d", user.Id)
	userkey := fmt.Sprintf("user:user:%s", user.User)

	repo.client.Set(idkey, user, repo.duration)
	repo.client.Set(userkey, user, repo.duration)

	println("se inserto user a cache id: ", idkey)
	println("se inserto user a cache user: ", userkey)

	return user, nil
}

func (repo Cache) GetUserById(Id int) (Model.User, e.ApiError) {
	key := fmt.Sprintf(keyFormat, Id)
	println("se busco usuario en cache")
	item := repo.client.Get(key)

	if item == nil {
		return Model.User{}, e.NewBadRequestApiError("Error al buscar el usuario")
	}
	if item.Expired() {
		return Model.User{}, e.NewBadRequestApiError("item expire")
	}
	userDAO, ok := item.Value().(Model.User)
	if !ok {
		return Model.User{}, e.NewBadRequestApiError("Error converting user")
	}
	return userDAO, nil
}

func (repo Cache) GetuserName(buscado int) (string, e.ApiError) {

	key := fmt.Sprintf(keyFormat, buscado)
	item := repo.client.Get(key)

	if item == nil {
		return "", e.NewBadRequestApiError("Error al buscar el usuario")
	}
	if item.Expired() {
		return "", e.NewBadRequestApiError("item expire")
	}
	userDAO, ok := item.Value().(Model.User)
	if !ok {
		return "", e.NewBadRequestApiError("Error converting user")
	}
	return userDAO.User, nil
}
