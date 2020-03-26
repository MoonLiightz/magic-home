package appcli

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/moonliightz/magic-home/internal/util"
	magichome "github.com/moonliightz/magic-home/pkg"
	"github.com/urfave/cli/v2"
)

type command struct {
	Color    *cli.Command
	State    *cli.Command
	Discover *cli.Command
}

// Command are the commands for the CLI
var Command = command{
	Color: &cli.Command{
		Name:      "color",
		Aliases:   []string{"c"},
		Usage:     "Set the color of the LED Strip",
		ArgsUsage: "<ip> <r> <g> <b> <w>",
		Flags:     []cli.Flag{Flag.Port},
		Action: func(c *cli.Context) error {
			ipArg := c.Args().Get(0)
			if ip := net.ParseIP(ipArg); ip != nil {

				controller, err := magichome.New(ip, uint16(c.Int("port")))
				if err != nil {
					return nil
				}

				err = controller.SetColor(magichome.Color{
					R: util.ParseStringToUint8(c.Args().Get(1)),
					G: util.ParseStringToUint8(c.Args().Get(2)),
					B: util.ParseStringToUint8(c.Args().Get(3)),
					W: util.ParseStringToUint8(c.Args().Get(4)),
				})
				if err != nil {
					return err
				}

				err = controller.Close()
				if err != nil {
					return err
				}

			} else {
				fmt.Println("Unvalid IP: ", ipArg)
				cli.ShowCommandHelpAndExit(c, "color", 1)
			}

			return nil
		},
	},
	State: &cli.Command{
		Name:      "state",
		Aliases:   []string{"s"},
		Usage:     "Switch the LED Strip state to on or off",
		ArgsUsage: "<ip> <state>",
		Flags:     []cli.Flag{Flag.Port},
		Action: func(c *cli.Context) error {
			ipArg := c.Args().Get(0)
			if ip := net.ParseIP(ipArg); ip != nil {
				stateArg := strings.ToLower(c.Args().Get(1))
				var status magichome.State

				if stateArg == "on" {
					status = magichome.On
				} else if stateArg == "off" {
					status = magichome.Off
				} else {
					fmt.Println("Invalid state: ", stateArg)
					cli.ShowCommandHelpAndExit(c, "state", 1)
					return nil
				}

				controller, err := magichome.New(ip, uint16(c.Int("port")))
				if err != nil {
					return err
				}

				err = controller.SetState(status)
				if err != nil {
					return err
				}

				err = controller.Close()
				if err != nil {
					return err
				}
			} else {
				fmt.Println("Invalid IP: ", ipArg)
				cli.ShowCommandHelpAndExit(c, "state", 1)
			}

			return nil
		},
	},
	Discover: &cli.Command{
		Name:      "discover",
		Aliases:   []string{"d"},
		Usage:     "Discover for Magic Home devices on the network",
		ArgsUsage: "",
		Flags:     []cli.Flag{Flag.BroadcastAddr, Flag.Timeout},
		Action: func(c *cli.Context) error {
			fmt.Print("Discovering")
			go func() {
				for {
					fmt.Print(".")
					time.Sleep(100 * time.Millisecond)
				}
			}()

			devices, err := magichome.Discover(magichome.DiscoverOptions{
				BroadcastAddr: c.String("broadcastaddr"),
				Timeout:       uint8(c.Int("timeout")),
			})
			if err != nil {
				return err
			}

			if len(*devices) >= 1 {
				fmt.Println()
				fmt.Println("Discovered the following devices:")
				fmt.Println()
				fmt.Println("Address    \t| ID         \t| Model")
				fmt.Println("---------------------------------------")
				for _, device := range *devices {
					fmt.Printf("%s\t| %s\t| %s\n", device.IP, device.ID, device.Model)
				}
			} else {
				fmt.Println()
				fmt.Println("No devices discovered.")
			}

			return nil
		},
	},
}
