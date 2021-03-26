```
andytsang@andy ex1.3 % go test -bench=.                                                                               
goos: darwin
goarch: arm64
pkg: github.com/andyytea/gobook/ch1/exercises/ex1.3
BenchmarkIter10-8       	1000000000	         0.0000044 ns/op
BenchmarkRange10-8      	1000000000	         0.0000026 ns/op
BenchmarkJoin10-8       	1000000000	         0.0000029 ns/op
BenchmarkIter1000-8     	1000000000	         0.0002419 ns/op
BenchmarkRange1000-8    	1000000000	         0.0002168 ns/op
BenchmarkJoin1000-8     	1000000000	         0.0001756 ns/op
BenchmarkIter10000-8    	1000000000	         0.002683 ns/op
BenchmarkRange10000-8   	1000000000	         0.002360 ns/op
BenchmarkJoin10000-8    	1000000000	         0.0008110 ns/op
PASS
ok  	github.com/andyytea/gobook/ch1/exercises/ex1.3	0.330s
```
