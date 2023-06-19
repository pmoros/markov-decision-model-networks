package util

import (
	"math/rand"
	"time"

	"github.com/pmoros/markov-decision-model-networks/pkg/model"
)

func CreateScenario(cells [][]model.Cell, policy model.Policy, agent model.Agent) model.Grid {
	return model.Grid{
		Cells:  cells,
		Policy: policy,
		Agent:  agent,
	}
}

func RunScenario(grid model.Grid) float64 {
	// TODO: info should be logged to different output
	// set initial position for agent
	grid.Agent.CurrentCell = grid.Agent.InitialCell

	// as long as the agent is not on a goal
	for model.GetCellFromCoords(grid.Agent.CurrentCell, grid).Type != model.Goal {
		// get the movement estabilished by policy
		intendedMovement := model.GetPolicyDirectionFromCoords(grid.Agent.CurrentCell, grid.Policy)

		// then calculate the final direction with the stochastic transition model
		nextDirection := calculateNextStochasticMovement(intendedMovement, grid.Agent.TransitionModel)

		// fmt.Printf("Current cell is %d and movement is %d\n",
		// 	model.GetCellFromCoords(grid.Agent.CurrentCell, grid).ID,
		// 	nextDirection)

		// deduct energy from agent
		grid.Agent.Energy += model.GetCellFromCoords(grid.Agent.CurrentCell, grid).Reward

		// shift agent's position
		grid.Agent.CurrentCell = moveAgent(grid.Agent.CurrentCell, nextDirection, grid)
	}

	// fmt.Printf("Final cell is %d\n",
	// 	model.GetCellFromCoords(grid.Agent.CurrentCell, grid).ID)

	// when simulation is done
	// fmt.Println(grid.Agent)
	return grid.Agent.Energy
}

func calculateNextStochasticMovement(intendedMovement *model.Direction, transitionModel model.Transition) model.Direction {
	// get array of samples made from transition model
	movementWeightedArray := GetWeightedArray(transitionModel)

	// set random seed with UNIX local time (probably should improve)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// then pick a random element from the sample array
	agentRotation := movementWeightedArray[rand.Intn(len(movementWeightedArray))]

	// add rotation to intended movement, so we can get the new stochastic direction
	return model.Direction((int(*intendedMovement) + int(agentRotation)) % 4)
}

func moveAgent(currentCell model.Coords, direction model.Direction, grid model.Grid) model.Coords {
	var destination model.Coords

	switch direction {
	case model.Up:
		destination = model.Coords{currentCell[0], currentCell[1] - 1}

		// if agent hits upper wall or if destination is blocked
		if currentCell[1] == 0 ||
			model.GetCellFromCoords(destination, grid).Type == model.Blocked {
			return currentCell
		}

		// if not, subtract one to Y coordinate
		return destination

	case model.Right:
		destination = model.Coords{currentCell[0] + 1, currentCell[1]}

		// if agent hits rightmost wall or if destination is blocked
		if currentCell[0] == len(grid.Cells[0])-1 ||
			model.GetCellFromCoords(destination, grid).Type == model.Blocked {
			return currentCell
		}

		// if not, add one to X coordinate
		return destination

	case model.Down:
		destination = model.Coords{currentCell[0], currentCell[1] + 1}

		// if agent hits lower wall or if destination is blocked
		if currentCell[1] == len(grid.Cells)-1 ||
			model.GetCellFromCoords(destination, grid).Type == model.Blocked {
			return currentCell
		}

		// if not, add one to Y coordinate
		return destination

	case model.Left:
		destination = model.Coords{currentCell[0] - 1, currentCell[1]}

		// if agent hits leftmost wall or if destination is blocked
		if currentCell[0] == 0 ||
			model.GetCellFromCoords(destination, grid).Type == model.Blocked {
			return currentCell
		}

		// if not, subtract one to X coordinate
		return destination
	}

	return currentCell
}
