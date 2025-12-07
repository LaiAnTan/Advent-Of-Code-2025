package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getOffsets(operators_line string) []int {

	offsets := []int{}
	
	counter := -1

	for _, c := range operators_line {

		if c == '+' || c == '*' {
			offsets = append(offsets, counter)
			counter = -1
		}

		counter += 1
	}
	offsets = append(offsets, counter + 1)

	return offsets[1:]
}

func parseOperands(lines []string, offsets []int) [][]int {

	problems := [][]int {}

	for i := 0; i < len(offsets); i += 1 {
		problems = append(problems, []int {})
		for j := 0; j < offsets[i]; j += 1 {
			problems[i] = append(problems[i], 0)
		}
	}

	fmt.Println(offsets)

	for _, line := range lines {
		begin := 0
		idx := 0

		for idx < len(offsets) {

			num := line[begin:begin + offsets[idx]]

			for j, c := range num {

				if c == ' ' {
					continue
				}

				n, err := strconv.Atoi(string(c))

				if err != nil {
					log.Fatal(err)
				}

				problems[idx][j] = problems[idx][j] * 10 + n
			}
			begin += offsets[idx] + 1
			idx += 1
		}
	}

	fmt.Println(problems)

	return problems
}

func main() {

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		if strings.ContainsAny(line, "*+") {
			break
		}
		lines = append(lines, line)
	}
	
	line := scanner.Text()
	offsets := getOffsets(line)
	problems := parseOperands(lines, offsets)

	total := 0

	idx := 0
	for _, op := range strings.Split(line, " ") {

		if op == "" {
			continue
		}

		result := problems[idx][0]
		for _, num := range problems[idx][1:] {
			switch op {
				case "*":
					result *= num
				case "+":
					result += num
			}
		}
	
		total += result
		idx += 1
	}

	fmt.Println("Total:", total)

}