package model

type Agent struct {
	TransitionModel Transition
	InitialCell     Coords
	CurrentCell     Coords
	Energy          float64
}
