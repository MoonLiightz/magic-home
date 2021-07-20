package appcli

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/moonliightz/magic-home/v2/internal/util"
	magichome "github.com/moonliightz/magic-home/v2/pkg"
	"github.com/urfave/cli/v2"
)

func parseOptions(c *cli.Context) magichome.Options {
	return magichome.Options{
		ApplyMasks:       c.Bool("masks"),
		ColdWhiteSupport: c.Bool("cw-support"),
		LogReceived:      c.Bool("verbose"),
		LogSending:       c.Bool("verbose"),
		Port:             uint16(c.Int("port")),
	}
}

func parseIpArg(c *cli.Context) net.IP {
	ipArg := c.Args().Get(0)
	if ip := net.ParseIP(ipArg); ip != nil {
		return ip
	} else {
		fmt.Printf("Invalid IP: %s\n\n", ipArg)
		cli.ShowCommandHelpAndExit(c, c.Command.Name, 1)
	}
	return nil
}

func parseStateArg(c *cli.Context) magichome.State {
	stateArg := strings.ToLower(c.Args().Get(1))
	var status magichome.State

	if stateArg == "on" {
		status = magichome.On
	} else if stateArg == "off" {
		status = magichome.Off
	} else {
		fmt.Println("Invalid state: ", stateArg)
		cli.ShowCommandHelpAndExit(c, c.Command.Name, 1)
	}

	return status
}

func colorAction(c *cli.Context) error {
	controller, err := magichome.New(parseIpArg(c), parseOptions(c))
	if err != nil {
		return err
	}
	defer controller.Close()

	_, err = controller.SetColor(
		util.ParseStringToUint8(c.Args().Get(1)),
		util.ParseStringToUint8(c.Args().Get(2)),
		util.ParseStringToUint8(c.Args().Get(3)),
	)
	if err != nil {
		return err
	}

	return nil
}

func colorAndWhiteAction(c *cli.Context) error {
	controller, err := magichome.New(parseIpArg(c), parseOptions(c))
	if err != nil {
		return err
	}
	defer controller.Close()

	_, err = controller.SetColorAndWarmWhite(
		util.ParseStringToUint8(c.Args().Get(1)),
		util.ParseStringToUint8(c.Args().Get(2)),
		util.ParseStringToUint8(c.Args().Get(3)),
		util.ParseStringToUint8(c.Args().Get(4)),
	)
	if err != nil {
		return err
	}

	return nil
}

func colorAndWhitesAction(c *cli.Context) error {
	controller, err := magichome.New(parseIpArg(c), parseOptions(c))
	if err != nil {
		return err
	}
	defer controller.Close()

	_, err = controller.SetColorAndWhites(
		util.ParseStringToUint8(c.Args().Get(1)),
		util.ParseStringToUint8(c.Args().Get(2)),
		util.ParseStringToUint8(c.Args().Get(3)),
		util.ParseStringToUint8(c.Args().Get(4)),
		util.ParseStringToUint8(c.Args().Get(5)),
	)
	if err != nil {
		return err
	}

	return nil
}

func warmWhiteAction(c *cli.Context) error {
	controller, err := magichome.New(parseIpArg(c), parseOptions(c))
	if err != nil {
		return err
	}
	defer controller.Close()

	_, err = controller.SetWarmWhite(
		util.ParseStringToUint8(c.Args().Get(1)),
	)
	if err != nil {
		return err
	}

	return nil
}

func whitesAction(c *cli.Context) error {
	controller, err := magichome.New(parseIpArg(c), parseOptions(c))
	if err != nil {
		return err
	}
	defer controller.Close()

	_, err = controller.SetWhites(
		util.ParseStringToUint8(c.Args().Get(1)),
		util.ParseStringToUint8(c.Args().Get(2)),
	)
	if err != nil {
		return err
	}

	return nil
}

func patternAction(c *cli.Context) error {
	controller, err := magichome.New(parseIpArg(c), parseOptions(c))
	if err != nil {
		return err
	}
	defer controller.Close()

	_, err = controller.SetPattern(
		magichome.Pattern(c.Args().Get(1)),
		util.ParseStringToUint8(c.Args().Get(2)),
	)
	if err != nil {
		return err
	}

	return nil
}

func stateAction(c *cli.Context) error {
	controller, err := magichome.New(parseIpArg(c), parseOptions(c))
	if err != nil {
		return err
	}
	defer controller.Close()

	_, err = controller.SetState(parseStateArg(c))
	if err != nil {
		return err
	}

	err = controller.Close()
	if err != nil {
		return err
	}

	return nil
}

func statusAction(c *cli.Context) error {
	controller, err := magichome.New(parseIpArg(c), parseOptions(c))
	if err != nil {
		return err
	}
	defer controller.Close()

	var deviceState *magichome.DeviceState
	deviceState, err = controller.GetDeviceState()
	if err != nil {
		return err
	}

	if c.Bool("json") {
		res, err := json.Marshal(deviceState)
		if err != nil {
			return err
		}
		fmt.Println(string(res))
	} else {
		fmt.Printf("Device is:\t\t\t")
		if deviceState.State == magichome.On {
			fmt.Println("On")
		} else {
			fmt.Println("Off")
		}
		fmt.Printf("Type:\t\t\t\t%2d (0x%.2x)\n", deviceState.DeviceType, deviceState.DeviceType)
		fmt.Printf("Version:\t\t\t%2d\n", deviceState.LedVersionNumber)
		fmt.Printf("Mode:\t\t\t\t%s\n", deviceState.Mode)
		fmt.Printf("Pattern:\t\t\t%s\n", deviceState.Pattern)
		fmt.Printf("Speed:\t\t\t\t%3d\n", deviceState.Speed)
		fmt.Printf("Color: \t\t\t R:\t%3d \n\t\t\t G:\t%3d \n\t\t\t B:\t%3d\n", deviceState.Color.R, deviceState.Color.G, deviceState.Color.B)
		fmt.Printf("Color Temperature:\tWW:\t%3d \n\t\t\tCW:\t%3d\n", deviceState.ColorTemperature.WW, deviceState.ColorTemperature.CW)
	}

	return nil
}

func discoverAction(c *cli.Context) error {
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
}
