package main

import (
	"bufio"
	"log"
	"os"
)

func first(scanner *bufio.Scanner) {
	scanner.Scan()
	s := scanner.Text()

	floor := 0
	for _, v := range s {
		if string(v) == "(" {
			floor++
		} else {
			floor--
		}
	}
	log.Println(floor)
}

func second(scanner *bufio.Scanner) {
	scanner.Scan()
	s := scanner.Text()

	floor := 0
	position := 0
	for _, v := range s {
		if floor == -1 {
			log.Println(position)
			break
		}
		if string(v) == "(" {
			floor++
		} else {
			floor--
		}

		position++

	}
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Can not open file")

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	//first(scanner)
	second(scanner)

}
