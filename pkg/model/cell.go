package model

const (
	Clear   CellType = "CLEAR"
	Blocked CellType = "BLOCKED"
	Goal    CellType = "GOAL"
)

type CellType string

type Cell struct {
	ID     int
	Type   CellType
	Reward float32
}

func NewCell(id int, cellType CellType, reward float32) Cell {
	return Cell{
		ID:     id,
		Type:   cellType,
		Reward: reward,
	}
}
