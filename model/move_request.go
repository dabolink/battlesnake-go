package model

import (
	"net/http"
	"encoding/json"
)

var (
	DIRECTIONS = []string{
		"up",
		"down",
		"left",
		"right",
	}
)

type MoveRequest struct {
	Food   PointList `json:"food"`
	Height int       `json:"height"`
	ID     int       `json:"id"`
	Snakes SnakeList `json:"snakes"`
	Turn   int       `json:"turn"`
	Width  int       `json:"width"`
	You    Snake     `json:"you"`
}

func NewMoveRequest(req *http.Request) (*MoveRequest, error) {
	decoded := MoveRequest{}
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return &decoded, err
}