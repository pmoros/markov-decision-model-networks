package util

import (
	"fmt"
	"github.com/pmoros/markov-decision-model-networks/pkg/model"
	"math/rand"
	"time"
)

func CreateScenario() model.Grid {
	return model.Grid{
		Cells: [][]model.Cell{
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
		},
		Policy: model.Policy{
			{model.Right, model.Right, model.Left, model.Idle},
			{model.Up, model.Idle, model.Down, model.Idle},
			{model.Right, model.Down, model.Right, model.Up},
		},
		Agent: model.Agent{
			TransitionModel: model.Transition{
				model.Forward:     0.8,
				model.RotateLeft:  0.1,
				model.RotateRight: 0.1,
				model.Back:        0.0, // no back movements on this example
			},
			InitialCell: model.Coords{0, 2},
			Energy:      0,
		},
	}
}

func RunScenario(grid model.Grid) {
	// set initial position for agent
	grid.Agent.CurrentCell = grid.Agent.InitialCell

	// as long as the agent is not on a goal
	for model.GetCellFromCoords(grid.Agent.CurrentCell, grid).Type != model.Goal {
		// get the movement estabilished by policy
		intendedMovement := model.GetPolicyDirectionFromCoords(grid.Agent.CurrentCell, grid.Policy)

		// then calculate the final direction with the stochastic transition model
		nextDirection := calculateNextStochasticMovement(intendedMovement, grid.Agent.TransitionModel)

		fmt.Printf("Current cell is %d and movement is %d\n",
			model.GetCellFromCoords(grid.Agent.CurrentCell, grid).ID,
			nextDirection)

		// deduct energy from agent
		grid.Agent.Energy += model.GetCellFromCoords(grid.Agent.CurrentCell, grid).Reward

		// shift agent's position
		grid.Agent.CurrentCell = moveAgent(grid.Agent.CurrentCell, nextDirection, grid)
	}

	fmt.Printf("Final cell is %d\n",
		model.GetCellFromCoords(grid.Agent.CurrentCell, grid).ID)

	// when simulation is done
	fmt.Println(grid.Agent)
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
