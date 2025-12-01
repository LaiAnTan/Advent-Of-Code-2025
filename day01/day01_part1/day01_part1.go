package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	
	dialPos := 50
	zeroCount := 0
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

		switch direction {
			case 'L':
				dialPos = (dialPos - magnitude) % 100
			case 'R':
				dialPos = (dialPos + magnitude) % 100
		}

		if dialPos == 0 {
			zeroCount += 1
		}
	}

	fmt.Println("Answer:", zeroCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}