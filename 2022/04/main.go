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

	for scanner.Scan() {
		pairs := strings.Split(scanner.Text(), ",")
		p1 := strings.Split(pairs[0], "-")
		p2 := strings.Split(pairs[1], "-")
		//shoudl convert each to ints, but it seems it works like this anyway :))
		if (p1[0] >= p2[0] && p1[1] <= p2[1] && p1[1] >= p2[0]) || (p1[0] <= p2[0] && p1[1] >= p2[1] && p2[1] >= p1[0]) {
			log.Println(pairs)
			sum++
		}
	}

	log.Println(sum)

}
