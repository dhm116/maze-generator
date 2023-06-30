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
	seed  int64
	Size  int
	Start *Cell
	End   *Cell
}

func (p *Point) ToArray() [2]int {
	return [2]int{p.X, p.Y}
}

func (p *Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (m *Maze) InitializeCells() {
	// Create the Cells
	for i := range m.Cells {
		x := int(math.Mod(float64(i), float64(m.Size)))
		y := i / m.Size

		m.Cells[i] = &Cell{
			Location:  &Point{X: x, Y: y},
			maze:      m,
			Neighbors: make([]*Cell, 0),
			Type:      Wall,
		}
	}

	// Populate each Cell's neighbors
	for i := range m.Cells {
		cell := m.Cells[i]
		x := cell.Location.X
		y := cell.Location.Y

		if x > 0 {
			cell.AddNeighbor(m.CellAt(x-1, y))
		}
		if x < m.Size-1 {
			cell.AddNeighbor(m.CellAt(x+1, y))
		}
		if y > 0 {
			cell.AddNeighbor(m.CellAt(x, y-1))
		}
		if y < m.Size-1 {
			cell.AddNeighbor(m.CellAt(x, y+1))
		}
	}
}

func (m *Maze) CellAt(x, y int) *Cell {
	return m.Cells[x+(y*m.Size)]
}

func (m *Maze) OpenCells() []*Cell {
	var results []*Cell
	results = make([]*Cell, 0)
	for i := range m.Cells {
		if m.Cells[i].Type != Wall {
			results = append(results, m.Cells[i])
		}
	}
	return results
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

func (c *Cell) IsStart() bool {
	return c.Type == Start
}

func (c *Cell) IsEnd() bool {
	return c.Type == End
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

func (c *Cell) AdjacentSpaces() []*Cell {
	var results []*Cell
	results = make([]*Cell, 0)
	for i := range c.Neighbors {
		if c.Neighbors[i].Type == Space {
			results = append(results, c.Neighbors[i])
		}
	}
	return results
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

	return int(diffX + diffY)
}

func (c *Cell) DirectionTo(other *Cell) Direction {
	var dir Direction
	diffX := c.Location.X - other.Location.X
	diffY := c.Location.Y - other.Location.Y
	switch {
	case diffX < 0:
		dir = Right
	case diffX > 0:
		dir = Left
	case diffY < 0:
		dir = Down
	case diffY > 0:
		dir = Up
	}
	return dir
}

func (c *Cell) String() string {
	return fmt.Sprintf("[%s] '%s'", c.Location.String(), c.Type)
}
