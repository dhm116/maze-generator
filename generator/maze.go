package generator

import (
	"fmt"
	"math"
	"strings"
)

type Point struct {
	X, Y int
}

type Cell struct {
	Location  *Point
	maze      *Maze
	Neighbors []*Cell
	Type      string
}

type Maze struct {
	Cells []*Cell
	Size  int
	Start *Cell
	End   *Cell
}

func (p *Point) ToArray() [2]int {
	return [2]int{p.X, p.Y}
}

// 0 [0,0] 1 [1,0] 2 [2,0] 3 [3,0] 4 [4,0]
// 5 [0,1] 6 [1,1] 7 [2,1] 8 [3,1] 9 [4,1]
func (m *Maze) CellAt(x, y int) *Cell {
	return m.Cells[x+(y*m.Size)]
}

func (m *Maze) ToString() string {
	var b strings.Builder

	for i := range m.Cells {
		cell := m.Cells[i]
		char := cell.Type
		if char == "X" {
			char = "▓"
		}
		fmt.Fprintf(&b, "%s", char)
		if cell.Location.X == m.Size-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}

func (c *Cell) IsBorder() bool {
	limit := c.maze.Size - 1
	return c.Location.X == 0 || c.Location.Y == 0 || c.Location.X == limit || c.Location.Y == limit
}

func (c *Cell) AddNeighbor(n *Cell) {
	c.Neighbors = append(c.Neighbors, n)
}

func (c *Cell) CanMoveTo() []*Cell {
	available := make([]*Cell, 0)
	for i := range c.Neighbors {
		if c.Neighbors[i].Type == Wall && !c.Neighbors[i].IsBorder() {
			available = append(available, c.Neighbors[i])
		}
	}
	return available
}

func (c *Cell) AdjacentSpacesCount() int {
	count := 0
	for i := range c.Neighbors {
		if c.Neighbors[i].Type == Space {
			count = count + 1
		}
	}
	return count
}

func (c *Cell) DistanceFrom(other *Cell) int {
	diffX := math.Abs(float64(c.Location.X - other.Location.X))
	diffY := math.Abs(float64(c.Location.Y - other.Location.Y))

	return int(math.Sqrt(math.Pow(diffX, 2) + math.Pow(diffY, 2)))
}
