package appcli

import (
	"github.com/urfave/cli/v2"
)

type command struct {
	Color          *cli.Command
	ColorAndWhite  *cli.Command
	ColorAndWhites *cli.Command
	Discover       *cli.Command
	Pattern        *cli.Command
	State          *cli.Command
	Status         *cli.Command
	WarmWhite      *cli.Command
	Whites         *cli.Command
}

// Command are the commands for the CLI
var Command = command{
	Color: &cli.Command{
		Name:      "color",
		Aliases:   []string{"c", "rgb"},
		Usage:     "Set the color values of the device",
		ArgsUsage: "<ip> <r> <g> <b>",
		Flags:     []cli.Flag{Flag.CWSupport, Flag.Masks, Flag.Port, Flag.Verbose},
		Action:    colorAction,
	},
	ColorAndWhite: &cli.Command{
		Name:      "colorandwhite",
		Aliases:   []string{"cww", "rgbw"},
		Usage:     "Set the color and warm white values of the device",
		ArgsUsage: "<ip> <r> <g> <b> <ww>",
		Flags:     []cli.Flag{Flag.CWSupport, Flag.Masks, Flag.Port, Flag.Verbose},
		Action:    colorAndWhiteAction,
	},
	ColorAndWhites: &cli.Command{
		Name:      "colorandwhites",
		Aliases:   []string{"cwwcw", "rgbww"},
		Usage:     "Set the color, warm white and cold white values op the device",
		ArgsUsage: "<ip> <r> <g> <b> <ww> <cw>",
		Flags:     []cli.Flag{Flag.CWSupport, Flag.Masks, Flag.Port, Flag.Verbose},
		Action:    colorAndWhitesAction,
	},
	WarmWhite: &cli.Command{
		Name:      "warmwhite",
		Aliases:   []string{"ww"},
		Usage:     "Set the warm white value of the device",
		ArgsUsage: "<ip> <ww>",
		Flags:     []cli.Flag{Flag.CWSupport, Flag.Masks, Flag.Port, Flag.Verbose},
		Action:    warmWhiteAction,
	},
	Whites: &cli.Command{
		Name:      "whites",
		Aliases:   []string{},
		Usage:     "Set the warm white and cold white value of the device",
		ArgsUsage: "<ip> <ww> <cw>",
		Flags:     []cli.Flag{Flag.CWSupport, Flag.Masks, Flag.Port, Flag.Verbose},
		Action:    whitesAction,
	},
	Pattern: &cli.Command{
		Name:      "pattern",
		Aliases:   []string{},
		Usage:     "Activate a built-in pattern",
		ArgsUsage: "<ip> <pattern> <speed>",
		Flags:     []cli.Flag{Flag.Port, Flag.Verbose},
		Action:    patternAction,
	},
	State: &cli.Command{
		Name:      "state",
		Aliases:   []string{"s"},
		Usage:     "Switch the device state to on or off",
		ArgsUsage: "<ip> <state>",
		Flags:     []cli.Flag{Flag.Port, Flag.Verbose},
		Action:    stateAction,
	},
	Status: &cli.Command{
		Name:      "status",
		Aliases:   []string{},
		Usage:     "Fetch and prints the current status of the device",
		ArgsUsage: "<ip>",
		Flags:     []cli.Flag{Flag.JSON, Flag.Port, Flag.Verbose},
		Action:    statusAction,
	},
	Discover: &cli.Command{
		Name:      "discover",
		Aliases:   []string{"d"},
		Usage:     "Discover for Magic Home devices on the network",
		ArgsUsage: "",
		Flags:     []cli.Flag{Flag.BroadcastAddr, Flag.Timeout},
		Action:    discoverAction,
	},
}
