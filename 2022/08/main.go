package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	forest := [][]int{}

	for scanner.Scan() {
		templ := strings.Split(scanner.Text(), "")

		line := []int{}
		for _, v := range templ {
			value, _ := strconv.Atoi(v)
			line = append(line, value)
		}

		forest = append(forest, line)
	}

	//	log.Println(forest)

	n := len(forest)
	sumv := 4*(n-2) + 4
	log.Println(sumv)
	for l := 1; l < n-1; l++ {
		for c := 1; c < n-1; c++ {
			visible := []int{1, 1, 1, 1}
			for i := 0; i < c; i++ {
				if forest[l][i] >= forest[l][c] {
					visible[0] = 0
					break
				}
			}

			for i := c + 1; i < n; i++ {
				if forest[l][i] >= forest[l][c] {
					visible[1] = 0
					break
				}
			}

			for i := 0; i < l; i++ {
				if forest[i][c] >= forest[l][c] {
					visible[2] = 0
				}
			}

			for i := l + 1; i < n; i++ {
				if forest[i][c] >= forest[l][c] {
					visible[3] = 0
				}
			}

			for _, v := range visible {
				if v == 1 {
					//log.Println(forest[l][c], visible)
					sumv++
					break
				}
			}
		}
	}

	log.Println(sumv)

}
