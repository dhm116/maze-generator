package api

import (
	"fmt"
	"strings"

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

func NewMazeResponseFromMaze(c *gin.Context, g *generator.Generator) *MazeResponse {
	maze := g.Maze()
	seedPath := strings.Replace(c.FullPath(), "random", fmt.Sprint(g.Seed()), 1)
	res := &MazeResponse{
		Name:             fmt.Sprintf("%dx%d", maze.Size, maze.Size),
		Path:             seedPath,
		StartingPosition: maze.Start.Location.ToArray(),
		EndingPosition:   maze.End.Location.ToArray(),
		Map:              make([][]string, maze.Size),
	}

	for x := range res.Map {
		res.Map[x] = make([]string, maze.Size)

		for y := range res.Map[x] {
			cell := maze.CellAt(x, y)
			cellType := cell.Type

			res.Map[x][y] = cellType

		}
	}

	return res
}
