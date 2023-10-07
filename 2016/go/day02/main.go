package day02

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

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	data := []string{}
	code := []int{}
	x := 1
	y := 1

	for scanner.Scan() {
		data = strings.Split(scanner.Text(), "")
		// log.Println(matrix)

		log.Println(data)

		for i, v := range data {
			if v == "U" {
				if x > 0 {
					x--
				}
			} else if v == "D" {
				if x < len(matrix)-1 {
					x++
				}
			} else if v == "L" {
				if y > 0 {
					y--
				}
			} else if v == "R" {
				if y < len(matrix)-1 {
					y++
				}
			}
			if i == len(data)-1 {
				code = append(code, matrix[x][y])
			}

			log.Println(matrix[x][y])
		}
		log.Println(code)

	}

}
