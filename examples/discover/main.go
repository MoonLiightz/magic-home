package main

import (
	"fmt"
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
	devices, err := magichome.Discover(magichome.DiscoverOptions{
		BroadcastAddr: "255.255.255.255",
		Timeout:       1, // (in seconds)
	})
	checkError(err)

	fmt.Println("Discovered the following devices:\n")
	fmt.Println("Address    \t| ID         \t| Model")
	fmt.Println("---------------------------------------")
	for _, device := range *devices {
		fmt.Printf("%s\t| %s\t| %s\n", device.IP, device.ID, device.Model)
	}
}
