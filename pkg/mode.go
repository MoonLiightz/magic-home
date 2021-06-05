package magichome

// Mode represents the mode of the Magic Home device
type Mode string

const (
	ColorMode     Mode = "color"
	CustomMode    Mode = "custom"
	IAPatternMode Mode = "ia_pattern"
	PatternMode   Mode = "pattern"
	SpecialMode   Mode = "special"
)

func determineMode(buffer []uint8) Mode {
	if buffer[3] == 0x61 || (buffer[3] == 0 && buffer[4] == 0x61) {
		return ColorMode
	} else if buffer[3] == 0x62 {
		return SpecialMode
	} else if buffer[3] == 0x60 {
		return CustomMode
	} else if buffer[3] >= 0x25 && buffer[3] <= 0x38 {
		return PatternMode
	} else if buffer[3] >= 0x64 && uint16(buffer[3]) <= 0x018f {
		return IAPatternMode
	}
	return "unknown"
}
