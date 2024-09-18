package flag3

import "os"

// Tree - Represents a command tree.
type Tree struct {
	command string
	next    []*Tree
}

// New - Create a new tree with the root command name set to the provided name.
func New(name string) *Tree {
	return &Tree{command: name}
}

func (t *Tree) Command() string {
	return t.command
}

func (t *Tree) Next() []*Tree {
	return t.next
}

func (t *Tree) MaxPathLen() int {
	if t == nil {
		return 0
	}

	maxSubPath := 0
	for _, next := range t.next {
		maxSubPath = max(maxSubPath, next.MaxPathLen())
	}

	return maxSubPath + 1
}

// NewCLI - Create a new tree with the root command name set to the current executable name.
// Just a convenience function to avoid having to pass os.Args[0] to New.
func NewCLI() *Tree {
	return &Tree{command: os.Args[0]}
}

func (t *Tree) Subcommand(name string) *Tree {
	next := &Tree{command: name}
	t.next = append(t.next, next)

	return next
}
