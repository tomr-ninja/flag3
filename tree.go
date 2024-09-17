package flag3

import (
	"errors"
	"os"
)

var ErrNoMatchedTree = errors.New("no tree matched")

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
