package api

import (
	"fmt"

	"github.com/dhm116/maze-generator/generator"
	"github.com/gin-gonic/gin"
)

type MazeResponse struct {
	Name             string     `json:"name"`
	Path             string     `json:"path"`
	StartingPosition [2]int     `json:"startingPosition"`
	EndingPosition   [2]int     `json:"endingPosition"`
	Map              [][]string `json:"map"`
}

func NewMazeResponseFromMaze(c *gin.Context, m *generator.Maze) *MazeResponse {
	res := &MazeResponse{
		Name:             fmt.Sprintf("%dx%d", m.Size, m.Size),
		Path:             c.FullPath(),
		StartingPosition: m.Start.Location.ToArray(),
		EndingPosition:   m.End.Location.ToArray(),
		Map:              make([][]string, m.Size),
	}

	for x := range res.Map {
		res.Map[x] = make([]string, m.Size)

		for y := range res.Map[x] {
			cell := m.CellAt(x, y)
			cellType := cell.Type

			res.Map[x][y] = cellType

		}
	}

	return res
}
