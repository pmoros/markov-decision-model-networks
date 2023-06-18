package model

const (
	Forward     Movement = 0
	RotateLeft  Movement = 1
	RotateRight Movement = 2
	Back        Movement = 3
)

type Movement int

type Transition map[Movement]Probability
