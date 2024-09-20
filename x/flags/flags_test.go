package flags

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFlags(t *testing.T) {
	type testCase struct {
		in               []string
		expectedFlags    []Flag
		expectedUnparsed []string
	}

	cases := []testCase{
		{
			in:               nil,
			expectedFlags:    nil,
			expectedUnparsed: nil,
		},
		{
			in:               []string{"hello"},
			expectedFlags:    nil,
			expectedUnparsed: []string{"hello"},
		},
		{
			in:               []string{"-i"},
			expectedFlags:    nil,
			expectedUnparsed: []string{"-i"}, // no value so still unable to parse
		},
		{
			in:               []string{"--kek"},
			expectedFlags:    nil,
			expectedUnparsed: []string{"--kek"}, // no value so still unable to parse
		},
		{
			in:               []string{"-i", "1"},
			expectedFlags:    []Flag{{Name: "-i", Value: "1"}},
			expectedUnparsed: nil,
		},
		{
			in:               []string{"--i=1"},
			expectedFlags:    []Flag{{Name: "--i", Value: "1"}},
			expectedUnparsed: nil,
		},
		{
			in:               []string{"-i", "1", "-b", "x"},
			expectedFlags:    []Flag{{Name: "-i", Value: "1"}, {Name: "-b", Value: "x"}},
			expectedUnparsed: nil,
		},
		{
			in:               []string{"--i=1", "--b=x"},
			expectedFlags:    []Flag{{Name: "--i", Value: "1"}, {Name: "--b", Value: "x"}},
			expectedUnparsed: nil,
		},
		{
			in:               []string{"-i", "1", "-b", "x", "hello"},
			expectedFlags:    []Flag{{Name: "-i", Value: "1"}, {Name: "-b", Value: "x"}},
			expectedUnparsed: []string{"hello"},
		},
		{
			in:               []string{"--i=1", "--b=x", "hello"},
			expectedFlags:    []Flag{{Name: "--i", Value: "1"}, {Name: "--b", Value: "x"}},
			expectedUnparsed: []string{"hello"},
		},
	}
	for _, tc := range cases {
		flags, unparsed := ExtractFlags(tc.in)
		assert.Equal(t, tc.expectedFlags, flags)
		assert.Equal(t, tc.expectedUnparsed, unparsed)
	}
}

func TestIsLatinLetter(t *testing.T) {
	assert.True(t, isLatinLetter('A'))
	assert.True(t, isLatinLetter('a'))
	assert.True(t, isLatinLetter('Z'))
	assert.True(t, isLatinLetter('z'))

	assert.False(t, isLatinLetter('1'))
	assert.False(t, isLatinLetter('-'))
	assert.False(t, isLatinLetter('*'))
	assert.False(t, isLatinLetter('!'))
}
