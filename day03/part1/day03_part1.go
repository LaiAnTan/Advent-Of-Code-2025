package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

		i := 0
		j := 1

		l_max := 0
		r_max := 0

		// brute force lmao
		for i < len(list) - 1 {

			if l_max < list[i] {
				l_max = list[i]
				r_max = 0
			}

			for ; j < len(list); j += 1 {

				if r_max < list[j] {
					r_max = list[j]
				}
			}

			i += 1
			j = i + 1
		}

		// fmt.Println("l_max:", l_max, "r_max:", r_max)

		sum += l_max * 10 + r_max
	}

	fmt.Println("Total Joltage:", sum)
}