package shef

type key string

const (
	Power       key = "power"
	PowerOn     key = "poweron"
	PowerOff    key = "poweroff"
	Format      key = "format"
	Pause       key = "pause"
	Rewind      key = "rew"
	Replay      key = "replay"
	Stop        key = "stop"
	Advance     key = "advance"
	FastForward key = "ffwd"
	Record      key = "record"
	Play        key = "play"
	Guide       key = "guide"
	Active      key = "active"
	List        key = "list"
	Exit        key = "exit"
	Back        key = "back"
	Menu        key = "menu"
	Info        key = "info"
	Up          key = "up"
	Down        key = "down"
	Left        key = "left"
	Right       key = "right"
	Select      key = "select"
	Red         key = "red"
	Green       key = "green"
	Yellow      key = "yellow"
	Blue        key = "blue"
	ChannelUp   key = "chanup"
	ChannelDown key = "chandown"
	Previous    key = "prev"
	Zero        key = "0"
	One         key = "1"
	Two         key = "2"
	Three       key = "3"
	Four        key = "4"
	Five        key = "5"
	Six         key = "6"
	Seven       key = "7"
	Eight       key = "8"
	Nine        key = "9"
	Dash        key = "dash"
	Enter       key = "enter"
)

// NewKey returns a new |key| who's value is |k|. This allows
// callers to generate new keys which did not exist at the time
// of implementation. This should be a temporary shim, as ideally
// new keys can simply be added as constants above.
func NewKey(k string) key {
	return key(k)
}

// keyAction represents an action on a key.
type keyAction string

const (
	// KeyUp simulates key being released only.
	KeyUp = "keyUp"
	// KeyDown simulates key being pressed only.
	KeyDown = "keyDown"
	// KeyPress simulates both press and release.
	KeyPress = "keyPress"
)
