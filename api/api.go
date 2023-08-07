package api

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"

	"github.com/dhm116/maze-generator/generator"
)

func RunServer(host string, port int) {
	app := iris.Default()
	app.Validator = validator.New()

	ac := makeAccessLog()
	defer ac.Close() // Close the underline file.

	app.UseRouter(ac.Handler)

	app.Get("/mazebot/random", randomMaze)
	app.Get("/mazebot/{seed:string}", seededMaze)

	app.Listen(fmt.Sprintf("%s:%d", host, port))
}

func makeAccessLog() *accesslog.AccessLog {
	// Initialize a new access log middleware.
	// ac := accesslog.File("./access.log")
	// Remove this line to disable logging to console:
	// ac.AddOutput(os.Stdout)
	ac := accesslog.New(os.Stdout)

	// The default configuration:
	ac.Delim = '|'
	ac.TimeFormat = "2006-01-02 15:04:05"
	ac.Async = false
	ac.IP = true
	ac.BytesReceivedBody = true
	ac.BytesSentBody = true
	ac.BytesReceived = false
	ac.BytesSent = false
	ac.BodyMinify = true
	ac.RequestBody = true
	ac.ResponseBody = false
	ac.KeepMultiLineError = true
	ac.PanicLog = accesslog.LogHandler

	// Default line format if formatter is missing:
	// Time|Latency|Code|Method|Path|IP|Path Params Query Fields|Bytes Received|Bytes Sent|Request|Response|
	//
	// Set Custom Formatter:
	// ac.SetFormatter(&accesslog.JSON{
	// 	Indent:    "  ",
	// 	HumanTime: true,
	// })
	// ac.SetFormatter(&accesslog.Template{Text: "{{.Code}}"})

	return ac
}

func wrapValidationErrors(errs validator.ValidationErrors) []validationError {
	validationErrors := make([]validationError, 0, len(errs))
	for _, validationErr := range errs {
		validationErrors = append(validationErrors, validationError{
			ActualTag: validationErr.ActualTag(),
			Namespace: validationErr.Namespace(),
			Kind:      validationErr.Kind().String(),
			Type:      validationErr.Type().String(),
			Value:     fmt.Sprintf("%v", validationErr.Value()),
			Param:     validationErr.Param(),
		})
	}

	return validationErrors
}

func randomMaze(ctx iris.Context) {
	var req MazeRequest

	err := ctx.ReadBody(&req)
	if err != nil {
		// Handle the error, below you will find the right way to do that...

		if errs, ok := err.(validator.ValidationErrors); ok {
			// Wrap the errors with JSON format, the underline library returns the errors as interface.
			validationErrors := wrapValidationErrors(errs)

			// Fire an application/json+problem response and stop the handlers chain.
			ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
				Title("Validation error").
				Detail("One or more fields failed to be validated").
				Type(ctx.GetCurrentRoute().Path()).
				Key("errors", validationErrors))

			return
		}

		// It's probably an internal JSON error, let's dont give more info here.
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return
	}

	gen := generator.NewMaze(req.MaxSize)

	res := NewMazeResponseFromMaze(ctx, gen)

	ctx.JSON(res)
}

func seededMaze(ctx iris.Context) {
	encodedSeed := ctx.Params().Get("seed")
	seed, size := generator.DecodeMapSeed(encodedSeed)
	gen := generator.FromSeed(size, seed)

	res := NewMazeResponseFromMaze(ctx, gen)

	ctx.JSON(res)
}
