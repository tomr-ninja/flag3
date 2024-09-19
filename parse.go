package flag3

import (
	"errors"
	"os"
)

const defaultCapacity = 3 // empirical optimal value

var (
	ErrNoMatchedTree = errors.New("no tree matched")
	ErrNoArgs        = errors.New("no args provided")
)

func ParseTo(chain *CommandsChain, args []string, trees ...*Tree) error {
	if len(args) < 1 {
		return ErrNoArgs
	}

	t := pickTree(args[0], trees)
	if t == nil {
		return ErrNoMatchedTree
	}

	lastCommandPos := 0
	if cap(chain.values) < defaultCapacity {
		chain.values = make([]commandWithArgs, defaultCapacity)
	}
	chain.values = chain.values[:1]
	chain.values[0].Command = args[0]

	for i := 1; i < len(args); i++ {
		next := pickTree(args[i], t.Next())
		if next == nil { // an argument that is not a command
			continue
		}

		lastCommandArgs := args[lastCommandPos+1 : i]
		if cap(chain.values[len(chain.values)-1].Args) < len(lastCommandArgs) {
			chain.values[len(chain.values)-1].Args = make([]string, len(lastCommandArgs))
		}
		chain.values[len(chain.values)-1].Args = chain.values[len(chain.values)-1].Args[:len(lastCommandArgs)]
		copy(chain.values[len(chain.values)-1].Args, lastCommandArgs)

		lastCommandPos = i
		chain.values = chain.values[:len(chain.values)+1]
		chain.values[len(chain.values)-1].Command = args[i]

		t = next
	}

	lastCommandArgs := args[lastCommandPos+1:]
	if cap(chain.values[len(chain.values)-1].Args) < len(lastCommandArgs) {
		chain.values[len(chain.values)-1].Args = make([]string, len(lastCommandArgs))
	}
	chain.values[len(chain.values)-1].Args = chain.values[len(chain.values)-1].Args[:len(lastCommandArgs)]
	copy(chain.values[len(chain.values)-1].Args, lastCommandArgs)

	chain.cur = -1

	return nil
}

func Parse(args []string, trees ...*Tree) (CommandsChain, error) {
	if len(args) < 1 {
		return CommandsChain{}, ErrNoArgs
	}

	res := make([]commandWithArgs, 0, defaultCapacity)

	t := pickTree(args[0], trees)
	if t == nil {
		return CommandsChain{}, ErrNoMatchedTree
	}

	lastCommandPos := 0
	res = append(res, commandWithArgs{Command: args[0]})

	for i := 1; i < len(args); i++ {
		next := pickTree(args[i], t.Next())
		if next == nil { // an argument that is not a command
			continue
		}

		if commandArg := args[lastCommandPos+1 : i]; len(commandArg) > 0 {
			res[len(res)-1].Args = commandArg
		}
		lastCommandPos = i
		res = append(res, commandWithArgs{Command: args[i]})

		t = next
	}

	if commandArgs := args[lastCommandPos+1:]; len(commandArgs) > 0 {
		res[len(res)-1].Args = commandArgs
	}

	return CommandsChain{values: res, cur: -1}, nil
}

// ParseCLI - Parse wrapper to simplify CLI parsing.
func ParseCLI(t *Tree) (CommandsChain, error) {
	return Parse(os.Args, t)
}

func pickTree(arg string, nodes []*Tree) *Tree {
	for _, t := range nodes {
		if t.Command() == arg {
			return t
		}
	}

	return nil
}
