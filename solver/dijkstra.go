package solver

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/dhm116/maze-generator/generator"
	"github.com/dhm116/maze-generator/queue"
)

type CellMap map[*generator.Cell]*generator.Cell

type CellPath struct {
	cells    CellMap
	dijkstra *Dijkstra
	path     []*generator.Cell
}

func NewCellPath(dijkstra *Dijkstra) *CellPath {
	return &CellPath{
		cells:    make(CellMap),
		dijkstra: dijkstra,
		path:     make([]*generator.Cell, 0),
	}
}

func (p *CellPath) Path() []*generator.Cell {
	sort.Sort(p)
	return p.path
}

func (p *CellPath) From(c *generator.Cell) *generator.Cell {
	return p.cells[c]
}

func (p *CellPath) Contains(c *generator.Cell) bool {
	_, ok := p.cells[c]
	return ok
}

func (p *CellPath) Append(from, to *generator.Cell) {
	// fmt.Printf("Adding %s to %s\n", from.String(), to.String())
	if len(p.path) == 0 || p.path[len(p.path)-1] != from {
		p.path = append(p.path, from)
	}
	p.path = append(p.path, to)
	p.cells[from] = to
}

func (p *CellPath) Len() int {
	return len(p.path)
}

func (p *CellPath) Swap(i, j int) {
	p.path[i], p.path[j] = p.path[j], p.path[i]
}

func (p *CellPath) Less(i, j int) bool {
	return p.dijkstra.distances[p.path[i]] > p.dijkstra.distances[p.path[j]]
}

func (p *CellPath) String() string {
	var b strings.Builder
	path := p.Path()
	for _, toCell := range path {
		fromCell := p.From(toCell)
		if fromCell == nil {
			break
		}
		dir := fromCell.DirectionTo(toCell).String()
		b.WriteString(dir)
	}
	reversed := b.String()
	chars := []rune(reversed)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

type Dijkstra struct {
	distances map[*generator.Cell]float64
	maze      *generator.Maze
	paths     CellMap
	queue     *queue.PriorityQueue
}

func NewDijkstra(m *generator.Maze) *Dijkstra {
	solver := &Dijkstra{
		distances: make(map[*generator.Cell]float64),
		maze:      m,
		paths:     make(CellMap),
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

func (d *Dijkstra) Paths() CellMap {
	return d.paths
}

func (d *Dijkstra) Solution() *CellPath {
	var cell *generator.Cell
	path := NewCellPath(d)
	cell = d.maze.End
	for cell != nil {
		prev := d.paths[cell]
		if prev != nil {
			// path[cell] = prev
			path.Append(cell, prev)
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
		fromCell := path.From(cell)
		inPath := path.Contains(cell)
		char := cell.Type
		switch {
		case char == generator.Wall:
			char = "▓"
		case inPath && cell != d.maze.Start && cell != d.maze.End:
			char = fromCell.DirectionTo(cell).Arrow()
		}
		fmt.Fprintf(&b, "%s", char)
		if cell.Location.X == d.maze.Size-1 {
			b.WriteString("\n")
		}
	}

	return b.String()
}
