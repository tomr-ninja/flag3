package flag3

import (
	"strings"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	benchPrompt := "first -a -b second --c=100 -d 42"
	args := strings.Split(benchPrompt, " ")

	t1 := New("first")
	t1.Subcommand("second")

	t2 := New("third")
	t2.Subcommand("fourth")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = Parse(args, t1, t2)
	}
}
