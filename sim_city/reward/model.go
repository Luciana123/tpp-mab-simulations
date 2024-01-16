package reward

import "time"

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
