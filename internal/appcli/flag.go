package appcli

import "github.com/urfave/cli/v2"

type flag struct {
	Port *cli.IntFlag
}

// Flag contains the options for the CLI
var Flag = flag{
	Port: &cli.IntFlag{
		Name:        "port",
		Aliases:     []string{"p"},
		Usage:       "socket port of the LED Strip",
		DefaultText: "5577",
		Value:       5577,
		Required:    false,
	},
}
