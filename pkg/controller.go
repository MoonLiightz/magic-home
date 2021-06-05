package magichome

import (
	"fmt"
	"net"
	"time"
)

// Options holds Magic Home device controller object options
type Options struct {
	ApplyMasks       bool   // Enables a special bitmask which is required for some devices of type 0x25, 0x35 or 0x44 (default: false)
	ColdWhiteSupport bool   // Enables support for setting cold white values (default: false)
	LogReceived      bool   // Enables logging of received bytes (default: false)
	LogSending       bool   // Enables logging of sending bytes (default: false)
	Port             uint16 // Port of the Magic Home device (default: 5577)
	ReadTimeout      uint32 // Time in seconds to wait for a response (default: 1)
	WriteTimeout     uint32 // Time in seconds to try to send the bytes (default: 1)
}

// Controller represents a Magic Home device
type Controller struct {
	ip          net.IP
	conn        net.Conn
	options     Options
	DeviceState DeviceState
}

// New initializes a new Magic Home device controller
func New(ip net.IP, options Options) (*Controller, error) {
	// check options
	if options.Port < 1 || options.Port > 65535 {
		// default port
		options.Port = 5577
	}
	if options.WriteTimeout <= 0 {
		// default one second
		options.WriteTimeout = 1
	}
	if options.ReadTimeout <= 0 {
		// default one second
		options.ReadTimeout = 1
	}

	mh := Controller{
		ip:          ip,
		options:     options,
		DeviceState: DeviceState{},
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", mh.ip, mh.options.Port))
	if err != nil {
		return nil, err
	}

	mh.conn = conn

	return &mh, nil
}

func (c *Controller) sendCommand(buffer []uint8) ([]uint8, error) {
	// calc checksum
	var sum uint = 0
	for _, value := range buffer {
		sum += uint(value)
	}
	buffer = append(buffer, byte(sum&0xff))

	if c.options.LogSending {
		fmt.Printf("Sending bytes:  ")
		for _, b := range buffer {
			fmt.Printf("0x%.2x ", b)
		}
		fmt.Println()
	}

	c.conn.SetDeadline(time.Now().Add(time.Duration(c.options.WriteTimeout) * time.Second))
	_, err := c.conn.Write(buffer)
	if err != nil {
		return nil, err
	}

	resBuffer := make([]byte, 1024)
	c.conn.SetReadDeadline(time.Now().Add(time.Duration(c.options.ReadTimeout) * time.Second))
	c.conn.Read(resBuffer)

	if c.options.LogReceived {
		fmt.Printf("Received bytes: ")
		for _, b := range resBuffer {
			if b == 0 {
				continue
			}
			fmt.Printf("0x%.2x ", b)
		}
		fmt.Println()
	}

	return resBuffer, nil
}

// Close closes the tcp connection to the LED Strip
func (c *Controller) Close() error {
	return c.conn.Close()
}
