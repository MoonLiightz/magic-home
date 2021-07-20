package main

import (
	"fmt"
	"os"

	"github.com/moonliightz/magic-home/v2/internal/appcli"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "magic-home",
		Usage: "A CLI for controlling Magic Home (Magic Hue) LED Strip Controller",
		Commands: []*cli.Command{
			appcli.Command.Color,
			appcli.Command.ColorAndWhite,
			appcli.Command.ColorAndWhites,
			appcli.Command.WarmWhite,
			appcli.Command.Whites,
			appcli.Command.Pattern,
			appcli.Command.State,
			appcli.Command.Status,
			appcli.Command.Discover,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
