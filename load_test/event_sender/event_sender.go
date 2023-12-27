package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

const TimeBetweenMessagesMs = 1000
const ConcurrentUsers = 100
const ExperimentId = "2-feats-eg"

var Arms = []string{"option1", "option2", "option3", "option4"}
var Weights = []int{50, 20, 10, 20}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func chooseRandomWeighted(arms []string, weights []int) string {
	totalWeight := 0
	for _, w := range weights {
		totalWeight += w
	}

	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(totalWeight)

	sum := 0
	for i, w := range weights {
		sum += w
		if r < sum {
			return arms[i]
		}
	}

	// Fallback, should not reach here
	return arms[0]
}

type Reward struct {
	ArmSelected  string            `json:"arm_selected"`
	ExperimentID string            `json:"experiment_id"`
	SessionID    string            `json:"session_id"`
	Reward       int               `json:"reward"`
	Context      map[string]string `json:"context"`
	Timestamp    time.Time         `json:"timestamp"`
}

type RewardMessage struct {
	Operation string `json:"operation"`
	Reward    Reward `json:"reward"`
}

func buildMessages() []RewardMessage {
	arm := chooseRandomWeighted(Arms, Weights)
	sessId := uuid.New().String()
	reward, _ := strconv.Atoi(chooseRandomWeighted([]string{"0", "1"}, []int{50, 50}))
	currentUTC := time.Now().UTC()

	// Context features:
	age := rand.Intn(100-13+1) + 13
	isFemale := chooseRandomWeighted([]string{"0", "1"}, []int{40, 60})

	r1 := RewardMessage{
		Operation: "create",
		Reward: Reward{
			ArmSelected:  arm,
			ExperimentID: ExperimentId,
			SessionID:    sessId,
			Context: map[string]string{
				"age":       strconv.Itoa(age),
				"is_female": isFemale,
			},
			Timestamp: currentUTC,
		},
	}

	if reward == 1 {
		r2 := RewardMessage{
			Operation: "update",
			Reward: Reward{
				ArmSelected:  arm,
				ExperimentID: ExperimentId,
				SessionID:    sessId,
				Timestamp:    currentUTC,
				Reward:       reward,
			},
		}
		return []RewardMessage{r1, r2}
	} else {

		return []RewardMessage{r1}

	}

}

func sendMessages(wg *sync.WaitGroup) {

	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://jackpotAdmin:jackpotPassword@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

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
	failOnError(err, "Failed to declare a queue")

	for {
		messages := buildMessages()

		for _, m := range messages {
			jsonMessage, _ := json.Marshal(m)
			err := ch.Publish(
				"",
				queueName,
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        jsonMessage,
				})

			failOnError(err, "Failed to publish a message")
			time.Sleep(TimeBetweenMessagesMs * time.Millisecond)
		}

	}

	wg.Done()
}

func main() {

	// not needed
	godotenv.Load(".env")

	var wg sync.WaitGroup

	for i := 1; i <= ConcurrentUsers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Simulate user activity
			sendMessages(&wg)
		}()
	}

	wg.Wait()

}
