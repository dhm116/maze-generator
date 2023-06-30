package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dhm116/maze-generator/generator"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:       "print size [seed]",
	Short:     "Generates a maze and prints it out",
	ValidArgs: []string{"size", "seed"},
	Run:       printMaze,
}

func init() {
	rootCmd.AddCommand(printCmd)
}

func printMaze(cmd *cobra.Command, args []string) {
	var err error
	var gen *generator.Generator
	var seed int64
	var size int
	hasSeed := false

	if len(args) == 0 {
		fmt.Println("Maze size not provided, exiting...")
		return
	}
	size, err = strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("The size '%v' is not a valid number, exiting...", args[0])
		return
	}
	if len(args) > 1 {
		// seed, err := strconv.Atoi(args[0])
		seed, err = strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Printf("The seed '%v' is not a valid number, exiting...", args[1])
			return
		}
		hasSeed = true
	}

	fmt.Printf("Generating a %dx%d maze\n", size, size)

	start := time.Now()
	if hasSeed {
		gen = generator.FromSeed(size, seed)
	} else {
		gen = generator.NewMaze(size)
	}
	duration := time.Since(start)

	fmt.Print(gen.Maze().ToString())
	fmt.Println("Took:", duration.String())
	fmt.Println("Seed:", gen.Seed())
}
