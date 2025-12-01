package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	
	dialPos := 50
	var timesPassedZero float64 = 0
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// max line length 65536 chars
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		direction := line[0]
		magnitude, err := strconv.Atoi(line[1:])

		if err != nil {
			log.Fatal(err)
		}

		start := dialPos
		
		/*
		Decompose total rotation of the dial into 
		-> (n * rotations from start back to start) + (rotation from start to end)

		The total times passed 0 would be
		-> n + (1 if rotation from start to end passes 0, else 0)
		*/

		n := math.Floor(float64(magnitude) / 100.)
		i := 0

		switch direction {
			case 'L':

				dialPos = (((dialPos - magnitude) % 100) + 100) % 100

				if dialPos == 0 {
					i = 1
					fmt.Println("hit 0")
				} else if start != 0 && start < dialPos {
					// if start < end after a left rotation, we passed 0
					i = 1
					fmt.Println("passed 0")
				}
			case 'R':
				dialPos = (dialPos + magnitude) % 100

				if dialPos == 0 {
					i = 1
					fmt.Println("hit 0")
				} else if dialPos != 0 && start > dialPos {
					// if start > end after a right rotation, we passed 0
					i = 1
					fmt.Println("passed 0")
				} 
		}

		timesPassedZero += n + float64(i)

		fmt.Println(line, start, "->", dialPos, ":", n, "+", i, "->", timesPassedZero)
	}

	fmt.Println("Answer: ", timesPassedZero)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}