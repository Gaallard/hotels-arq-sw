package reservas

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repository SQL) CheckHotelExists(idHotel string) (bool, error) {
	return false, nil
}

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
	result, err := repository.GetReservaById(reserva.ID)
	log.Println("ID: ", reserva.ID)
	if err != nil {
		log.Panic("Error reversa no exists")
		return reserva, fmt.Errorf("error reserva doesnt exists:")
	}
	Newreserva := repository.db.Model(&result).Update("noches", reserva.Noches)
	if Newreserva.Error != nil {
		log.Panic("Error updating the hotel")
		return reserva, fmt.Errorf("error updating document:")
	}
	return result, nil
}

const (
	connectionURI = "mongodb://%s:%s" // %s es marcador de puesto para el host y el puerto
)

type MongoConfig struct {
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Collection string
}

type Mongo struct {
	client     *mongo.Client //cliente que contiene la conexion
	database   string        //nombre de la base de datos
	collection string        //coleccion
}

func NewMongo(config MongoConfig) Mongo {
	credentials := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}

	ctx := context.Background()                                 // para manejar cancelaciones o límites de tiempo en las operaciones.
	uri := fmt.Sprintf(connectionURI, config.Host, config.Port) //Construye la URI de conexión utilizando el host y el puerto.
	cfg := options.Client().ApplyURI(uri).SetAuth(credentials)  //Configura las opciones del cliente de MongoDB, incluyendo la URI y las credenciales de autenticación.

	client, err := mongo.Connect(ctx, cfg)
	if err != nil {
		log.Panicf("error connecting to mongo DB: %v", err)
	}

	return Mongo{
		client:     client,
		database:   config.Database,
		collection: config.Collection,
	}
}

func (m Mongo) GetReservaById(id int64) (dao.Reserva, error) {
	return dao.Reserva{}, nil
}

func (m Mongo) InsertReserva(ctx context.Context, reserva dao.Reserva) (dao.Reserva, error) {
	return dao.Reserva{}, nil
}

func (m Mongo) UpdateReserva(ctx context.Context, reserva dao.Reserva) (dao.Reserva, error) {
	return dao.Reserva{}, nil
}

func (m Mongo) CheckHotelExists(idHotel string) (bool, error) {
	collection := m.client.Database(m.database).Collection(m.collection)
	objectID, err := primitive.ObjectIDFromHex(idHotel)
	if err != nil {
		return false, fmt.Errorf("Hotel not found: %s", err)
	}

	filter := bson.M{"_id": objectID}
	var result bson.M
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, fmt.Errorf("error finding hotel: %s", err)
	}
	return true, nil
}
