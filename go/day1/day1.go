package day1

import (
	"errors"
	"slices"
	"strconv"
	"strings"
)

func filterSpace(lStrings []string) []string {
	res := []string{}
	for i := range lStrings {
		if lStrings[i] != "" {
			res = append(res, lStrings[i])
		}
	}
	return res
}

func slicer(text string) ([]int, []int, error) {
	lines := strings.Split(text, "\n")
	nb_lines := len(lines)
	slice1 := make([]int, nb_lines)
	slice2 := make([]int, nb_lines)
	var err1 error
	var err2 error
	for i := 0; i < nb_lines; i++ {
		line := lines[i]
		splitted := filterSpace(strings.Split(line, " "))
		slice1[i], err1 = strconv.Atoi(strings.TrimSpace(splitted[0]))
		slice2[i], err2 = strconv.Atoi(strings.TrimSpace(splitted[1]))
	}
	return slice1, slice2, errors.Join(err1, err2)
}

func part1Compute(slice1 []int, slice2 []int) (res int) {
	for i := range len(slice1) {
		if slice1[i] < slice2[i] {
			res += slice2[i] - slice1[i]
		} else {
			res += slice1[i] - slice2[i]
		}
	}
	return res
}

func part2Compute(slice1 []int, slice2 []int) (res int) {
	dictSlice2 := make(map[int]int)

	for _, val := range slice2 {
		dictSlice2[val] += val
	}

	for _, val := range slice1 {
		res += dictSlice2[val]
	}

	return res
}

func Day1(text string) (part1 int, part2 int, err error) {

	slice1, slice2, error := slicer(text)

	if error != nil {
		return part1, part2, error
	}

	if len(slice1) != len(slice2) {
		return part1, part2, errors.New("different length slices")
	}

	slices.Sort(slice1)
	slices.Sort(slice2)
	part1 = part1Compute(slice1, slice2)
	part2 = part2Compute(slice1, slice2)
	return part1, part2, nil

}
