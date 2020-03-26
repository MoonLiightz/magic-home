package magichome

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// Device represents a Magic Home device that answered the udp broadcast
type Device struct {
	IP    net.IP
	ID    string
	Model string
}

// DiscoverOptions are used to configure the discovering
type DiscoverOptions struct {
	BroadcastAddr string
	Timeout       uint8
}

// Discover searches for Magic Home devices on the network
func Discover(options DiscoverOptions) (*[]Device, error) {
	if options.BroadcastAddr == "" {
		options.BroadcastAddr = "255.255.255.255"
	}

	broadcastAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:48899", options.BroadcastAddr))
	if err != nil {
		return nil, err
	}

	localAddr, err := net.ResolveUDPAddr("udp4", ":0")
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP("udp4", localAddr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	receivedDevice := make(chan *Device)
	done := make(chan bool, 1)
	var devices []Device

	go receive(conn, receivedDevice)
	go send(conn, broadcastAddr)
	go timeout(options.Timeout, done)

	for {
		select {
		case device, ok := <-receivedDevice:
			if ok {
				devices = append(devices, *device)
			}
		case <-done:
			done = nil
		}

		if done == nil {
			break
		}
	}

	close(receivedDevice)

	return &devices, nil
}

func send(conn *net.UDPConn, broadcastAddr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("HF-A11ASSISTHREAD"), broadcastAddr)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func receive(conn *net.UDPConn, receivedDevice chan<- *Device) {
	for {
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			break
		}

		data := strings.Split(string(buf[:n]), ",")
		if len(data) == 3 {
			receivedDevice <- &Device{
				IP:    net.ParseIP(data[0]),
				ID:    data[1],
				Model: data[2],
			}
		}
	}
}

func timeout(seconds uint8, done chan<- bool) {
	if seconds < 1 {
		seconds = 1
	} else if seconds > 30 {
		seconds = 30
	}

	time.Sleep(time.Duration(seconds) * time.Second)

	done <- true
}
