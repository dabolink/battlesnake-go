package model

import (
	"net/http"
	"encoding/json"
)

type StartRequest struct {
	GameID int `json:"game_id"`
}

func NewStartRequest(req *http.Request) (*StartRequest, error) {
	decoded := StartRequest{}
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return &decoded, err
}
