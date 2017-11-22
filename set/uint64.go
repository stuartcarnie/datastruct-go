package set

import (
	"math"
)

const freeKey = 0

// Map is a map-like data-structure for int64s
type Uint64 struct {
	data       []uint64
	fillFactor float64
	threshold  int

	mask uint64

	hasFreeKey bool
	size       int
}

// New returns a map initialized with n spaces and uses the stated fillFactor.
// The map will grow as needed.
func NewUint64(n int, fillFactor float64) *Uint64 {
	capacity := arraySize(n, fillFactor)
	return &Uint64{
		data:       make([]uint64, capacity),
		fillFactor: fillFactor,
		threshold:  int(math.Floor(float64(capacity) * fillFactor)),
		mask:       uint64(capacity - 1),
	}
}

func (s *Uint64) Len() int { return s.size }

func (s *Uint64) Contains(key uint64) bool {
	if key == freeKey {
		return s.hasFreeKey
	}

	ptr := phiMix(key) & s.mask
	k := s.data[ptr]

	if k == key {
		return true
	}

	for k != key && k != freeKey {
		ptr = (ptr + 1) & s.mask
		k = s.data[ptr]
	}

	return k != freeKey
}

// Add adds the specified key to the set.
func (s *Uint64) Add(key uint64) {
	if key == freeKey {
		if !s.hasFreeKey {
			s.size += 1
		}
		s.hasFreeKey = true
		return
	}

	ptr := phiMix(key) & s.mask
	k := s.data[ptr]

	for {
		if k == freeKey {
			s.data[ptr] = key
			if s.size >= s.threshold {
				s.rehash()
			} else {
				s.size += 1
			}
			return
		} else if k == key {
			return
		}
		ptr = (ptr + 1) & s.mask
		k = s.data[ptr]
	}
}

func (s *Uint64) rehash() {
	newCapacity := len(s.data) * 2
	s.threshold = int(math.Floor(float64(newCapacity) * s.fillFactor))
	s.mask = uint64(newCapacity - 1)

	oldData := s.data
	s.data = make([]uint64, newCapacity)

	if s.hasFreeKey {
		s.size = 1
	} else {
		s.size = 0
	}

	for _, o := range oldData {
		if o != freeKey {
			s.Add(o)
		}
	}
}
