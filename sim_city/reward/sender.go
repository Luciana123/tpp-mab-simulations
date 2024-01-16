package reward

import (
	"github.com/google/uuid"
	"log"
	"sim-city/config"
	"time"
)

func FailOnError(err error,
	msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func buildMessages(arm string, reward int, context map[string]string) []RewardMessage {

	sessId := uuid.New().String()
	currentUTC := time.Now().UTC()

	r1 := RewardMessage{
		Operation: "create",
		Reward: Reward{
			ArmSelected:  arm,
			ExperimentID: config.ExperimentId,
			SessionID:    sessId,
			Context:      context,
			Timestamp:    currentUTC,
		},
	}

	if reward == 1 {
		r2 := RewardMessage{
			Operation: "update",
			Reward: Reward{
				ArmSelected:  arm,
				ExperimentID: config.ExperimentId,
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
