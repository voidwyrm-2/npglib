package npglib

// An alias for all the built-in number types
type Number interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}
