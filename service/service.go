package service

import "github.com/dabolink/battlesnake-go/model"

func DetermineBestMove(board *model.GameBoard, directions model.Directions) string {
	point := directions.CurrentLocation

	northTile := board.GetTileAt(point.X, point.Y+1)
	if northTile == nil || !northTile.IsValid {
		directions.Up = false
	}
	southTile := board.GetTileAt(point.X, point.Y-1)
	if southTile == nil || !southTile.IsValid {
		directions.Down = false
	}
	eastTile := board.GetTileAt(point.X+1, point.Y)
	if eastTile == nil || !eastTile.IsValid {
		directions.Left = false
	}
	westTile := board.GetTileAt(point.X-1, point.Y)
	if westTile == nil || !westTile.IsValid {
		directions.Right = false
	}

	var bestDirection string
	bestValue := -9999999999999

	if directions.Up {
		if northTile.Value > bestValue {
			bestDirection = model.UP
			bestValue = northTile.Value
		}
	} else if directions.Down {
		if northTile.Value > bestValue {
			bestDirection = model.DOWN
			bestValue = southTile.Value
		}
	} else if directions.Left {
		if directions.Left {
			bestDirection = model.LEFT
			bestValue = westTile.Value
		}
	} else if directions.Right {
		if directions.Right {
			bestDirection = model.UP
			bestValue = eastTile.Value
		}
	} else {
		return model.UP
	}
	return bestDirection
}
