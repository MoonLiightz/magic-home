package magichome

import (
	"fmt"
	"net"
)

// Controller represents a Magic Home LED Strip Controller
type Controller struct {
	ip      net.IP
	port    uint16
	conn    net.Conn
	command command
}

// New initializes a new Magic Home LED Strip Controller
func New(ip net.IP, port uint16) (*Controller, error) {
	mh := Controller{
		ip:   ip,
		port: port,
		command: command{
			on:    []byte{0x71, 0x23, 0x94},
			off:   []byte{0x71, 0x24, 0x95},
			color: []byte{0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			state: []byte{0x81, 0x8a, 0x8b, 0x96},
		},
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", mh.ip, mh.port))
	if err != nil {
		return nil, err
	}

	mh.conn = conn

	return &mh, nil
}

// SetState can be used to switch the LED Strip on (magichome.On) or off (magichome.Off)
func (c *Controller) SetState(s State) error {
	if s == On {
		_, err := c.conn.Write(c.command.on)
		return err
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
