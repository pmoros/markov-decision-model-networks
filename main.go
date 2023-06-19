package main

import (
	"flag"
	"fmt"

	"github.com/pmoros/markov-decision-model-networks/pkg/model"
	"github.com/pmoros/markov-decision-model-networks/pkg/policy"
	"github.com/pmoros/markov-decision-model-networks/pkg/util"
)

func main() {
	// Define a flag for the policyType
	policyTypeFlag := flag.Int("policyType", 0, "policy type (0-Default, 1-TypeA, 2-TypeB, 3-TypeC)")
	timesToRunFlag := flag.Int("timesToRun", 1, "times to run the simulation")

	// Parse the command-line flags
	flag.Parse()

	// Convert the policyType from int to PolicyType enum
	policyType := model.PolicyType(*policyTypeFlag)

	// Print whether the policy is fixed, randomized or iterated
	switch policyType {
	case model.Fixed:
		fmt.Println("Fixed policy")
	case model.Randomized:
		fmt.Println("Randomized policy")
	case model.Iterated:
		fmt.Println("Iterated policy")
	default:
		fmt.Println("Unknown policy")
	}

	for i := 0; i < *timesToRunFlag; i++ {

		// policyType := model.Fixed // model.Fixed, model.Randomized or model.Iterated
		policy := policy.GeneratePolicy(policyType)
		cells := [][]model.Cell{
			{
				model.NewCell(0, model.Clear, -0.04),
				model.NewCell(1, model.Clear, -0.04),
				model.NewCell(2, model.Clear, -0.04),
				model.NewCell(3, model.Goal, 1.0),
			},
			{
				model.NewCell(4, model.Clear, -0.04),
				model.NewCell(5, model.Blocked, 0),
				model.NewCell(6, model.Clear, -0.04),
				model.NewCell(7, model.Goal, -1.0),
			},
			{
				model.NewCell(8, model.Clear, -0.04),
				model.NewCell(9, model.Clear, -0.04),
				model.NewCell(10, model.Clear, -0.04),
				model.NewCell(11, model.Clear, -0.04),
			},
		}
		agent := model.Agent{
			TransitionModel: model.Transition{
				model.Forward:     0.8,
				model.RotateLeft:  0.1,
				model.RotateRight: 0.1,
				model.Back:        0.0, // no back movements on this example
			},
			InitialCell: model.Coords{0, 2},
			Energy:      0,
		}

		scenarioGrid := util.CreateScenario(cells, policy, agent)
		energyLevel := util.RunScenario(scenarioGrid)

		fmt.Println(policy)
		fmt.Println(energyLevel)
	}
}
