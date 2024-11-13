package reservas

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	dao "reserva-api/dao"
	domain "reserva-api/domain"
)

type Repository interface {
	GetReservaById(id int64) (dao.Reserva, error)
	InsertReserva(ctx context.Context, reserva dao.Reserva) (dao.Reserva, error)
	UpdateReserva(ctx context.Context, reserva dao.Reserva) (dao.Reserva, error)
	DeleteReserva(ctx context.Context, reserva dao.Reserva) error
	GetMisReservasById(id int64) ([]dao.Reserva, error)
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

func (service Service) GetMisReservasById(ctx context.Context, id int64) ([]domain.Hotel, error) {
	reservaDAO, err := service.mainRepo.GetMisReservasById(id)

	if err != nil {
		return []domain.Hotel{}, fmt.Errorf("error getting hotel from repository: %v", err)
	}

	result := make([]domain.Hotel, 0)
	for _, hotel := range reservaDAO {
		if hotel.Estado == 1 {
			urlHotel := fmt.Sprintf("http://localhost:8081/hotels/%s ", hotel.Hotel)
			response, err := http.Get(urlHotel)

			if err != nil {
				return []domain.Hotel{}, fmt.Errorf("error getting hotel from server: %v", err)
			}
			if response.StatusCode != http.StatusOK {
				return []domain.Hotel{}, fmt.Errorf("Unexpected error with status code: %d", response.Status)
			}

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal("Error al ller el body de hotel", err)
			}
			var hotelBuscado domain.Hotel
			err = json.Unmarshal(body, &hotelBuscado)
			result = append(result, domain.Hotel{
				Id:     hotelBuscado.Id,
				Name:   hotelBuscado.Name,
				Noches: int64(hotel.Noches),
			})
		} else {
			continue
		}
	}
	return result, nil

}

func (service Service) InsertReserva(ctx context.Context, reserva domain.Reserva) (domain.Reserva, error) {
	var Reserva dao.Reserva
	Reserva.User = int(reserva.User)
	Reserva.Noches = int(reserva.Noches)
	Reserva.Hotel = reserva.Hotel
	Reserva.Estado = int(reserva.Estado)
	println("Recibe user: ", int(reserva.User))
	println("Recibe hotel: ", reserva.Hotel)

	//comprobamos existencia del hotel en Mongo llamando a hotels-api
	urlHotel := fmt.Sprintf("http://localhost:8081/hotels/%s ", reserva.Hotel)
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

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error al ller el body de hotel", err)
	}
	var hotel domain.Hotel
	err = json.Unmarshal(body, &hotel)
	if err != nil {
		log.Fatal("error al cargar el hotel para reservas: ", err)
	}

	if hotel.Available_rooms > 0 {

		hotel.Available_rooms = hotel.Available_rooms - 1
		hotelSend, err := json.Marshal(hotel)

		if err != nil {
			return domain.Reserva{}, fmt.Errorf("error to marshal: %v", err)
		}

		urlHotel := fmt.Sprintf("http://localhost:8081/hotels/%s ", hotel.Id)
		response, err := http.NewRequest(http.MethodPut, urlHotel, bytes.NewBuffer(hotelSend))

		if err != nil {
			return domain.Reserva{}, fmt.Errorf("error getting hotel from server: %v", err)
		}

		client := &http.Client{}
		resp, err := client.Do(response)
		if err != nil {
			return domain.Reserva{}, fmt.Errorf("error al enviar la solicitud PUT: %w", err)
		}
		defer resp.Body.Close()

		// Verificar si la respuesta fue exitosa
		if resp.StatusCode != http.StatusOK {
			return domain.Reserva{}, fmt.Errorf("error en la respuesta del servidor, código de estado: %d", resp.StatusCode)
		}

		reservaDomain, err := service.mainRepo.InsertReserva(ctx, Reserva)
		if err != nil {

			return reserva, fmt.Errorf("Error creando la reserva")
		}
		reserva.ID = reservaDomain.ID
		return reserva, nil
	} else {
		return domain.Reserva{}, fmt.Errorf("hotel with no rooms: %s", reserva.Hotel)
	}

}

func (service Service) UpdateReserva(ctx context.Context, reserva domain.Reserva) (domain.Reserva, error) {
	var Reserva dao.Reserva
	Reserva.ID = reserva.ID
	Reserva.User = int(reserva.User)
	Reserva.Noches = int(reserva.Noches)
	Reserva.Hotel = reserva.Hotel
	Reserva.Estado = int(reserva.Estado)

	println("user recibido service: ", reserva.User)
	println("hotel recibido service: ", reserva.Hotel)

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
	println("service: ", daoReserva.ID)
	err := service.mainRepo.DeleteReserva(ctx, daoReserva)
	if err != nil {
		return fmt.Errorf("Error eliminando reserva service", reserva.ID, err)
	}

	return nil
}
