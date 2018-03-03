package main

import (
	"fmt"
	"github.com/dabolink/battlesnake-go/model"
	"github.com/dabolink/battlesnake-go/service"
	"net/http"
)

func Start(res http.ResponseWriter, req *http.Request) {
	fmt.Println("START REQUEST")

	data, err := model.NewStartRequest(req)
	if err != nil {
		fmt.Println("Bad start request: ", err)
	}
	dump(data)

	respond(res, model.StartResponse{
		Taunt:          model.START_TAUNT,
		Color:          model.PRIMARY_COLOR,
		Name:           model.NAME,
		HeadType:       model.HEAD_PIXEL,
		TailType:       model.TAIL_ROUND_BUM,
		SecondaryColor: model.SECONDARY_COLOR,
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	fmt.Println("MOVE REQUEST")

	moveRequest, err := model.NewMoveRequest(req)

	if err != nil {
		fmt.Println("Bad move request: ", err)
	}
	dump(moveRequest)

	gameBoard := model.GenerateGameBoard(*moveRequest)

	directions := model.NewDirections(moveRequest.You.Head())

	move := service.DetermineBestMove(gameBoard, directions)

	respond(res, model.MoveResponse{
		Move: move,
	})
}
