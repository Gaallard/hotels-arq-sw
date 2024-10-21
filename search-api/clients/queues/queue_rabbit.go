package queues

import (
	"fmt"

	// "hotels-api/domain/hotels"
	"log"

	"github.com/streadway/amqp"
)

type RabbitConfig struct {
	Username  string
	Password  string
	Host      string
	Port      string
	QueueName string
}

type Rabbit struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewRabbit(config RabbitConfig) Rabbit {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", config.Username, config.Password, config.Host, config.Port))
	if err != nil {
		log.Fatalf("error getting Rabbit connection: %w", err)
	}
	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("error creating Rabbit channel: %w", err)
	}
	queue, err := channel.QueueDeclare(config.QueueName, false, false, false, false, nil)

	return Rabbit{
		connection: connection,
		channel:    channel,
		queue:      queue,
	}
}

// prueba para ver si llega mensaje

func (p Rabbit) ConsumeCola() {
	msgs, err := p.channel.Consume(
		p.queue.Name,
		"",
		true,
		false,
		false,
		true,
		nil,
	)

	if err != nil {
		log.Printf("Error al recibir mensages")
	}

	d := <-msgs
	log.Printf("Mensage recibido: %s", d.Body)

}

//en un publisher, ver
/*
func (queue Rabbit) Receive(hotelNew hotels.HotelNew) error {

	bytes, err := json.Marshal(hotelNew)
	if err != nil {
		return fmt.Errorf("error marshaling Rabbit hotelNew: %w", err)
	}

	if err := queue.channel(
		"", queue.queue.Name, false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		}); err != nil {
		return fmt.Errorf("error publishing to Rabbit: %w", err)
	}

	return nil
}
*/
