package util

import (
	"github.com/pmoros/markov-decision-model-networks/pkg/model"
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
		},
	}
}

func RunScenario(grid model.Grid) {
	// as long as the agent is not on a goal
	grid.Agent.CurrentCell = grid.Agent.InitialCell

	for model.GetCellFromCoords(grid.Agent.CurrentCell, grid).Type == model.Goal {
		intendedMovement := model.GetPolicyDirectionFromCoords(grid.Agent.CurrentCell, grid.Policy)
	}
}

func calculateNextMovement(agent model.Agent, grid model.Grid, policy model.Policy) model.Direction {

}
