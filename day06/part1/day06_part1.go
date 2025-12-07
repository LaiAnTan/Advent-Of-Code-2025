package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	operands := [][]int{}
	operators := []string{}

	// init operands array with first line of operands
	// one array for each math problem

	scanner.Scan()
	line := scanner.Text()
	parts := strings.Split(line, " ")

	curr_idx := 0
	for _, part := range parts {

		if len(part) != 0 {
			operands = append(operands, []int{})

			n, err := strconv.Atoi(part)

			if err != nil {
				log.Fatal(err)
			}

			operands[curr_idx] = append(operands[curr_idx], n)
			curr_idx += 1
		}
	}

	// handle all other operands
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		if strings.ContainsAny(line, "*+") {
			break
		}

		curr_idx := 0
		for _, part := range parts {

			if len(part) == 0 {
				continue
			}

			n, err := strconv.Atoi(part)

			if err != nil {
				log.Fatal(err)
			}

			operands[curr_idx] = append(operands[curr_idx], n)
			curr_idx += 1
		}
	}

	// handle operators
	line = scanner.Text()
	parts = strings.Split(line, " ")
	for _, op := range parts {

		if len(op) == 0 {
			continue
		}

		operators = append(operators, op)
	}
	
	fmt.Println("Operands:", operands)
	fmt.Println("Operators:", operators)

	sum := 0
	result := 0

	for i, op := range operators {

		switch op {
			case "+":
				result = 0
				for _, num := range operands[i] {
					result += num
				}

			case "*":
				result = 1
				for _, num := range operands[i] {
					result *= num
				}
		}

		sum += result
	}

	fmt.Println("Sum:", sum)
}