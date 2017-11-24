package set

import "math"

const freeKey = 0

// Uint64 represents a distinct set of uint64 values.
type Uint64 struct {
	mask       uint64
	fillFactor float64
	threshold  int
	size       int

	data []uint64

	hasFreeKey bool
}

// NewUint64 returns a set initialized to store n elements without allocating.
func NewUint64(n int, fillFactor float64) *Uint64 {
	capacity := arraySize(n, fillFactor)
	s := &Uint64{fillFactor: fillFactor}
	s.newData(capacity)
	return s
}

func (s *Uint64) newData(n int) {
	s.threshold = int(math.Floor(float64(n) * s.fillFactor))
	s.mask = uint64(n - 1)
	s.data = make([]uint64, n)
}

func (s *Uint64) Len() int { return s.size }

func (s *Uint64) Contains(key uint64) bool {
	if key == freeKey {
		return s.hasFreeKey
	}

	ptr := phiMix(key) & s.mask
	k := s.data[ptr]

	for {
		if k == key {
			return true
		}
		if k == freeKey {
			return false
		}
		ptr = (ptr + 1) & s.mask
		k = s.data[ptr]
	}
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
				s.rehash(len(s.data) * 2)
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

func (s *Uint64) Remove(key uint64) bool {
	if key == freeKey {
		if !s.hasFreeKey {
			return false
		}
		s.hasFreeKey = false
		s.size--
		return true
	}

	ptr := phiMix(key) & s.mask
	k := s.data[ptr]

	for {
		if k == key {
			s.shiftKeys(ptr)
			s.size--
			return true
		} else if k == freeKey {
			return false
		}
		ptr = (ptr + 1) & s.mask
		k = s.data[ptr]
	}
}

func (s *Uint64) shiftKeys(ptr uint64) uint64 {
	var last, slot uint64
	var key uint64
	for {
		last, ptr = ptr, (ptr+1)&s.mask
		for {
			key = s.data[ptr]
			if key == freeKey {
				s.data[last] = freeKey
				return last
			}

			slot = phiMix(key) & s.mask // current key starting slot
			lastIsSameBucket := last >= slot
			keyIsDifferentBucket := slot > ptr

			if last <= ptr {
				if lastIsSameBucket || keyIsDifferentBucket {
					break
				}
			} else if lastIsSameBucket && keyIsDifferentBucket {
				break
			}
			ptr = (ptr + 1) & s.mask
		}
		s.data[last] = key
	}
}

func (s *Uint64) rehash(size int) {
	old := s.data
	s.newData(size)

	if s.hasFreeKey {
		s.size = 1
	} else {
		s.size = 0
	}

	for _, o := range old {
		if o != freeKey {
			s.Add(o)
		}
	}
}
