package hotels

import (
	"context"
	"fmt"
	hotelsDAO "hotels-api/dao/hotels"
	hotelsDomain "hotels-api/domain/hotels"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

const (
	connectionURI = "mongodb://%s:%s" // %s es marcador de puesto para el host y el puerto
)

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

func (repository Mongo) GetHotelByID(ctx context.Context, id int64) (hotelsDAO.Hotel, error) {
	// Get from MongoDB
	result := repository.client.Database(repository.database).Collection(repository.collection).FindOne(ctx, bson.M{"id": id})
	if result.Err() != nil {
		return hotelsDAO.Hotel{}, fmt.Errorf("error finding document: %w", result.Err())
	}

	// Convert document to DAO
	var hotelDAO hotelsDAO.Hotel
	if err := result.Decode(&hotelDAO); err != nil {
		return hotelsDAO.Hotel{}, fmt.Errorf("error decoding result: %w", err)
	}
	return hotelDAO, nil
}

func (repository Mongo) InsertHotel(ctx context.Context, hotel hotelsDAO.Hotel) error {
	result, err := repository.client.Database(repository.database).Collection(repository.collection).InsertOne(ctx, hotel)
	if err != nil {
		return fmt.Errorf("Error inserting new hotel: %w", err)
	}

	// Opcionalmente, puedes usar el resultado aquí
	fmt.Printf("Inserted hotel with ID: %v\n", result.InsertedID)

	return nil
}

func (repository Mongo) DeleteHotel(ctx context.Context, id int64) error {
	_, err := repository.client.Database(repository.database).Collection(repository.collection).DeleteOne(ctx, bson.M{"id": id})

	if err != nil {
		return fmt.Errorf("Error finding hotel: %w", err)
	}

	return nil
}

func (repository Mongo) UpdateHotel(ctx context.Context, id int64, hotelDomain hotelsDomain.Hotel) (hotelsDomain.Hotel, error) {
	_, err := repository.client.Database(repository.database).Collection(repository.collection).UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": hotelDomain})

	if err != nil {
		return hotelsDomain.Hotel{}, fmt.Errorf("Error finding hotel: %w", err)
	}

	return hotelDomain, nil
}
