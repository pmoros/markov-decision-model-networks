package util

import (
	"github.com/pmoros/markov-decision-model-networks/pkg/model"
)

func GetWeightedArray(transition model.Transition) []model.Movement {
	weightedArray := make([]model.Movement, 100)
	sampleCounter := 0

	for movement, prob := range transition {
		numberOfSamples := int(prob * 100)

		for ; sampleCounter < numberOfSamples; sampleCounter++ {
			weightedArray[sampleCounter] = movement
		}
	}

	return weightedArray
}
