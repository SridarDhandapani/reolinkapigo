package enum

type OsdPosition uint

const (
	UPPER_LEFT OsdPosition = iota
	TOP_CENTER
	UPPER_RIGHT
	LOWER_LEFT
	BOTTOM_CENTER
	LOWER_RIGHT
)

func (op OsdPosition) Value() string {
	return []string{"Upper Left", "Top Center", "Upper Right", "Lower Left", "Bottom Center", "Lower Right"}[op]
}

type Toggle uint

const (
	Disabled Toggle = iota
	Enabled
)

func (t Toggle) String() string {
	return [...]string{"disabled", "enabled"}[t]
}
