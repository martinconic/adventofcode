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

	// rock , paper , scissors
	eq := map[string]string{"A": "R", "B": "P", "C": "S"}

	beats := map[string]string{"A": "S", "B": "R", "C": "P"}

	beated := map[string]string{"A": "P", "B": "S", "C": "R"}

	score := map[string]int{"R": 1, "P": 2, "S": 3}

	for scanner.Scan() {
		v := strings.Fields(scanner.Text())

		if v[1] == "X" {
			sum += score[beats[v[0]]]
		} else if v[1] == "Y" {
			sum += score[eq[v[0]]] + 3
		} else {
			sum += score[beated[v[0]]] + 6
		}

	}

	log.Println(sum)
}
