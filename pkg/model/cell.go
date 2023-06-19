package model

import "fmt"

const (
	Clear   CellType = "CLEAR"
	Blocked CellType = "BLOCKED"
	Goal    CellType = "GOAL"
)

type CellType string

type Cell struct {
	ID     int
	Type   CellType
	Reward float64
}

func NewCell(id int, cellType CellType, reward float64) Cell {
	return Cell{
		ID:     id,
		Type:   cellType,
		Reward: reward,
	}
}

func GetCellFromCoords(coords Coords, grid Grid) *Cell {
	if coords[0] >= len(grid.Cells[0]) || coords[1] >= len(grid.Cells) {
		fmt.Println("cell exceed grid bounds")
		return nil
	}

	for vCoord, row := range grid.Cells {
		for hCoord, cell := range row {
			if coords[0] == hCoord && coords[1] == vCoord {
				return &cell
			}
		}
	}
	return nil
}
