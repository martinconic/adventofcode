package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	bingo := false
	numbers := []int{}
	matrix := [][]int{}
	lmatrix := [][][]int{}
	sum := 0
	bingon := 0

	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("Can not open file")

	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	slist := strings.Split(scanner.Text(), ",")
	for _, s := range slist {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	log.Println(numbers)

	scanner.Scan()
	for scanner.Scan() {
		if scanner.Text() == "" {
			lmatrix = append(lmatrix, matrix)
			matrix = [][]int{}
		} else {
			numb := []int{}
			slist := strings.Fields(scanner.Text())
			for _, s := range slist {
				n, _ := strconv.Atoi(s)
				numb = append(numb, n)
			}
			matrix = append(matrix, numb)

		}
	}

	for _, numb := range numbers {
		for _, mat := range lmatrix {
			for _, list := range mat {
				for i, n := range list {
					if numb == n {
						list[i] = -1
					}
				}
			}

			// log.Println(mat)

			var sum1, sum2 int
			for i := 0; i < len(mat); i++ {
				for j := 0; j < len(mat); j++ {
					sum1 += mat[i][j]
					sum2 += mat[j][i]
				}

				if sum1 == -5 || sum2 == -5 {
					bingo = true
					break
				}
			}

			if bingo {
				log.Println(mat)

				for _, list := range mat {
					for _, n := range list {
						if n != -1 {
							sum += n
						}
					}
				}
				log.Println(sum)
				break
			}
		}
		if bingo {
			bingon = numb
			log.Println(numb)
			break
		}
	}

	log.Println(bingon * sum)

}
