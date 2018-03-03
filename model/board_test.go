package model

import "testing"

var (
	mySnake = Snake{
		Body: PointList{
			{
				X: 0,
				Y: 0,
			},
			{
				X: 0,
				Y: 1,
			},
		},
		Health: 1,
		ID:     "joy",
		Name:   "Snakey",
		Taunt:  "Taunty",
	}
	enemySnake = Snake{
		Body: PointList{
			{
				X: 9,
				Y: 0,
			},
			{
				X: 9,
				Y: 1,
			},
		},
		Health: 1,
		ID:     "??????",
		Name:   "Joe",
		Taunt:  "a squirrel thing",
	}
	moveRequest = MoveRequest{
		Food: PointList{
			{
				X: 7,
				Y: 7,
			},
		},
		Height: 10,
		Width:  10,
		ID:     73,
		Snakes: SnakeList{
			mySnake,
			enemySnake,
		},
		Turn: 0,
		You:  mySnake,
	}
)

func TestGenerateGameBoard(t *testing.T) {
	_ = GenerateGameBoard(moveRequest)
}
