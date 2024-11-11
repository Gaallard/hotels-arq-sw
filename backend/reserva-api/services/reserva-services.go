package reservas

import (
	"context"
	"fmt"
	"net/http"
	dao "reserva-api/dao"
	domain "reserva-api/domain"
)

type Repository interface {
	GetReservaById(id int64) (dao.Reserva, error)
	InsertReserva(ctx context.Context, reserva dao.Reserva) (dao.Reserva, error)
	UpdateReserva(ctx context.Context, reserva dao.Reserva) (dao.Reserva, error)
	DeleteReserva(ctx context.Context, reserva dao.Reserva) error
}
type Service struct {
	mainRepo Repository
}

func NewService(mainRepo Repository) Service {
	return Service{
		mainRepo: mainRepo,
	}
}

func (service Service) GetReservaById(ctx context.Context, id int64) (domain.Reserva, error) {
	reservaDAO, err := service.mainRepo.GetReservaById(id)

	if err != nil {
		return domain.Reserva{}, fmt.Errorf("error getting hotel from repository: %v", err)
	}

	return domain.Reserva{
		ID:     reservaDAO.ID,
		User:   int64(reservaDAO.User),
		Hotel:  reservaDAO.Hotel,
		Noches: int64(reservaDAO.Noches),
		Estado: int64(reservaDAO.Estado),
	}, nil

}

func (service Service) InsertReserva(ctx context.Context, reserva domain.Reserva) (domain.Reserva, error) {
	var Reserva dao.Reserva
	Reserva.User = int(reserva.User)
	Reserva.Noches = int(reserva.Noches)
	Reserva.Hotel = reserva.Hotel
	Reserva.Estado = int(reserva.Estado)

	//comprobamos existencia del hotel en Mongo llamando a hotels-api
	urlHotel := fmt.Sprintf("localhost:8080/hotels/%s ", reserva.Hotel)
	response, err := http.Get(urlHotel)
	if err != nil {
		return domain.Reserva{}, fmt.Errorf("error getting hotel from server: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		return domain.Reserva{}, fmt.Errorf("Unexpected error with status code: %d", response.Status)
	}
	if response.StatusCode == http.StatusNotFound {
		return domain.Reserva{}, fmt.Errorf("hotel not found with ID: %s", reserva.Hotel)
	}

	//comprobamos existencia de usuario llamando a users-api
	urlUser := fmt.Sprintf("localhost:8080/users/%s ", reserva.User)
	response, err = http.Get(urlUser)
	if err != nil {
		return domain.Reserva{}, fmt.Errorf("error getting user from server: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		return domain.Reserva{}, fmt.Errorf("Unexpected error with status code: %d", response.Status)
	}
	if response.StatusCode == http.StatusNotFound {
		return domain.Reserva{}, fmt.Errorf("user not found with ID: %s", reserva.User)
	}

	reservaDomain, err := service.mainRepo.InsertReserva(ctx, Reserva)
	if err != nil {

		return reserva, fmt.Errorf("Error creando la reserva")
	}
	reserva.ID = reservaDomain.ID
	return reserva, nil
}

func (service Service) UpdateReserva(ctx context.Context, reserva domain.Reserva) (domain.Reserva, error) {
	var Reserva dao.Reserva
	Reserva.ID = reserva.ID
	Reserva.User = int(reserva.User)
	Reserva.Noches = int(reserva.Noches)
	Reserva.Hotel = reserva.Hotel
	Reserva.Estado = int(reserva.Estado)

	reservaDomain, err := service.mainRepo.UpdateReserva(ctx, Reserva)
	if err != nil {
		return reserva, fmt.Errorf("Error insertar reserva service")
	}

	reserva.Noches = int64(reservaDomain.Noches)

	return reserva, nil
}

func (service Service) DeleteReserva(ctx context.Context, reserva domain.Reserva) error {
	daoReserva := dao.Reserva{
		ID:     reserva.ID,
		User:   int(reserva.User),
		Noches: int(reserva.Noches),
		Hotel:  reserva.Hotel,
		Estado: int(reserva.Estado),
	}

	err := service.mainRepo.DeleteReserva(ctx, daoReserva)
	if err != nil {
		return fmt.Errorf("Error eliminando reserva service", reserva.ID, err)
	}

	return nil
}
