package config

var AgeToBucket = ageBuckets()
var HourToBuckets = hourBuckets()

var ProbabilityOfClick = map[string]map[string]float32{
	"teen-morning|day": {
		"option1": 0.9,
		"option2": 0.1,
		"option3": 0.1,
	},
	"teen-evening|afternoon|night": {
		"option1": 0.5,
		"option2": 0.01,
		"option3": 0.01,
	},
	"young-morning|day|evening": {
		"option1": 0.01,
		"option2": 0.7,
		"option3": 0.1,
	},
	"young-afternoon|night": {
		"option1": 0.9,
		"option2": 0.1,
		"option3": 0.1,
	},
	"adult-morning|day|evening|afternoon|night": {
		"option1": 0.01,
		"option2": 0.01,
		"option3": 0.9,
	},
	"senior-morning": {
		"option1": 0.01,
		"option2": 0.1,
		"option3": 0.5,
	},
	"senior-day|evening|afternoon|night": {
		"option1": 0.01,
		"option2": 0.1,
		"option3": 0.9,
	},
}

func ageBuckets() (ageBucketsMap map[float32]string) {
	ageBucketsMap = map[float32]string{}

	for i := 0; i <= 13; i++ {
		ageBucketsMap[float32(i)] = "teen"
	}

	for i := 13; i <= 25; i++ {
		ageBucketsMap[float32(i)] = "young"
	}

	for i := 26; i <= 50; i++ {
		ageBucketsMap[float32(i)] = "adult"
	}

	for i := 51; i <= 200; i++ {
		ageBucketsMap[float32(i)] = "senior"
	}
	return

}

func hourBuckets() (hourBuckets map[float32]string) {
	hourBuckets = map[float32]string{}

	for i := 7; i <= 10; i++ {
		hourBuckets[float32(i)] = "morning"
	}

	for i := 11; i <= 15; i++ {
		hourBuckets[float32(i)] = "day"
	}

	for i := 16; i <= 20; i++ {
		hourBuckets[float32(i)] = "evening"
	}

	for i := 21; i <= 23; i++ {
		hourBuckets[float32(i)] = "afternoon"
	}

	for i := 0; i <= 6; i++ {
		hourBuckets[float32(i)] = "night"
	}

	return

}
