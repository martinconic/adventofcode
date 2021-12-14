package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var count int
var co int
var temp [][]int
var matrix [][]int
var maxx, maxy int

var folding = [][]string{{"x", "655"}, {"y", "447"}, {"x", "327"}, {"y", "223"}, {"x", "163"}, {"y", "111"},
	{"x", "81"}, {"y", "55"}, {"x", "40"}, {"y", "27"}, {"y", "13"}, {"y", "6"}}

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
		matrix[list[1]][list[0]] = 1
	}

	for _, fold := range folding {
		if fold[0] == "x" {
			foldx, _ := strconv.Atoi(fold[1])
			for i := 0; i < maxx; i++ {
				k := 0
				for j := maxy - 1; j >= foldx; j-- {
					if matrix[k][i] == 1 || matrix[j][i] == 1 {
						matrix[j][i] = 0
						matrix[k][i] = 1
					}
					k++
				}
			}
			maxx -= foldx
		} else {
			foldy, _ := strconv.Atoi(fold[1])
			for i := 0; i < maxy; i++ {
				k := 0
				for j := maxx - 1; j >= foldy; j-- {
					if matrix[i][k] == 1 || matrix[i][j] == 1 {
						matrix[i][j] = 0
						matrix[i][k] = 1
					}
					k++
				}
			}
			maxy -= foldy
		}
	}

	n := 80
	smatrix := make([][]string, n)
	for i := range smatrix {
		smatrix[i] = make([]string, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[j][i] == 1 {
				smatrix[j][i] = "#"
			} else {
				smatrix[j][i] = "."
			}
		}
	}

	for i, value := range smatrix {
		val := value[:n]
		log.Println(val)
		if i == 50 {
			break
		}
	}

	for i, value := range matrix {
		val := value[:n]
		log.Println(val)
		if i == 20 {
			break
		}

	}
}
