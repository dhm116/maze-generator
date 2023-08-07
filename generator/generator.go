package generator

import (
	"math/rand"
	"time"
)

type Direction int

const (
	Up = Direction(1 << iota)
	Down
	Left
	Right
)

const (
	Wall  = "X"
	Space = " "
	Start = "A"
	End   = "B"
)

func (d Direction) String() string {
	var result string
	switch d {
	case Up:
		result = "N"
	case Down:
		result = "S"
	case Left:
		result = "W"
	case Right:
		result = "E"
	}
	return result
}

func (d Direction) Arrow() string {
	var result string
	switch d {
	case Up:
		result = "↑"
	case Down:
		result = "↓"
	case Left:
		result = "←"
	case Right:
		result = "→"
	}
	return result
}

type Generator struct {
	maze *Maze
	rand rand.Rand
	seed int64
}

func NewMaze(size int) *Generator {
	return FromSeed(size, time.Now().UnixNano())
}

func FromSeed(size int, seed int64) *Generator {
	maze := &Maze{
		Size:  size,
		Cells: make([]*Cell, size*size),
	}

	maze.InitializeCells()

	gen := &Generator{
		maze: maze,
		rand: *rand.New(rand.NewSource(seed)),
		seed: seed,
	}

	gen.initializeMaze()

	gen.setStartingPosition()
	gen.setEndingPosition()

	return gen
}

func (g *Generator) Maze() *Maze {
	return g.maze
}

func (g *Generator) Seed() int64 {
	return g.seed
}

func (g *Generator) initializeMaze() {
	// 1. Start with a grid full of walls.
	// 2. Pick a cell, mark it as part of the maze. Add the walls of the cell to the wall list.
	// 3. While there are walls in the list:
	// 		1. Pick a random wall from the list. If only one of the cells that the wall divides is visited, then:
	// 			1. Make the wall a passage and mark the unvisited cell as part of the maze.
	// 			2. Add the neighboring walls of the cell to the wall list.
	// 		2. Remove the wall from the list.
	var cell *Cell
	var index int
	walls := make([]*Cell, 0)
	cell = g.randomCell()
	cell.Type = Space

	walls = append(walls, cell.CanMoveTo()...)

	for len(walls) > 0 {
		cell, index = g.pickDirection(walls)
		// Avoid creating wide spaces
		if cell.AdjacentSpacesCount() <= 1 {
			// Turn the cell into a moveable space
			cell.Type = Space
			moves := cell.CanMoveTo()
			walls = append(walls, moves...)
		}
		// Remove the wall from the list
		walls = append(walls[:index], walls[index+1:]...)
	}
}

func (g *Generator) randomCell() *Cell {
	// cell := g.maze.Cells[g.rand.Intn(len(g.maze.Cells))]
	// if cell.IsBorder() {
	// 	return g.randomCell()
	// }
	// return cell
	playable := g.maze.PlayableCells()
	return playable[g.rand.Intn(len(playable))]
}

func (g *Generator) randomSpaceCell() *Cell {
	spaces := g.maze.OpenCells()
	return spaces[g.rand.Intn(len(spaces))]
}

func (g *Generator) pickDirection(choices []*Cell) (*Cell, int) {
	index := g.rand.Intn(len(choices))
	return choices[index], index
}

func (g *Generator) setStartingPosition() {
	cell := g.randomSpaceCell()
	cell.Type = Start
	g.maze.Start = cell
}

func (g *Generator) setEndingPosition() {
	var cell *Cell
	var distance int
	start := g.maze.Start
	bound := g.maze.Size / 4

	cell = g.randomSpaceCell()
	distance = start.DistanceFrom(cell)

	// fmt.Printf("%s distance is %d, bound is %d\n", cell.String(), distance, bound)
	// fmt.Printf("\tdistance < bound [%+v]\n", distance < bound)
	for distance <= bound {
		// fmt.Println("Picking a new cell")
		cell = g.randomCell()
		distance = start.DistanceFrom(cell)
		// fmt.Printf("%s distance is %d, bound is %d\n", cell.String(), distance, bound)
		// fmt.Printf("\tdistance < bound [%+v]\n", distance < bound)
	}
	// fmt.Printf("Final %s distance is %d, bound is %d\n", cell.String(), distance, bound)
	cell.Type = End
	g.maze.End = cell
}
