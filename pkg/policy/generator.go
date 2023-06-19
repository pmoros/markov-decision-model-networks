package policy

import (
	"math/rand"

	"github.com/pmoros/markov-decision-model-networks/pkg/model"
)

func GeneratePolicy(policyType model.PolicyType) model.Policy {
	switch policyType {
	case model.Fixed:
		return GenerateFixedPolicy()
	case model.Randomized:
		return GenerateRandomizedPolicy()
	case model.Iterated:
		return GenerateIteratedPolicy()
	default:
		return nil
	}
}

func GenerateFixedPolicy() model.Policy {
	return model.Policy{
		{model.Right, model.Right, model.Left, model.Idle},
		{model.Up, model.Idle, model.Down, model.Idle},
		{model.Right, model.Down, model.Right, model.Up},
	}
}

func GenerateRandomizedPolicy() model.Policy {
	rows := 3
	columns := 4
	policy := make(model.Policy, rows)
	for i := 0; i < rows; i++ {
		policy[i] = make([]model.Direction, columns)
		for j := 0; j < columns; j++ {
			if (i == 0 && j == 3) || (i == 1 && j == 3) { // Specify the position where model.Idle should appear
				policy[i][j] = model.Direction(model.Idle)
			} else {
				randomValue := model.Direction(rand.Intn(int(model.Left))) // Generate a random value within the enum range excluding model.Idle
				policy[i][j] = randomValue
			}
		}
	}
	return policy
}

func GenerateIteratedPolicy() model.Policy {
	return model.Policy{
		{model.Right, model.Right, model.Left, model.Idle},
		{model.Up, model.Idle, model.Down, model.Idle},
		{model.Right, model.Down, model.Right, model.Up},
	}
}
