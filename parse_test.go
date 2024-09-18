package flag3_test

import (
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
					Args:    []string{},
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
					Args:    []string{},
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
}
