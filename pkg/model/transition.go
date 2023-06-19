package model

const (
	Forward     Movement = 0
	RotateLeft  Movement = -1
	RotateRight Movement = 1
	Back        Movement = 2
)

type Movement int

type Transition map[Movement]Probability
