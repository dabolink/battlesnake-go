package model

import (
	"testing"
	"reflect"
	"encoding/json"
)

func TestNewMoveRequest(t *testing.T) {
	s := `
	{
	  "food": {
		"data": [
		  {
			"object": "point",
			"x": 0,
			"y": 9
		  }
		],
		"object": "list"
	  },
	  "height": 20,
	  "id": 5,
	  "object": "world",
	  "snakes": {
		"data": [
		  {
			"body": {
			  "data": [
				{
				  "object": "point",
				  "x": 13,
				  "y": 19
				},
				{
				  "object": "point",
				  "x": 13,
				  "y": 19
				},
				{
				  "object": "point",
				  "x": 13,
				  "y": 19
				}
			  ],
			  "object": "list"
			},
			"health": 100,
			"id": "58a0142f-4cd7-4d35-9b17-815ec8ff8e70",
			"length": 3,
			"name": "Sonic Snake",
			"object": "snake",
			"taunt": "Gotta go fast"
		  },
		  {
			"body": {
			  "data": [
				{
				  "object": "point",
				  "x": 8,
				  "y": 15
				},
				{
				  "object": "point",
				  "x": 8,
				  "y": 15
				},
				{
				  "object": "point",
				  "x": 8,
				  "y": 15
				}
			  ],
			  "object": "list"
			},
			"health": 100,
			"id": "48ca23a2-dde8-4d0f-b03a-61cc9780427e",
			"length": 3,
			"name": "Typescript Snake",
			"object": "snake",
			"taunt": ""
		  }
		],
		"object": "list"
	  },
	  "turn": 10,
	  "width": 20,
	  "you": {
		"body": {
		  "data": [
			{
			  "object": "point",
			  "x": 8,
			  "y": 15
			},
			{
			  "object": "point",
			  "x": 8,
			  "y": 15
			},
			{
			  "object": "point",
			  "x": 8,
			  "y": 15
			}
		  ],
		  "object": "list"
		},
		"health": 100,
		"id": "48ca23a2-dde8-4d0f-b03a-61cc9780427e",
		"length": 3,
		"name": "Typescript Snake",
		"object": "snake",
		"taunt": ""
	  }
	}
	`

	req := requestWithBody(s)
	result, err := NewMoveRequest(req)

	if err != nil {
		t.Fatal(err)
	}
	expected := MoveRequest{
		Food:   PointList{Point{X: 0, Y: 9}},
		Height: 20,
		ID:     5,
		Snakes: SnakeList{
			{
				Body:   PointList{{X: 13, Y: 19}, {X: 13, Y: 19}, {X: 13, Y: 19}},
				Health: 100,
				ID:     "58a0142f-4cd7-4d35-9b17-815ec8ff8e70",
				Length: 3,
				Name:   "Sonic Snake",
				Taunt:  "Gotta go fast",
			},
			{
				Body:   PointList{{X: 8, Y: 15}, {X: 8, Y: 15}, {X: 8, Y: 15}},
				Health: 100,
				ID:     "48ca23a2-dde8-4d0f-b03a-61cc9780427e",
				Length: 3,
				Name:   "Typescript Snake",
				Taunt:  "",
			},
		},
		Turn:  10,
		Width: 20,
		You: Snake{
			Body:   PointList{{X: 8, Y: 15}, {X: 8, Y: 15}, {X: 8, Y: 15}},
			Health: 100,
			ID:     "48ca23a2-dde8-4d0f-b03a-61cc9780427e",
			Length: 3,
			Name:   "Typescript Snake",
			Taunt:  "",
		},
	}

	if !reflect.DeepEqual(*result, expected) {
		t.Errorf("Expected:\n%#v\nGot:\n%#v", expected, result)
	}
}

func TestMoveResponse(t *testing.T) {
	res := MoveResponse{
		Move: "up",
	}

	result, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	expected := `{"move":"up"}`
	if string(result) != expected {
		t.Errorf("Expected:\n%#v\nGot:\n%#v", expected, string(result))
	}
}

