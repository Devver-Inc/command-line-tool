// main.go
package main

import (
	"github.com/Devver-Inc/cli/internal/app"
)

func main() {
	config := app.ParseArgs()
	model := app.InitialModel()

	if config.Interactive {
		app.RunInteractive(model)
	} else if config.Command != "" {
		app.RunCommand(model, config.Command, config.Args)
	} else {
		app.RunInteractive(model)
	}
}
