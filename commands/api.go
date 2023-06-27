package commands

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/dhm116/maze-generator/api"
	"github.com/dhm116/maze-generator/generator"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Starts the mazebot API",
	Run:   runAPI,
}

func init() {
	rootCmd.AddCommand(apiCmd)
	apiCmd.Flags().Int("port", 8080, "Changes the server port")
	apiCmd.Flags().IP("host", net.ParseIP("127.0.0.1"), "Specifies the host addresses to respond to")
}

func runAPI(cmd *cobra.Command, args []string) {
	host, _ := cmd.Flags().GetIP("host")
	port, _ := cmd.Flags().GetInt("port")

	route := gin.Default()

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	route.GET("/mazebot/random", mazeRequest)

	route.Run(fmt.Sprintf("%v:%v", host, port))
}

func mazeRequest(c *gin.Context) {
	var req api.MazeRequest

	err := c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Request validation failed!",
			"error":   err.Error(),
		})
		return
	}

	maze := generator.NewMaze(req.MaxSize)

	res := api.NewMazeResponseFromMaze(c, maze)

	c.JSON(http.StatusOK, res)
}
