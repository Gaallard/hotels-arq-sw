package reservas

import (
	"context"
	"fmt"
	dao "reserva-api/dao"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

type SQLConfig struct {
	Name string
	User string
	Pass string
	Host string
}

type SQL struct {
	db       *gorm.DB
	Database string
}

// DeleteReserva implements reservas.Repository.

func NewSql(config SQLConfig) SQL {
	db, err := gorm.Open("mysql", config.User+":"+config.Pass+"@tcp("+config.Host+":3306)/"+config.Name+"?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Println("Connection Established gg")
	}
	db.AutoMigrate(&dao.Reserva{})
	return SQL{
		db:       db,
		Database: config.Name,
	}

}

func (repository SQL) GetReservaById(id int64) (dao.Reserva, error) {
	var buscado dao.Reserva
	log.Println("ID: ", id)
	result := repository.db.Where("ID = ?", id).First(&buscado)
	log.Println("resultado: ", result)
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return dao.Reserva{}, fmt.Errorf("no reservation found with ID: %d", id)
		}
		return dao.Reserva{}, fmt.Errorf("error finding document: %v", result.Error)
	}
	return buscado, nil
}

func (repository SQL) InsertReserva(ctx context.Context, reserva dao.Reserva) (dao.Reserva, error) {
	result := repository.db.Create(&reserva)
	if result.Error != nil {
		log.Panic("Error creating the hotel")
		return reserva, fmt.Errorf("error inserting document:")
	}
	return reserva, nil
}

func (repository SQL) UpdateReserva(ctx context.Context, reserva dao.Reserva) (dao.Reserva, error) {
	var buscado dao.Reserva
	println("noches recibido: ", reserva.Noches)

	result := repository.db.Where("user = ? AND hotel = ?", reserva.User, reserva.Hotel).First(&buscado)

	if result.Error != nil {
		return dao.Reserva{}, fmt.Errorf("error finding document: %v", result.Error)
	}

	Newreserva := repository.db.Model(&buscado).Update("noches", reserva.Noches)
	if Newreserva.Error != nil {
		log.Panic("Error updating the hotel")
		return reserva, fmt.Errorf("error updating document:")
	}
	println("hotel recibido: ", buscado.Noches)

	return buscado, nil
}

func (repository SQL) DeleteReserva(ctx context.Context, reserva dao.Reserva) error {
	var buscado dao.Reserva
	println("user recibido: ", reserva.User)
	println("hotel recibido: ", reserva.Hotel)

	result := repository.db.Where("user = ? AND hotel = ?", reserva.User, reserva.Hotel).First(&buscado)

	if result.Error != nil {
		println("Error buscando la reserva")
		return fmt.Errorf("error buscando la reserva: %v", result.Error)
	}
	resul := repository.db.Delete(&buscado)
	if resul.Error != nil {
		println("Error deleting the hotel")
		return fmt.Errorf("error deleting document:")
	}
	return nil
}

func (repository SQL) GetMisReservasById(id int64) ([]dao.Reserva, error) {
	var buscado []dao.Reserva
	log.Println("ID: ", id)
	result := repository.db.Where("user = ?", id).Find(&buscado)
	log.Println("resultado: ", result)
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			return []dao.Reserva{}, fmt.Errorf("no reservation found with ID: %d", id)
		}
		return []dao.Reserva{}, fmt.Errorf("error finding document: %v", result.Error)
	}
	return buscado, nil
}
