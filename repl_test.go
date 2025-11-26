package main

import (
	"testing"
)

var cases = []struct {
	input    string
	expected []string
}{
	{
		input:    "  ",
		expected: []string{},
	},
	{
		input:    "  hello  ",
		expected: []string{"hello"},
	},
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "  HellO  World  ",
		expected: []string{"hello", "world"},
	},
}

func TestCleanInput(t *testing.T) {
	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("expected: %v, got: %v", expectedWord, word)
			}
		}
	}
}
