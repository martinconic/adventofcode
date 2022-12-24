package main

import (
	"bufio"
	"log"
	"os"
)

var hm [][]int
var sx, sy, ex, ey int

var steps int

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}
}
