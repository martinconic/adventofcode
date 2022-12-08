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
	score := 1

	for l := 1; l < n-1; l++ {
		for c := 1; c < n-1; c++ {
			temp := 1
			for i := c - 1; i >= 0; i-- {
				if forest[l][i] >= forest[l][c] || i == 0 {
					temp = temp * (c - i)
					break
				}
			}

			for i := c + 1; i < n; i++ {
				if forest[l][i] >= forest[l][c] || i == n-1 {
					temp = temp * (i - c)
					break
				}
			}

			for i := l - 1; i >= 0; i-- {
				if forest[i][c] >= forest[l][c] || i == 0 {
					temp = temp * (l - i)
					break
				}
			}

			for i := l + 1; i < n; i++ {
				if forest[i][c] >= forest[l][c] || i == n-1 {
					temp = temp * (i - l)
					break
				}
			}

			if temp > score {
				score = temp
				//log.Println(l, c, forest[l][c], score)
			}

		}
	}

	log.Println(score)

}
