package scheduler

type Flag uint8

const (
	OpenFlag Flag = 1 << iota
	HolidayFlag
)

// NewFlagsMask returns an empty flags bitmask.
func NewFlagsMask() Flag {
	return 0
}

// SetFlag adds a flag in the flags bitmask.
func SetFlag(mask, flag Flag) Flag {
	return mask | flag
}

// ClearFlag removes a flag from the flags bitmask.
func ClearFlag(mask, flag Flag) Flag {
	return mask &^ flag
}

// ToggleFlag toggles a flag in the flags bitmask.
func ToggleFlag(mask, flag Flag) Flag {
	return mask ^ flag
}

// HasFlag returns whether a flag is present in the flags bitmask.
func HasFlag(mask, flag Flag) bool {
	return (mask & flag) != 0
}
