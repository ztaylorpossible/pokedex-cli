package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "here is some text",
			expected: []string{"here", "is", "some", "text"},
		},
		{
			input: "SHOULD BE LOWERCASED",
			expected: []string{"should", "be", "lowercased"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length mismatch - expected: %v (%v), actual: %v (%v)",
			len(c.expected), c.expected, len(actual), actual);
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Fatalf("word mismatch - expected: %v, actual: %v", expectedWord, word)
			}
		}
	}
}
