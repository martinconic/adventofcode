package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var matrix [][]int
var basins map[string][]int
var basin map[string]int
var l1, l2 int

func main() {
	lowpoints := []int{}
	list := []int{}
	basins = make(map[string][]int)
	basin = make(map[string]int)
	lb := []int{}

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

	l1 = len(matrix)
	l2 = len(matrix[0])

	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
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
				key := fmt.Sprintf("%d,%d", i, j)
				basins[key] = []int{}
				lowpoints = append(lowpoints, matrix[i][j])
			}
		}
	}

	// log.Println(lowpoints)
	// log.Println(len(lowpoints))

	for key := range basins {
		keys := strings.Split(key, ",")
		k0, _ := strconv.Atoi(keys[0])
		k1, _ := strconv.Atoi(keys[1])

		isInBassin(k0, k1, key)

		// log.Println(basin)
		// log.Println("-------------------------")
		lb = append(lb, len(basin))
		basin = make(map[string]int)
	}

	sort.Ints(lb)

	log.Println(lb)

	n := len(lb)
	log.Println(n)

	log.Println(lb[n-1] * lb[n-2] * lb[n-3])
}

func isInBassin(i, j int, k string) {

	if matrix[i][j] != 9 {

		basin[k] = matrix[i][j]

		if i-1 >= 0 && matrix[i][j] < matrix[i-1][j] {
			key := fmt.Sprintf("%d,%d", i-1, j)
			isInBassin(i-1, j, key)

		}

		if i+1 < l1 && matrix[i][j] < matrix[i+1][j] {
			key := fmt.Sprintf("%d,%d", i+1, j)
			isInBassin(i+1, j, key)
		}

		if j-1 >= 0 && matrix[i][j] < matrix[i][j-1] {
			key := fmt.Sprintf("%d,%d", i, j-1)
			isInBassin(i, j-1, key)

		}

		if j+1 < l2 && matrix[i][j] < matrix[i][j+1] {
			key := fmt.Sprintf("%d,%d", i, j+1)
			isInBassin(i, j+1, key)
		}
	}
}
