package api

type MazeRequest struct {
	MinSize int `form:"minSize" url:"minSize" json:"minSize" validate:"required,number,min=12,ltefield=MaxSize"`
	MaxSize int `form:"maxSize" url:"maxSize" json:"maxSize" validate:"required,number,gtefield=MinSize"`
}
