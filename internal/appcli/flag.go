package appcli

import "github.com/urfave/cli/v2"

type flag struct {
	Port          *cli.IntFlag
	BroadcastAddr *cli.StringFlag
	Timeout       *cli.IntFlag
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
	BroadcastAddr: &cli.StringFlag{
		Name:        "broadcastaddr",
		Aliases:     []string{"b"},
		Usage:       "broadcast address of the netwerk",
		DefaultText: "255.255.255.255",
		Value:       "255.255.255.255",
		Required:    false,
	},
	Timeout: &cli.IntFlag{
		Name:        "timeout",
		Aliases:     []string{"t"},
		Usage:       "discover search timeout",
		DefaultText: "1 second",
		Value:       1,
		Required:    false,
	},
}
