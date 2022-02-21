package magichome

// State represents the status of the device (on or off)
type State uint8

const (
	// On represents the switched-on state of the device
	On State = 1

	// Off represents the switched-off state of the device
	Off State = 0
)

// DeviceState represents the state of the device
type DeviceState struct {
	Color            Color
	ColorTemperature ColorTemperature
	DeviceType       uint8
	LedVersionNumber uint8
	Mode             Mode
	Pattern          Pattern
	Speed            uint8
	State            State
}

func determineState(stateByte uint8) State {
	if stateByte == 0x23 {
		return On
	}
	return Off
}

// SetState can be used to switch the LED Strip on (magichome.On) or off (magichome.Off)
func (c *Controller) SetState(s State) ([]uint8, error) {
	if s == On {
		return c.sendCommand([]uint8{0x71, 0x23, 0x0f})
	}
	return c.sendCommand([]uint8{0x71, 0x24, 0x0f})
}

// TurnOn turns the Magic Home device on
func (c *Controller) TurnOn() ([]uint8, error) {
	return c.SetState(On)
}

// TurnOff turns the Magic Home device off
func (c *Controller) TurnOff() ([]uint8, error) {
	return c.SetState(Off)
}

// GetDeviceState queries and returns the current device state
func (c *Controller) GetDeviceState() (*DeviceState, error) {
	buffer := []uint8{0x81, 0x8a, 0x8b}
	resBuffer, err := c.sendCommand(buffer)
	if err != nil {
		return nil, err
	}

	c.DeviceState.DeviceType = resBuffer[1]
	c.DeviceState.State = determineState(resBuffer[2])
	c.DeviceState.Mode = determineMode(resBuffer)
	c.DeviceState.Pattern = determinePattern(resBuffer[3])
	c.DeviceState.Speed = delayToSpeed(resBuffer[5])
	c.DeviceState.Color.R = resBuffer[6]
	c.DeviceState.Color.G = resBuffer[7]
	c.DeviceState.Color.B = resBuffer[8]
	c.DeviceState.ColorTemperature.WW = resBuffer[9]
	c.DeviceState.ColorTemperature.CW = resBuffer[11]
	c.DeviceState.LedVersionNumber = resBuffer[10]

	if c.DeviceState.DeviceType == 0x25 {
		c.options.ApplyMasks = true
	} else if c.DeviceState.DeviceType == 0x35 {
		c.options.ApplyMasks = true
		c.options.ColdWhiteSupport = true
	} else if c.DeviceState.DeviceType == 0x44 {
		c.options.ApplyMasks = true
	}

	return &c.DeviceState, nil
}
