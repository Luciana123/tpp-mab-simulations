# Jackpot Sim City

Simulation of client using jackpot framework with some rules:

This experiment must be loaded:
```
Experiment id: sim-city-experiment
Arms: 'option1', 'option2', 'option3'
Features: "age", "hour_sin", "hour_cos"
```

Features:
1. Age: The age of the user.
2. hour_sin / hour_cos: Given the circular nature of the hour of the day, we can't use the hour directly
as a feature (for example 12 is closer to 0), so the trick to encode it is treated them as a circle, using cos and
sin functions.

Buckets:
1. Age buckets
   * Age 0 to 13: teen
   * Age 14 to 25: young
   * Age 25 to 50: adult
   * Age 50 to 100: senior

2. Hour Buckets:
   * 7 - 10: morning
   * 10 - 15: day
   * 15 - 20: evening
   * 20 - 22: afternoon
   * 22 - 7: night

The probability of clicking each option follow these rules:
* teen - morning/day : 
  * option1: 0.9
  * option2: 0.1
  * option3: 0.1
* teen - evening/afternoon/night:
  * option1: 0.5
  * option2: 0.01
  * option3: 0.01
* young - morning/day/evening:
  * option1: 0.01
  * option2: 0.7
  * option3: 0.1
* young - afternoon/night:
  * option1: 0.01
  * option2: 0.9
  * option3: 0.1
* adult - morning/day/evening/afternoon/night:
  * option1: 0.01
  * option2: 0.01
  * option3: 0.9
* senior - morning:
  * option1: 0.01
  * option2: 0.1
  * option3: 0.5
* senior day/evening/afternoon/night:
  * option1: 0.01
  * option2: 0.1
  * option3: 0.9
  