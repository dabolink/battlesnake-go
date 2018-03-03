package model

import (
	"math/rand"
	"time"
)

const (
	UP    = "up"
	DOWN  = "down"
	LEFT  = "left"
	RIGHT = "right"
)

type Directions struct {
	CurrentLocation Point
	Up              bool
	Down            bool
	Left            bool
	Right           bool
}

func NewDirections(location Point) Directions {
	return Directions{
		Up:              true,
		Down:            true,
		Left:            true,
		Right:           true,
		CurrentLocation: location,
	}
}

func (directions Directions) GetMove() (string, error) {
	var d []string
	if directions.Up {
		d = append(d, UP)
	} else if directions.Down {
		d = append(d, DOWN)
	} else if directions.Left {
		d = append(d, LEFT)
	} else if directions.Right {
		d = append(d, RIGHT)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return d[r.Intn(4)], nil
}
