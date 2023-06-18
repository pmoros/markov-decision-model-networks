package model

type Probability float32

type Transition map[Direction]Probability

type Policy map[int]Direction
