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
	}, nil

}

func (service Service) InsertReserva(ctx context.Context, reserva domain.Reserva) (domain.Reserva, error) {
	var Reserva dao.Reserva
	Reserva.User = int(reserva.User)
	Reserva.Noches = int(reserva.Noches)
	Reserva.Hotel = reserva.Hotel

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

	reservaDomain, err := service.mainRepo.UpdateReserva(ctx, Reserva)
	if err != nil {
		return reserva, fmt.Errorf("Error insertar reserva service")
	}

	reserva.Noches = int64(reservaDomain.Noches)

	return reserva, nil
}
