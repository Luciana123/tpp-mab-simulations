# simulations
Simulations of multi armed bandit models made in Go (proof of concept)

# Goals
This project test the Jackpot framework in two separated ways:
1. Load testing (load_test)
   Run a load_test with k6 to benchmark the endpoints of Jackpot framework.
2. Simulation (sim_city)
    Simulate the users with a certain probability of clicking the arm, to test 
    Jackpot ability to find patterns and select the correct arms.
