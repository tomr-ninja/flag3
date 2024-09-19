package flag3_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tomr-ninja/flag3"
)

func TestParseArgs(t *testing.T) {
	t1 := flag3.New("first")
	t1.Subcommand("second")

	t2 := flag3.New("third")
	t2.Subcommand("fourth")

	type (
		commandWithArgs struct {
			Command string
			Args    []string
		}
		testCase struct {
			input string
			chain []commandWithArgs
		}
	)

	testCases := []testCase{
		{
			input: "first",
			chain: []commandWithArgs{
				{
					Command: "first",
					Args:    nil,
				},
			},
		},
		{
			input: "first -a -b",
			chain: []commandWithArgs{
				{
					Command: "first",
					Args:    []string{"-a", "-b"},
				},
			},
		},
		{
			input: "first -a -b second",
			chain: []commandWithArgs{
				{
					Command: "first",
					Args:    []string{"-a", "-b"},
				},
				{
					Command: "second",
					Args:    nil,
				},
			},
		},
		{
			input: "first -a -b second -c -d",
			chain: []commandWithArgs{
				{
					Command: "first",
					Args:    []string{"-a", "-b"},
				},
				{
					Command: "second",
					Args:    []string{"-c", "-d"},
				},
			},
		},
		{
			input: "first -a -b third -c -d", // third is not a valid subcommand of first
			chain: []commandWithArgs{
				{
					Command: "first",
					Args:    []string{"-a", "-b", "third", "-c", "-d"},
				},
			},
		},
		{
			input: "first -a second second",
			chain: []commandWithArgs{
				{
					Command: "first",
					Args:    []string{"-a"},
				},
				{
					Command: "second",
					Args:    []string{"second"},
				},
			},
		},
	}

	t.Run("Parse", func(t *testing.T) {
		for _, tc := range testCases {
			t.Run(tc.input, func(t *testing.T) {
				chain, err := flag3.Parse(strings.Split(tc.input, " "), t1, t2)
				require.NoError(t, err)
				require.True(t, chain.Next())

				for i, expected := range tc.chain {
					require.Equal(t, expected.Command, chain.Command())
					require.Equal(t, expected.Args, chain.Args())

					ok := chain.Next()
					if i < len(tc.chain)-1 {
						require.True(t, ok)
					} else {
						require.False(t, ok)
					}
				}
			})
		}
	})

	t.Run("ParseTo", func(t *testing.T) {
		chain := flag3.CommandsChain{}

		for _, tc := range testCases {
			t.Run(tc.input, func(t *testing.T) {
				err := flag3.ParseTo(&chain, strings.Split(tc.input, " "), t1, t2)
				require.NoError(t, err)
				require.True(t, chain.Next())

				for i, expected := range tc.chain {
					require.Equal(t, expected.Command, chain.Command())
					require.Equal(t, expected.Args, chain.Args())

					ok := chain.Next()
					if i < len(tc.chain)-1 {
						require.True(t, ok)
					} else {
						require.False(t, ok)
					}
				}
			})
		}
	})
}

func TestParseArgsError(t *testing.T) {
	tree := flag3.New("first").Subcommand("second")

	_, err := flag3.Parse([]string{"unknown"}, tree)
	require.Error(t, err)

	_, err = flag3.Parse([]string{}, tree)
	require.Error(t, err)

	chain := flag3.CommandsChain{}

	err = flag3.ParseTo(&chain, []string{"unknown"}, tree)
	require.Error(t, err)

	err = flag3.ParseTo(&chain, []string{}, tree)
	require.Error(t, err)
}

func TestParseArgsCLI(t *testing.T) {
	osArgs := os.Args
	os.Args = []string{"first", "-a", "second", "-b"}
	defer func() {
		os.Args = osArgs
	}()

	tree := flag3.New("first")
	tree.Subcommand("second")

	cmd, err := flag3.ParseCLI(tree)
	require.NoError(t, err)

	require.True(t, cmd.Next())
	require.Equal(t, "first", cmd.Command())
	require.Equal(t, []string{"-a"}, cmd.Args())

	require.True(t, cmd.Next())
	require.Equal(t, "second", cmd.Command())
	require.Equal(t, []string{"-b"}, cmd.Args())
}
