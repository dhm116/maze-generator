package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dhm116/maze-generator/generator"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "print size",
	Short: "Generates a maze and prints it out",
	Args:  cobra.ExactArgs(1),
	Run:   printMaze,
}

func init() {
	rootCmd.AddCommand(printCmd)
}

func printMaze(cmd *cobra.Command, args []string) {
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

	fmt.Print(maze.ToString())
	fmt.Printf("Took %s\n", duration.String())
}
