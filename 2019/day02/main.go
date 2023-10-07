package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	log.Println("Day2:")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	strvalues := strings.Split(scanner.Text(), ",")
	values := []int{}
	for _, v := range strvalues {
		intValue, _ := strconv.Atoi(v)
		values = append(values, intValue)
	}
	values[1] = 12
	values[2] = 2
	log.Println(values)

	for i := 0; i < len(values)-3; i += 4 {
		if values[i] == 99 {
			log.Println(values[0])
			return
		}
		if values[i] == 1 {
			values[values[i+3]] = values[values[i+1]] + values[values[i+2]]
		}
		if values[i] == 2 {
			values[values[i+3]] = values[values[i+1]] * values[values[i+2]]
		}
		log.Println(values)
	}
}

func partTwo() {
	log.Println("Day2:")

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	strvalues := strings.Split(scanner.Text(), ",")
	firstValues := []int{}
	values := []int{}
	for _, v := range strvalues {
		intValue, _ := strconv.Atoi(v)
		firstValues = append(firstValues, intValue)
	}
	firstValues[1] = 12
	firstValues[2] = 2

	values = firstValues
	for i := 0; i < len(firstValues)-3; {
		values = firstValues

		if values[i] == 99 {
			// log.Println(values[0])
			i++
			continue
		}
		if values[i] == 1 {
			values[values[i+3]] = values[values[i+1]] + values[values[i+2]]
		}
		if values[i] == 2 {
			values[values[i+3]] = values[values[i+1]] * values[values[i+2]]
		}

		if values[0] == 19690720 {
			log.Println(values[0])
			log.Println(100*values[i+1] + values[i+2])
			log.Println(100*values[1] + values[2])
			return
		}

		i += 4
	}
}

func main() {
	// partOne()
	partTwo()
}
