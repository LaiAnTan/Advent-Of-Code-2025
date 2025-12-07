package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// function to merge two ranges into one range, if possible
// we return the output ranges in a list,
// if there is 1 element -> merged, if 2 elements -> not merged.
func mergeRanges(a []int, b []int) [][]int {

	if a[0] <= b[0] && a[1] >= b[1] {
		return [][]int{{a[0], a[1]}}
	}

	if b[0] <= a[0] && b[1] >= a[1] {
		return [][]int{{b[0], b[1]}}
	}

	if b[0] <= a[1] && b[0] > a[0] {
		return [][]int{{a[0], b[1]}}
	} 

	return [][]int{a, b}
}

func main() {

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	ranges := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		parts := strings.Split(line, "-")

		start, err := strconv.Atoi(parts[0])

		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.Atoi(parts[1])

		if err != nil {
			log.Fatal(err)
		}

		ranges = append(ranges, []int{start, end})
	}

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] != ranges[j][0] {
			return ranges[i][0] < ranges[j][0]
		}
		return ranges[i][1] < ranges[j][1]
	})

	fmt.Println("Sorted:", ranges)

	i := 0
	j := 1
	for j < len(ranges) {

		fmt.Println(i, j)

		result := mergeRanges(ranges[i], ranges[j])

		if len(result) == 1 {
			// insert into list, remove the original 2 elements
			ranges = slices.Insert(ranges, i, result[0])
			ranges = slices.Delete(ranges, i + 1, i + 3)
		} else if len(result) == 2 {
			i += 1
		}
		
		j = i + 1

	} 

	fmt.Println("Merged into", len(ranges), "elements:", ranges)

	sum := 0
	for _, r := range ranges {
		sum += (r[1] - r[0] + 1)
	}

	fmt.Println("Total:", sum)
}