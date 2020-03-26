package main

import (
	"fmt"
	"net"
	"os"
	"time"

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
	controller, err := magichome.New(net.ParseIP("192.168.1.42"), 5577)
	checkError(err)

	// Turn LED Strip Controller on
	err = controller.SetState(magichome.On)
	checkError(err)

	// Sleep a few seconds to avoid spam
	time.Sleep(5 * time.Second)

	// Set color of LED Strip to white
	err = controller.SetColor(magichome.Color{
		R: 255,
		G: 255,
		B: 255,
		W: 0,
	})
	checkError(err)

	// Sleep again a few seconds to avoid spam
	time.Sleep(5 * time.Second)

	// Tun LED Strip Controller off
	err = controller.SetState(magichome.Off)
	checkError(err)

	// And finaly close the connection to LED Strip Controller
	err = controller.Close()
	checkError(err)
}
