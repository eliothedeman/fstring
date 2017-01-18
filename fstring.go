package fstring

import "unsafe"

const (
	lenSize              = 64
	flagBits             = 8
	lenMask              = ((1 << lenSize) - 1) >> flagBits
	flagsMask            = 0 ^ lenMask
	smallStringOptomized = 1
	smallStringThreshold = 15
)

// FString is a faster string implimentation.
// Uses a small string optomization.
type FString struct {
	str unsafe.Pointer
	len int
}

func fromString(s string) FString {
	if len(s) < smallStringThreshold {
		x := *(*rawFString)(unsafe.Pointer(&s))
		f := *(*FString)(unsafe.Pointer(&x))
		f.setFlag(smallStringOptomized)
		return f
	}
	return *(*FString)(unsafe.Pointer(&s))
}

func (f FString) cat(x FString) FString {
	if f.realLen()+x.realLen() < smallStringThreshold {
		start := f.realLen()
		out := [16]uint8{}
		raw := x.toRaw()
		for i := 0; i < x.realLen(); i++ {
			out[i+start] = raw[i]
		}

		return *(*FString)(unsafe.Pointer(&raw))
	}

	return fromString(f.String() + x.String())
}

func (f *FString) hasFlag(flag uint8) bool {
	return (f.flags() & flag) > 0
}
func (f *FString) setFlag(flag uint8) {
	f.len = (f.len | int(flag))
}

func (f *FString) realLen() int {

	return f.len & lenMask
}

func (f FString) String() string {
	var str unsafe.Pointer
	var l int
	if f.hasFlag(smallStringOptomized) {
		str = unsafe.Pointer(&f)
		l = f.realLen()
	} else {
		str = f.str
		l = f.realLen()
	}

	y := FString{
		str: str,
		len: l,
	}

	return *(*string)(unsafe.Pointer(&y))
}

func (f *FString) flags() uint8 {
	flags := uint64(f.len & flagsMask)

	return uint8(flags >> (lenSize - flagBits))
}

func (f *FString) toRaw() *rawFString {
	return (*rawFString)(unsafe.Pointer(f))
}

type rawFString [16]byte

func (r *rawFString) len() int {
	x := *(*int)(unsafe.Pointer(&r[8]))

	// take off the high 4 bits
	return x & lenMask
}

func (r *rawFString) str() unsafe.Pointer {
	return unsafe.Pointer(&r[0])
}
