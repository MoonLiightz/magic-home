package magichome

import (
	"fmt"
	"net"
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
	if options.Port <= 0 {
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

		sum += uint(value)
	}
	_, err := c.conn.Write(c.command.off)
	return err
}

// SetColor can be used to change the color of the LED Strip
func (c *Controller) SetColor(color Color) error {
	bytesToSend := c.command.color

	bytesToSend[1] = byte(color.R)
	bytesToSend[2] = byte(color.G)
	bytesToSend[3] = byte(color.B)
	bytesToSend[4] = byte(color.W)

	var sum int
	for _, value := range bytesToSend {
		sum += int(value)
	}

	bytesToSend = append(bytesToSend, byte(sum&0xff))

	_, err := c.conn.Write(bytesToSend)
	if err != nil {
		return err
	}

	return nil
}

// GetDeviceState can be used to get information about the state of the LED Strip
func (c *Controller) GetDeviceState() (*DeviceState, error) {

	_, err := c.conn.Write(c.command.state)
	if err != nil {
		return nil, err
	}

	response := make([]byte, 14)
	_, err = c.conn.Read(response)
	if err != nil {
		return nil, err
	}

	deviceState := DeviceState{}
	deviceState.DeviceType = response[1]
	deviceState.Mode = response[3]
	deviceState.Slowness = response[5]
	deviceState.Color.R = response[6]
	deviceState.Color.G = response[7]
	deviceState.Color.B = response[8]
	deviceState.Color.W = response[9]
	deviceState.LedVersionNumber = response[10]

	if response[2] == 0x23 {
		deviceState.State = On
	} else {
		deviceState.State = Off
	}

	return &deviceState, nil
}

// Close closes the tcp connection to the LED Strip
func (c *Controller) Close() error {
	return c.conn.Close()
}
