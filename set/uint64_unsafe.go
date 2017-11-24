// +build fast

package set

import (
	"math"
	"reflect"
	"unsafe"
)

func (s *Uint64) newData(n int) {
	s.threshold = int(math.Floor(float64(n) * s.fillFactor))
	s.mask = uint64(n*8 - 1)
	s.data = make([]uint64, n)
	s.dataptr = (*reflect.SliceHeader)(unsafe.Pointer(&s.data)).Data
}

func (s *Uint64) ptr(i uint64) uint64 {
	return (i * 8) & s.mask
}

func (s *Uint64) inc(base, p uint64) uint64 {
	return (base + (p * 8)) & s.mask
}

func (s *Uint64) get(base uint64) uint64 {
	return *(*uint64)(unsafe.Pointer(s.dataptr + uintptr(base)))
}

func (s *Uint64) set(base, key uint64) {
	*(*uint64)(unsafe.Pointer(s.dataptr + uintptr(base))) = key
}
