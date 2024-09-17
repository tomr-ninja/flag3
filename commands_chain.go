package flag3

type (
	// CommandsChain - Represents a sorted chain of commands.
	CommandsChain struct {
		values []commandWithArgs
		cur    int
	}
	commandWithArgs struct {
		Command string
		Args    []string
	}
)

func (p *CommandsChain) Next() bool {
	if p.cur >= len(p.values)-1 {
		return false
	}

	p.cur++

	return true
}

func (p *CommandsChain) Command() string {
	return p.values[p.cur].Command
}

func (p *CommandsChain) Args() []string {
	return p.values[p.cur].Args
}
