package model

import "encoding/json"

type Snake struct {
	Body   PointList `json:"body"`
	Health int       `json:"health"`
	ID     string    `json:"id"`
	Length int       `json:"length"`
	Name   string    `json:"name"`
	Taunt  string    `json:"taunt"`
}

func (snake Snake) Head() Point { return snake.Body[0] }

type HeadType string

const (
	HEAD_BENDR     HeadType = "bendr"
	HEAD_DEAD               = "dead"
	HEAD_FANG               = "fang"
	HEAD_PIXEL              = "pixel"
	HEAD_REGULAR            = "regular"
	HEAD_SAFE               = "safe"
	HEAD_SAND_WORM          = "sand-worm"
	HEAD_SHADES             = "shades"
	HEAD_SMILE              = "smile"
	HEAD_TONGUE             = "tongue"
)

type TailType string

const (
	TAIL_BLOCK_BUM    TailType = "block-bum"
	TAIL_CURLED                = "curled"
	TAIL_FAT_RATTLE            = "fat-rattle"
	TAIL_FRECKLED              = "freckled"
	TAIL_PIXEL                 = "pixel"
	TAIL_REGULAR               = "regular"
	TAIL_ROUND_BUM             = "round-bum"
	TAIL_SKINNY                = "skinny"
	TAIL_SMALL_RATTLE          = "small-rattle"
)

// Parses List<Point> into []Point
type PointList []Point

func (list *PointList) UnmarshalJSON(data []byte) error {
	var obj struct {
		Data []Point `json:"data"`
	}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	*list = obj.Data
	return nil
}

// Parses List<Snake> into []Snake
type SnakeList []Snake

func (list *SnakeList) UnmarshalJSON(data []byte) error {
	var obj struct {
		Data []Snake `json:"data"`
	}
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}
	*list = obj.Data
	return nil
}

