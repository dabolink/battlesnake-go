package model

import (
	"testing"
	"bytes"
	"net/http"
	"encoding/json"
)

func requestWithBody(body string) *http.Request {
	req, err := http.NewRequest("", "", bytes.NewBufferString(body))
	if err != nil {
		panic(err)
	}
	return req
}

func TestStartRequest(t *testing.T) {
	s := `
	{
	  "game_id": 1234
	}
	`
	req := requestWithBody(s)
	result, err := NewStartRequest(req)

	if err != nil {
		t.Fatal(err)
	}
	expected := StartRequest{
		GameID: 1234,
	}
	if *result != expected {
		t.Errorf("Expected:\n%#v\nGot:\n%#v", expected, result)
	}
}

func TestNewStartResponse(t *testing.T) {
	res := StartResponse{}

	result, err := json.Marshal(res)
	if err != nil {
		t.Fatal(err)
	}
	expected := "{}"
	if string(result) != expected {
		t.Errorf("Expected:\n%#v\nGot:\n%#v", expected, string(result))
	}

	res =StartResponse{
		Color:          "#FF0000",
		SecondaryColor: "#00FF00",
		HeadURL:        "http://placecage.com/c/100/100",
		Name:           "Cage Snake",
		Taunt:          "OH GOD NOT THE BEES",
		HeadType:       HEAD_SAND_WORM,
		TailType:       TAIL_FAT_RATTLE,
	}

	result, err = json.MarshalIndent(res, "", "    ")
	if err != nil {
		t.Fatal(err)
	}
	expected = `{
    "color": "#FF0000",
    "name": "Cage Snake",
    "head_url": "http://placecage.com/c/100/100",
    "taunt": "OH GOD NOT THE BEES",
    "head_type": "sand-worm",
    "tail_type": "fat-rattle",
    "secondary_color": "#00FF00"
}`

	if string(result) != expected {
		t.Errorf("Expected:\n%#v\nGot:\n%#v", expected, string(result))
	}
}
