package main

import (
	"fmt"
	"github.com/dabolink/battlesnake-go/model"
	"github.com/dabolink/battlesnake-go/service"
	"testing"
)

var (
	mySnake = model.Snake{
		Body: model.PointList{
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
	enemySnake = model.Snake{
		Body: model.PointList{
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
	moveRequest = model.MoveRequest{
		Food: model.PointList{
			{
				X: 7,
				Y: 7,
			},
		},
		Height: 10,
		Width:  10,
		ID:     73,
		Snakes: model.SnakeList{
			mySnake,
			enemySnake,
		},
		Turn: 0,
		You:  mySnake,
	}
)

func TestMove(t *testing.T) {
	gameBoard := model.GenerateGameBoard(moveRequest)

	directions := model.NewDirections(moveRequest.You.Head())

	move := service.DetermineBestMove(gameBoard, directions)

	fmt.Println(move)
}
