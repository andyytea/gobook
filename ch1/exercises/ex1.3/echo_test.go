package echo

import (
	"strings"
	"testing"
)

var exampleString = strings.Fields("go run andy is really cool")

func benchmark(b *testing.B, size int, fn func([]string) string) {
	for i := 0; i < size; i++ {
		fn(exampleString)
	}
}

func BenchmarkIter10(b *testing.B)  { benchmark(b, 10, EchoAddIter) }
func BenchmarkRange10(b *testing.B) { benchmark(b, 10, EchoAddRange) }
func BenchmarkJoin10(b *testing.B)  { benchmark(b, 10, EchoJoin) }

func BenchmarkIter1000(b *testing.B)  { benchmark(b, 1000, EchoAddIter) }
func BenchmarkRange1000(b *testing.B) { benchmark(b, 1000, EchoAddRange) }
func BenchmarkJoin1000(b *testing.B)  { benchmark(b, 1000, EchoJoin) }

func BenchmarkIter10000(b *testing.B)  { benchmark(b, 10000, EchoAddIter) }
func BenchmarkRange10000(b *testing.B) { benchmark(b, 10000, EchoAddRange) }
func BenchmarkJoin10000(b *testing.B)  { benchmark(b, 10000, EchoJoin) }
