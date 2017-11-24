package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPowerOf2(t *testing.T) {
	tests := []struct {
		name string
		v, x uint64
	}{
		{"1", 1, 2},
		{"3", 3, 4},
		{"17", 17, 32},
		{"100", 100, 128},
		{"200000", 200000, 262144},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.x, nextPowerOf2(test.v))
		})
	}
}
