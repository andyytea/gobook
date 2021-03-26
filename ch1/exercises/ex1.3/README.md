```
andytsang@andy ex1.3 % go test -bench=.                                                                               
goos: darwin
goarch: arm64
pkg: github.com/andyytea/gobook/ch1/exercises/ex1.3
<<<<<<< HEAD
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
=======
BenchmarkIter10-8       	1000000000	         0.0000022 ns/op	       0 B/op	       0 allocs/op
BenchmarkRange10-8      	1000000000	         0.0000057 ns/op	       0 B/op	       0 allocs/op
BenchmarkJoin10-8       	1000000000	         0.0000012 ns/op	       0 B/op	       0 allocs/op
BenchmarkIter1000-8     	1000000000	         0.0004250 ns/op	       0 B/op	       0 allocs/op
BenchmarkRange1000-8    	1000000000	         0.0001948 ns/op	       0 B/op	       0 allocs/op
BenchmarkJoin1000-8     	1000000000	         0.0000841 ns/op	       0 B/op	       0 allocs/op
BenchmarkIter10000-8    	1000000000	         0.002353 ns/op	            0 B/op	       0 allocs/op
BenchmarkRange10000-8   	1000000000	         0.002587 ns/op	            0 B/op	       0 allocs/op
BenchmarkJoin10000-8    	1000000000	         0.0008995 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/andyytea/gobook/ch1/exercises/ex1.3	0.241s
```
>>>>>>> d4f3c50ac9f2b5545a218ecc33eb72429490639e
