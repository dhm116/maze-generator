package api

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dhm116/maze-generator/generator"
)

func RunServer(host string, port int) {
	route := gin.Default()

	route.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	route.GET("/mazebot/random", randomMaze)
	route.GET("/mazebot/:seed", seededMaze)

	route.Run(fmt.Sprintf("%s:%d", host, port))
}

func randomMaze(c *gin.Context) {
	var req MazeRequest

	err := c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Request validation failed!",
			"error":   err.Error(),
		})
		return
	}

	gen := generator.NewMaze(req.MaxSize)

	res := NewMazeResponseFromMaze(c, gen)

	c.JSON(http.StatusOK, res)
}

func seededMaze(c *gin.Context) {
	encodedSeed := c.Param("seed")
	seed, size := generator.DecodeMapSeed(encodedSeed)
	gen := generator.FromSeed(size, seed)

	res := NewMazeResponseFromMaze(c, gen)

	c.JSON(http.StatusOK, res)
}
