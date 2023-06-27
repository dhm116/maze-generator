package api

type MazeRequest struct {
	MinSize int `form:"minSize"`
	MaxSize int `form:"maxSize"`
}
