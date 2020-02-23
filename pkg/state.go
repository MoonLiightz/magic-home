package magichome

// State represents the status of the device (on or off)
type State uint8

const (
	// On represents the switched-on state of the device
	On State = 1

	// Off represents the switched-off state of the device
	Off State = 0
)
