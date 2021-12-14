package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var count int
var temp [][]int
var matrix [][]int
var maxx, maxy int
var foldx = 655

func main() {

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(data[0])
		y, _ := strconv.Atoi(data[1])
		list := []int{x, y}

		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
		temp = append(temp, list)
	}
	maxx++
	maxy++
	log.Println(maxx, maxy)

	matrix = make([][]int, maxy)
	for i := range matrix {
		matrix[i] = make([]int, maxx)
	}

	for _, list := range temp {
		// log.Println(list[0], list[1])
		matrix[list[1]][list[0]] = 1
	}

	for i := 0; i < maxy; i++ {
		k := 0

		for j := maxx - 1; j >= foldx; j-- {
			if matrix[i][k] == 1 || matrix[i][j] == 1 {
				matrix[i][k] = 1
				count++
			}
			k++
		}
	}

	log.Println(count)

}
