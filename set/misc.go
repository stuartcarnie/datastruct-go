package set

import (
	"math"
	"math/bits"
)

const phi = 0x9E3779B9

func phiMix(x uint64) uint64 {
	h := x * phi
	return h ^ (h >> 16)
}

func nextPowerOf2(x uint64) uint64 {
	return 1 << uint(bits.Len64(x))
}

func arraySize(exp int, fill float64) int {
	s := nextPowerOf2(uint64(math.Ceil(float64(exp) / fill)))
	if s < 2 {
		s = 2
	}
	return int(s)
}
