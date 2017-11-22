package set

import "testing"

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
			if g := nextPowerOf2(test.v); g != test.x {
				t.Errorf("got=%d, exp=%d", g, test.x)
			}
		})
	}
}

func TestUint64Simple(t *testing.T) {
	m := NewUint64(10, 0.99)
	var i uint64
	for i = 2; i < 20000; i += 2 {
		m.Add(i)
	}
	for i = 2; i < 20000; i += 2 {
		if !m.Contains(i) {
			t.Errorf("didn't contain key")
		}
		if m.Contains(i + 1) {
			t.Errorf("didn't expected key")
		}
	}

	for i = 2; i < 20000; i += 2 {
		m.Add(i)
		if !m.Contains(i) {
			t.Errorf("didn't get expected value")
		}
		if m.Contains(i + 1) {
			t.Errorf("didn't get expected NO-VALUE value")
		}
	}
}

func TestUint64(t *testing.T) {
	m := NewUint64(10, 0.6)

	step := uint64(61)

	var i uint64
	m.Add(0)
	for i = 1; i < 100000000; i += step {
		m.Add(i)

		if !m.Contains(i) {
			t.Errorf("expected set to contain key %d", i)
		}
	}

	for i = 1; i < 100000000; i += step {
		if !m.Contains(i) {
			t.Errorf("expected set to contain key %d", i)
		}

		for j := i + 1; j < i+step; j++ {
			if m.Contains(j) {
				t.Errorf("expected no key for %d", j)
			}
		}
	}

	if !m.Contains(0) {
		t.Errorf("expected key 0")
	}
}

const (
	max  = 999999999
	step = 9534
)

func fillUint64(m *Uint64) {
	var j uint64
	for j = 0; j < max; j += step {
		m.Add(j)
		for k := j; k < j+16; k++ {
			m.Add(k)
		}

	}
}

func fillStdMap(m map[int64]struct{}) {
	var j int64
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
		m := NewUint64(2048, 0.60)
		fillUint64(m)
	}
}

func BenchmarkStdMapFill(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m := make(map[int64]struct{}, 2048)
		fillStdMap(m)
	}
}

func BenchmarkUint64Contains10PercentHitRate(b *testing.B) {
	var j uint64
	m := NewUint64(2048, 0.60)
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
	var j int64
	m := make(map[int64]struct{}, 2048)
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
	m := NewUint64(2048, 0.60)
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
	var j int64
	m := make(map[int64]struct{}, 2048)
	fillStdMap(m)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sum := int64(0)
		for j = 0; j < max; j += step {
			if _, ok := m[j]; ok {
				sum += 1
			}
		}
	}
}
