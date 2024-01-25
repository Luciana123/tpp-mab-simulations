package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"sim-city/config"
	"sim-city/jackpot"
	"sim-city/model"
	"sim-city/reward"
	"sim-city/sim_results"
	"strconv"
	"strings"
	"time"
)

func runSimulation(duration time.Duration) {

	sim_results.CreateFile("sim_city_results.csv")

	sim_results.AddLine("sim_city_results.csv", "time,time_sim,age_sim,hour_sim,arm_selected,reward")

	r := jackpot.NewRewardQueue()
	jackpotClient := jackpot.NewHTTPClient("http://localhost:8090")
	//jackpotClient := jackpot.NewHTTPClient("https://arm-selector.onrender.com")

	startTime := time.Now()
	for time.Since(startTime) < duration {

		sessId := uuid.New().String()
		simVisit := reward.SimulateVisit()

		armRequest := map[string]float32{
			"age":      simVisit.Context[0],
			"hour_sin": simVisit.Context[1],
			"hour_cos": simVisit.Context[2],
		}

		payload, _ := json.Marshal(armRequest)

		result, err := jackpotClient.Post(fmt.Sprintf("/api/v1/arm/selection/%s", config.ExperimentId), payload)

		var armSelectorResponse model.ArmSelectorResponse
		err = json.Unmarshal([]byte(result), &armSelectorResponse)
		if err != nil {
			return
		}

		messages := model.RewardMessage{
			Operation: "create",
			Reward: model.Reward{
				ArmSelected:  armSelectorResponse.Arm,
				ExperimentID: config.ExperimentId,
				SessionID:    sessId,
				Context: map[string]string{
					"age":      strconv.FormatFloat(float64(simVisit.Context[0]), 'f', -1, 32),
					"hour_sin": strconv.FormatFloat(float64(simVisit.Context[1]), 'f', -1, 32),
					"hour_cos": strconv.FormatFloat(float64(simVisit.Context[2]), 'f', -1, 32),
				},
				Timestamp: simVisit.Time,
			},
		}

		jsonMessage, _ := json.Marshal(messages)
		r.SendMessage(jsonMessage)

		reward := reward.SimulateVisitReward(simVisit, armSelectorResponse.Arm)

		if reward == 1 {
			messages := model.RewardMessage{
				Operation: "update",
				Reward: model.Reward{
					ArmSelected:  armSelectorResponse.Arm,
					ExperimentID: config.ExperimentId,
					SessionID:    sessId,
					Context: map[string]string{
						"age":      strconv.FormatFloat(float64(simVisit.Context[0]), 'f', -1, 32),
						"hour_sin": strconv.FormatFloat(float64(simVisit.Context[1]), 'f', -1, 32),
						"hour_cos": strconv.FormatFloat(float64(simVisit.Context[2]), 'f', -1, 32),
					},
					Timestamp: simVisit.Time,
					Reward:    reward,
				},
			}

			jsonMessage, _ := json.Marshal(messages)
			r.SendMessage(jsonMessage)
		}

		dateFormat := "2006-01-02 15:04:05"

		currentDate := time.Now()
		iterationResult := strings.Join([]string{currentDate.Format(dateFormat), simVisit.Time.Format(dateFormat),
			strconv.Itoa(simVisit.Age), strconv.Itoa(simVisit.Hour),
			armSelectorResponse.Arm, strconv.Itoa(reward)}, ",")

		sim_results.AddLine("sim_city_results.csv", iterationResult)

	}

}

func main() {

	// not needed
	godotenv.Load(".env")

	// Should I create the experiment here?

	print("sim_city starte running, running time: 10 minutes...")
	runSimulation(10 * time.Minute)

	// And delete the experiment here?
}
