package cos418_hw1_1

import (
	"fmt"
	"testing"
)

func equal(counts1, counts2 []WordCount) bool {
	if len(counts1) != len(counts2) {
		return false
	}
	for i := range counts1 {
		if counts1[i] != counts2[i] {
			return false
		}
	}
	return true
}

func assertEqual(t *testing.T, answer, expected []WordCount) {
	if !equal(answer, expected) {
		t.Fatal(fmt.Sprintf(
			"Word counts did not match...\nExpected: %v\nActual: %v",
			expected,
			answer))
	}
}

func TestSimple(t *testing.T) {
	answer1 := topWords("simple.txt", 4, 0)
	answer2 := topWords("simple.txt", 5, 4)
	expected1 := []WordCount{
		{"hello", 5},
		{"you", 3},
		{"and", 2},
		{"dont", 2},
	}
	expected2 := []WordCount{
		{"hello", 5},
		{"dont", 2},
		{"everyone", 2},
		{"look", 2},
		{"again", 1},
	}
	assertEqual(t, answer1, expected1)
	assertEqual(t, answer2, expected2)
}

func TestDeclarationOfIndependence(t *testing.T) {
	answer := topWords("declaration_of_independence.txt", 5, 6)
	expected := []WordCount{
		{"people", 10},
		{"states", 8},
		{"government", 6},
		{"powers", 5},
		{"assent", 4},
	}
	assertEqual(t, answer, expected)
}
