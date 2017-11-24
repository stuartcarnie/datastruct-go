package set

const freeKey = 0

// Map is a map-like data-structure for int64s
type Uint64 struct {
	mask       uint64
	fillFactor float64
	threshold  int
	size       int

	data    []uint64
	dataptr uintptr

	hasFreeKey bool
}

// New returns a map initialized with n spaces and uses the stated fillFactor.
// The map will grow as needed.
func NewUint64(n int, fillFactor float64) *Uint64 {
	capacity := arraySize(n, fillFactor)
	s := &Uint64{fillFactor: fillFactor}
	s.newData(capacity)
	return s
}

func (s *Uint64) Len() int { return s.size }

func (s *Uint64) Contains(key uint64) bool {
	if key == freeKey {
		return s.hasFreeKey
	}

	ptr := s.ptr(phiMix(key))
	k := s.get(ptr)

	for {
		if k == key {
			return true
		}
		if k == freeKey {
			return false
		}
		ptr = s.inc(ptr, 1)
		k = s.get(ptr)
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

	ptr := s.ptr(phiMix(key))
	k := s.get(ptr)

	for {
		if k == freeKey {
			s.set(ptr, key)
			if s.size >= s.threshold {
				s.rehash(len(s.data) * 2)
			} else {
				s.size += 1
			}
			return
		} else if k == key {
			return
		}
		ptr = s.inc(ptr, 1)
		k = s.get(ptr)
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

	ptr := s.ptr(phiMix(key))
	k := s.get(ptr)

	for {
		if k == key {
			s.shiftKeys(ptr)
			s.size--
			return true
		} else if k == freeKey {
			return false
		}
		ptr = s.inc(ptr, 1)
		k = s.get(ptr)
	}
}

func (s *Uint64) shiftKeys(ptr uint64) uint64 {
	var last, slot uint64
	var key uint64
	for {
		last, ptr = ptr, s.inc(ptr, 1)
		for {
			key = s.get(ptr)
			if key == freeKey {
				s.set(last, freeKey)
				return last
			}

			slot = s.ptr(phiMix(key)) // current key starting slot
			lastIsSameBucket := last >= slot
			keyIsDifferentBucket := slot > ptr

			if last <= ptr {
				if lastIsSameBucket || keyIsDifferentBucket {
					break
				}
			} else if lastIsSameBucket && keyIsDifferentBucket {
				break
			}
			ptr = s.inc(ptr, 1)
		}
		s.set(last, key)
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
