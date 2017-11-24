// +build !fast

package set

import "math"

func (s *Uint64) newData(n int) {
	s.threshold = int(math.Floor(float64(n) * s.fillFactor))
	s.mask = uint64(n - 1)
	s.data = make([]uint64, n)
}

func (s *Uint64) ptr(i uint64) uint64 {
	return i & s.mask
}

func (s *Uint64) inc(base, p uint64) uint64 {
	return (base + p) & s.mask
}

func (s *Uint64) get(base uint64) uint64 {
	return s.data[base]
}

func (s *Uint64) set(base, key uint64) {
	s.data[base] = key
}
