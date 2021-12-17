package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var count int
var matrix [][]int
var l1, l2 int

func naive() {
	matrix = [][]int{}

	file, err := os.Open("../sample.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "")
		list := []int{}
		for _, value := range data {
			val, _ := strconv.Atoi(value)
			list = append(list, val)
		}

		matrix = append(matrix, list)

	}

	l1 = len(matrix)
	l2 = len(matrix[0])

	count = findLowestPath(0, 0, 0)

	log.Println(count)
}

func findLowestPath(i, j, cost int) int {

	if i == l1-1 && j == l2-1 {
		cost += matrix[i][j]
	} else if i == l1-1 && j < l2-1 {
		cost += findLowestPath(i, j+1, matrix[i][j+1])
	} else if j == l2-1 && i < l1-1 {
		cost += findLowestPath(i+1, j, matrix[i+1][j])
	} else {
		cost += min(findLowestPath(i+1, j, matrix[i+1][i]), findLowestPath(i, j+1, matrix[i][j+1]))
	}
	return cost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
