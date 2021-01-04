package main

import (
	"fmt"
	"net"
	"os"

	magichome "github.com/moonliightz/magic-home/pkg"
)

func checkError(e error) {
	if e != nil {
		fmt.Println("Error: ", e)
		os.Exit(1)
	}
}

func main() {
	// Create a new Magic Home LED Strip Controller
	controller, err := magichome.New(net.ParseIP("192.168.42.105"), 5577)
	checkError(err)

	// Get device state
	var deviceState *magichome.DeviceState
	deviceState, err = controller.GetDeviceState()
	checkError(err)

	// Work with the response (print it)
	fmt.Printf("Device is: ")
	if deviceState.State == magichome.On {
		fmt.Println("On")
	} else {
		fmt.Println("Off")
	}
	fmt.Printf("Color: \tR: %d \n\tG: %d \n\tB: %d \n\tW: %d\n", deviceState.Color.R, deviceState.Color.G, deviceState.Color.B, deviceState.Color.W)

	// And finaly close the connection to LED Strip Controller
	err = controller.Close()
	checkError(err)
}
