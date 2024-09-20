package flags

import (
	"flag"
	"strconv"
	"testing"
)

func BenchmarkExtractFlags(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var (
			iFlag int64
			aFlag string
		)
		flags, _ := ExtractFlags([]string{"-i", "42", "-a", "hello"})
		for _, f := range flags {
			switch f.Name {
			case "-i":
				iFlag, _ = strconv.ParseInt(f.Value, 10, 64)
			case "-a":
				aFlag = f.Value
			}
		}
		_ = iFlag
		_ = aFlag
	}
}

func BenchmarkFlagSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fs := flag.NewFlagSet("test", flag.ContinueOnError)
		var (
			iFlag int
			aFlag string
		)
		fs.IntVar(&iFlag, "i", 0, "")
		fs.StringVar(&aFlag, "a", "", "")

		_ = fs.Parse([]string{"-i", "42", "-a", "hello"})
	}
}
