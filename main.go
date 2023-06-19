package main

import (
	"fmt"

	"github.com/pmoros/markov-decision-model-networks/pkg/model"
	"github.com/pmoros/markov-decision-model-networks/pkg/util"
)

func main() {
	policyType := model.Randomized
	fmt.Println("MDP")
	scenarioGrid := util.CreateScenario(policyType)
	util.RunScenario(scenarioGrid)
}
