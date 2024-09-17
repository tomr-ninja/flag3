package main

import (
	"flag"
	"strings"

	"github.com/tomr-ninja/flag3"
	"github.com/tomr-ninja/flag3/tree"
)

// try 'go run example/cli/main.go --capitalized=true fizz -i 3'
func main() {
	t := tree.NewCLI()
	t.Subcommand("foo").Subcommand("bar")
	t.Subcommand("fizz")

	cmd, err := flag3.ParseCLI(t)
	if err != nil || !cmd.Next() {
		panic("impossible")
	}

	rootCfg, err := parseRootLevelConfig(cmd.Args())
	if err != nil {
		panic(err)
	}

	if !cmd.Next() {
		panic("no command")
	}

	switch cmd.Command() {
	case "foo":
		err = execFoo(rootCfg, cmd)
	case "fizz":
		var cfg *fizzConfig
		if cfg, err = parseFizzConfig(cmd.Args()); err == nil {
			err = execFizz(rootCfg, cfg)
		}
	}

	if err != nil {
		panic(err)
	}
}

type rootConfig struct {
	capitalized bool
}

func parseRootLevelConfig(args []string) (*rootConfig, error) {
	cfg := &rootConfig{
		capitalized: false,
	}
	fs := flag.NewFlagSet("root", flag.ExitOnError)
	fs.BoolVar(&cfg.capitalized, "capitalized", cfg.capitalized, "capitalize the output")
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	return cfg, nil
}

func execFoo(rootCfg *rootConfig, cmd flag3.CommandsChain) error {
	val := "foo"
	hasSub := cmd.Next()
	if hasSub {
		switch cmd.Command() {
		case "bar":
			val += "bar"
		}
	}

	if rootCfg.capitalized {
		val = strings.ToUpper(val)
	}

	println(val)

	return nil
}

type fizzConfig struct {
	i int
}

func parseFizzConfig(args []string) (*fizzConfig, error) {
	cfg := &fizzConfig{
		i: 0,
	}
	fs := flag.NewFlagSet("fizz", flag.ExitOnError)
	fs.IntVar(&cfg.i, "i", cfg.i, "number of repetitions")
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	return cfg, nil
}

func execFizz(rootCfg *rootConfig, cfg *fizzConfig) error {
	val := "fizzbuzz"
	if rootCfg.capitalized {
		val = strings.ToUpper(val)
	}

	for j := 0; j < cfg.i; j++ {
		println(val)
	}

	return nil
}
