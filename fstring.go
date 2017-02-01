package fstring

import (
	"log"
	"unsafe"
)

const (
	pointerSize     = 8
	smallDataSize   = (pointerSize * 2) - 1
	smallStringFlag = 1 << 7
	lenMask         = 0xFF >> 4
)

type stringStructSmall struct {
	str   [smallDataSize]byte
	flags byte
}

func fromString(s string) stringStructSmall {
	b := []byte(s)
	l := len(b)
	var x stringStructSmall
	if l <= smallDataSize {
		copy(x.str[:], b)
		x.flags = setSmallStringFlag(toLen(l))
	} else {
		x = *(*stringStructSmall)(unsafe.Pointer(&s))
	}

	return x

}

func eq(a, b stringStructSmall) bool {
	if isSmallStringSet(a.flags) {
		ua := *(*uint64)(unsafe.Pointer(&a))
		ub := *(*uint64)(unsafe.Pointer(&b))
		if ua != ub {
			return false
		}

		ua = *(*uint64)(unsafe.Pointer(&a.str[8]))
		ub = *(*uint64)(unsafe.Pointer(&b.str[8]))

		return ua == ub

	}

	ax := toString(a)
	bx := toString(b)
	log.Println(ax, bx)
	return toString(a) == toString(b)
}

func eqs(a, b string) bool {
	return a == b
}

func toString(s stringStructSmall) string {
	if isSmallStringSet(s.flags) {
		return string(s.str[:lenInFlags(s.flags)])
	}
	log.Println(s)

	return *(*string)(unsafe.Pointer(&struct {
		p unsafe.Pointer
		l int
	}{
		p: unsafe.Pointer(&s.str),
		l: getLargeLen(&s),
	}))
}

func getLargeLen(s *stringStructSmall) int {
	return *(*int)(unsafe.Pointer(&s.str[pointerSize]))
}

func canBeSmallString(size int) bool {
	return size <= smallDataSize
}

func setSmallStringFlag(flags byte) byte {
	return flags | smallStringFlag
}

func isSmallStringSet(flags byte) bool {
	return (flags & smallStringFlag) > 0
}

func stringStructSmallOf(sp *string) *stringStructSmall {
	return (*stringStructSmall)(unsafe.Pointer(sp))
}

func isSmallString(s *string) bool {
	return isSmallStringSet(stringStructSmallOf(s).flags)
}

func lenInFlags(b byte) int {
	return int(b & lenMask)
}

func toLen(i int) byte {
	return byte(i)
}
