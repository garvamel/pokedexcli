package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  catch       squirtle  ",
			expected: []string{"catch", "squirtle"},
		},
		{
			input:    "  inspect pidgey",
			expected: []string{"inspect", "pidgey"},
		},
		{
			input:    "explore basement   ",
			expected: []string{"explore", "basement"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Words don't match")
				return
			}
		}
	}
}
