package flag3_test

import (
	"fmt"
	"testing"

	"github.com/tomr-ninja/flag3"

	"github.com/stretchr/testify/require"
)

func TestTree(t *testing.T) {
	t.Run("level 1", func(t *testing.T) {
		t1 := flag3.New("1")
		t1.Subcommand("1.1")

		require.Equal(t, "1", t1.Command())
		require.Equal(t, "1.1", t1.Next()[0].Command())
	})

	t.Run("level 2", func(t *testing.T) {
		t2 := flag3.New("2")
		t2.Subcommand("2.1")
		t2.Subcommand("2.2")

		require.Equal(t, "2", t2.Command())
		require.Equal(t, "2.1", t2.Next()[0].Command())
		require.Equal(t, "2.2", t2.Next()[1].Command())
	})

	t.Run("level 3", func(t *testing.T) {
		t3 := flag3.New("3")
		t3.Subcommand("3.1")
		t32 := t3.Subcommand("3.2")
		t32.Subcommand("3.2.1")
		t32.Subcommand("3.2.2")

		require.Equal(t, "3", t3.Command())
		require.Equal(t, "3.1", t3.Next()[0].Command())
		require.Equal(t, "3.2", t3.Next()[1].Command())
		require.Equal(t, "3.2.1", t3.Next()[1].Next()[0].Command())
		require.Equal(t, "3.2.2", t3.Next()[1].Next()[1].Command())
	})
}

func TestTree_MaxPathLen(t *testing.T) {
	type testCase struct {
		tree func() *flag3.Tree
		want int
	}

	testCases := []testCase{
		{
			tree: func() *flag3.Tree {
				return nil
			},
			want: 0,
		},
		{
			tree: func() *flag3.Tree {
				return flag3.New("1")
			},
			want: 1,
		},
		{
			tree: func() *flag3.Tree {
				tree := flag3.New("1")
				tree.Subcommand("1.1")

				return tree
			},
			want: 2,
		},
		{
			tree: func() *flag3.Tree {
				tree := flag3.New("1")
				tree.Subcommand("1.1").Subcommand("1.1.1")
				tree.Subcommand("1.2")

				return tree
			},
			want: 3,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, tc.tree().MaxPathLen())
		})
	}
}
