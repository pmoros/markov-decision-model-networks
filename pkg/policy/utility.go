package policy

import (
	"fmt"
	"strconv"

	"github.com/pmoros/markov-decision-model-networks/pkg/model"
)

type MDP struct {
	Policy     model.Policy
	Cells      [][]model.Cell
	Agent      model.Agent
	Discount   float64
	Iterations int
}

// PolicyIteration runs the policy iteration algorithm
func PolicyIteration(mdp MDP) model.Policy {
	// Initialization
	V := make(map[string]float64)
	policy := mdp.Policy

	// Policy evaluation
	for i := 0; i < mdp.Iterations; i++ {
		V = EvaluatePolicy(mdp, policy, V)
	}

	// Policy improvement
	newPolicy := ImprovePolicy(mdp, policy, V)

	for !IsEqualPolicy(policy, newPolicy) {
		policy = newPolicy
		V = EvaluatePolicy(mdp, policy, V)
		newPolicy = ImprovePolicy(mdp, policy, V)
	}

	return newPolicy
}

// EvaluatePolicy performs policy evaluation
func EvaluatePolicy(mdp MDP, policy model.Policy, V map[string]float64) map[string]float64 {
	// Reset V for the new iteration
	for _, row := range mdp.Cells {
		for _, cell := range row {
			if cell.Type == model.Blocked {
				V[coordsToString(model.Coords{cell.ID})] = cell.Reward
			} else {
				V[coordsToString(model.Coords{cell.ID})] = 0.0
			}
		}
	}

	for i := 0; i < mdp.Iterations; i++ {
		newV := make(map[string]float64)
		for _, row := range mdp.Cells {
			for _, cell := range row {
				if cell.Type == model.Blocked || cell.Type == model.Goal {
					continue
				}

				state := model.Coords{cell.ID}
				action := policy[state[0]][state[1]]
				reward := cell.Reward
				Vs := CalculateExpectedValue(mdp, state, action, V)
				newV[coordsToString(state)] = reward + mdp.Discount*Vs
			}
		}

		V = newV
	}

	return V
}

// CalculateExpectedValue calculates the expected value for a given state-action pair
func CalculateExpectedValue(mdp MDP, state model.Coords, action model.Direction, V map[string]float64) float64 {
	expectedValue := 0.0

	transitions := mdp.Agent.TransitionModel

	// Calculate the expected value for each possible transition
	for movement, probability := range transitions {
		nextState := GetNextState(state, movement)
		value := V[coordsToString(nextState)]
		expectedValue += float64(probability) * value
	}

	return expectedValue
}

// GetNextState returns the next state based on the current state and movement
func GetNextState(state model.Coords, movement model.Movement) model.Coords {
	switch movement {
	case model.Forward:
		return model.Coords{state[0] - 1, state[1]}
	case model.RotateLeft:
		return model.Coords{state[0], state[1] - 1}
	case model.RotateRight:
		return model.Coords{state[0], state[1] + 1}
	default:
		return state
	}
}

// ImprovePolicy performs policy improvement
func ImprovePolicy(mdp MDP, policy model.Policy, V map[string]float64) model.Policy {
	newPolicy := make(model.Policy, len(policy))
	for i := range policy {
		newPolicy[i] = make([]model.Direction, len(policy[i]))
		copy(newPolicy[i], policy[i])
	}

	for _, row := range mdp.Cells {
		for _, cell := range row {
			if cell.Type == model.Blocked || cell.Type == model.Goal {
				continue
			}

			state := model.Coords{cell.ID}
			bestAction := GetBestAction(mdp, state, V)
			newPolicy[state[0]][state[1]] = bestAction
		}
	}

	return newPolicy
}

// GetBestAction returns the best action for a given state based on its expected values
func GetBestAction(mdp MDP, state model.Coords, V map[string]float64) model.Direction {
	bestAction := mdp.Policy[state[0]][state[1]]
	bestValue := CalculateExpectedValue(mdp, state, bestAction, V)

	transitions := mdp.Agent.TransitionModel

	// Find the action with the highest expected value
	for movement, _ := range transitions {
		action := model.Direction(movement)
		value := CalculateExpectedValue(mdp, state, action, V)
		if value > bestValue {
			bestAction = action
			bestValue = value
		}
	}

	return bestAction
}

// IsEqualPolicy checks if two policies are equal
func IsEqualPolicy(policy1 model.Policy, policy2 model.Policy) bool {
	for i := range policy1 {
		for j := range policy1[i] {
			if policy1[i][j] != policy2[i][j] {
				return false
			}
		}
	}

	return true
}

// PrintPolicy prints the given policy
func PrintPolicy(policy model.Policy) {
	for _, row := range policy {
		for _, action := range row {
			fmt.Printf("%v\t", action)
		}
		fmt.Println()
	}
}

// Converts Coords to string representation
func coordsToString(coords model.Coords) string {
	return strconv.Itoa(coords[0]) + "-" + strconv.Itoa(coords[1])
}
