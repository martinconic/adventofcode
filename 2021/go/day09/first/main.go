package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var count int
var prevcount int

func main() {
	lowpoints := []int{}
	matrix := [][]int{}
	list := []int{}

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "")
		// log.Println(bytesa)
		for _, value := range data {
			number, _ := strconv.Atoi(value)
			list = append(list, number)
		}

		matrix = append(matrix, list)
		list = []int{}
	}

	// log.Println(matrix)
	l1 := len(matrix)
	l2 := len(matrix[0])

	for i := 0; i < l1; i++ {

		for j := 0; j < l2; j++ {
			// log.Println("--", matrix[i][j], "--")
			smallest := true
			if i-1 >= 0 && matrix[i-1][j] <= matrix[i][j] {
				smallest = false
			} else if i+1 < l1 && matrix[i+1][j] <= matrix[i][j] {
				smallest = false
			} else if j-1 >= 0 && matrix[i][j-1] <= matrix[i][j] {
				smallest = false
			} else if j+1 < l2 && matrix[i][j+1] <= matrix[i][j] {
				smallest = false
			}

			if smallest {
				lowpoints = append(lowpoints, matrix[i][j])
			}
		}
		// log.Println(lowpoints)
		// log.Println(matrix[i])
	}

	log.Println(lowpoints)

	count := 0
	for _, value := range lowpoints {
		count += 1 + value
	}

	log.Println(count)

}
