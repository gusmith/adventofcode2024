package day2

import (
	"testing"
)

func TestExample(t *testing.T) {
	text := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	expected1 := 2
	expected2 := 4
	p1, p2, err := Day2(text)

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
