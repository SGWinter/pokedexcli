package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "charmander bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    " pikachu CHARIZARD squirtLE ",
			expected: []string{"pikachu", "charizard", "squirtle"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		// add more tests
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths do not match: '%v' vs '%v'", actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}
