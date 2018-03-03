package model

import (
	"fmt"
)

type GameBoard struct {
	BoardValues [][]BoardTile
	Width       int
	MySnakeID   string
	Height      int
}

func (board *GameBoard) Print() {
	fmt.Println(board.Width, board.Height)
	for _, row := range board.BoardValues {
		fmt.Println()
		for _, tile := range row {
			fmt.Print(tile.Value, " ")
		}
	}
	fmt.Println()
}
func (board GameBoard) GetTileAt(x int, y int) *BoardTile {
	if x < 0 || board.Width <= x {
		return nil
	}
	if y < 0 || board.Height <= y {
		return nil
	}
	return &board.BoardValues[x][y]
}

type TileType int

const (
	SNAKE_VALUE          = -999
	SNAKE_GRADIENT_VALUE = -5
	FOOD_GRADIENT_VALUE  = 10
	SNAKE_TAIL_VALUE     = -100
)

const (
	NONE TileType = iota
	OTHER_TAIL
	OTHER_HEAD
	OTHER_BODY
	ME_TAIL
	ME_HEAD
	ME_BODY
	FOOD
)

type BoardTile struct {
	Type    TileType
	Value   int
	IsValid bool
}

func GenerateGameBoard(request MoveRequest) *GameBoard {
	boardVals := make([][]BoardTile, request.Width) // initialize a slice of dy slices
	for i := 0; i < request.Width; i++ {
		boardVals[i] = make([]BoardTile, request.Height) // initialize a slice of dx unit8 in each of dy slices
		for j := 0; j < request.Height; j++ {
			(&boardVals[i][j]).Init()
		}
	}
	gameBoard := &GameBoard{
		Width:       request.Width,
		Height:      request.Height,
		BoardValues: boardVals,
		MySnakeID:   request.You.ID,
	}

	AddSnakesToBoard(request.Snakes, gameBoard, gameBoard.MySnakeID)
	AddFoodToBoard(request.Food, gameBoard)

	gameBoard.Print()
	return gameBoard
}

func AddSnakesToBoard(snakes SnakeList, gameBoard *GameBoard, mySnakeID string) {
	for _, snake := range snakes {
		AddSnakeToBoard(snake, gameBoard, mySnakeID)
	}
}
func AddSnakeToBoard(snake Snake, gameBoard *GameBoard, mySnakeID string) {
	isMe := mySnakeID == snake.ID
	if snake.Health == 0 {
		return
	}
	for i, bodyPart := range snake.Body {
		var tileType TileType
		var value int
		if i == 0 {
			if isMe {
				tileType = ME_HEAD
			} else {
				tileType = OTHER_HEAD
			}
			value = SNAKE_VALUE
		} else if snake.Length-1 == i {
			if isMe {
				tileType = ME_TAIL
			} else {
				tileType = OTHER_TAIL
			}
			value = SNAKE_TAIL_VALUE
		} else {
			if isMe {
				tileType = ME_BODY
			} else {
				tileType = OTHER_BODY
			}
			value = SNAKE_VALUE
		}
		boardTile := &(gameBoard.BoardValues)[bodyPart.X][bodyPart.Y]
		boardTile.Set(tileType, false, value)

		if isMe {
			continue
		}
		addGradientAtTile(gameBoard, snake.Head(), SNAKE_GRADIENT_VALUE)
	}
}

func addGradientAtTile(board *GameBoard, point Point, value int) {
	absValue := abs(value)
	for x := point.X - absValue; x <= absValue+point.X; x++ {
		for y := point.Y - absValue; y <= absValue+point.Y; y++ {
			tile := board.GetTileAt(x, y)
			if tile != nil {
				val := calcValue(point, x, y, value)
				tile.Value = tile.Value + val
			}
		}
	}
}
func calcValue(point Point, x int, y int, value int) int {
	distance := abs(point.X-x) + abs(point.Y-y)
	if value < 0 {
		return value + distance
	} else {
		return value - distance
	}
}

func abs(value int) int {
	if value >= 0 {
		return value
	} else {
		return -1 * value
	}
}

func AddFoodToBoard(foodList PointList, gameBoard *GameBoard) {
	for _, food := range foodList {
		tile := gameBoard.GetTileAt(food.X, food.Y)
		tile.Set(FOOD, true, 0)
		addGradientAtTile(gameBoard, food, FOOD_GRADIENT_VALUE)
	}
}

func (tile *BoardTile) Set(tileType TileType, isValid bool, value int) {
	(*tile).IsValid = isValid
	(*tile).Type = tileType
	(*tile).Value = value

}

func (tile *BoardTile) Init() {
	(*tile).IsValid = true
	(*tile).Type = NONE
	(*tile).Value = 0
}
