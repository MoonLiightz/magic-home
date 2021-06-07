package magichome

import (
	"github.com/moonliightz/magic-home/v2/internal/util"
)

// Color represents a RGB color
type Color struct {
	R uint8
	G uint8
	B uint8
}

// ColorTemperature represents the color impression warm white and cold white
type ColorTemperature struct {
	WW uint8
	CW uint8
}

func (c *Controller) colorCommand(r uint8, g uint8, b uint8, ww uint8, cw uint8, mask uint8) ([]uint8, error) {
	r = util.Clamp(r, 0, 255)
	g = util.Clamp(g, 0, 255)
	b = util.Clamp(b, 0, 255)
	ww = util.Clamp(ww, 0, 255)

	var buffer []uint8
	if c.options.ColdWhiteSupport {
		cw = util.Clamp(cw, 0, 255)
		buffer = []uint8{0x31, r, g, b, ww, cw, mask, 0x0f}
	} else {
		buffer = []uint8{0x31, r, g, b, ww, mask, 0x0f}
	}

	resBuffer, err := c.sendCommand(buffer)
	if err == nil {
		c.DeviceState.Color.R = r
		c.DeviceState.Color.G = g
		c.DeviceState.Color.B = b
		c.DeviceState.ColorTemperature.WW = ww
		c.DeviceState.ColorTemperature.CW = cw
	}

	return resBuffer, err
}

// SetColor can be used to change the color values of the Magic Home device
func (c *Controller) SetColor(r uint8, g uint8, b uint8) ([]uint8, error) {
	if c.options.ApplyMasks {
		return c.colorCommand(r, g, b, 0, 0, 0x0f)
	}
	return c.SetColorAndWhites(r, g, b, c.DeviceState.ColorTemperature.WW, c.DeviceState.ColorTemperature.CW)
}

// SetColorAndWhites can be used to change the color and white values of the Magic Home device
func (c *Controller) SetColorAndWhites(r uint8, g uint8, b uint8, ww uint8, cw uint8) ([]uint8, error) {
	return c.colorCommand(r, g, b, ww, cw, 0)
}

// SetColorAndWarmWhite can be used to change the color and warm white values of the Magic Home device
func (c *Controller) SetColorAndWarmWhite(r uint8, g uint8, b uint8, ww uint8) ([]uint8, error) {
	return c.colorCommand(r, g, b, ww, c.DeviceState.ColorTemperature.CW, 0)
}

// SetWarmWhite can be used to change the warm white value of the Magic Home device
func (c *Controller) SetWarmWhite(ww uint8) ([]uint8, error) {
	if c.options.ApplyMasks {
		return c.colorCommand(0, 0, 0, ww, c.DeviceState.ColorTemperature.CW, 0x0f)
	}
	return c.SetColorAndWarmWhite(c.DeviceState.Color.R, c.DeviceState.Color.G, c.DeviceState.Color.B, ww)
}

// SetWhites can be used to change the white values of the Magic Home device
func (c *Controller) SetWhites(ww uint8, cw uint8) ([]uint8, error) {
	if c.options.ApplyMasks {
		return c.colorCommand(0, 0, 0, ww, cw, 0x0f)
	}
	return c.SetColorAndWhites(c.DeviceState.Color.R, c.DeviceState.Color.G, c.DeviceState.Color.B, ww, cw)
}
