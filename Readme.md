# conclusion from the benchmarks of (proto and flat buffers)
- flat buffer over the network is so light and perfectly serilized for higher throughput

## flat buffer benchmark results:

go test ./cmd/bench/... -run=Benchmark -bench=. -benchtime=2s -count=5 -benchmem
goos: darwin
goarch: arm64
pkg: shiv/cmd/bench
BenchmarkFlatBuffer-8   	   56452	     40466 ns/op	    7254 B/op	      99 allocs/op
BenchmarkFlatBuffer-8   	   59589	     40295 ns/op	    7254 B/op	      99 allocs/op
BenchmarkFlatBuffer-8   	   59715	     40193 ns/op	    7252 B/op	      99 allocs/op
BenchmarkFlatBuffer-8   	   60014	     40194 ns/op	    7253 B/op	      99 allocs/op
BenchmarkFlatBuffer-8   	   60003	     40085 ns/op	    7253 B/op	      99 allocs/op
PASS
ok  	shiv/cmd/bench	14.327s

----------------------------------------------------------------
## proto buffer benchmark results:

go test ./bench/... -run=Benchmark -bench=. -benchtime=2s -count=5 -benchmem
goos: darwin
goarch: arm64
pkg: shiva/bench
BenchmarkProtoBuffer-8   	   50097	     46096 ns/op	   11784 B/op	     106 allocs/op
BenchmarkProtoBuffer-8   	   52015	     45915 ns/op	   11782 B/op	     106 allocs/op
BenchmarkProtoBuffer-8   	   52359	     46037 ns/op	   11784 B/op	     106 allocs/op
BenchmarkProtoBuffer-8   	   52330	     45803 ns/op	   11786 B/op	     106 allocs/op
BenchmarkProtoBuffer-8   	   52185	     46172 ns/op	   11783 B/op	     106 allocs/op
PASS
ok  	shiva/bench	14.593s

