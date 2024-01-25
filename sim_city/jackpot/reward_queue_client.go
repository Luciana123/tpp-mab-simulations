package jackpot

import (
	"github.com/streadway/amqp"
	"sim-city/reward"
)

type RewardQueueClient struct {
	Channel    *amqp.Channel
	connection *amqp.Connection
}

func NewRewardQueue() *RewardQueueClient {
	// Connect to RabbitMQ server
	//conn, err := amqp.Dial("amqps://jackpotAdmin:jackpotPassword@b-af17891f-ed2f-4168-a831-6a41ce175372.mq.us-east-2.amazonaws.com:5671/")
	conn, err := amqp.Dial("amqp://jackpotAdmin:jackpotPassword@localhost:5672/")

	reward.FailOnError(err, "Failed to connect to RabbitMQ")

	// Create a channel
	ch, err := conn.Channel()
	reward.FailOnError(err, "Failed to open a channel")

	// Declare a queue
	queueName := "reward-queue"
	_, err = ch.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,
	)
	reward.FailOnError(err, "Failed to declare a queue")

	return &RewardQueueClient{
		Channel:    ch,
		connection: conn,
	}
}

func (r *RewardQueueClient) SendMessage(jsonMessage []byte) {
	err := r.Channel.Publish(
		"",
		"reward-queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonMessage,
		})

	reward.FailOnError(err, "Failed to publish a message")
}

func (r *RewardQueueClient) Close() {
	r.Channel.Close()
	r.connection.Close()
}
