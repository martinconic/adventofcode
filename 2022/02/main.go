package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	eq := map[string]string{"X": "A", "Y": "B", "Z": "C"}

	beats := map[string]string{"X": "C", "Y": "A", "Z": "B"}

	score := map[string]int{"X": 1, "Y": 2, "Z": 3}

	for scanner.Scan() {
		v := strings.Fields(scanner.Text())

		if v[0] == eq[v[1]] {
			sum += score[v[1]] + 3
		} else if v[0] == beats[v[1]] {
			sum += score[v[1]] + 6
		} else {
			sum += score[v[1]]
		}
	}

	log.Println(sum)
}
