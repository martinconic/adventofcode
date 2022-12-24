package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var hm [][]int
var sx, sy, ex, ey int

var steps int

func is_visited(s string, visited []string) bool {
	for i := 0; i < len(visited); i++ {
		if visited[i] == s {
			return true
		}
	}

	return false
}

func findE(cx, cy int, visited []string) {
	s := strconv.Itoa(cx) + strconv.Itoa(cy)
	visited = append(visited, s)
	ln := len(visited)

	//	log.Println(visited)
	if cx+1 < len(hm) && (hm[cx+1][cy]-hm[cx][cy] == 1 || hm[cx+1][cy]-hm[cx][cy] == 0) {
		if cx+1 == ex && cy == ey {
			if ln < steps {
				steps = ln
			}

		} else {
			if ln < steps {
				ns := strconv.Itoa(cx+1) + strconv.Itoa(cy)
				if !is_visited(ns, visited) {
					findE(cx+1, cy, visited)
				}
			}
		}
	}
	if cx-1 >= 0 && (hm[cx-1][cy]-hm[cx][cy] == 1 || hm[cx-1][cy]-hm[cx][cy] == 0) {
		if cx-1 == ex && cy == ey {
			if ln < steps {
				steps = ln
			}

		} else {
			if ln < steps {
				ns := strconv.Itoa(cx-1) + strconv.Itoa(cy)
				if !is_visited(ns, visited) {

					findE(cx-1, cy, visited)

				}
			}
		}
	}
	if cy+1 < len(hm[0]) && (hm[cx][cy+1]-hm[cx][cy] == 1 || hm[cx][cy+1]-hm[cx][cy] == 0) {
		if cx == ex && cy+1 == ey {
			if ln < steps {
				steps = ln
			}

		} else {
			if ln < steps {
				ns := strconv.Itoa(cx) + strconv.Itoa(cy+1)
				if !is_visited(ns, visited) {

					findE(cx, cy+1, visited)

				}
			}
		}
	}
	if cy-1 >= 0 && (hm[cx][cy-1]-hm[cx][cy] == 1 || hm[cx][cy-1]-hm[cx][cy] == 0) {
		if cx == ex && cy-1 == ey {
			if ln < steps {
				steps = ln
			}

		} else {
			if ln < steps {
				ns := strconv.Itoa(cx) + strconv.Itoa(cy-1)
				if !is_visited(ns, visited) {
					findE(cx, cy-1, visited)
				}
			}
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("can not open file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	a := map[string]int{"S": 0, "a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14,
		"o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26, "E": 27}

	count := 0
	for scanner.Scan() {
		heights := strings.Split(scanner.Text(), "")

		line := []int{}
		for _, h := range heights {
			if h == "S" {
				sy = len(line)
				sx = count
			} else if h == "E" {
				ey = len(line)
				ex = count
			}

			line = append(line, a[h])

		}

		hm = append(hm, line)

		count++
	}

	steps = 10000000

	visited := []string{}

	findE(sx, sy, visited)

	//	log.Println(hm)
	log.Println(steps)
}
