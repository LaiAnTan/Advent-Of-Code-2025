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

	fmt.Println(ranges)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		n, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		for _, r := range ranges {
			if r[0] <= n && n <= r[1] {
				sum += 1
				break
			}
		}

	}

	fmt.Println("Total fresh:", sum)
	
}