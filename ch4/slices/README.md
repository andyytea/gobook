# Arrays

Arrays and strucures are `aggregate` types meaning that their values are concatenations of other values in memory.
Additionally, arrays and structures are fixed in size, while slices and maps are dynamic data structures that can grow as values are added.

In Go, all elements of an array must have the same type.

Since we're talking about Go, there are a lot of convenience features that helop us create and define composite types.

```go
a := [3]int{}           // [ 0, 0, 0 ]
b := [...]int{1, 2 ,3}  // [ 1, 2, 3 ]
c := [3]int{1, 2}       // [ 1, 2, 0 ]
d := [...]int{0:1, 2:3} // [ 1, 0, 3 ]
```
Note: In Go, arrays are passed by value instead of passed by reference, meaning if we passed an array to a function, all values of the array would be copied to another array. 

# Slices

A slice is written in the form `[]T`. Observe the effect of the following assignment

```go
months := [...]string{"", "January", /* ... */, "December"}

s1 = months[4:7]    // creates slice with len = 3, cap = 9
s2 = months[6:9]    // creates slice with len = 3, cap = 7
```
The underlying implementation of a slice is just a light weight data structure that gives access to a subsequence of the elements of an underlying array. 
A slice has three components: (1) a pointer to the first element of the underlying array that is accessible to the slice, (2) a length, (3) a capacity. 
If we try to slice beyond the capacity of a slice, it causes a panic. Although a slice within capacity and beyond the length extends the slice.

Note: arrays are comparable, but since slices can be considered ligtweight data structures, they are not comparable. Byte slices ([]bytes) can be compared using `bytes.Equal`, but for other types we must write our own comparison.

Inherently, this function would be "deep" in nature as it must check length and compare each value. However, deep equivalence is problematic for two main reasons: 
(1) becase elements of a slice are indirect, making it possible for a slice to contain itself
(2) since they are indirect, a fixed slice value may contain different elements at different times as the contents of the underlying array are modified. (slices contain references/pointers to mutable values in an underlying array)

Note: For reference types (pointers and channels), the `==` operator tests *reference identity*, that is, whether the two entites refer to the same place in memory. A "shallow" equality test for slices could be useful, but the inconsistent treatment of slices and arrays by the `==` operator would be confusing. 

In practice, the safest choice is to disallow slice comparisons altogether. The only legal slice comparison is against `nil` which has 0 capacity and 0 length when initialized. 

Note: not the same as an empty slice

```
goos: darwin
goarch: arm64
pkg: github.com/aydtsang/gobook/ch4/slices
BenchmarkAppend
BenchmarkAppend/manual-128
BenchmarkAppend/manual-128-8         	 1677187	       642.2 ns/op	    2560 B/op	       0 allocs/op
BenchmarkAppend/standard-128
BenchmarkAppend/standard-128-8       	 4602044	      1801 ns/op	    6392 B/op	       0 allocs/op
BenchmarkAppend/manual-1024
BenchmarkAppend/manual-1024-8        	  142672	      8139 ns/op	   30103 B/op	       0 allocs/op
BenchmarkAppend/standard-1024
BenchmarkAppend/standard-1024-8      	  348940	      8797 ns/op	   43164 B/op	       0 allocs/op
BenchmarkAppend/manual-65536
BenchmarkAppend/manual-65536-8       	    1731	    591596 ns/op	 1240602 B/op	       0 allocs/op
BenchmarkAppend/standard-65536
BenchmarkAppend/standard-65536-8     	    7069	    687506 ns/op	 2663372 B/op	       0 allocs/op
BenchmarkAppend/manual-1048576
BenchmarkAppend/manual-1048576-8     	     134	   7470829 ns/op	32051994 B/op	       0 allocs/op
BenchmarkAppend/standard-1048576
BenchmarkAppend/standard-1048576-8   	     585	  15091026 ns/op	50287413 B/op	       0 allocs/op

PASS
ok  	github.com/aydtsang/gobook/ch4/slices	32.934s

LOG OUTPUT:

    append_test.go:36: (standard) cap: 1
    ...
    append_test.go:36: (standard) cap: 1024
    append_test.go:36: (standard) cap: 1280
    append_test.go:36: (standard) cap: 1696
    append_test.go:36: (standard) cap: 2304
    append_test.go:36: (standard) cap: 3072
    append_test.go:36: (standard) cap: 4096
    append_test.go:36: (standard) cap: 5120
    append_test.go:36: (standard) cap: 7168
    append_test.go:36: (standard) cap: 9216
    ...
    append_test.go:36: (standard) cap: 464896
    append_test.go:36: (standard) cap: 581632
    append_test.go:36: (standard) cap: 727040
    append_test.go:36: (standard) cap: 909312
    append_test.go:36: (standard) cap: 1136640

    append_test.go:24: (manual) cap: 1
    ...
    append_test.go:24: (manual) cap: 1024
    append_test.go:24: (manual) cap: 2048
    append_test.go:24: (manual) cap: 4096
    append_test.go:24: (manual) cap: 8192
    append_test.go:24: (manual) cap: 16384
    append_test.go:24: (manual) cap: 32768
    append_test.go:24: (manual) cap: 65536
    append_test.go:24: (manual) cap: 131072
    append_test.go:24: (manual) cap: 262144
    append_test.go:24: (manual) cap: 524288
    append_test.go:24: (manual) cap: 1048576
```

Clearly, the standard implementation of append is making the performance tradeoff of reallocating more frequently, rather than our manual doubling strategy which ends up allocating ridiculous chunks of memory away. 

As a consequence of picking nice exponents of 2 for the size of each array, it appears that the manual method managed to allocate less memory. However, if there were 1 more element, we would have doubled the total capacity just for the single element. Despite making substantially more reallocations, the standard implementation clearly makes dynamic optimizations since it is surprisingly performant.
