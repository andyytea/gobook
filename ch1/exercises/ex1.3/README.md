```
goos: darwin
goarch: arm64
pkg: github.com/andyytea/gopl/ch1/exercises/ex1.3
BenchmarkIterSmall-8             4106612               264.8 ns/op           344 B/op          8 allocs/op
BenchmarkRangeSmall-8            4478161               268.5 ns/op           344 B/op          8 allocs/op
BenchmarkJoinSmall-8            14639874                80.79 ns/op           64 B/op          1 allocs/op
BenchmarkIterLarge-8              337269              3532 ns/op           15768 B/op         67 allocs/op
BenchmarkRangeLarge-8             336606              3533 ns/op           15768 B/op         67 allocs/op
BenchmarkJoinLarge-8             2175598               558.2 ns/op           448 B/op          1 allocs/op
PASS
ok      github.com/andyytea/gopl/ch1/exercises/ex1.3    8.541s
```