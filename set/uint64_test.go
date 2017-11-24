package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stuartcarnie/datastruct/set"
)

func fillRange(t *testing.T, s *set.Uint64, min, max, step uint64) {
	t.Helper()
	for key := min; key < max; key += step {
		s.Add(key)
	}
}

func testRange(t *testing.T, min, max, step uint64, fn func(uint64) bool) {
	for key := min; key < max; key += step {
		if !fn(key) {
			return
		}
	}
}

func TestUint64_Remove(t *testing.T) {
	s := set.NewUint64(16, 0.99)

	fillRange(t, s, 100, 200, 10)
	assert.True(t, s.Contains(100))

	s.Remove(110)
	testRange(t, 100, 200, 10, func(key uint64) bool {
		if key == 110 {
			assert.False(t, s.Contains(key))
		} else {
			assert.True(t, s.Contains(key))
		}
		return true
	})
}

func TestUint64Simple(t *testing.T) {
	s := set.NewUint64(10, 0.99)
	var i uint64
	for i = 2; i < 20000; i += 2 {
		s.Add(i)
	}
	for i = 2; i < 20000; i += 2 {
		if !s.Contains(i) {
			t.Errorf("didn't contain key")
		}
		if s.Contains(i + 1) {
			t.Errorf("didn't expect key")
		}
	}

	for i = 2; i < 20000; i += 2 {
		s.Add(i)
		if !s.Contains(i) {
			t.Errorf("didn't contain key")
		}
		if s.Contains(i + 1) {
			t.Errorf("didn't expect key")
		}
	}
}

func TestUint64(t *testing.T) {
	s := set.NewUint64(10, 0.6)

	step := uint64(61)

	var i uint64
	s.Add(0)
	for i = 1; i < 100000000; i += step {
		s.Add(i)

		if !s.Contains(i) {
			t.Errorf("expected set to contain key %d", i)
		}
	}

	for i = 1; i < 100000000; i += step {
		if !s.Contains(i) {
			t.Errorf("expected set to contain key %d", i)
		}

		for j := i + 1; j < i+step; j++ {
			if s.Contains(j) {
				t.Errorf("expected no key for %d", j)
			}
		}
	}

	if !s.Contains(0) {
		t.Errorf("expected key 0")
	}
}

const (
	max  = 999999999
	step = 9534
)

func fillUint64(m *set.Uint64) {
	var j uint64
	for j = 0; j < max; j += step {
		m.Add(j)
		for k := j; k < j+16; k++ {
			m.Add(k)
		}
	}
}

func fillStdMap(m map[uint64]struct{}) {
	var j uint64
	for j = 0; j < max; j += step {
		m[j] = struct{}{}
		for k := j; k < j+16; k++ {
			m[k] = struct{}{}
		}
	}
}

func BenchmarkUint64Fill(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m := set.NewUint64(2048, 0.60)
		fillUint64(m)
	}
}

func BenchmarkStdMapFill(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m := make(map[uint64]struct{}, 2048)
		fillStdMap(m)
	}
}

func BenchmarkUint64_Fill10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s := set.NewUint64(16, 0.99)
		for j := uint64(200); j < 210; j++ {
			s.Add(j)
		}
	}
}

func BenchmarkStdMap_Fill10(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s := make(map[uint64]struct{}, 16)
		for j := uint64(200); j < 210; j++ {
			s[j] = struct{}{}
		}
	}
}

func BenchmarkUint64Contains10PercentHitRate(b *testing.B) {
	var j uint64
	m := set.NewUint64(2048, 0.60)
	fillUint64(m)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sum := int64(0)
		for j = 0; j < max; j += step {
			for k := j; k < 10; k++ {
				if m.Contains(k) {
					sum += 1
				}
			}
		}
	}
}

func BenchmarkStdMapGet10PercentHitRate(b *testing.B) {
	var j uint64
	m := make(map[uint64]struct{}, 2048)
	fillStdMap(m)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sum := int64(0)
		for j = 0; j < max; j += step {
			for k := j; k < 10; k++ {
				if _, ok := m[k]; ok {
					sum += 1
				}
			}
		}
	}
}

func BenchmarkUint64Contains100PercentHitRate(b *testing.B) {
	var j uint64
	m := set.NewUint64(2048, 0.60)
	fillUint64(m)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sum := int64(0)
		for j = 0; j < max; j += step {
			if m.Contains(j) {
				sum += 1
			}
		}
	}
}

func BenchmarkStdMapGet100PercentHitRate(b *testing.B) {
	var j uint64
	m := make(map[uint64]struct{}, 2048)
	fillStdMap(m)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sum := uint64(0)
		for j = 0; j < max; j += step {
			if _, ok := m[j]; ok {
				sum += 1
			}
		}
	}
}

func BenchmarkUint64AddRemove(b *testing.B) {
	s := set.NewUint64(10000, 0.99)
	for i := uint64(0); i < 10000; i++ {
		s.Add(i)
	}
	for i := uint64(0); i < 10000; i += 4 {
		s.Remove(i)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for i := uint64(0); i < 10000; i++ {
			s.Add(i)
		}
		for i := uint64(0); i < 10000; i += 4 {
			s.Remove(i)
		}
	}
}

func BenchmarkStdMapAddRemove(b *testing.B) {
	m := make(map[uint64]struct{}, 10000)
	for i := uint64(0); i < 10000; i++ {
		m[i] = struct{}{}
	}
	for i := uint64(0); i < 10000; i += 4 {
		delete(m, i)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for i := uint64(0); i < 10000; i++ {
			m[i] = struct{}{}
		}
		for i := uint64(0); i < 10000; i += 4 {
			delete(m, i)
		}
	}
}
