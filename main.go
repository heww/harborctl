package main

import (
	"fmt"
	"os"

	"github.com/heww/harborctl/command"
)

func main() {
	app := command.App()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "harborctl: %s\n", err)
		os.Exit(1)
	}
}