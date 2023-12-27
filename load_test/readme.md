# Load test for Jackpot

Uses k6, needs an experiment to be previously created.

```
Experiment id: 2-feats-eg
Arms: 'option1', 'option2', 'option3', 'option4'
Features: "age", "is_female"
```

# Description

We need a process that sends events during the hole load test to be realistic,
we use a process called event_sender, which is a program written in Go, that emulates 
concurrent users sending reward messages for an experiment like the one described above.


Then, the load test run with k6, which is a tool created by grafana to run load tests,
we have written two scenarios
* reward_predictor: test the model prediction, makes requests to the load_tests service.
So in this case you only need the reward-predictor service up and running.
* arm_selector: test the entire framework from the arm_selector service.