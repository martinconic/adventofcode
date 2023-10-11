package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	sum := 0
	for scanner.Scan() {
		matrix := [][]string{}

		data1 := strings.Fields(scanner.Text())
		scanner.Scan()
		data2 := strings.Fields(scanner.Text())
		scanner.Scan()
		data3 := strings.Fields(scanner.Text())

		matrix = append(matrix, data1)
		matrix = append(matrix, data2)
		matrix = append(matrix, data3)

		for i := 0; i < 3; i++ {
			x1, _ := strconv.Atoi(matrix[0][i])
			x2, _ := strconv.Atoi(matrix[1][i])
			x3, _ := strconv.Atoi(matrix[2][i])

			if x1+x2 > x3 && x1+x3 > x2 && x2+x3 > x1 {
				sum++
			}
		}
	}
	log.Println(sum)
}

func main() {
	partOne()
}
