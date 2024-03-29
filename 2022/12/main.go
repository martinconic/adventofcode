package main

import (
	"log"
	"os"
	"bufio"
	"strings"
	"strconv"
	
)

var hm [][]int
var sx, sy, ex, ey int


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
		//log.Println(line)

		count++

	}
	n := len(hm)
	visited := map[string]int{}
	dist : = map[string]int{} 
	q := []string{}

	visited["00"] = 1
	dist["00"] = hm[0][0]
	q = append(q, "00")
	height := 0
	
	for {
		l := len(q)
		if l == 0 {
			break
		}

		
		
		
		//		s := strconv.Itoa(v) + strconv.Itoa(i)
			
	
	}

	
	log.Println(visited)

}
