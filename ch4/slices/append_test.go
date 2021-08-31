package slices

import (
    "testing"
    "fmt"
)

func benchmarkAppend(b *testing.B, size int, appendFn func ([]int, int) []int) {
    s := make([]int, 0, 0)
    for i := 0; i < b.N; i++ {
        for n := 0; n < size; n++ {
            s = appendFn(s, n)
        }
    }
}

func logGrowthManual (b *testing.B, size int) {
    var temp_cap int
    s := make([]int, 0, 0)
    for n := 0; n < size; n++ {
        temp_cap = cap(s)
        s = appendInt(s, n)
        if temp_cap != cap(s) {
            b.Logf("(manual) cap: %d", cap(s))
        }
    }
} 

func logGrowthStandard (b *testing.B, size int) {
    var temp_cap int
    s := make([]int, 0, 0)
    for n := 0; n < size; n++ {
        temp_cap = cap(s)
        s = append(s, n)
        if temp_cap != cap(s) {
            b.Logf("(standard) cap: %d", cap(s))
        }
    }
}


func BenchmarkAppend(b *testing.B) {
    sizes := []int{1<<7, 1<<10, 1<<16, 1<<20}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("manual-%d", size), func (b *testing.B) {
            benchmarkAppend(b, size, appendInt)
        })
        b.Run(fmt.Sprintf("standard-%d", size), func (b *testing.B) {
            s := make([]int, 0, 0)
            for i := 0; i < b.N; i++ {
                for n := 0; n < size; n++ {
                    s = append(s, n)
                }
            }
        })    
    }
    last := len(sizes) - 1
    logGrowthStandard(b, sizes[last])
    logGrowthManual(b, sizes[last])
}
