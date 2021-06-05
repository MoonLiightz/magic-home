package magichome

import (
	"errors"
	"fmt"

	"github.com/moonliightz/magic-home/v2/internal/util"
)

// Pattern represents a Magic Home pattern
type Pattern string

var patterns = map[Pattern]uint8{
	"SevenColorCrossFade":   0x25,
	"RedGradualChange":      0x26,
	"GreenGradualChange":    0x27,
	"BlueGradualChange":     0x28,
	"YellowGradualChange":   0x29,
	"CyanGradualChange":     0x2a,
	"PurpleGradualChange":   0x2b,
	"WhiteGradualChange":    0x2c,
	"RedGreenCrossFade":     0x2d,
	"RedBlueCrossFade":      0x2e,
	"GreenBlueCrossFade":    0x2f,
	"SevenColorStrobeFlash": 0x30,
	"RedStrobeFlasg":        0x31,
	"GreenStrobeFlash":      0x32,
	"BlueStrobeFlash":       0x33,
	"YellowStrobeFlash":     0x34,
	"CyanStrobeFlash":       0x35,
	"PurpleStrobeFlash":     0x36,
	"WhiteStrobeFlash":      0x37,
	"SevenColorJumping":     0x38,
}

func determinePattern(patternByte uint8) Pattern {
	if patternByte >= 0x25 && patternByte <= 0x38 {
		for key := range patterns {
			if patterns[key] == patternByte {
				return key
			}
		}
	}
	return "unknown"
}

func delayToSpeed(delay uint8) uint8 {
	delay = util.Clamp(delay, 1, 31)
	delay--
	return uint8(100 - (float32(delay) / 30 * 100))
}

func speedToDelay(speed uint8) uint8 {
	speed = util.Clamp(speed, 0, 100)
	return uint8((30 - ((float32(speed) / 100) * 30)) + 1)
}

func (c *Controller) SetPattern(pattern Pattern, speed uint8) ([]uint8, error) {
	patternBytes := patterns[pattern]
	if patternBytes == 0 {
		return nil, errors.New("invalid pattern")
	}

	delay := speedToDelay(speed)
	fmt.Println(delay)
	buffer := []uint8{0x61, patternBytes, delay, 0x0f}

	return c.sendCommand(buffer)
}
