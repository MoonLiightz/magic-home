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
	DeviceType       uint8
	State            State
	LedVersionNumber uint8
	Mode             uint8
	Slowness         uint8
	Color            Color
}
