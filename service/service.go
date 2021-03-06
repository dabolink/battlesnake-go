package service

import (
	"fmt"
	"github.com/dabolink/battlesnake-go/model"
)

func DetermineBestMove(board *model.GameBoard, directions model.Directions) string {
	point := directions.CurrentLocation

	northTile := board.GetTileAt(point.X-1, point.Y)
	if northTile == nil || !northTile.IsValid {
		directions.Up = false
	}
	southTile := board.GetTileAt(point.X+1, point.Y)
	if southTile == nil || !southTile.IsValid {
		directions.Down = false
	}
	eastTile := board.GetTileAt(point.X, point.Y-1)
	if eastTile == nil || !eastTile.IsValid {
		directions.Right = false
	}
	westTile := board.GetTileAt(point.X, point.Y+1)
	if westTile == nil || !westTile.IsValid {
		directions.Left = false
	}

	var bestDirection string
	bestValue := -9999999999999

	if directions.Up {
		fmt.Println("up", northTile.Value)
		if northTile.Value > bestValue {
			bestDirection = model.UP
			bestValue = northTile.Value
		}
	} else if directions.Down {
		fmt.Println("down", southTile.Value)
		if southTile.Value > bestValue {
			bestDirection = model.DOWN
			bestValue = southTile.Value
		}
	} else if directions.Left {
		fmt.Println("left", westTile.Value)
		if westTile.Value > bestValue {
			bestDirection = model.LEFT
			bestValue = westTile.Value
		}
	} else if directions.Right {
		fmt.Println("right", eastTile.Value)
		if eastTile.Value > bestValue {
			bestDirection = model.RIGHT
			bestValue = eastTile.Value
		}
	} else {
		return model.UP
	}
	return bestDirection
}
