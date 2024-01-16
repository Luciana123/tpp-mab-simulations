package reward

import (
	"fmt"
	"sim-city/config"
	"testing"
)

func TestSimulation(t *testing.T) {

	results := map[string]map[string]int{}

	for i := 0; i < 1000000; i++ {
		v := SimulateVisit()
		r1 := SimulateVisitReward(v, "option1")
		r2 := SimulateVisitReward(v, "option2")
		r3 := SimulateVisitReward(v, "option3")

		ageBuc := config.AgeToBucket[float32(v.age)]
		hourBuc := config.HourToBuckets[float32(v.hour)]

		key := fmt.Sprintf("%s-%s", ageBuc, hourBuc)

		rewardByArmMap, ok := results[key]

		if !ok {
			results[key] = map[string]int{
				"option1": r1,
				"option2": r2,
				"option3": r3,
			}
		} else {
			rewardByArmMap["option1"] = results[key]["option1"] + r1
			rewardByArmMap["option2"] = results[key]["option2"] + r2
			rewardByArmMap["option3"] = results[key]["option3"] + r3
		}
	}

}
