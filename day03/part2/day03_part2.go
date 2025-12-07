package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"math"
)

func main() {

	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		list := []int{}

		for i := 0; i < len(line); i += 1 {

			curr, err := strconv.Atoi(string(line[i]))

			if err != nil {
				log.Fatal(err)
			}
			list = append(list, curr)
		}

		max_joltage := 0
		max := 0
		max_pos := -1

		fmt.Println("List:", list)

		for i := 11; i >= 0; i -= 1 {
			
			max = 0
			k := len(list) - i
			fmt.Println("Checking:", list[max_pos + 1:k])

			for j := max_pos + 1; j < k; j += 1 {
				if max < list[j] {
					max = list[j]
					max_pos = j
				}
			}

			fmt.Println("Max:", max, "at pos", max_pos)
			max_joltage += max * int(math.Pow(10., float64(i)))
			
		}

		fmt.Println("Joltage:", max_joltage)

		sum += max_joltage
	}

	fmt.Println("Total Joltage:", sum)
}