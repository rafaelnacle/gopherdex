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
			input: "hello boy",
			expected: []string{
				"hello",
				"boy",
			},
		},
	}

	for _, cs := range cases {
		current := cleanInput(cs.input)
		if len(current) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v",
				len(current),
				len(cs.expected),
			)
			continue
		}
		for i := range current {
			currentWord := current[i]
			expectedWord := cs.expected[i]
			if currentWord != expectedWord {
				t.Errorf("%v doest not equal %v", currentWord, expectedWord)
			}
		}
	}
}
