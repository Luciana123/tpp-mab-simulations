package reward

import (
	"math"
	"math/rand"
	"sim-city/config"
	"strings"
	"time"
)

type Visit struct {
	Context []float32
	Age     int
	Hour    int
	Time    time.Time
}

func SimulateVisit() (visit Visit) {

	age := rand.Intn(120)

	currentDate := time.Now().UTC().Truncate(24 * time.Hour)
	randomTime := currentDate.Add(time.Duration(rand.Intn(24)) * time.Hour)
	randomDuration := time.Duration(rand.Intn(24))*time.Hour +
		time.Duration(rand.Intn(60))*time.Second +
		time.Duration(rand.Intn(1000))*time.Millisecond
	randomTime = randomTime.Add(randomDuration)

	hourSin, hourCos := encodeHourAsSinCos(randomTime)

	context := []float32{float32(age), float32(hourSin), float32(hourCos)}

	visit = Visit{
		Context: context,
		Hour:    randomTime.Hour(),
		Age:     age,
		Time:    randomTime,
	}

	return
}

func SimulateVisitReward(visit Visit, armSelected string) (reward int) {
	ageBucket := config.AgeToBucket[float32(visit.Age)]
	hourBucket := config.HourToBuckets[float32(visit.Hour)]

	for key, val := range config.ProbabilityOfClick {
		if strings.Contains(key, ageBucket) && strings.Contains(key, hourBucket) {
			armProbabilityMap := val
			reward = simulateReward(float64(armProbabilityMap[armSelected]))
		}
	}
	return
}

func encodeHourAsSinCos(timestamp time.Time) (float64, float64) {
	hour := float64(timestamp.Hour())
	hourSin := math.Sin(2 * math.Pi * hour / 24)
	hourCos := math.Cos(2 * math.Pi * hour / 24)
	return hourSin, hourCos
}

func simulateReward(probability float64) int {
	if rand.Float64() < probability {
		return 1
	}
	return 0
}
