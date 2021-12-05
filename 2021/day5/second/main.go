package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const n = 10

func Split(r rune) bool {
	return r == ' ' || r == '-' || r == '>' || r == ','
}

func main() {
	var lmatrix [][]string
	var matrix [n][n]int
	count := 0

	file, err := os.Open("../sample.txt")

	if err != nil {
		log.Fatalf("Can not open file")

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list := strings.FieldsFunc(scanner.Text(), Split)
		lmatrix = append(lmatrix, list)

	}

	log.Println(lmatrix)

	for _, m := range lmatrix {
		x1, _ := strconv.Atoi(m[0])
		y1, _ := strconv.Atoi(m[1])
		x2, _ := strconv.Atoi(m[2])
		y2, _ := strconv.Atoi(m[3])

		if x1 == x2 {
			if y1 >= y2 {
				for i := y1; i >= y2; i-- {
					matrix[i][x1] += 1
				}
			} else {
				for i := y1; i <= y2; i++ {
					matrix[i][x1] += 1
				}
			}
		} else if y1 == y2 {
			if x1 <= x2 {
				for i := x1; i <= x2; i++ {
					matrix[y1][i] += 1
				}
			} else {
				for i := x1; i >= x2; i-- {
					matrix[y1][i] += 1
				}
			}
		} else if x1 == y1 && x2 == y2 {
			if x1 > x2 {
				for i := x1; i >= x2; i-- {
					matrix[i][i] += 1
				}
			} else {
				for i := x1; i <= x2; i++ {
					matrix[i][i] += 1
				}
			}
		} else if x1 == y2 && x2 == y1 {
			if x1 > x2 {
				for i := x1; i >= x2; i-- {
					matrix[y2-i][i] += 1

				}
			} else {
				for i := x1; i <= x2; i++ {
					matrix[y2-i][i] += 1
				}
			}
		} else if x1-y1 == x2-y2 {
			if x1 > x2 {
				j := 0
				for i := x1; i >= x2; i-- {
					matrix[y1-j][x1-j] += 1
					j++
				}
			} else {
				j := 0
				for i := x1; i <= x2; i++ {
					matrix[y1+j][x1+j] += 1
					j++
				}
			}
		} else if x1+y1 == x2+y2 {
			if x1 < x2 {
				j := 0
				for i := x1; i <= x2; i++ {
					matrix[y1-j][i] += 1
					j++
				}
			} else {
				j := 0
				for i := x1; i >= x2; i-- {
					matrix[y1+j][i] += 1
					j++
				}
			}
		}
	}

	log.Println(matrix)
	for _, row := range matrix {
		log.Println(row)
	}

	for _, row := range matrix {
		for _, value := range row {
			if value > 1 {
				count++
			}
		}
	}

	log.Println(count)

}
