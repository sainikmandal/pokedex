package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{"hello world", []string{"hello", "world"}},
		{"  hello  world  ", []string{"hello", "world"}},
		{"Charmander Bulbasaur PIKACHU", []string{"charmander", "bulbasaur", "pikachu"}},
		{"   spaced   out   ", []string{"spaced", "out"}},
		{"Mixed CASE   Example ", []string{"mixed", "case", "example"}},
		{"", []string{}},     // Edge case: empty string
		{"    ", []string{}}, // Edge case: only spaces
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("For input %q, expected %v but got %v", c.input, c.expected, actual)
		}
	}
}
