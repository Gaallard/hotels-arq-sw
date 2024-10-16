package reservas

import (
	"context"
	"fmt"
	dao "reserva-api/dao"
	domain "reserva-api/domain"
)

type Repository interface {
	GetReservaById(id int64) (dao.Reserva, error)
	InsertReserva(ctx context.Context, reserva dao.Reserva) (dao.Reserva, error)
	UpdateReserva(ctx context.Context, reserva dao.Reserva) (dao.Reserva, error)
	CheckHotelExists(idHotel string) (bool, error)
}
type Service struct {
	mainRepo  Repository
	mongoRepo Repository
}

func NewService(mainRepo Repository, mongoRepo Repository) Service {
	return Service{
		mainRepo:  mainRepo,
		mongoRepo: mongoRepo,
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

	hotelExists, err := service.mongoRepo.CheckHotelExists(Reserva.Hotel)

	if err != nil {
		return domain.Reserva{}, fmt.Errorf("error checking hotel exists: %v", err)
	}

	if !hotelExists {
		return domain.Reserva{}, fmt.Errorf("hotel does not exist")
	}

	reservaDomain, err := service.mainRepo.InsertReserva(ctx, Reserva)
	if err != nil {

		return reserva, fmt.Errorf("Error insertar reserva service")
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

func (service Service) CheckHotelExists(idHotel string) (bool, error) {
	exists, err := service.mongoRepo.CheckHotelExists(idHotel)
	if err != nil {
		return false, fmt.Errorf("Error checking hotel from repository: %v", err)
	}
	return exists, nil
}
