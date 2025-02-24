package day1

import (
	"testing"
)

func TestExample(t *testing.T) {
	text := `3   4
4   3
2   5
1   3
3   9
3   3`

	expected1 := 11
	expected2 := 31
	p1, p2, err := Day1(text)

	if err != nil {
		t.Fatalf("Error %v", err)
	}

	if expected1 != p1 {
		t.Fatalf("Part 1, got %v expected %v", p1, expected1)
	}

	if expected2 != p2 {
		t.Fatalf("Part 2, got %v expected %v", p2, expected2)
	}

}
