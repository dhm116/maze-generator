package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dhm116/maze-generator/generator"
	"github.com/dhm116/maze-generator/solver"
	"github.com/spf13/cobra"
)

var solveCmd = &cobra.Command{
	Use:   "solve size",
	Short: "Generates a maze, calculates a solution and prints it out",
	Args:  cobra.ExactArgs(1),
	Run:   solveMaze,
}

func init() {
	rootCmd.AddCommand(solveCmd)
}

func solveMaze(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Maze size not provided, exiting...")
		return
	}
	size, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("The size '%v' is not a valid number, exiting...", args[0])
		return
	}

	fmt.Printf("Generating a %dx%d maze\n", size, size)

	start := time.Now()
	maze := generator.NewMaze(size)
	duration := time.Since(start)
	fmt.Printf("Took %s\n", duration.String())

	d := solver.NewDijkstra(maze)
	fmt.Print(d.ToString())
}
