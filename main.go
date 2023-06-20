package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/pmoros/markov-decision-model-networks/pkg/model"
	"github.com/pmoros/markov-decision-model-networks/pkg/policy"
	"github.com/pmoros/markov-decision-model-networks/pkg/util"
)

func main() {
	// Define a flag for the policyType
	policyTypeFlag := flag.Int("policyType", 0, "policy type (0-Default, 1-TypeA, 2-TypeB, 3-TypeC)")
	timesToRunFlag := flag.Int("timesToRun", 1, "times to run the simulation")
	showPolicyFlag := flag.Bool("showPolicy", false, "show the policy")

	// Parse the command-line flags
	flag.Parse()

	// Convert the policyType from int to PolicyType enum
	policyType := model.PolicyType(*policyTypeFlag)

	// Print whether the policy is fixed, randomized or iterated
	if *showPolicyFlag {
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
	}

	var wg sync.WaitGroup
	wg.Add(*timesToRunFlag)

	for i := 0; i < *timesToRunFlag; i++ {
		go func() {
			// policyType := model.Fixed // model.Fixed, model.Randomized or model.Iterated
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

			policyModel := model.Policy{}
			if policyType == model.Iterated {
				policyAux := policy.GeneratePolicy(model.Fixed)
				policyModel = policy.GenerateIteratedPolicy(cells, agent, policyAux)
			} else {
				policyModel = policy.GeneratePolicy(policyType)
			}

			scenarioGrid := util.CreateScenario(cells, policyModel, agent)
			energyLevel := util.RunScenario(scenarioGrid)

			if *showPolicyFlag {
				fmt.Println(policyModel)
			}
			fmt.Println(energyLevel)

			wg.Done()
		}()
	}

	wg.Wait()
}
