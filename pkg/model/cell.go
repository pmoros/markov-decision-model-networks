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
	Reward int
}
