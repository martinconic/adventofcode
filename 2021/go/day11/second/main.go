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

func main() {

	matrix = [][]int{}

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "")
		ld := []int{}

		for _, d := range data {
			n, _ := strconv.Atoi(d)
			ld = append(ld, n)
		}

		matrix = append(matrix, ld)

	}
	step := 0
	for {
		step++
		for i, list := range matrix {
			for j := range list {
				matrix[i][j] += 1
				if matrix[i][j] == 10 {
					increaseEnergy(i-1, j)
					increaseEnergy(i+1, j)
					increaseEnergy(i, j-1)
					increaseEnergy(i, j+1)
					increaseEnergy(i-1, j-1)
					increaseEnergy(i-1, j+1)
					increaseEnergy(i+1, j-1)
					increaseEnergy(i+1, j+1)
				}
			}
		}

		allzero := 0
		for i, list := range matrix {
			for j, value := range list {
				if value > 9 {
					matrix[i][j] = 0
					allzero++
				}
			}
		}
		if allzero == 100 {
			log.Println(step)
			break
		}
	}

	// log.Println(matrix)
	// log.Println(count)

}

func increaseEnergy(i, j int) {
	if i >= 0 && i < 10 && j >= 0 && j < 10 && matrix[i][j] <= 9 {
		matrix[i][j] += 1
		if matrix[i][j] == 10 {
			increaseEnergy(i-1, j)
			increaseEnergy(i+1, j)
			increaseEnergy(i, j-1)
			increaseEnergy(i, j+1)
			increaseEnergy(i-1, j-1)
			increaseEnergy(i-1, j+1)
			increaseEnergy(i+1, j-1)
			increaseEnergy(i+1, j+1)
		}
	}
}
