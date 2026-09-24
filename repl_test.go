package main

import (
	"testing"
	
)


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "  GoLang  is  awesome  ",
			expected: []string{"golang", "is", "awesome"},
		},
		{
			input: " THIS IS    ALL CAPS",
			expected: []string{"this", "is", "all", "caps"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.expected) != len(actual) {
			t.Errorf("Error: Length doesnt match: expected %v, actual %v", len(c.expected), len(actual))
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error: Values does not match - expected = %v actual = %v", expectedWord, word)

			}
		}
	}
}

