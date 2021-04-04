package echo

import (
	"strings"
	"testing"
)

const exampleString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"

var example = strings.Fields(exampleString)

func benchmark(b *testing.B, size int, fn func([]string) string) {
	for i := 0; i < b.N; i++ {
		fn(example[1:size])
	}
}

func test(t *testing.T, fn func([]string) string) {
	got := fn(example)
	want := exampleString
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func TestEchoIter(t *testing.T)  { test(t, EchoIter) }
func TestEchoRange(t *testing.T) { test(t, EchoRange) }
func TestEchoJoin(t *testing.T)  { test(t, EchoJoin) }

func BenchmarkIterSmall(b *testing.B)  { benchmark(b, 10, EchoIter) }
func BenchmarkRangeSmall(b *testing.B) { benchmark(b, 10, EchoRange) }
func BenchmarkJoinSmall(b *testing.B)  { benchmark(b, 10, EchoJoin) }

func BenchmarkIterLarge(b *testing.B)  { benchmark(b, len(example), EchoIter) }
func BenchmarkRangeLarge(b *testing.B) { benchmark(b, len(example), EchoRange) }
func BenchmarkJoinLarge(b *testing.B)  { benchmark(b, len(example), EchoJoin) }
