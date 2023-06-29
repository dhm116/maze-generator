package generator

import (
	"math/rand"
)

const (
	Up = 1 << iota
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

func NewMaze(size int) *Maze {
	maze := &Maze{
		Size:  size,
		Cells: make([]*Cell, size*size),
	}

	maze.InitializeCells()

	initializeMaze(maze)

	pickStartingPosition(maze)
	pickEndingPosition(maze)

	return maze
}

func initializeMaze(m *Maze) {
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
	cell = randomCell(m)
	cell.Type = Space

	walls = append(walls, cell.CanMoveTo()...)

	for len(walls) > 0 {
		// fmt.Printf("At [%d,%d] with %d walls\n", cell.Location.X, cell.Location.Y, len(walls))
		cell, index = pickDirection(walls)
		// fmt.Printf("\tPicked [%d,%d]\n", cell.Location.X, cell.Location.Y)
		if cell.AdjacentSpacesCount() <= 1 {
			// fmt.Printf("\tSwitching [%d,%d] to a space\n", cell.Location.X, cell.Location.Y)
			cell.Type = Space
			moves := cell.CanMoveTo()
			walls = append(walls, moves...)
			// fmt.Printf("\tAdded available walls - now %d walls\n", len(walls))
			// fmt.Print(m.ToString())
		}
		walls = append(walls[:index], walls[index+1:]...)
		// fmt.Printf("\tRemoved wall from list - now %d walls\n", len(walls))
	}
}

func randomCell(m *Maze) *Cell {
	cell := m.Cells[rand.Intn(len(m.Cells))]
	if cell.IsBorder() {
		return randomCell(m)
	}
	return cell
}

func pickDirection(choices []*Cell) (*Cell, int) {
	index := rand.Intn(len(choices))
	return choices[index], index
}

func pickStartingPosition(m *Maze) {
	var cell *Cell
	cell = randomCell(m)
	for cell.Type != Space {
		cell = randomCell(m)
	}
	cell.Type = Start
	m.Start = cell
}

func pickEndingPosition(m *Maze) {
	var cell *Cell
	var distance int
	start := m.Start
	bound := m.Size / 4

	cell = randomCell(m)
	distance = start.DistanceFrom(cell)

	for cell.Type != Space && distance < bound {
		cell = randomCell(m)
		distance = start.DistanceFrom(cell)
	}
	cell.Type = End
	m.End = cell
}
