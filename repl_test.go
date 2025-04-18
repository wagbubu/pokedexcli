package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	}{
		{
			input: "   hello world   ",
			expected: []string{"hello", "world"},
		},
	}


	for _, c := range cases {
		actual := cleanInput(c.input)

		for i := range actual {
			got := actual[i]
			expected := c.expected[i]

			if got != expected {
				t.Errorf("Expected '%s', but got '%s'", expected, got)
			}
		
		}
	}
	
}