# Go Data Structures

Some experiments with alternate data structures in Go

## datastruct/set

Based on [this article](http://java-performance.info/implementing-world-fastest-java-int-to-int-hash-map/), implemented
`uint64` set data structure.

```
± go test -bench=. -run=- ./set
goos: darwin
goarch: amd64
pkg: github.com/stuartcarnie/datastruct/set
BenchmarkUint64Fill-8                        	      10	 102256594 ns/op	67076179 B/op	      12 allocs/op
BenchmarkStdMapFill-8                        	       5	 288558604 ns/op	54751081 B/op	   73621 allocs/op
BenchmarkUint64Contains10PercentHitRate-8    	   20000	     66783 ns/op	       0 B/op	       0 allocs/op
BenchmarkStdMapGet10PercentHitRate-8         	   20000	     79737 ns/op	       0 B/op	       0 allocs/op
BenchmarkUint64Contains100PercentHitRate-8   	    1000	   2362402 ns/op	       0 B/op	       0 allocs/op
BenchmarkStdMapGet100PercentHitRate-8        	     200	   8232297 ns/op	       0 B/op	       0 allocs/op
```
