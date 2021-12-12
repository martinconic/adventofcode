package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var count int
var matrix [][]int
var cmap map[string][]string
var smallonce bool = true

func main() {

	matrix = [][]int{}
	cmap = make(map[string][]string)

	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Can not open file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "-")
		if _, ok := cmap[data[0]]; ok {
			cmap[data[0]] = append(cmap[data[0]], data[1])
		} else {
			cmap[data[0]] = []string{data[1]}
		}

		if _, ok := cmap[data[1]]; ok {
			cmap[data[1]] = append(cmap[data[1]], data[0])
		} else {
			cmap[data[1]] = []string{data[0]}
		}
	}
	log.Println(cmap)

	for _, val := range cmap["start"] {
		path := "start" + "," + val
		findpath(val, path, []string{val})
	}

	log.Println(count)

}

func findpath(value, path string, twice []string) {
	for _, val := range cmap[value] {
		nottwice := !istwice(path, twice)
		if !strings.Contains(path, strings.ToLower(val)) || (nottwice && val != "start") {
			t := []string{}
			pth := path + "," + val
			t = append(twice, val)

			if val == "end" {
				log.Println(pth)
				count++
			} else {
				findpath(val, pth, t)
			}
		}
	}
}

func istwice(path string, twice []string) bool {
	for _, v := range twice {
		if strings.Contains(path, strings.ToLower(v)) {
			if strings.Count(path, v) == 2 {
				return true
			}
		}
	}
	return false
}
