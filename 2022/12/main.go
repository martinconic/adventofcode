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

func main() {
	file, err := os.Open("sample.txt")
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

	steps = 100000000

	visited := map[string]int{}

	cx := sx
	cy := sy
	s := strconv.Itoa(cx) + strconv.Itoa(cy)
	visited[s] += 1
	q := []string{}
	q = append(q, s)

	for {
		ln := len(q)
		if ln == 0 {
			break
		}
		log.Println(q)
		if cx+1 < len(hm) && (hm[cx+1][cy]-hm[cx][cy] == 1 || hm[cx+1][cy]-hm[cx][cy] == 0) {
			if cx+1 == ex && cy == ey {
				if ln < steps {
					steps = ln
				} else {
					q = q[:ln-1]
				}
			} else if count < steps {
				s := strconv.Itoa(cx+1) + strconv.Itoa(cy)
				if _, ok := visited[s]; !ok {
					visited[s] += 1
					q = append(q, s)
				}
			}
		}
		if cx-1 >= 0 && (hm[cx-1][cy]-hm[cx][cy] == 1 || hm[cx-1][cy]-hm[cx][cy] == 0) {
			if cx-1 == ex && cy == ey {
				if ln < steps {
					steps = ln
				} else {
					q = q[:ln-1]
				}

			} else if count < steps {
				s := strconv.Itoa(cx-1) + strconv.Itoa(cy)
				if _, ok := visited[s]; !ok {
					visited[s] += 1
					q = append(q, s)
				}

			}
		}
		if cy+1 < len(hm[0]) && (hm[cx][cy+1]-hm[cx][cy] == 1 || hm[cx][cy+1]-hm[cx][cy] == 0) {
			if cx == ex && cy+1 == ey {
				if ln < steps {
					steps = ln
				} else {
					q = q[:ln-1]

				}

			} else if count < steps {
				s := strconv.Itoa(cx) + strconv.Itoa(cy+1)
				if _, ok := visited[s]; !ok {
					visited[s] += 1
					q = append(q, s)
				}

			}
		}
		if cy-1 >= 0 && (hm[cx][cy-1]-hm[cx][cy] == 1 || hm[cx][cy-1]-hm[cx][cy] == 0) {
			if cx == ex && cy-1 == ey {
				if ln < steps {
					steps = ln
				} else {
					q = q[:ln-1]

				}

			} else if count < steps {
				s := strconv.Itoa(cx) + strconv.Itoa(cy-1)
				if _, ok := visited[s]; !ok {
					visited[s] += 1
					q = append(q, s)
				}

			}
		}
	}
	//	log.Println(hm)
	log.Println(steps)
}
