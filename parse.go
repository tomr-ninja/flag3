package flag3

import (
	"os"

	"github.com/tomr-ninja/flag3/tree"
)

const defaultCapacity = 3 // empirical optimal value

func Parse(args []string, trees ...*tree.Tree) (CommandsChain, error) {
	res := make([]commandWithArgs, 0, defaultCapacity)

	t := pickTree(args[0], trees)
	if t == nil {
		return CommandsChain{}, tree.ErrNoMatchedTree
	}

	lastCommandPos := 0
	res = append(res, commandWithArgs{Command: t.Command()})

	for i := 1; i < len(args); i++ {
		next := pickTree(args[i], t.Next())
		if next == nil { // an argument that is not a command
			continue
		}

		res[len(res)-1].Args = args[lastCommandPos+1 : i]
		lastCommandPos = i
		res = append(res, commandWithArgs{Command: args[i]})

		t = next
	}

	res[len(res)-1].Args = args[lastCommandPos+1:]

	return CommandsChain{values: res, cur: -1}, nil
}

// ParseCLI - Parse wrapper to simplify CLI parsing.
func ParseCLI(t *tree.Tree) (CommandsChain, error) {
	return Parse(os.Args, t)
}

func pickTree(arg string, nodes []*tree.Tree) *tree.Tree {
	for _, t := range nodes {
		if t.Command() == arg {
			return t
		}
	}

	return nil
}
