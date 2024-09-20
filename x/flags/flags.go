package flags

import (
	"strings"
)

type Flag struct {
	Name  string
	Value string
}

// ExtractFlags - decompose a slice of strings into a slice of flags - with names and values.
// Only supports flags like '-i 1' and '--hello=world'. Any other format would not be considered a flag.
//
// It's only extraction, not parsing. The values are still just strings.
func ExtractFlags(args []string) (res []Flag, unparsed []string) {
	var (
		flag   Flag
		parsed bool
	)
	for len(args) > 0 {
		flag, args, parsed = extractOne(args)
		if parsed {
			res = append(res, flag)
		} else {
			break
		}
	}

	return res, args
}

func extractOne(args []string) (Flag, []string, bool) {
	f := Flag{}

	switch {
	case len(args) > 1 && len(args[0]) == 2 && args[0][0] == '-' && isLatinLetter(args[0][1]):
		f.Name = args[0]
		f.Value = args[1]
		args = args[2:]

	case len(args[0]) >= 5 && args[0][:2] == "--" && isLatinLetter(args[0][2]):
		div := strings.IndexByte(args[0], '=')
		if div == -1 {
			return Flag{}, args, false
		}
		f.Name = args[0][:div]
		f.Value = args[0][div+1:]
		args = args[1:]

	default:
		return Flag{}, args, false
	}

	if len(args) == 0 {
		args = nil // just seems more semantically correct
	}

	return f, args, true
}

func isLatinLetter(b byte) bool {
	return b >= 65 && b <= 90 || b >= 97 && b <= 122
}
