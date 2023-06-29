package solver

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/dhm116/maze-generator/generator"
	"github.com/dhm116/maze-generator/queue"
)

type Dijkstra struct {
	distances map[*generator.Cell]float64
	maze      *generator.Maze
	paths     map[*generator.Cell]*generator.Cell
	queue     *queue.PriorityQueue
}

func NewDijkstra(m *generator.Maze) *Dijkstra {
	solver := &Dijkstra{
		distances: make(map[*generator.Cell]float64),
		maze:      m,
		paths:     make(map[*generator.Cell]*generator.Cell),
		queue:     queue.NewPriorityQueue(),
	}
	took := solver.CalculateShortestPath()
	fmt.Println("Solved the maze in:", took.String())
	return solver
}

func (d *Dijkstra) CalculateShortestPath() time.Duration {
	now := time.Now()
	cells := d.maze.OpenCells()

	for i := range cells {
		cell := cells[i]
		dist := math.Inf(1)
		if cell == d.maze.Start {
			dist = 0
		}
		d.distances[cell] = dist
		d.paths[cell] = nil
		d.queue.Push(cell, int(dist))
	}

	for d.queue.Len() > 0 {
		cell := d.queue.Pop().(*generator.Cell)
		if cell.IsEnd() {
			break
		}
		for i := range cell.Neighbors {
			neighbor := cell.Neighbors[i]
			dist := d.distances[cell] + float64(cell.DistanceFrom(neighbor))
			if dist < d.distances[neighbor] {
				d.distances[neighbor] = dist
				d.paths[neighbor] = cell
				d.queue.Update(neighbor, int(dist))
			}
		}
	}
	duration := time.Since(now)
	return duration
}

func (d *Dijkstra) Distances() map[*generator.Cell]float64 {
	return d.distances
}

func (d *Dijkstra) Paths() map[*generator.Cell]*generator.Cell {
	return d.paths
}

func (d *Dijkstra) Solution() map[*generator.Cell]*generator.Cell {
	var cell *generator.Cell
	path := make(map[*generator.Cell]*generator.Cell)
	cell = d.maze.End
	for cell != nil {
		prev := d.paths[cell]
		if prev != nil {
			path[cell] = prev
		}
		cell = prev
	}
	return path
}

func (d *Dijkstra) ToString() string {
	var b strings.Builder
	path := d.Solution()

	for i := range d.maze.Cells {
		cell := d.maze.Cells[i]
		_, inPath := path[cell]
		char := cell.Type
		switch {
		case char == generator.Wall:
			char = "▓"
		case inPath && cell != d.maze.Start && cell != d.maze.End:
			char = "*"
		}
		fmt.Fprintf(&b, "%s", char)
		if cell.Location.X == d.maze.Size-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}
