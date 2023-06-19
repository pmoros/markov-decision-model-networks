package model

import "fmt"

type Probability float32

type Policy [][]Direction

type Coords []int

func GetPolicyDirectionFromCoords(coords Coords, policy Policy) *Direction {
	if coords[0] >= len(policy[0]) || coords[1] >= len(policy) {
		fmt.Println("cell exceed policy bounds")
		return nil
	}

	for vCoord, row := range policy {
		for hCoord, cell := range row {
			if coords[0] == hCoord && coords[1] == vCoord {
				return &cell
			}
		}
	}
	return nil
}
