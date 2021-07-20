package appcli

import "github.com/urfave/cli/v2"

type flag struct {
	Port          *cli.IntFlag
	BroadcastAddr *cli.StringFlag
	Timeout       *cli.IntFlag
	JSON          *cli.BoolFlag
	Verbose       *cli.BoolFlag
	CWSupport     *cli.BoolFlag
	Masks         *cli.BoolFlag
}

// Flag contains the options for the CLI
var Flag = flag{
	Port: &cli.IntFlag{
		Name:        "port",
		Aliases:     []string{"p"},
		Usage:       "socket port of the device",
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
	JSON: &cli.BoolFlag{
		Name:        "json",
		Aliases:     []string{},
		Usage:       "return output as a JSON-String",
		DefaultText: "false",
		Value:       false,
		Required:    false,
	},
	Verbose: &cli.BoolFlag{
		Name:        "verbose",
		Aliases:     []string{},
		Usage:       "verbose mode to output things like sent and received bytes",
		DefaultText: "false",
		Value:       false,
		Required:    false,
	},
	CWSupport: &cli.BoolFlag{
		Name:        "cw-support",
		Aliases:     []string{},
		Usage:       "enables support for setting cold white values",
		DefaultText: "false",
		Value:       false,
		Required:    false,
	},
	Masks: &cli.BoolFlag{
		Name:        "masks",
		Aliases:     []string{},
		Usage:       "enables a special bitmask which is required for some devices of type 0x25, 0x35 or 0x44",
		DefaultText: "false",
		Value:       false,
		Required:    false,
	},
}
